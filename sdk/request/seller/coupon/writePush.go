package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponWritePushRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponWritePushRequest() (req *SellerCouponWritePushRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.write.pushCoupon", Params: make(map[string]interface{}, 6)}
	req = &SellerCouponWritePushRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponWritePushRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponWritePushRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponWritePushRequest) SetRequestId(requestId string) {
	req.Request.Params["requestId"] = requestId
}

func (req *SellerCouponWritePushRequest) GetRequestId() string {
	requestId, found := req.Request.Params["requestId"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerCouponWritePushRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *SellerCouponWritePushRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SellerCouponWritePushRequest) SetDistrTime(distrTime string) {
	req.Request.Params["distrTime"] = distrTime
}

func (req *SellerCouponWritePushRequest) GetDistrTime() string {
	distrTime, found := req.Request.Params["distrTime"]
	if found {
		return distrTime.(string)
	}
	return ""
}

func (req *SellerCouponWritePushRequest) SetUuid(uuid string) {
	req.Request.Params["uuid"] = uuid
}

func (req *SellerCouponWritePushRequest) GetUuid() string {
	uuid, found := req.Request.Params["uuid"]
	if found {
		return uuid.(string)
	}
	return ""
}

func (req *SellerCouponWritePushRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerCouponWritePushRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}
