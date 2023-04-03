package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type FullCouponGetPromoSkusRequest struct {
	Request *sdk.Request
}

// create new request
func NewFullCouponGetPromoSkusRequest() (req *FullCouponGetPromoSkusRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.getPromoSkus", Params: make(map[string]interface{}, 3)}
	req = &FullCouponGetPromoSkusRequest{
		Request: &request,
	}
	return
}

func (req *FullCouponGetPromoSkusRequest) SetWareId(wareId uint64) {
	req.Request.Params["wareId"] = wareId
}

func (req *FullCouponGetPromoSkusRequest) GetWareId() uint64 {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.(uint64)
	}
	return 0
}

func (req *FullCouponGetPromoSkusRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *FullCouponGetPromoSkusRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *FullCouponGetPromoSkusRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *FullCouponGetPromoSkusRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}
