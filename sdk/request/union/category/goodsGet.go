package category

import (
	"github.com/XiBao/jos/sdk"
)

type GoodsGetReq struct {
	ParentId uint64 `json:"parentId"`   // 父类目id(一级父类目为0)
	Grade    uint   `json:"subUnionId"` // 类目级别(类目级别 0，1，2 代表一、二、三级类目)
}

type GoodsGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewGoodsGetRequest() (req *GoodsGetRequest) {
	request := sdk.Request{MethodName: "jd.union.open.category.goods.get", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &GoodsGetRequest{
		Request: &request,
	}
	return
}

func (req *GoodsLinkQueryRequest) SetReq(goodsReq *GoodsGetReq) {
	req.Request.Params["req"] = goodsReq
}

func (req *GoodsLinkQueryRequest) GetGoodsReq() *GoodsGetReq {
	goodsReq, found := req.Request.Params["req"]
	if found {
		return goodsReq.(*GoodsGetReq)
	}
	return nil
}
