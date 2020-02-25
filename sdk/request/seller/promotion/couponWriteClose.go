package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionCouponWriteCloseRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionCouponWriteCloseRequest() (req *SellerPromotionCouponWriteCloseRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.write.close", Params: make(map[string]interface{}, 3)}
	req = &SellerPromotionCouponWriteCloseRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionCouponWriteCloseRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionCouponWriteCloseRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCloseRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionCouponWriteCloseRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCloseRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerPromotionCouponWriteCloseRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}
