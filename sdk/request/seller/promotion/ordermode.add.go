package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionOrdermodeAddRequest struct {
	Request *sdk.Request
}

// ordermode add new request
func NewSellerPromotionOrdermodeAddRequest() (req *SellerPromotionOrdermodeAddRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.ordermode.add", Params: make(map[string]interface{}, 7)}
	req = &SellerPromotionOrdermodeAddRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionOrdermodeAddRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionOrdermodeAddRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionOrdermodeAddRequest) SetFavorMode(favorMode uint8) {
	req.Request.Params["favor_mode"] = favorMode
}

func (req *SellerPromotionOrdermodeAddRequest) GetFavorMode() uint8 {
	favorMode, found := req.Request.Params["favor_mode"]
	if found {
		return favorMode.(uint8)
	}
	return 0
}

func (req *SellerPromotionOrdermodeAddRequest) SetQuota(quota string) {
	req.Request.Params["quota"] = quota
}

func (req *SellerPromotionOrdermodeAddRequest) GetQuota() string {
	quota, found := req.Request.Params["quota"]
	if found {
		return quota.(string)
	}
	return ""
}

func (req *SellerPromotionOrdermodeAddRequest) SetRate(rate string) {
	req.Request.Params["rate"] = rate
}

func (req *SellerPromotionOrdermodeAddRequest) GetRate() string {
	rate, found := req.Request.Params["rate"]
	if found {
		return rate.(string)
	}
	return ""
}

func (req *SellerPromotionOrdermodeAddRequest) SetPlus(plus string) {
	req.Request.Params["plus"] = plus
}

func (req *SellerPromotionOrdermodeAddRequest) GetPlus() string {
	plus, found := req.Request.Params["plus"]
	if found {
		return plus.(string)
	}
	return ""
}

func (req *SellerPromotionOrdermodeAddRequest) SetMinus(minus string) {
	req.Request.Params["minus"] = minus
}

func (req *SellerPromotionOrdermodeAddRequest) GetMinus() string {
	minus, found := req.Request.Params["minus"]
	if found {
		return minus.(string)
	}
	return ""
}

func (req *SellerPromotionOrdermodeAddRequest) SetLink(link string) {
	req.Request.Params["link"] = link
}

func (req *SellerPromotionOrdermodeAddRequest) GetLink() string {
	link, found := req.Request.Params["link"]
	if found {
		return link.(string)
	}
	return ""
}

func (req *SellerPromotionOrdermodeAddRequest) SetFreePostage(freePostage string) {
	req.Request.Params["free_postage"] = freePostage
}

func (req *SellerPromotionOrdermodeAddRequest) GetFreePostage() string {
	freePostage, found := req.Request.Params["free_postage"]
	if found {
		return freePostage.(string)
	}
	return ""
}
