package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type FullCouponGetLastDataRequest struct {
	Request *sdk.Request
}

// create new request
func NewFullCouponGetLastDataRequest() (req *FullCouponGetLastDataRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.getLastData", Params: make(map[string]interface{}, 4)}
	req = &FullCouponGetLastDataRequest{
		Request: &request,
	}
	return
}

func (req *FullCouponGetLastDataRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *FullCouponGetLastDataRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *FullCouponGetLastDataRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *FullCouponGetLastDataRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *FullCouponGetLastDataRequest) SetDate(date string) {
	req.Request.Params["date"] = date
}

func (req *FullCouponGetLastDataRequest) GetDate() string {
	date, found := req.Request.Params["date"]
	if found {
		return date.(string)
	}
	return ""
}

func (req *FullCouponGetLastDataRequest) SetShopId(shopId uint64) {
	req.Request.Params["shopId"] = shopId
}

func (req *FullCouponGetLastDataRequest) GetShopId() uint64 {
	shopId, found := req.Request.Params["shopId"]
	if found {
		return shopId.(uint64)
	}
	return 0
}
