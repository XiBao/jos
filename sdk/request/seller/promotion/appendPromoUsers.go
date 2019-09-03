package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionAppendPromoUsersRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionAppendPromoUsersRequest() (req *SellerPromotionAppendPromoUsersRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.appendPromoUsers", Params: make(map[string]interface{}, 7)}
	req = &SellerPromotionAppendPromoUsersRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionAppendPromoUsersRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionAppendPromoUsersRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionAppendPromoUsersRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionAppendPromoUsersRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionAppendPromoUsersRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *SellerPromotionAppendPromoUsersRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerPromotionAppendPromoUsersRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promoId"] = promoId
}

func (req *SellerPromotionAppendPromoUsersRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promoId"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionAppendPromoUsersRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *SellerPromotionAppendPromoUsersRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SellerPromotionAppendPromoUsersRequest) SetBeginTime(beginTime string) {
	req.Request.Params["begin_time"] = beginTime
}

func (req *SellerPromotionAppendPromoUsersRequest) GetBeginTime() string {
	beginTime, found := req.Request.Params["begin_time"]
	if found {
		return beginTime.(string)
	}
	return ""
}

func (req *SellerPromotionAppendPromoUsersRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *SellerPromotionAppendPromoUsersRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}
