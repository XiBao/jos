package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponReadGetCouponCountRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponReadGetCouponCountRequest() (req *SellerCouponReadGetCouponCountRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.read.getCouponCount", Params: make(map[string]interface{}, 13)}
	req = &SellerCouponReadGetCouponCountRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponReadGetCouponCountRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerCouponReadGetCouponCountRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponCountRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponReadGetCouponCountRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponCountRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerCouponReadGetCouponCountRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}

func (req *SellerCouponReadGetCouponCountRequest) SetType(cType string) {
	req.Request.Params["type"] = cType
}

func (req *SellerCouponReadGetCouponCountRequest) GetType() string {
	cType, found := req.Request.Params["type"]
	if found {
		return cType.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponCountRequest) SetGrantType(grantType uint8) {
	req.Request.Params["grantType"] = grantType
}

func (req *SellerCouponReadGetCouponCountRequest) GetGrantType() uint8 {
	grantType, found := req.Request.Params["grantType"]
	if found {
		return grantType.(uint8)
	}
	return 0
}

func (req *SellerCouponReadGetCouponCountRequest) SetGrantWay(grantWay uint8) {
	req.Request.Params["grantWay"] = grantWay
}

func (req *SellerCouponReadGetCouponCountRequest) GetGrantWay() uint8 {
	grantWay, found := req.Request.Params["grantWay"]
	if found {
		return grantWay.(uint8)
	}
	return 0
}

func (req *SellerCouponReadGetCouponCountRequest) SetBindType(bindType uint8) {
	req.Request.Params["bindType"] = bindType
}

func (req *SellerCouponReadGetCouponCountRequest) GetBindType() uint8 {
	bindType, found := req.Request.Params["bindType"]
	if found {
		return bindType.(uint8)
	}
	return 0
}

func (req *SellerCouponReadGetCouponCountRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *SellerCouponReadGetCouponCountRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponCountRequest) SetCreateMonth(createMonth string) {
	req.Request.Params["createMonth"] = createMonth
}

func (req *SellerCouponReadGetCouponCountRequest) GetCreateMonth() string {
	createMonth, found := req.Request.Params["createMonth"]
	if found {
		return createMonth.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponCountRequest) SetCreatorType(creatorType string) {
	req.Request.Params["creatorType"] = creatorType
}

func (req *SellerCouponReadGetCouponCountRequest) GetCreatorType() string {
	creatorType, found := req.Request.Params["creatorType"]
	if found {
		return creatorType.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponCountRequest) SetClosed(closed string) {
	req.Request.Params["closed"] = closed
}

func (req *SellerCouponReadGetCouponCountRequest) GetClosed() string {
	closed, found := req.Request.Params["closed"]
	if found {
		return closed.(string)
	}
	return ""
}
