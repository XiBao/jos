package user

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/user"
)

type GetUserBaseInfoByEncryPinRequest struct {
	api.BaseRequest
	Pin      string `json:"pin,omitempty" codec:"pin,omitempty"`           // 用户标识
	LoadType int    `json:"loadType,omitempty" codec:"loadType,omitempty"` // 加载类型
}

type GetUserBaseInfoByEncryPinResponse struct {
	ErrorResp *api.ErrorResponnse                   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetUserBaseInfoByEncryPinSubResponse `json:"jingdong_jos_get_user_base_info_responce,omitempty" codec:"jingdong_jos_get_user_base_info_responce,omitempty"`
}

func (r GetUserBaseInfoByEncryPinResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetUserBaseInfoByEncryPinResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetUserBaseInfoByEncryPinSubResponse struct {
	Code          string    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	UserJosResult *UserInfo `json:"getuserbaseinfobypin_result,omitempty" codec:"getuserbaseinfobypin_result,omitempty"`
}

func (r GetUserBaseInfoByEncryPinSubResponse) IsError() bool {
	return r.Code != "0"
}

func (r GetUserBaseInfoByEncryPinSubResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 店铺信息查询
func GetUserBaseInfoByEncryPin(ctx context.Context, req *GetUserBaseInfoByEncryPinRequest) (*UserInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := user.NewGetUserBaseInfoByEncryPinRequest()
	r.SetPin(req.Pin)
	r.SetLoadType(req.LoadType)

	var response GetUserBaseInfoByEncryPinResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	user := response.Data.UserJosResult
	user.EncryPin = req.Pin
	return user, nil
}
