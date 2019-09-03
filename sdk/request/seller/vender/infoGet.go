package vender

import (
	"github.com/XiBao/jos/sdk"
)

type SellerVenderInfoGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerVenderInfoGetRequest() (req *SellerVenderInfoGetRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.vender.info.get", Params: nil}
	req = &SellerVenderInfoGetRequest{
		Request: &request,
	}
	return
}
