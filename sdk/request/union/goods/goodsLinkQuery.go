package goods

import (
	"github.com/XiBao/jos/sdk"
)

type LinkGoodsReq struct {
	Url        string `json:"url"`        // 链接
	SubUnionId string `json:"subUnionId"` // 子联盟ID（需要联系运营开通权限才能拿到数据）
}

type GoodsLinkQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewGoodsLinkQueryRequest() (req *GoodsLinkQueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.goods.link.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &GoodsLinkQueryRequest{
		Request: &request,
	}
	return
}

func (req *GoodsLinkQueryRequest) SetGoodsReq(goodsReq *LinkGoodsReq) {
	req.Request.Params["goodsReq"] = goodsReq
}

func (req *GoodsLinkQueryRequest) GetGoodsReq() *LinkGoodsReq {
	goodsReq, found := req.Request.Params["goodsReq"]
	if found {
		return goodsReq.(*LinkGoodsReq)
	}
	return nil
}
