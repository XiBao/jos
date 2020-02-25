package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponReadGetCouponListRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponReadGetCouponListRequest() (req *SellerCouponReadGetCouponListRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.read.getCouponList", Params: make(map[string]interface{}, 13)}
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

func (req *SellerCouponReadGetCouponListRequest) SetType(cType uint) {
	req.Request.Params["type"] = cType
}

func (req *SellerCouponReadGetCouponListRequest) GetType() uint {
	cType, found := req.Request.Params["type"]
	if found {
		return cType.(uint)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetGrantType(grantType uint) {
	req.Request.Params["grantType"] = grantType
}

func (req *SellerCouponReadGetCouponListRequest) GetGrantType() uint {
	grantType, found := req.Request.Params["grantType"]
	if found {
		return grantType.(uint)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetGrantWay(grantWay uint) {
	req.Request.Params["grantWay"] = grantWay
}

func (req *SellerCouponReadGetCouponListRequest) GetGrantWay() uint {
	grantWay, found := req.Request.Params["grantWay"]
	if found {
		return grantWay.(uint)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetBindType(bindType uint) {
	req.Request.Params["bindType"] = bindType
}

func (req *SellerCouponReadGetCouponListRequest) GetBindType() uint {
	bindType, found := req.Request.Params["bindType"]
	if found {
		return bindType.(uint)
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

func (req *SellerCouponReadGetCouponListRequest) SetCreatorType(creatorType uint) {
	req.Request.Params["creatorType"] = creatorType
}

func (req *SellerCouponReadGetCouponListRequest) GetCreatorType() uint {
	creatorType, found := req.Request.Params["creatorType"]
	if found {
		return creatorType.(uint)
	}
	return 0
}

func (req *SellerCouponReadGetCouponListRequest) SetClosed(closed uint) {
	req.Request.Params["closed"] = closed
}

func (req *SellerCouponReadGetCouponListRequest) GetClosed() uint {
	closed, found := req.Request.Params["closed"]
	if found {
		return closed.(uint)
	}
	return 0
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

func (req *SellerCouponReadGetCouponListRequest) SetPageSize(pageSize uint) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *SellerCouponReadGetCouponListRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
