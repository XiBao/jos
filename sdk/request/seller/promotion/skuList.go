package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionSkuListRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionSkuListRequest() (req *SellerPromotionSkuListRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.sku.list", Params: make(map[string]interface{}, 9)}
	req = &SellerPromotionSkuListRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionSkuListRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionSkuListRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionSkuListRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionSkuListRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionSkuListRequest) SetPromoId(promoId uint64) {
	req.Request.Params["promo_id"] = promoId
}

func (req *SellerPromotionSkuListRequest) GetPromoId() uint64 {
	promoId, found := req.Request.Params["promo_id"]
	if found {
		return promoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionSkuListRequest) SetWareId(wareId uint64) {
	req.Request.Params["ware_id"] = wareId
}

func (req *SellerPromotionSkuListRequest) GetWareId() uint64 {
	wareId, found := req.Request.Params["ware_id"]
	if found {
		return wareId.(uint64)
	}
	return 0
}

func (req *SellerPromotionSkuListRequest) SetSkuId(skuId uint64) {
	req.Request.Params["sku_id"] = skuId
}

func (req *SellerPromotionSkuListRequest) GetSkuId() uint64 {
	skuId, found := req.Request.Params["sku_id"]
	if found {
		return skuId.(uint64)
	}
	return 0
}

func (req *SellerPromotionSkuListRequest) SetBindType(bindType uint8) {
	req.Request.Params["bind_type"] = bindType
}

func (req *SellerPromotionSkuListRequest) GetBindType() uint8 {
	bindType, found := req.Request.Params["bind_type"]
	if found {
		return bindType.(uint8)
	}
	return 0
}

func (req *SellerPromotionSkuListRequest) SetPromoType(promoType int) {
	req.Request.Params["promo_type"] = promoType
}

func (req *SellerPromotionSkuListRequest) GetPromoType() int {
	promoType, found := req.Request.Params["promo_type"]
	if found {
		return promoType.(int)
	}
	return 0
}

func (req *SellerPromotionSkuListRequest) SetPage(page int) {
	req.Request.Params["page"] = page
}

func (req *SellerPromotionSkuListRequest) GetPage() int {
	page, found := req.Request.Params["page"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *SellerPromotionSkuListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageS_size"] = pageSize
}

func (req *SellerPromotionSkuListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageS_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}
