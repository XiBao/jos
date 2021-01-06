package goods

import (
	"github.com/XiBao/jos/sdk"
)

type BigFieldGoodsReq struct {
	SkuIds []uint64 `json:"skuIds`
	Fields []string `json:"fields"`
}

type BigFieldQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewBigFieldQueryRequest() (req *BigFieldQueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.goods.bigfield.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &BigFieldQueryRequest{
		Request: &request,
	}
	return
}

func (req *BigFieldQueryRequest) SetGoodsReq(goodsReq *BigFieldGoodsReq) {
	req.Request.Params["goodsReq"] = goodsReq
}

func (req *BigFieldQueryRequest) GetGoodsReq() *BigFieldGoodsReq {
	goodsReq, found := req.Request.Params["goodsReq"]
	if found {
		return goodsReq.(*BigFieldGoodsReq)
	}
	return nil
}
