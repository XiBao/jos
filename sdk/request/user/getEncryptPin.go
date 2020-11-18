package user

import (
	"github.com/XiBao/jos/sdk"
)

type GetEncryptPinRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEncryptPinRequest() (req *GetEncryptPinRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.getEncryptPin", Params: make(map[string]interface{}, 1)}
	req = &GetEncryptPinRequest{
		Request: &request,
	}
	return
}

func (req *GetEncryptPinRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *GetEncryptPinRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}
