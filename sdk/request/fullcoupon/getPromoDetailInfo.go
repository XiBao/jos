package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type FullCouponGetPromoDetailInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewFullCouponGetPromoDetailInfoRequest() (req *FullCouponGetPromoDetailInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.getPromoDetailInfo", Params: make(map[string]interface{}, 2)}
	req = &FullCouponGetPromoDetailInfoRequest{
		Request: &request,
	}
	return
}

func (req *FullCouponGetPromoDetailInfoRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *FullCouponGetPromoDetailInfoRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *FullCouponGetPromoDetailInfoRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *FullCouponGetPromoDetailInfoRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}
