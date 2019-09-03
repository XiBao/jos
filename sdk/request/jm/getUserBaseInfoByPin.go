package jm

import (
	"github.com/XiBao/jos/sdk"
)

type GetUserBaseInfoByPinRequest struct {
	Request *sdk.Request
}

// 与GetJmUserBaseInfoByEncryPin相同
// create new request
func NewGetUserBaseInfoByPinRequest() (req *GetUserBaseInfoByPinRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.jm.getUserBaseInfoByPin", Params: make(map[string]interface{}, 2)}
	req = &GetUserBaseInfoByPinRequest{
		Request: &request,
	}
	return
}

func (req *GetUserBaseInfoByPinRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *GetUserBaseInfoByPinRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *GetUserBaseInfoByPinRequest) SetLoadType(loadType int) {
	req.Request.Params["loadType"] = loadType
}

func (req *GetUserBaseInfoByPinRequest) GetLoadType() int {
	loadType, found := req.Request.Params["loadType"]
	if found {
		return loadType.(int)
	}
	return 0
}
