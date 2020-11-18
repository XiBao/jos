package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionCheckRequest struct {
	Request *sdk.Request
}

// check new request
func NewSellerPromotionCheckRequest() (req *SellerPromotionCheckRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.check", Params: make(map[string]interface{}, 1)}
	req = &SellerPromotionCheckRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionCheckRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionCheckRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionCheckRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *SellerPromotionCheckRequest) GetStatus() uint8 {
	status, found := req.Request.Params["status"]
	if found {
		return status.(uint8)
	}
	return 0
}
