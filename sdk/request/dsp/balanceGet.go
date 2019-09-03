package dsp

import (
	"github.com/XiBao/jos/sdk"
)

type BalanceGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewBalanceGetRequest() (req *DspBalanceRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.balance.get", Params: nil}
	req = &DspBalanceRequest{
		Request: &request,
	}
	return
}
