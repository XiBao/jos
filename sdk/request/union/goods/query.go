package goods

import (
	"github.com/XiBao/jos/sdk"
)

type QueryReq struct {
	Cid1                 uint64   `json:"cid1,omitempty"`                 // 一级类目id
	Cid2                 uint64   `json:"cid2,omitempty"`                 // 二级类目id
	Cid3                 uint64   `json:"cid3,omitempty"`                 // 三级类目id
	SkuIds               []uint64 `json:"skuIds,omitempty"`               // skuid集合(一次最多支持查询100个sku)，数组类型开发时记得加[]
	keyword              string   `json:"keyword,omitempty"`              // 关键词，字数同京东商品名称一致，目前未限制
	PriceFrom            float64  `json:"pricefrom,omitempty"`            // 商品价格下限
	PriceTo              float64  `json:"priceto,omitempty"`              // 商品价格上限
	CommissionShareStart uint     `json:"commissionShareStart,omitempty"` // 佣金比例区间开始
	CommissionShareEnd   uint     `json:"commissionShareEnd,omitempty"`   // 佣金比例区间结束
	Owner                string   `json:"owner,omitempty"`                // 商品类型：自营[g]，POP[p]
	IsCoupon             uint     `json:"isCoupon,omitempty"`             // 是否是优惠券商品，1：有优惠券，0：无优惠券
	IsPG                 uint     `json:"isPG,omitempty"`                 // 是否是拼购商品，1：拼购商品，0：非拼购商品
	IsHot                uint     `json:"isHot,omitempty"`                // 是否是爆款，1：爆款商品，0：非爆款商品
	PingouPriceStart     float64  `json:"pingouPriceStart,omitempty"`     // 拼购价格区间开始
	PingouPriceEnd       float64  `json:"pingouPriceEnd,omitempty"`       // 拼购价格区间结束
	BrandCode            string   `json:"brandCode,omitempty"`            // 品牌code
	ShopId               uint64   `json:"shopId,omitempty"`               // 店铺Id
	HasContent           uint     `json:"hasContent,omitempty"`           // 1：查询内容商品；其他值过滤掉此入参条件。
	HasBestCoupon        uint     `json:"hasBestCoupon,omitempty"`        // 1：查询有最优惠券商品；其他值过滤掉此入参条件。
	Pid                  string   `json:"pid,omitempty"`                  // 联盟id_应用iD_推广位id
	Fields               string   `json:"fields,omitempty"`               // 支持出参数据筛选，逗号','分隔，目前可用：videoInfo
	PageIndex            uint     `json:"pageIndex,omitempty"`            // 页码，默认1
	PageSize             uint     `json:"pageSize,omitempty"`             // 每页数量，默认20，上限50
	SortName             string   `json:"sortName,omitempty"`             // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30DaysSku：sku维度30天引单量，comments：评论数，goodComments：好评数)
	Sort                 string   `json:"sort,omitempty"`                 // asc,desc升降序,默认降序
}

type QueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewQueryRequest() (req *QueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.goods.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &QueryRequest{
		Request: &request,
	}
	return
}

func (req *QueryRequest) SetGoodsReqDTO(goodsReq *QueryReq) {
	req.Request.Params["goodsReqDTO"] = goodsReq
}

func (req *QueryRequest) GetGoodsReqDTO() *QueryReq {
	goodsReq, found := req.Request.Params["goodsReqDTO"]
	if found {
		return goodsReq.(*QueryReq)
	}
	return nil
}
