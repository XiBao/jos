package token

import (
	"encoding/json"
	"errors"

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

type TokenToPinData struct {
	Code      string               `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string               `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *GetencryptPinResult `json:"getencryptpin_result,omitempty" codec:"getencryptpin_result,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("No result.")
	}
	var response TokenToPinResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	if response.ErrorResp != nil {
		return "", response.ErrorResp
	}
	if response.Data.Code != "0" {
		return "", errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return "", errors.New("No pin.")
	}

	return response.Data.Result.Pin, nil
}
