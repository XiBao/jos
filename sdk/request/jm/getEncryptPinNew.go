package jm

import (
	"github.com/XiBao/jos/sdk"
)

type GetEncryptPinNewRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetEncryptPinNewRequest() (req *GetEncryptPinNewRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.jm.center.user.getEncryptPinNew", Params: make(map[string]interface{}, 2)}
	req = &GetEncryptPinNewRequest{
		Request: &request,
	}
	return
}

func (req *GetEncryptPinNewRequest) SetSource(source string) {
	req.Request.Params["source"] = source
}

func (req *GetEncryptPinNewRequest) GetSource() string {
	source, found := req.Request.Params["source"]
	if found {
		return source.(string)
	}
	return ``
}

func (req *GetEncryptPinNewRequest) SetToken(token string) {
	req.Request.Params["token"] = token
}

func (req *GetEncryptPinNewRequest) GetToken() string {
	token, found := req.Request.Params["token"]
	if found {
		return token.(string)
	}
	return ""
}
