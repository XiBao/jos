package user

import (
	"github.com/XiBao/jos/sdk"
)

type GetUserBaseInfoByEncryPinRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetUserBaseInfoByEncryPinRequest() (req *GetUserBaseInfoByEncryPinRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.get.user.base.info", Params: make(map[string]interface{}, 2)}
	req = &GetUserBaseInfoByEncryPinRequest{
		Request: &request,
	}
	return
}

func (req *GetUserBaseInfoByEncryPinRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *GetUserBaseInfoByEncryPinRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *GetUserBaseInfoByEncryPinRequest) SetLoadType(loadType int) {
	req.Request.Params["loadType"] = loadType
}

func (req *GetUserBaseInfoByEncryPinRequest) GetLoadType() int {
	loadType, found := req.Request.Params["loadType"]
	if found {
		return loadType.(int)
	}
	return 0
}
