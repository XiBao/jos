package user

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/user"
)

type GetEncryptPinRequest struct {
	api.BaseRequest
	Pin string `json:"pin,omitempty" codec:"pin,omitempty"` // 明文PIN
}

type GetEncryptPinResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetEncryptPinData  `json:"jingdong_jos_getEncryptPin_responce,omitempty" codec:"jingdong_jos_getEncryptPin_responce,omitempty"`
}

func (r GetEncryptPinResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetEncryptPinResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetEncryptPinData struct {
	Code   string               `json:"code,omitempty" codec:"code,omitempty"`
	Result *GetEncryptPinResult `json:"result,omitempty" codec:"result,omitempty"`
}

func (r GetEncryptPinData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r GetEncryptPinData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type GetEncryptPinResult struct {
	Code      int    `json:"code,omitempty" codec:"code,omitempty"`
	Data      string `json:"data,omitempty" codec:"data,omitempty"`
	RequestId string `json:"requestId,omitempty" codec:"requestId,omitempty"`
}

func (r GetEncryptPinResult) IsError() bool {
	return r.Code != 200
}

func (r GetEncryptPinResult) Error() string {
	return sdk.ErrorString(r.Code, r.Data)
}

// 明文PIN转加密PIN
func GetEncryptPin(req *GetEncryptPinRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := user.NewGetEncryptPinRequest()
	r.SetPin(req.Pin)

	var response GetEncryptPinResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Data.Result.Data, nil
}
