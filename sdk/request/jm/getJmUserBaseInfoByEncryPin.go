package jm

import (
	"github.com/XiBao/jos/sdk"
)

type GetJmUserBaseInfoByEncryPinRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetJmUserBaseInfoByEncryPinRequest() (req *GetJmUserBaseInfoByEncryPinRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.jm.getUserBaseInfoByPin", Params: make(map[string]interface{}, 2)}
	req = &GetJmUserBaseInfoByEncryPinRequest{
		Request: &request,
	}
	return
}

func (req *GetJmUserBaseInfoByEncryPinRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *GetJmUserBaseInfoByEncryPinRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *GetJmUserBaseInfoByEncryPinRequest) SetLoadType(loadType int) {
	req.Request.Params["loadType"] = loadType
}

func (req *GetJmUserBaseInfoByEncryPinRequest) GetLoadType() int {
	loadType, found := req.Request.Params["loadType"]
	if found {
		return loadType.(int)
	}
	return 0
}
