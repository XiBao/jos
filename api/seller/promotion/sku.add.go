package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
)

type SkuAddRequest struct {
	api.BaseRequest
	PromoId     uint64 `json:"promo_id" codec:"promo_id"`         // 促销编号
	SkuIds      string `json:"sku_ids" codec:"sku_ids"`           // SKU编号
	JdPrices    string `json:"jd_prices" codec:"jd_prices"`       // 京东价，以元为单位，最高可精确到小数点后两位
	PromoPrices string `json:"promo_prices" codec:"promo_prices"` // 促销价，以元为单位，最高可精确到小数点后两位，且必须小于京东价（注：1.套装促销必填；2.fbp商家设置赠品促销时，赠品SKU的该字段必填，将作为订单拆分价使用；3.单品促销选填，不填时系统将默认设置为京东价，此时单品促销必须添加促销道具；4.其它促销不需要设置。）
	Seq         string `json:"seq" codec:"seq"`                   // 套装商品展示次序，相同商品的SKU上次序必须一致。（只对套装促销、总价满赠类促销有效）
	Num         string `json:"num" codec:"num"`                   // 赠品赠送数量，只能送1-3个。(只对赠品促销有效)
	BindType    string `json:"bind_type" codec:"bind_type"`       // 绑定类型, 可选值：主商品（1），赠品（2), bound == 3 部分商品不参与(3)。(赠品促销、满减送促销中的赠品需要设置为2，其余均设置为1)

}
type SkuAddResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SkuAddResponseData `json:"jingdong_seller_promotion_sku_add_responce,omitempty" codec:"jingdong_seller_promotion_sku_add_responce,omitempty"`
}

type SkuAddResponseData struct {
	Code      string   `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string   `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Ids       []uint64 `json:"ids,omitempty" codec:"ids,omitempty"`
}

// 添加参加促销的sku，单次最多添加100个SKU，一个促销最多支持1000个SKU，当基于套装促销添加SKU时，最多可设置3个商品的SKU，并且相同商品的次序要一致；当基于赠品促销添加SKU时，赠品SKU只能是1-5个，每个赠品只能赠送1-3个，赠品的总价应低于主商品中的最小京东价。
func SkuAdd(req *SkuAddRequest) ([]uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionSkuAddRequest()
	r.SetPromoId(req.PromoId)
	r.SetSkuIds(req.SkuIds)
	r.SetJdPrices(req.JdPrices)
	if req.PromoPrices != "" {
		r.SetPromoPrices(req.PromoPrices)
	}
	if req.Seq != "" {
		r.SetSeq(req.Seq)
	}
	if req.Num != "" {
		r.SetNum(req.Num)
	}
	if req.BindType != "" {
		r.SetBindType(req.BindType)
	}
	result, err := client.PostExecute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response SkuAddResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Ids, nil
}
