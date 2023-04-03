package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type FullCouponGetTrendDataRequest struct {
	Request *sdk.Request
}

// create new request
func NewFullCouponGetTrendDataRequest() (req *FullCouponGetTrendDataRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.getTrendData", Params: make(map[string]interface{}, 5)}
	req = &FullCouponGetTrendDataRequest{
		Request: &request,
	}
	return
}

func (req *FullCouponGetTrendDataRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *FullCouponGetTrendDataRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *FullCouponGetTrendDataRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *FullCouponGetTrendDataRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *FullCouponGetTrendDataRequest) SetStartDate(startDate string) {
	req.Request.Params["startDate"] = startDate
}

func (req *FullCouponGetTrendDataRequest) GetStartDate() string {
	startDate, found := req.Request.Params["startDate"]
	if found {
		return startDate.(string)
	}
	return ""
}

func (req *FullCouponGetTrendDataRequest) SetEndDate(endDate string) {
	req.Request.Params["endDate"] = endDate
}

func (req *FullCouponGetTrendDataRequest) GetEndDate() string {
	endDate, found := req.Request.Params["endDate"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *FullCouponGetTrendDataRequest) SetShopId(shopId uint64) {
	req.Request.Params["shopId"] = shopId
}

func (req *FullCouponGetTrendDataRequest) GetShopId() uint64 {
	shopId, found := req.Request.Params["shopId"]
	if found {
		return shopId.(uint64)
	}
	return 0
}
