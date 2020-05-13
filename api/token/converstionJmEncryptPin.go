package token

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/token"
	"github.com/daviddengcn/ljson"
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

type ConverstionJmEncryptPinData struct {
	Code      string                         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ConverstionJmEncryptPinResult `json:"result,omitempty" codec:"result,omitempty"`
}

type ConverstionJmEncryptPinResult struct {
	Code         int    `json:"code,omitempty" codec:"code,omitempty"`
	RequestId    string `json:"requestId,omitempty" codec:"requestId,omitempty"`
	Message      string `json:"message,omitempty" codec:"message,omitempty"`
	Data         string `json:"data,omitempty" codec:"data,omitempty"`
	OpenIdSeller string `json:"open_id_seller,omitempty" codec:"open_id_seller,omitempty"`
}

// 输入单个订单id，得到所有相关订单信息
func ConverstionJmEncryptPin(req *ConverstionJmEncryptPinRequest) (*ConverstionJmEncryptPinResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := token.NewConverstionJmEncryptPinRequest()
	r.SetAppKey(req.AppKey)
	r.SetEncryptPin(req.EncryptPin)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result.")
	}
	var response ConverstionJmEncryptPinResponse
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
	if response.Data.Result == nil {
		return nil, errors.New("No pin.")
	}

	return response.Data.Result, nil
}
