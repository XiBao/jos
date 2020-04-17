package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionSuspendRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionSuspendRequest() (req *SellerPromotionSuspendRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.suspend", Params: make(map[string]interface{}, 5)}
	req = &SellerPromotionSuspendRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionSuspendRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionSuspendRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionSuspendRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionSuspendRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionSuspendRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *SellerPromotionSuspendRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerPromotionSuspendRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionSuspendRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionSuspendRequest) SetPromoType(promoType uint8) {
	req.Request.Params["promo_type"] = promoType
}

func (req *SellerPromotionSuspendRequest) GetPromoType() uint8 {
	promoType, found := req.Request.Params["promo_type"]
	if found {
		return promoType.(uint8)
	}
	return 0
}
