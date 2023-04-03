package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type FullCouponGetPromoListInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewFullCouponGetPromoListInfoRequest() (req *FullCouponGetPromoListInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.getPromoListInfo", Params: make(map[string]interface{}, 10)}
	req = &FullCouponGetPromoListInfoRequest{
		Request: &request,
	}
	return
}

func (req *FullCouponGetPromoListInfoRequest) SetWareId(wareId uint64) {
	req.Request.Params["wareId"] = wareId
}

func (req *FullCouponGetPromoListInfoRequest) GetWareId() uint64 {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.(uint64)
	}
	return 0
}

func (req *FullCouponGetPromoListInfoRequest) SetPageIndex(pageIndex int) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *FullCouponGetPromoListInfoRequest) GetPageIndex() int {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(int)
	}
	return 0
}

func (req *FullCouponGetPromoListInfoRequest) SetEvtStatus(evtStatus int) {
	req.Request.Params["evtStatus"] = evtStatus
}

func (req *FullCouponGetPromoListInfoRequest) GetEvtStatus() int {
	evtStatus, found := req.Request.Params["evtStatus"]
	if found {
		return evtStatus.(int)
	}
	return 0
}

func (req *FullCouponGetPromoListInfoRequest) SetEvtName(evtName string) {
	req.Request.Params["evtName"] = evtName
}

func (req *FullCouponGetPromoListInfoRequest) GetEvtName() string {
	evtName, found := req.Request.Params["evtName"]
	if found {
		return evtName.(string)
	}
	return ""
}

func (req *FullCouponGetPromoListInfoRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *FullCouponGetPromoListInfoRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *FullCouponGetPromoListInfoRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *FullCouponGetPromoListInfoRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *FullCouponGetPromoListInfoRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *FullCouponGetPromoListInfoRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *FullCouponGetPromoListInfoRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *FullCouponGetPromoListInfoRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *FullCouponGetPromoListInfoRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuId"] = skuId
}

func (req *FullCouponGetPromoListInfoRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(uint64)
	}
	return 0
}

func (req *FullCouponGetPromoListInfoRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *FullCouponGetPromoListInfoRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}
