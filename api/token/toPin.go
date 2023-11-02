package token

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/token"
)

type TokenToPinRequest struct {
	api.BaseRequest
	Token  string `json:"token,omitempty" codec:"token,omitempty"`
	Source string `json:"source,omitempty" codec:"source,omitempty"`
}

type TokenToPinResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *TokenToPinData     `json:"jingdong_jos_token_source_to_pin_responce,omitempty" codec:"jingdong_jos_token_source_to_pin_responce,omitempty"`
}

func (r TokenToPinResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r TokenToPinResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type TokenToPinData struct {
	Code      string               `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string               `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *GetencryptPinResult `json:"getencryptpin_result,omitempty" codec:"getencryptpin_result,omitempty"`
}

func (r TokenToPinData) IsError() bool {
	return r.Code != "0" || r.Result == nil
}

func (r TokenToPinData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type GetencryptPinResult struct {
	Pin string `json:"getencryptpin_result,omitempty" codec:"getencryptpin_result,omitempty"`
}

// 输入单个订单id，得到所有相关订单信息
func TokenToPin(req *TokenToPinRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := token.NewTokenToPinRequest()
	r.SetToken(req.Token)
	r.SetSource(req.Source)

	var response TokenToPinResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Data.Result.Pin, nil
}
