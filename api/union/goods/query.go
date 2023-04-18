package goods

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/goods"
)

type QueryRequest struct {
	api.BaseRequest
	Cid1                 uint64   `json:"cid1,omitempty"`                 // 一级类目id
	Cid2                 uint64   `json:"cid2,omitempty"`                 // 二级类目id
	Cid3                 uint64   `json:"cid3,omitempty"`                 // 三级类目id
	SkuIds               []uint64 `json:"skuIds,omitempty"`               // skuid集合(一次最多支持查询100个sku)，数组类型开发时记得加[]
	Keyword              string   `json:"keyword,omitempty"`              // 关键词，字数同京东商品名称一致，目前未限制
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
	SortName             string   `json:"sortName,omitempty"`             // 排序字段(price：单价, commissionShare：佣金比例, commission：佣金， inOrderCount30Days：30天引单量， inOrderComm30Days：30天支出佣金)
	Sort                 string   `json:"sort,omitempty"`                 // asc,desc升降序,默认降序
}

type QueryResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty"`
	Data      *QueryResponseData  `json:"jd_union_open_goods_query_responce,omitempty"`
}

type QueryResponseData struct {
	Result string `json:"queryResult,omitempty"`
}

// 查询商品及优惠券信息，返回的结果可调用转链接口生成单品或二合一推广链接。支持按SKUID、关键词、优惠券基本属性、是否拼购、是否爆款等条件查询，建议不要同时传入SKUID和其他字段，以获得较多的结果。支持按价格、佣金比例、佣金、引单量等维度排序。用优惠券链接调用转链接口时，需传入搜索接口link字段返回的原始优惠券链接，切勿对链接进行任何encode、decode操作，否则将导致转链二合一推广链接时校验失败。
func Query(req *QueryRequest) (*JingfenQueryResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewQueryRequest()
	goodsReq := &goods.QueryReq{
		Cid1:                 req.Cid1,
		Cid2:                 req.Cid2,
		Cid3:                 req.Cid3,
		SkuIds:               req.SkuIds,
		Keyword:              req.Keyword,
		PriceFrom:            req.PriceFrom,
		PriceTo:              req.PriceTo,
		CommissionShareStart: req.CommissionShareStart,
		CommissionShareEnd:   req.CommissionShareEnd,
		Owner:                req.Owner,
		IsCoupon:             req.IsCoupon,
		IsPG:                 req.IsPG,
		IsHot:                req.IsHot,
		PingouPriceStart:     req.PingouPriceStart,
		PingouPriceEnd:       req.PingouPriceEnd,
		BrandCode:            req.BrandCode,
		ShopId:               req.ShopId,
		HasContent:           req.HasContent,
		HasBestCoupon:        req.HasBestCoupon,
		Pid:                  req.Pid,
		Fields:               req.Fields,
		PageIndex:            req.PageIndex,
		PageSize:             req.PageSize,
		SortName:             req.SortName,
		Sort:                 req.Sort,
	}
	r.SetGoodsReqDTO(goodsReq)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	var response QueryResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.Data == nil || response.Data.Result == "" {
		return nil, errors.New("no result")
	}
	var resp JingfenQueryResult
	err = json.Unmarshal([]byte(response.Data.Result), &resp)
	if err != nil {
		return nil, err
	}
	if resp.Code != 200 {
		return nil, &api.ErrorResponnse{Code: strconv.FormatInt(resp.Code, 10), ZhDesc: resp.Message}
	}

	return &resp, nil
}
