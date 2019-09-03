package ware

import (
	"github.com/XiBao/jos/sdk"
)

type WarePriceGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewWarePriceGetRequest() (req *WarePriceGetRequest) {
	request := sdk.Request{MethodName: "jingdong.ware.price.get", Params: make(map[string]interface{}, 1)}
	req = &WarePriceGetRequest{
		Request: &request,
	}
	return
}

func (req *WarePriceGetRequest) SetSkuId(skuId string) {
	req.Request.Params["sku_id"] = skuId
}

func (req *WarePriceGetRequest) GetSkuId() string {
	skuId, found := req.Request.Params["sku_id"]
	if found {
		return skuId.(string)
	}
	return ""
}
