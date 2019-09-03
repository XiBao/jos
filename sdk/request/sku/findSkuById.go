package sku

import (
	"github.com/XiBao/jos/sdk"
)

type FindSkuByIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewFindSkuByIdRequest() (req *FindSkuByIdRequest) {
	request := sdk.Request{MethodName: "jingdong.sku.read.findSkuById", Params: make(map[string]interface{}, 2)}
	req = &FindSkuByIdRequest{
		Request: &request,
	}
	return
}

func (req *FindSkuByIdRequest) SetFields(fields string) {
	req.Request.Params["field"] = fields
}

func (req *FindSkuByIdRequest) GetFields() string {
	fields, found := req.Request.Params["field"]
	if found {
		return fields.(string)
	}
	return ""
}

func (req *FindSkuByIdRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuId"] = skuId
}

func (req *FindSkuByIdRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(uint64)
	}
	return 0
}
