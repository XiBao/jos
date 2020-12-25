package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionListRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionListRequest() (req *SellerPromotionListRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.list", Params: make(map[string]interface{}, 14)}
	req = &SellerPromotionListRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionListRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionListRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionListRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionListRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionListRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionListRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *SellerPromotionListRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *SellerPromotionListRequest) SetType(pType int) {
	req.Request.Params["type"] = pType
}

func (req *SellerPromotionListRequest) GetType() int {
	pType, found := req.Request.Params["type"]
	if found {
		return pType.(int)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetFavorMode(favorMode int) {
	req.Request.Params["favor_mode"] = favorMode
}

func (req *SellerPromotionListRequest) GetFavorMode() int {
	favorMode, found := req.Request.Params["favor_mode"]
	if found {
		return favorMode.(int)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetBeginTime(beginTime string) {
	req.Request.Params["begin_time"] = beginTime
}

func (req *SellerPromotionListRequest) GetBeginTime() string {
	beginTime, found := req.Request.Params["begin_time"]
	if found {
		return beginTime.(string)
	}
	return ""
}

func (req *SellerPromotionListRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *SellerPromotionListRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *SellerPromotionListRequest) SetPromoStatus(promoStatus string) {
	req.Request.Params["promo_status"] = promoStatus
}

func (req *SellerPromotionListRequest) GetPromoStatus() string {
	promoStatus, found := req.Request.Params["promo_status"]
	if found {
		return promoStatus.(string)
	}
	return ""
}

func (req *SellerPromotionListRequest) SetWareId(wareId uint64) {
	req.Request.Params["ware_id"] = wareId
}

func (req *SellerPromotionListRequest) GetWareId() uint64 {
	wareId, found := req.Request.Params["ware_id"]
	if found {
		return wareId.(uint64)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetSkuId(skuId uint64) {
	req.Request.Params["sku_id"] = skuId
}

func (req *SellerPromotionListRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["sku_id"]
	if found {
		return skuId.(uint64)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetSrcType(srcType int) {
	req.Request.Params["src_type"] = srcType
}

func (req *SellerPromotionListRequest) GetSrcType() int {
	srcType, found := req.Request.Params["src_type"]
	if found {
		return srcType.(int)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetPage(page int) {
	req.Request.Params["page"] = page
}

func (req *SellerPromotionListRequest) GetPage() int {
	page, found := req.Request.Params["page"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageS_size"] = pageSize
}

func (req *SellerPromotionListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageS_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *SellerPromotionListRequest) SetStartId(startId uint64) {
	req.Request.Params["start_id"] = startId
}

func (req *SellerPromotionListRequest) GetStartId() uint64 {
	startId, found := req.Request.Params["start_id"]
	if found {
		return startId.(uint64)
	}
	return 0
}
