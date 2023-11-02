package vender

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type CheckAuthForVenderIdRequest struct {
	api.BaseRequest
	PermCode string `json:"permCode"`
}

type CheckAuthForVenderIdResponse struct {
	ErrorResp *api.ErrorResponnse              `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CheckAuthForVenderIdSubResponse `json:"jingdong_vender_auth_checkAuthForVenderId_responce,omitempty" codec:"jingdong_vender_auth_checkAuthForVenderId_responce,omitempty"`
}

func (r CheckAuthForVenderIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CheckAuthForVenderIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CheckAuthForVenderIdSubResponse struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *AuthResult `json:"AuthResult,omitempty" codec:"AuthResult,omitempty"`
}

func (r CheckAuthForVenderIdSubResponse) IsError() bool {
	return r.Code != "0" || r.Result == nil
}

func (r CheckAuthForVenderIdSubResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type AuthResult struct {
	Success bool `json:"success,omitempty" codec:"success,omitempty"`
	Auth    bool `json:"auth,omitempty" codec:"auth,omitempty"`
}

func CheckAuthForVenderId(req *CheckAuthForVenderIdRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewCheckAuthForVenderIdRequest()
	r.SetPermCode(req.PermCode)

	var response CheckAuthForVenderIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}

	return response.Data.Result.Auth, nil
}
