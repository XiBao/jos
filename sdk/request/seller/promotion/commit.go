package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionCommitRequest struct {
	Request *sdk.Request
}

// commit new request
func NewSellerPromotionCommitRequest() (req *SellerPromotionCommitRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.commit", Params: make(map[string]interface{}, 1)}
	req = &SellerPromotionCommitRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionCommitRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionCommitRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}
