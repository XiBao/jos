package vender

import (
	"github.com/XiBao/jos/sdk"
)

type CheckAuthForVenderIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewCheckAuthForVenderIdRequest() (req *CheckAuthForVenderIdRequest) {
	request := sdk.Request{MethodName: "jingdong.vender.auth.checkAuthForVenderId", Params: make(map[string]interface{}, 1)}
	req = &CheckAuthForVenderIdRequest{
		Request: &request,
	}
	return
}

func (req *CheckAuthForVenderIdRequest) SetPermCode(permCode string) {
	req.Request.Params["permCode"] = permCode
}

func (req *CheckAuthForVenderIdRequest) GetPermCode() string {
	permCode, found := req.Request.Params["permCode"]
	if found {
		return permCode.(string)
	}
	return ""
}
