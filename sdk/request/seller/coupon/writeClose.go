package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponWriteCloseRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponWriteCloseRequest() (req *SellerCouponWriteCloseRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.write.close", Params: make(map[string]interface{}, 3)}
	req = &SellerCouponWriteCloseRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponWriteCloseRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerCouponWriteCloseRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerCouponWriteCloseRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponWriteCloseRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponWriteCloseRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerCouponWriteCloseRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}
