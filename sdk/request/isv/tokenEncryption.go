package isv

import (
	"github.com/XiBao/jos/sdk"
)

type IsvTokenEncryptionRequest struct {
	Request *sdk.Request
}

// create new request
func NewIsvTokenEncryptionRequest() (req *IsvTokenEncryptionRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.isv.token.encryption", Params: make(map[string]interface{}, 1)}
	req = &IsvTokenEncryptionRequest{
		Request: &request,
	}
	return
}

func (req *IsvTokenEncryptionRequest) SetTokenStr(tokenStr string) {
	req.Request.Params["tokenStr"] = tokenStr
}

func (req *IsvTokenEncryptionRequest) GetTokenStr() string {
	tokenStr, found := req.Request.Params["tokenStr"]
	if found {
		return tokenStr.(string)
	}
	return ""
}
