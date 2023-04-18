package vender

import (
	"encoding/json"
	"errors"

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

type CheckAuthForVenderIdSubResponse struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *AuthResult `json:"AuthResult,omitempty" codec:"AuthResult,omitempty"`
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
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}
	var response CheckAuthForVenderIdResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return false, errors.New(response.Data.ErrorDesc)
	}

	if response.Data.Result == nil {
		return false, errors.New("no result.")
	}

	return response.Data.Result.Auth, nil
}
