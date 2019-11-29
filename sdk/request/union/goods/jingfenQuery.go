package goods

import (
	"github.com/XiBao/jos/sdk"
)

type GoodsReq struct {
	EliteId   uint   `json:"eliteId"`             // 频道id：1-好券商品,2-超级大卖场,10-9.9专区,22-热销爆品,24-数码家电,25-超市,26-母婴玩具,27-家具日用,28-美妆穿搭,29-医药保健,30-图书文具,31-今日必推,32-王牌好货,33-秒杀商品,34-拼购商品
	PageIndex uint   `json:"pageIndex,omitempty"` // 页码，默认1
	PageSize  uint   `json:"pageSize,omitempty"`  // 每页数量，默认20，上限50
	SortName  string `json:"sortName,omitempty"`  // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	Sort      string `json:"sort,omitempty"`      // asc,desc升降序,默认降序
}

type JingfenQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewJingfenQueryRequest() (req *JingfenQueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.goods.jingfen.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &JingfenQueryRequest{
		Request: &request,
	}
	return
}

func (req *JingfenQueryRequest) SetGoodsReq(goodsReq *GoodsReq) {
	req.Request.Params["goodsReq"] = goodsReq
}

func (req *JingfenQueryRequest) GetGoodsReq() *GoodsReq {
	goodsReq, found := req.Request.Params["goodsReq"]
	if found {
		return goodsReq.(*GoodsReq)
	}
	return nil
}
