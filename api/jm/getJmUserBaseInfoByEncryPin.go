package jm

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/user"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
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

func (r GetJmUserBaseInfoByEncryPinResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetJmUserBaseInfoByEncryPinResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetJmUserBaseInfoByEncryPinSubResponse struct {
	Code          string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	UserJosResult *user.UserInfo `json:"getuserbaseinfobypin_result,omitempty" codec:"getuserbaseinfobypin_result,omitempty"`
}

func (r GetJmUserBaseInfoByEncryPinSubResponse) IsError() bool {
	return r.Code != "0"
}

func (r GetJmUserBaseInfoByEncryPinSubResponse) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

// 用户信息查询
func GetJmUserBaseInfoByEncryPin(req *GetJmUserBaseInfoByEncryPinRequest) (*user.UserInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetJmUserBaseInfoByEncryPinRequest()
	r.SetPin(req.Pin)
	r.SetLoadType(req.LoadType)

	var response GetJmUserBaseInfoByEncryPinResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	user := response.Data.UserJosResult
	user.EncryPin = req.Pin
	return user, nil
}
