package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionResumeRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionResumeRequest() (req *SellerPromotionResumeRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.resume", Params: make(map[string]interface{}, 5)}
	req = &SellerPromotionResumeRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionResumeRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionResumeRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionResumeRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionResumeRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionResumeRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *SellerPromotionResumeRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerPromotionResumeRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionResumeRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionResumeRequest) SetPromoType(promoType uint8) {
	req.Request.Params["promo_type"] = promoType
}

func (req *SellerPromotionResumeRequest) GetPromoType() uint8 {
	promoType, found := req.Request.Params["promo_type"]
	if found {
		return promoType.(uint8)
	}
	return 0
}
