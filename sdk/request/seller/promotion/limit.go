package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionLimitRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionLimitRequest() (req *SellerPromotionLimitRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.getPromoLimit", Params: make(map[string]interface{}, 5)}
	req = &SellerPromotionLimitRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionLimitRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionLimitRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionLimitRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionLimitRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionLimitRequest) SetCategoryId(categoryId uint64) {
	req.Request.Params["category_id"] = categoryId
}

func (req *SellerPromotionLimitRequest) GetCategoryId() uint64 {
	categoryId, found := req.Request.Params["category_id"]
	if found {
		return categoryId.(uint64)
	}
	return 0
}

func (req *SellerPromotionLimitRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *SellerPromotionLimitRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *SellerPromotionLimitRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *SellerPromotionLimitRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}
