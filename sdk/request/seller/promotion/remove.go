package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionRemoveRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionRemoveRequest() (req *SellerPromotionRemoveRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.remove", Params: make(map[string]interface{}, 5)}
	req = &SellerPromotionRemoveRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionRemoveRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionRemoveRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionRemoveRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionRemoveRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionRemoveRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *SellerPromotionRemoveRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerPromotionRemoveRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionRemoveRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionRemoveRequest) SetPromoType(promoType uint8) {
	req.Request.Params["promo_type"] = promoType
}

func (req *SellerPromotionRemoveRequest) GetPromoType() uint8 {
	promoType, found := req.Request.Params["promo_type"]
	if found {
		return promoType.(uint8)
	}
	return 0
}
