package jm

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
)

type GetEncryptPinNewRequest struct {
	api.BaseRequest
	Token  string `json:"token,omitempty" codec:"token,omitempty"`   //	京东或者微信token
	Source string `json:"source,omitempty" codec:"source,omitempty"` // 	01:京东App，02：微信
}

type GetEncryptPinNewResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEncryptPinNewData `json:"jingdong_pop_jm_center_user_getEncryptPinNew_responce,omitempty" codec:"jingdong_pop_jm_center_user_getEncryptPinNew_responce,omitempty"`
}

func (r GetEncryptPinNewResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetEncryptPinNewResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetEncryptPinNewData struct {
	Code       string                      `json:"code,omitempty" codec:"code,omitempty"`
	ReturnType *GetEncryptPinNewReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r GetEncryptPinNewData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r GetEncryptPinNewData) Error() string {
	if r.ReturnType != nil && r.ReturnType.IsError() {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type GetEncryptPinNewReturnType struct {
	Message   string `json:"message,omitempty" codec:"message,omitempty"`     //接口的执行信息
	Pin       string `json:"pin,omitempty" codec:"pin,omitempty"`             //用户pin
	Code      int64  `json:"code,omitempty" codec:"code,omitempty"`           //状态码
	RequestId string `json:"requestId,omitempty" codec:"requestId,omitempty"` //请求id
}

func (r GetEncryptPinNewReturnType) IsError() bool {
	return r.Code != 0
}

func (r GetEncryptPinNewReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

func GetEncryptPinNew(req *GetEncryptPinNewRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetEncryptPinNewRequest()
	r.SetToken(req.Token)
	r.SetSource(req.Source)

	var response GetEncryptPinNewResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Data.ReturnType.Pin, nil

}
