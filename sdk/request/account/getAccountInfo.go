package account

import (
	"github.com/XiBao/jos/sdk"
)

type GetAccountInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetAccountInfoRequest() (req *GetAccountInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.account.getAccountInfo", Params: make(map[string]interface{}, 2)}
	req = &GetAccountInfoRequest{
		Request: &request,
	}
	return
}

func (req *GetAccountInfoRequest) SetAccountType(accountType uint8) {
	req.Request.Params["accountType"] = accountType
}

func (req *GetAccountInfoRequest) GetAccountType() uint8 {
	accountType, found := req.Request.Params["accountType"]
	if found {
		return accountType.(uint8)
	}
	return 0
}

func (req *GetAccountInfoRequest) SetAccountCode(accountCode string) {
	req.Request.Params["accountCode"] = accountCode
}

func (req *GetAccountInfoRequest) GetAccountCode() string {
	accountCode, found := req.Request.Params["accountCode"]
	if found {
		return accountCode.(string)
	}
	return ""
}
