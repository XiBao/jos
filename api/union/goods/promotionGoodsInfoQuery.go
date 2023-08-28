package goods

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/goods"
)

type PromotionGoodsInfoQueryRequest struct {
	api.BaseRequest
	SkuIds []uint64 `json:"skuIds"` // skuId集合
}

type PromotionGoodsInfoQueryResponse struct {
	ErrorResp *api.ErrorResponnse                  `json:"error_response,omitempty"`
	Data      *PromotionGoodsInfoQueryResponseData `json:"jd_union_open_goods_promotiongoodsinfo_query_response,omitempty"`
}

type PromotionGoodsInfoQueryResponseData struct {
	Result string `json:"queryResult,omitempty"`
}

type PromotionQueryResult struct {
	Code    int64                `json:"code,omitempty"`
	Message string               `json:"message,omitempty"`
	Data    []PromotionGoodsResp `json:"data,omitempty"`
}

type PromotionGoodsResp struct {
	SkuId             uint64  `json:"skuId,omitempty"`             // skuId
	UnitPrice         float64 `json:"unitPrice,omitempty"`         // 商品单价即京东价
	MaterialUrl       string  `json:"materialUrl,omitempty"`       // 商品落地页
	IsFreeFreightRisk uint    `json:"isFreeFreightRisk,omitempty"` // 是否支持运费险(1:是,0:否)
	IsFreeShipping    uint    `json:"isFreeShipping,omitempty"`    // 是否包邮(1:是,0:否,2:自营商品遵从主站包邮规则)
	IsSeckill         uint    `json:"isSeckill,omitempty"`         // 是否秒杀(1:是,0:否)
	IsJdSale          uint    `json:"isJdSale,omitempty"`          // 是否自营(1:是,0:否)
	CommisionRatioWl  float64 `json:"commisionRatioWl,omitempty"`  // 无线佣金比例
	CommisionRatioPc  float64 `json:"commisionRatioPc,omitempty"`  // PC佣金比例
	ImgUrl            string  `json:"imgUrl,omitempty"`            // 图片地址
	Vid               uint64  `json:"vid,omitempty"`               // 商家ID
	ShopId            uint64  `json:"shopId,omitempty"`            // 店铺ID
	Cid               uint64  `json:"cid,omitempty"`               // 一级类目ID
	CidName           string  `json:"cidName,omitempty"`           // 一级类目名称
	Cid2              uint64  `json:"cid2,omitempty"`              // 二级类目ID
	Cid2Name          string  `json:"cid2Name,omitempty"`          // 二级类目名称
	Cid3              uint64  `json:"cid3,omitempty"`              // 三级类目ID
	Cid3Name          string  `json:"cid3Name,omitempty"`          // 三级类目名称
	WlUnitPrice       float64 `json:"wlUnitPrice,omitempty"`       // 商品无线京东价（单价为-1表示未查询到该商品单价）
	InOrderCount      uint    `json:"inOrderCount,omitempty"`      // 30天引单数量
	GoodsName         string  `json:"goodsName,omitempty"`         // 商品名称
	EndDate           int64   `json:"endDate,omitempty"`           // 推广结束日期(时间戳，毫秒)
	StartDate         int64   `json:"startDate,omitempty"`         // 推广开始日期（时间戳，毫秒）
}

// 大字段商品查询接口
func PromotionGoodsInfoQuery(req *PromotionGoodsInfoQueryRequest) ([]PromotionGoodsResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := goods.NewPromotionGoodsInfoQueryRequest()
	var skuIds []string
	for _, i := range req.SkuIds {
		skuIds = append(skuIds, strconv.FormatUint(i, 10))
	}
	r.SetSkuIds(strings.Join(skuIds, ","))

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	var response PromotionGoodsInfoQueryResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.Data == nil {
		return nil, errors.New("no result")
	}

	var ret PromotionQueryResult
	err = json.Unmarshal([]byte(response.Data.Result), &ret)
	if err != nil {
		return nil, err
	}

	if ret.Code != 200 {
		return nil, &api.ErrorResponnse{Code: strconv.FormatInt(ret.Code, 10), ZhDesc: ret.Message}
	}

	return ret.Data, nil
}
