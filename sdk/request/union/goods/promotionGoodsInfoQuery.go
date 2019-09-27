package goods

import (
	"github.com/XiBao/jos/sdk"
)

type PromotionGoodsInfoQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewPromotionGoodsInfoQueryRequest() (req *PromotionGoodsInfoQueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.goods.promotiongoodsinfo.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &PromotionGoodsInfoQueryRequest{
		Request: &request,
	}
	return
}

func (req *PromotionGoodsInfoQueryRequest) SetSkuIds(skuIds string) {
	req.Request.Params["skuIds"] = skuIds
}

func (req *PromotionGoodsInfoQueryRequest) GetSkuIds() string {
	skuIds, found := req.Request.Params["skuIds"]
	if found {
		return skuIds.(string)
	}
	return ""
}
