package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponReadGetCouponByIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponReadGetCouponByIdRequest() (req *SellerCouponReadGetCouponByIdRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.read.getCouponById", Params: make(map[string]interface{}, 3)}
	req = &SellerCouponReadGetCouponByIdRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponReadGetCouponByIdRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerCouponReadGetCouponByIdRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponByIdRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponReadGetCouponByIdRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponByIdRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerCouponReadGetCouponByIdRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}
