package jm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/user"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request"
	"github.com/daviddengcn/ljson"
)

type GetJmUserBaseInfoByEncryPinRequest struct {
	api.BaseRequest
	Pin      string `json:"pin,omitempty" codec:"pin,omitempty"`           // 用户标识
	LoadType int    `json:"loadType,omitempty" codec:"loadType,omitempty"` // 加载类型
}

type GetJmUserBaseInfoByEncryPinResponse struct {
	ErrorResp *api.ErrorResponnse                     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetJmUserBaseInfoByEncryPinSubResponse `json:"jingdong_pop_jm_getUserBaseInfoByPin_responce,omitempty" codec:"jingdong_pop_jm_getUserBaseInfoByPin_responce,omitempty"`
}

type GetJmUserBaseInfoByEncryPinSubResponse struct {
	Code          string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	UserJosResult *user.UserInfo `json:"getuserbaseinfobypin_result,omitempty" codec:"getuserbaseinfobypin_result,omitempty"`
}

// 用户信息查询
func GetJmUserBaseInfoByEncryPin(req *GetJmUserBaseInfoByEncryPinRequest) (*user.UserInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetJmUserBaseInfoByEncryPinRequest()
	r.SetPin(req.Pin)
	r.SetLoadType(req.LoadType)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}
	var response GetJmUserBaseInfoByEncryPinResponse
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
