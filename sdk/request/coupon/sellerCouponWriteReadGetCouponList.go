package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponReadGetCouponListRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponReadGetCouponListRequest() (req *SellerCouponReadGetCouponListRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.read.getCouponList", Params: make(map[string]interface{}, 18)}
	req = &SellerCouponReadGetCouponListRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponReadGetCouponListRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerCouponReadGetCouponListRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponListRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponReadGetCouponListRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponListRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerCouponReadGetCouponListRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetType(cType string) {
	req.Request.Params["type"] = cType
}

func (req *SellerCouponReadGetCouponListRequest) GetType() string {
	cType, found := req.Request.Params["type"]
	if found {
		return cType.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponListRequest) SetGrantType(grantType uint8) {
	req.Request.Params["grantType"] = grantType
}

func (req *SellerCouponReadGetCouponListRequest) GetGrantType() uint8 {
	grantType, found := req.Request.Params["grantType"]
	if found {
		return grantType.(uint8)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetBindType(bindType uint8) {
	req.Request.Params["bindType"] = bindType
}

func (req *SellerCouponReadGetCouponListRequest) GetBindType() uint8 {
	bindType, found := req.Request.Params["bindType"]
	if found {
		return bindType.(uint8)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetGrantWay(grantWay uint8) {
	req.Request.Params["grantWay"] = grantWay
}

func (req *SellerCouponReadGetCouponListRequest) GetGrantWay() uint8 {
	grantWay, found := req.Request.Params["grantWay"]
	if found {
		return grantWay.(uint8)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *SellerCouponReadGetCouponListRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponListRequest) SetCreateMonth(createMonth string) {
	req.Request.Params["createMonth"] = createMonth
}

func (req *SellerCouponReadGetCouponListRequest) GetCreateMonth() string {
	createMonth, found := req.Request.Params["createMonth"]
	if found {
		return createMonth.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponListRequest) SetCreatorType(creatorType string) {
	req.Request.Params["creatorType"] = creatorType
}

func (req *SellerCouponReadGetCouponListRequest) GetCreatorType() string {
	creatorType, found := req.Request.Params["creatorType"]
	if found {
		return creatorType.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponListRequest) SetClosed(closed string) {
	req.Request.Params["closed"] = closed
}

func (req *SellerCouponReadGetCouponListRequest) GetClosed() string {
	closed, found := req.Request.Params["closed"]
	if found {
		return closed.(string)
	}
	return ""
}

func (req *SellerCouponReadGetCouponListRequest) SetPage(page uint) {
	req.Request.Params["page"] = page
}

func (req *SellerCouponReadGetCouponListRequest) GetPage() uint {
	page, found := req.Request.Params["page"]
	if found {
		return page.(uint)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetPageSize(pageSize uint8) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *SellerCouponReadGetCouponListRequest) GetPageSize() uint8 {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint8)
	}
	return 0
}
