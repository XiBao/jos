package jm

import (
	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
)

type GetOpenIdRequest struct {
	BaseRequest
	Source string `json:"source" codec:"source"` //  01:京东App，02：微信
	Token  string `json:"token" codec:"token"`
}

type GetOpenIdResponse struct {
	ErrorResp *ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      GetOpenIdResponseData `json:"jingdong_pop_jm_center_user_getOpenId_responce"`
}

func (r GetOpenIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data.IsError()
}

func (r GetOpenIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return r.Data.Error()
}

type GetOpenIdResponseData struct {
	ReturnType GetOpenIdReturnType `json:"returnType"`
}

func (r GetOpenIdResponseData) IsError() bool {
	return r.ReturnType.IsError()
}

func (r GetOpenIdResponseData) Error() string {
	return r.ReturnType.Error()
}

type GetOpenIdReturnType struct {
	Message   string `json:"message"`
	OpenID    string `json:"open_id"`
	Pin       string `json:"pin"`
	RequestID string `json:"requestId"`
	Code      int    `json:"code"`
}

func (r GetOpenIdReturnType) IsError() bool {
	return r.Code != 0 && r.Code != 200
}

func (r GetOpenIdReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

func GetOpenId(req GetOpenIdRequest) (GetOpenIdReturnType, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewGetOpenIdRequest()
	r.SetSource(req.Source)
	r.SetToken(req.Token)

	var response GetOpenIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return GetOpenIdReturnType{}, err
	}
	return response.Data.ReturnType, nil

}
