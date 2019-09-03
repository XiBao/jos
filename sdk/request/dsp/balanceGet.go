package dsp

import (
	"github.com/XiBao/jos/sdk"
)

type BalanceGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewBalanceGetRequest() (req *BalanceGetRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.balance.get", Params: nil}
	req = &BalanceGetRequest{
		Request: &request,
	}
	return
}
