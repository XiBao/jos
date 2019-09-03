package ware

import (
	"github.com/XiBao/jos/sdk"
)

type MobileBigFieldRequest struct {
	Request *sdk.Request
}

// create new request
func NewMobileBigFieldRequest() (req *MobileBigFieldRequest) {
	request := sdk.Request{MethodName: "jingdong.new.ware.mobilebigfield.get", Params: make(map[string]interface{}, 1)}
	req = &MobileBigFieldRequest{
		Request: &request,
	}
	return
}

func (req *MobileBigFieldRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuid"] = skuId
}

func (req *MobileBigFieldRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuid"]
	if found {
		return skuId.(uint64)
	}
	return 0
}
