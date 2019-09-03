package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionUserListRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionUserListRequest() (req *SellerPromotionUserListRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.market.retrieve.promotion.getPromoUserList", Params: make(map[string]interface{}, 7)}
	req = &SellerPromotionUserListRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionUserListRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionUserListRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionUserListRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionUserListRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionUserListRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *SellerPromotionUserListRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerPromotionUserListRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionUserListRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionUserListRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *SellerPromotionUserListRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SellerPromotionUserListRequest) SetPage(page int) {
	req.Request.Params["page_index"] = page
}

func (req *SellerPromotionUserListRequest) GetPage() int {
	page, found := req.Request.Params["page_index"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *SellerPromotionUserListRequest) SetPageSize(pageSize int) {
	req.Request.Params["page_size"] = pageSize
}

func (req *SellerPromotionUserListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}
