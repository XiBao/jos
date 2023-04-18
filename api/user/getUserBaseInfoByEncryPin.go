package user

import (
	"encoding/json"
	"errors"

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

type GetUserBaseInfoByEncryPinSubResponse struct {
	Code          string    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	UserJosResult *UserInfo `json:"getuserbaseinfobypin_result,omitempty" codec:"getuserbaseinfobypin_result,omitempty"`
}

// 店铺信息查询
func GetUserBaseInfoByEncryPin(req *GetUserBaseInfoByEncryPinRequest) (*UserInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := user.NewGetUserBaseInfoByEncryPinRequest()
	r.SetPin(req.Pin)
	r.SetLoadType(req.LoadType)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}
	var response GetUserBaseInfoByEncryPinResponse
	err = json.Unmarshal(result, &response)
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
