package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponWriteLockRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponWriteLockRequest() (req *SellerCouponWriteLockRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.write.lockCoupon", Params: make(map[string]interface{}, 6)}
	req = &SellerCouponWriteLockRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponWriteLockRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponWriteLockRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponWriteLockRequest) SetRequestId(requestId string) {
	req.Request.Params["requestId"] = requestId
}

func (req *SellerCouponWriteLockRequest) GetRequestId() string {
	requestId, found := req.Request.Params["requestId"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerCouponWriteLockRequest) SetTime(cTime uint64) {
	req.Request.Params["time"] = cTime
}

func (req *SellerCouponWriteLockRequest) GetTime() uint64 {
	cTime, found := req.Request.Params["time"]
	if found {
		return cTime.(uint64)
	}
	return 0
}

func (req *SellerCouponWriteLockRequest) SetPurpose(purpose string) {
	req.Request.Params["purpose"] = purpose
}

func (req *SellerCouponWriteLockRequest) GetPurpose() string {
	purpose, found := req.Request.Params["purpose"]
	if found {
		return purpose.(string)
	}
	return ""
}

func (req *SellerCouponWriteLockRequest) SetOperateTime(operateTime string) {
	req.Request.Params["operateTime"] = operateTime
}

func (req *SellerCouponWriteLockRequest) GetOperateTime() string {
	operateTime, found := req.Request.Params["operateTime"]
	if found {
		return operateTime.(string)
	}
	return ""
}

func (req *SellerCouponWriteLockRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerCouponWriteLockRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}
