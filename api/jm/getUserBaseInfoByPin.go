package jm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/user"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
	"github.com/daviddengcn/ljson"
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

type GetUserBaseInfoByPinSubResponse struct {
	Code          string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	UserJosResult *user.UserInfo `json:"shop_jos_result,omitempty" codec:"shop_jos_result,omitempty"`
}

// 店铺信息查询
func GetUserBaseInfoByPin(req *GetUserBaseInfoByPinRequest) (*user.UserInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetUserBaseInfoByPinRequest()
	r.SetPin(req.Pin)
	r.SetLoadType(req.LoadType)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}
	var response GetUserBaseInfoByPinResponse
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

	user := response.Data.UserJosResult
	user.EncryPin = req.Pin
	return user, nil
}
