package jm

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/user"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
)

type GetUserBaseInfoByPinRequest struct {
	api.BaseRequest
	Pin      string `json:"pin,omitempty" codec:"pin,omitempty"`           // 用户标识
	LoadType int    `json:"loadType,omitempty" codec:"loadType,omitempty"` // 加载类型
}

type GetUserBaseInfoByPinResponse struct {
	ErrorResp *api.ErrorResponnse              `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetUserBaseInfoByPinSubResponse `json:"jingdong_vender_shop_query_responce,omitempty" codec:"jingdong_vender_shop_query_responce,omitempty"`
}

func (r GetUserBaseInfoByPinResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetUserBaseInfoByPinResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetUserBaseInfoByPinSubResponse struct {
	Code          string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	UserJosResult *user.UserInfo `json:"shop_jos_result,omitempty" codec:"shop_jos_result,omitempty"`
}

func (r GetUserBaseInfoByPinSubResponse) IsError() bool {
	return r.Code != "0"
}

func (r GetUserBaseInfoByPinSubResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 店铺信息查询
func GetUserBaseInfoByPin(ctx context.Context, req *GetUserBaseInfoByPinRequest) (*user.UserInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetUserBaseInfoByPinRequest()
	r.SetPin(req.Pin)
	r.SetLoadType(req.LoadType)

	var response GetUserBaseInfoByPinResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	user := response.Data.UserJosResult
	user.EncryPin = req.Pin
	return user, nil
}
