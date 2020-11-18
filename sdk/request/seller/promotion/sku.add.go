package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionSkuAddRequest struct {
	Request *sdk.Request
}

// sku add new request
func NewSellerPromotionSkuAddRequest() (req *SellerPromotionSkuAddRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.sku.add", Params: make(map[string]interface{}, 7)}
	req = &SellerPromotionSkuAddRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionSkuAddRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionSkuAddRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionSkuAddRequest) SetSkuIds(skuIds string) {
	req.Request.Params["sku_ids"] = skuIds
}

func (req *SellerPromotionSkuAddRequest) GetSkuIds() string {
	skuIds, found := req.Request.Params["sku_ids"]
	if found {
		return skuIds.(string)
	}
	return ""
}

func (req *SellerPromotionSkuAddRequest) SetJdPrices(jdPrices string) {
	req.Request.Params["jd_prices"] = jdPrices
}

func (req *SellerPromotionSkuAddRequest) GetJdPrices() string {
	jdPrices, found := req.Request.Params["jd_prices"]
	if found {
		return jdPrices.(string)
	}
	return ""
}

func (req *SellerPromotionSkuAddRequest) SetPromoPrices(promoPrices string) {
	req.Request.Params["promo_prices"] = promoPrices
}

func (req *SellerPromotionSkuAddRequest) GetPromoPrices() string {
	promoPrices, found := req.Request.Params["promo_prices"]
	if found {
		return promoPrices.(string)
	}
	return ""
}

func (req *SellerPromotionSkuAddRequest) SetSeq(seq string) {
	req.Request.Params["seq"] = seq
}

func (req *SellerPromotionSkuAddRequest) GetSeq() string {
	seq, found := req.Request.Params["seq"]
	if found {
		return seq.(string)
	}
	return ""
}

func (req *SellerPromotionSkuAddRequest) SetNum(num string) {
	req.Request.Params["num"] = num
}

func (req *SellerPromotionSkuAddRequest) GetNum() string {
	num, found := req.Request.Params["num"]
	if found {
		return num.(string)
	}
	return ""
}

func (req *SellerPromotionSkuAddRequest) SetBindType(bindType string) {
	req.Request.Params["bind_type"] = bindType
}

func (req *SellerPromotionSkuAddRequest) GetBindType() string {
	bindType, found := req.Request.Params["bind_type"]
	if found {
		return bindType.(string)
	}
	return ""
}
