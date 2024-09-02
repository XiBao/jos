package token

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/token"
)

type ConverstionJmEncryptPinRequest struct {
	api.BaseRequest
	AppKey     string `json:"appKey,omitempty" codec:"appKey,omitempty"`
	EncryptPin string `json:"encryptPin,omitempty" codec:"encryptPin,omitempty"`
}

type ConverstionJmEncryptPinResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ConverstionJmEncryptPinData `json:"jingdong_TokenToPinCenter_converstionJmEncryptPin_responce,omitempty" codec:"jingdong_TokenToPinCenter_converstionJmEncryptPin_responce,omitempty"`
}

func (r ConverstionJmEncryptPinResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ConverstionJmEncryptPinResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ConverstionJmEncryptPinData struct {
	Code      string                         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ConverstionJmEncryptPinResult `json:"result,omitempty" codec:"result,omitempty"`
}

func (r ConverstionJmEncryptPinData) IsError() bool {
	return r.Code != "0" || r.Result == nil
}

func (r ConverstionJmEncryptPinData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type ConverstionJmEncryptPinResult struct {
	Code         int    `json:"code,omitempty" codec:"code,omitempty"`
	RequestId    string `json:"requestId,omitempty" codec:"requestId,omitempty"`
	Message      string `json:"message,omitempty" codec:"message,omitempty"`
	Data         string `json:"data,omitempty" codec:"data,omitempty"`
	OpenIdSeller string `json:"open_id_seller,omitempty" codec:"open_id_seller,omitempty"`
}

func (r ConverstionJmEncryptPinResult) IsError() bool {
	return r.Code != 200
}

func (r ConverstionJmEncryptPinResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

// 输入单个订单id，得到所有相关订单信息
func ConverstionJmEncryptPin(ctx context.Context, req *ConverstionJmEncryptPinRequest) (*ConverstionJmEncryptPinResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := token.NewConverstionJmEncryptPinRequest()
	r.SetAppKey(req.AppKey)
	r.SetEncryptPin(req.EncryptPin)

	var response ConverstionJmEncryptPinResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
