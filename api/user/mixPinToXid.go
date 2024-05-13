package user

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/user"
)

type MixPinToXidRequest struct {
	api.BaseRequest
	AppKey string `json:"appKey,omitempty" codec:"appKey,omitempty"` // 应用appKey
	MixPin string `json:"mixPin,omitempty" codec:"mixPin,omitempty"` // 加密后的pin
}

type MixPinToXidResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *MixPinToXidData    `json:"jingdong_jos_oauth_rpc_xid_mixPin2Xid_responce,omitempty" codec:"jingdong_jos_oauth_rpc_xid_mixPin2Xid_responce,omitempty"`
}

func (r MixPinToXidResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r MixPinToXidResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type MixPinToXidData struct {
	Code   string             `json:"code,omitempty" codec:"code,omitempty"`
	Result *MixPinToXidResult `json:"result,omitempty" codec:"result,omitempty"`
}

func (r MixPinToXidData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r MixPinToXidData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type MixPinToXidResult struct {
	Code      int    `json:"code,omitempty" codec:"code,omitempty"`
	Data      string `json:"data,omitempty" codec:"data,omitempty"`
	RequestId string `json:"requestId,omitempty" codec:"requestId,omitempty"`
}

func (r MixPinToXidResult) IsError() bool {
	return r.Code != 200
}

func (r MixPinToXidResult) Error() string {
	return sdk.ErrorString(r.Code, r.Data)
}

// 加密pin转为xid
func MixPinToXid(req *MixPinToXidRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := user.NewMixPinToXidRequest()
	r.SetAppKey(req.AppKey)
	r.SetMixPin(req.MixPin)

	var response MixPinToXidResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Data.Result.Data, nil
}
