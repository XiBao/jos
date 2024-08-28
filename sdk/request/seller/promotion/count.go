package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionCountRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionCountRequest() (req *SellerPromotionCountRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.count", Params: make(map[string]interface{}, 12)}
	req = &SellerPromotionCountRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionCountRequest) SetIp(Ip string) {
	req.Request.Params["ip"] = Ip
}

func (req *SellerPromotionCountRequest) GetIp() string {
	Ip, found := req.Request.Params["ip"]
	if found {
		return Ip.(string)
	}
	return ""
}

func (req *SellerPromotionCountRequest) SetPort(Port string) {
	req.Request.Params["port"] = Port
}

func (req *SellerPromotionCountRequest) GetPort() string {
	Port, found := req.Request.Params["port"]
	if found {
		return Port.(string)
	}
	return ""
}

func (req *SellerPromotionCountRequest) SetPromoId(PromoId uint64) {
	req.Request.Params["promo_id"] = PromoId
}

func (req *SellerPromotionCountRequest) GetPromoId() uint64 {
	PromoId, found := req.Request.Params["promo_id"]
	if found {
		return PromoId.(uint64)
	}
	return 0
}

func (req *SellerPromotionCountRequest) SetName(Name string) {
	req.Request.Params["name"] = Name
}

func (req *SellerPromotionCountRequest) GetName() string {
	Name, found := req.Request.Params["name"]
	if found {
		return Name.(string)
	}
	return ""
}

func (req *SellerPromotionCountRequest) SetType(Type uint8) {
	req.Request.Params["type"] = Type
}

func (req *SellerPromotionCountRequest) GetType() uint8 {
	Type, found := req.Request.Params["type"]
	if found {
		return Type.(uint8)
	}
	return 0
}

func (req *SellerPromotionCountRequest) SetFavorMode(FavorMode uint8) {
	req.Request.Params["favor_mode"] = FavorMode
}

func (req *SellerPromotionCountRequest) GetFavorMode() uint8 {
	FavorMode, found := req.Request.Params["favor_mode"]
	if found {
		return FavorMode.(uint8)
	}
	return 0
}

func (req *SellerPromotionCountRequest) SetBeginTime(BeginTime string) {
	req.Request.Params["begin_time"] = BeginTime
}

func (req *SellerPromotionCountRequest) GetBeginTime() string {
	BeginTime, found := req.Request.Params["begin_time"]
	if found {
		return BeginTime.(string)
	}
	return ""
}

func (req *SellerPromotionCountRequest) SetEndTime(EndTime string) {
	req.Request.Params["end_time"] = EndTime
}

func (req *SellerPromotionCountRequest) GetEndTime() string {
	EndTime, found := req.Request.Params["end_time"]
	if found {
		return EndTime.(string)
	}
	return ""
}

func (req *SellerPromotionCountRequest) SetPromoStatus(PromoStatus uint8) {
	req.Request.Params["promo_status"] = PromoStatus
}

func (req *SellerPromotionCountRequest) GetPromoStatus() uint8 {
	PromoStatus, found := req.Request.Params["promo_status"]
	if found {
		return PromoStatus.(uint8)
	}
	return 0
}

func (req *SellerPromotionCountRequest) SetWareId(WareId string) {
	req.Request.Params["ware_id"] = WareId
}

func (req *SellerPromotionCountRequest) GetWareId() string {
	WareId, found := req.Request.Params["ware_id"]
	if found {
		return WareId.(string)
	}
	return ""
}

func (req *SellerPromotionCountRequest) SetSkuId(SkuId string) {
	req.Request.Params["sku_id"] = SkuId
}

func (req *SellerPromotionCountRequest) GetSkuId() string {
	SkuId, found := req.Request.Params["sku_id"]
	if found {
		return SkuId.(string)
	}
	return ""
}

func (req *SellerPromotionCountRequest) SetSrcType(SrcType uint64) {
	req.Request.Params["src_type"] = SrcType
}

func (req *SellerPromotionCountRequest) GetSrcType() uint64 {
	SrcType, found := req.Request.Params["src_type"]
	if found {
		return SrcType.(uint64)
	}
	return 0
}
