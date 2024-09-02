package user

import (
	"github.com/XiBao/jos/sdk"
)

type MixPinToXidRequest struct {
	Request *sdk.Request
}

// create new request
func NewMixPinToXidRequest() (req *MixPinToXidRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.oauth.rpc.xid.mixPin2Xid", Params: make(map[string]interface{}, 2)}
	req = &MixPinToXidRequest{
		Request: &request,
	}
	return
}

func (req *MixPinToXidRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *MixPinToXidRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *MixPinToXidRequest) SetMixPin(mixPin string) {
	req.Request.Params["mixPin"] = mixPin
}

func (req *MixPinToXidRequest) GetMixPin() string {
	mixPin, found := req.Request.Params["mixPin"]
	if found {
		return mixPin.(string)
	}
	return ""
}
