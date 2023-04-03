package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type FullCouponGetPromoWaresRequest struct {
	Request *sdk.Request
}

// create new request
func NewFullCouponGetPromoWaresRequest() (req *FullCouponGetPromoWaresRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.getPromoWares", Params: make(map[string]interface{}, 3)}
	req = &FullCouponGetPromoWaresRequest{
		Request: &request,
	}
	return
}

func (req *FullCouponGetPromoWaresRequest) SetPageIndex(pageIndex uint) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *FullCouponGetPromoWaresRequest) GetPageIndex() uint {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(uint)
	}
	return 0
}

func (req *FullCouponGetPromoWaresRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *FullCouponGetPromoWaresRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *FullCouponGetPromoWaresRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *FullCouponGetPromoWaresRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}
