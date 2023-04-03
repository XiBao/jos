package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type FullCouponGetFlowPromoInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewFullCouponGetFlowPromoInfoRequest() (req *FullCouponGetFlowPromoInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.getFlowPromoInfo", Params: make(map[string]interface{}, 2)}
	req = &FullCouponGetFlowPromoInfoRequest{
		Request: &request,
	}
	return
}

func (req *FullCouponGetFlowPromoInfoRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *FullCouponGetFlowPromoInfoRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *FullCouponGetFlowPromoInfoRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *FullCouponGetFlowPromoInfoRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}
