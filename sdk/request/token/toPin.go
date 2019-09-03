package token

import (
	"github.com/XiBao/jos/sdk"
)

type TokenToPinRequest struct {
	Request *sdk.Request
}

// create new request
func NewTokenToPinRequest() (req *TokenToPinRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.token.source.to.pin", Params: make(map[string]interface{}, 2)}
	req = &TokenToPinRequest{
		Request: &request,
	}
	return
}

func (req *TokenToPinRequest) SetToken(token string) {
	req.Request.Params["token"] = token
}

func (req *TokenToPinRequest) GetToken() string {
	token, found := req.Request.Params["token"]
	if found {
		return token.(string)
	}
	return ""
}

func (req *TokenToPinRequest) SetSource(source string) {
	req.Request.Params["source"] = source
}

func (req *TokenToPinRequest) GetSource() string {
	source, found := req.Request.Params["source"]
	if found {
		return source.(string)
	}
	return ""
}
