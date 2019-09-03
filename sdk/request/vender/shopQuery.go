package vender

import (
	"github.com/XiBao/jos/sdk"
)

type VenderShopQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewVenderShopQueryRequest() (req *VenderShopQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.vender.shop.query", Params: nil}
	req = &VenderShopQueryRequest{
		Request: &request,
	}
	return
}
