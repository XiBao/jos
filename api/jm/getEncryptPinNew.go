package jm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
	"github.com/daviddengcn/ljson"
)

type GetEncryptPinNewRequest struct {
	api.BaseRequest
	Token  string `json:"token,omitempty" codec:"token,omitempty"`   //	京东或者微信token
	Source string `json:"source,omitempty" codec:"source,omitempty"` // 	01:京东App，02：微信
}

type GetEncryptPinNewResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEncryptPinNewData `json:"jingdong_pop_jm_center_user_getEncryptPinNew_responce",omitempty" codec:"jingdong_pop_jm_center_user_getEncryptPinNew_responce",omitempty"`
}

type GetEncryptPinNewData struct {
	Code       string                      `json:"code,omitempty" codec:"code,omitempty"`
	ReturnType *GetEncryptPinNewReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type GetEncryptPinNewReturnType struct {
	Message   string `json:"message,omitempty" codec:"message,omitempty"`     //接口的执行信息
	Pin       string `json:"pin,omitempty" codec:"pin,omitempty"`             //用户pin
	Code      uint64 `json:"code,omitempty" codec:"code,omitempty"`           //状态码
	RequestId string `json:"requestId,omitempty" codec:"requestId,omitempty"` //请求id
}

func GetEncryptPinNew(req *GetEncryptPinNewRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetEncryptPinNewRequest()
	r.SetToken(req.Token)
	r.SetSource(req.Source)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return ``, err
	}
	if len(result) == 0 {
		return ``, errors.New("No result info.")
	}
	var response GetEncryptPinNewResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return ``, err
	}

	if response.ErrorResp != nil {
		return ``, response.ErrorResp
	}

	if response.Data.ReturnType.Code != 0 {
		return ``, errors.New(response.Data.ReturnType.Message)
	}

	return response.Data.ReturnType.Pin, nil

}
