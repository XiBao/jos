package jm

import "github.com/XiBao/jos/sdk"

type GetPurchaseInfoRequest struct {
	Request *sdk.Request
}

func NewGetPurchaseInfoRequest() (req *GetPurchaseInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.getPurchaseInfo", Params: make(map[string]interface{})}
	req = &GetPurchaseInfoRequest{
		Request: &request,
	}
	return
}
