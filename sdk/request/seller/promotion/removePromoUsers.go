package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionRemovePromoUsersRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionRemovePromoUsersRequest() (req *SellerPromotionRemovePromoUsersRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.removePromoUsers", Params: make(map[string]interface{}, 7)}
	req = &SellerPromotionRemovePromoUsersRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionRemovePromoUsersRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionRemovePromoUsersRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionRemovePromoUsersRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionRemovePromoUsersRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionRemovePromoUsersRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *SellerPromotionRemovePromoUsersRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerPromotionRemovePromoUsersRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *SellerPromotionRemovePromoUsersRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionRemovePromoUsersRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *SellerPromotionRemovePromoUsersRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SellerPromotionRemovePromoUsersRequest) SetBeginTime(beginTime string) {
	req.Request.Params["begin_time"] = beginTime
}

func (req *SellerPromotionRemovePromoUsersRequest) GetBeginTime() string {
	beginTime, found := req.Request.Params["begin_time"]
	if found {
		return beginTime.(string)
	}
	return ""
}

func (req *SellerPromotionRemovePromoUsersRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *SellerPromotionRemovePromoUsersRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}
