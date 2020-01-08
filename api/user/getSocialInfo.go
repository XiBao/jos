package user

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/user"
	"github.com/daviddengcn/ljson"
)

type GetSocialInfoRequest struct {
	api.BaseRequest
	Pin string `json:"pin,omitempty" codec:"pin,omitempty"` // 用户标识
}

type GetSocialInfoResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetSocialInfoData  `json:"jingdong_user_getUserSocialInfo_responce,omitempty" codec:"jingdong_user_getUserSocialInfo_responce,omitempty"`
}

type GetSocialInfoData struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Info      *SocialInfo `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

// 店铺信息查询
func GetSocialInfo(req *GetSocialInfoRequest) (*SocialInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := user.NewGetSocialInfoRequest()
	r.SetPin(req.Pin)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}
	var response GetSocialInfoResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	user := response.Data.Info
	user.EncryPin = req.Pin
	return user, nil
}
