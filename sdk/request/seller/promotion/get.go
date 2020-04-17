package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionGetRequest() (req *SellerPromotionGetRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.get", Params: make(map[string]interface{}, 4)}
	req = &SellerPromotionGetRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionGetRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionGetRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionGetRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionGetRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionGetRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionGetRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionGetRequest) SetPromoType(promoType uint8) {
	req.Request.Params["promo_type"] = promoType
}

func (req *SellerPromotionGetRequest) GetPromoType() uint8 {
	promoType, found := req.Request.Params["promo_type"]
	if found {
		return promoType.(uint8)
	}
	return 0
}
