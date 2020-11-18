package user

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/user"
	"github.com/daviddengcn/ljson"
)

type GetEncryptPinRequest struct {
	api.BaseRequest
	Pin string `json:"pin,omitempty" codec:"pin,omitempty"` // 明文PIN
}

type GetEncryptPinResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEncryptPinData  `json:"jingdong_jos_getEncryptPin_responce,omitempty" codec:"jingdong_jos_getEncryptPin_responce,omitempty"`
}

type GetEncryptPinData struct {
	Code   string               `json:"code,omitempty" codec:"code,omitempty"`
	Result *GetEncryptPinResult `json:"result,omitempty" codec:"result,omitempty"`
}

type GetEncryptPinResult struct {
	Code      int    `json:"code,omitempty" codec:"code,omitempty"`
	Data      string `json:"data,omitempty" codec:"data,omitempty"`
	RequestId string `json:"requestId,omitempty" codec:"requestId,omitempty"`
}

// 明文PIN转加密PIN
func GetEncryptPin(req *GetEncryptPinRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := user.NewGetEncryptPinRequest()
	r.SetPin(req.Pin)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("no result.")
	}
	var response GetEncryptPinResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	if response.ErrorResp != nil {
		return "", response.ErrorResp
	}
	if response.Data == nil || response.Data.Result == nil || response.Data.Result.Data == "" {
		return "", errors.New("no encrypt pin.")
	}

	return response.Data.Result.Data, nil
}
