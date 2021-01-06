package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponSkuListRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponSkuListRequest() (req *SellerCouponSkuListRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.read.getCouponSkuList", Params: make(map[string]interface{}, 9)}
	req = &SellerCouponSkuListRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponSkuListRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerCouponSkuListRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerCouponSkuListRequest) SetAppKey(appKey string) {
	req.Request.Params["appKey"] = appKey
}

func (req *SellerCouponSkuListRequest) GetAppKey() string {
	appKey, found := req.Request.Params["appKey"]
	if found {
		return appKey.(string)
	}
	return ""
}

func (req *SellerCouponSkuListRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponSkuListRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponSkuListRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *SellerCouponSkuListRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}

func (req *SellerCouponSkuListRequest) SetWareId(wareId uint64) {
	req.Request.Params["wareId"] = wareId
}

func (req *SellerCouponSkuListRequest) GetWareId() uint64 {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.(uint64)
	}
	return 0
}

func (req *SellerCouponSkuListRequest) SetSkuId(skuId uint64) {
	req.Request.Params["skuId"] = skuId
}

func (req *SellerCouponSkuListRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(uint64)
	}
	return 0
}

func (req *SellerCouponSkuListRequest) SetBindType(bindType uint8) {
	req.Request.Params["bind_type"] = bindType
}

func (req *SellerCouponSkuListRequest) GetBindType() uint8 {
	bindType, found := req.Request.Params["bind_type"]
	if found {
		return bindType.(uint8)
	}
	return 0
}

func (req *SellerCouponSkuListRequest) SetSearchStatus(searchStatus int) {
	req.Request.Params["searchStatus"] = searchStatus
}

func (req *SellerCouponSkuListRequest) GetSearchStatus() int {
	searchStatus, found := req.Request.Params["searchStatus"]
	if found {
		return searchStatus.(int)
	}
	return 0
}

func (req *SellerCouponSkuListRequest) SetPage(page int) {
	req.Request.Params["page"] = page
}

func (req *SellerCouponSkuListRequest) GetPage() int {
	page, found := req.Request.Params["page"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *SellerCouponSkuListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageS_size"] = pageSize
}

func (req *SellerCouponSkuListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageS_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}
