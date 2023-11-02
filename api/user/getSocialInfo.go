package user

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/user"
)

type GetSocialInfoRequest struct {
	api.BaseRequest
	Pin string `json:"pin,omitempty" codec:"pin,omitempty"` // 用户标识
}

type GetSocialInfoResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetSocialInfoData  `json:"jingdong_user_getUserSocialInfo_responce,omitempty" codec:"jingdong_user_getUserSocialInfo_responce,omitempty"`
}

func (r *GetSocialInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r *GetSocialInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetSocialInfoData struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Info      *SocialInfo `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r GetSocialInfoData) IsError() bool {
	return r.Code != "0"
}

func (r GetSocialInfoData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 店铺信息查询
func GetSocialInfo(req *GetSocialInfoRequest) (*SocialInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := user.NewGetSocialInfoRequest()
	r.SetPin(req.Pin)

	var response GetSocialInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	user := response.Data.Info
	user.EncryPin = req.Pin
	return user, nil
}
