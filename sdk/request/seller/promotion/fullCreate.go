package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionFullCreateRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionFullCreateRequest() (req *SellerPromotionFullCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.v2.unit.full.create", Params: make(map[string]interface{}, 29)}
	req = &SellerPromotionFullCreateRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionFullCreateRequest) SetIp(Ip string) {
	req.Request.Params["ip"] = Ip
}

func (req *SellerPromotionFullCreateRequest) GetIp() string {
	Ip, found := req.Request.Params["ip"]
	if found {
		return Ip.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetPort(Port string) {
	req.Request.Params["port"] = Port
}

func (req *SellerPromotionFullCreateRequest) GetPort() string {
	Port, found := req.Request.Params["port"]
	if found {
		return Port.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetRequestId(RequestId string) {
	req.Request.Params["request_id"] = RequestId
}

func (req *SellerPromotionFullCreateRequest) GetRequestId() string {
	RequestId, found := req.Request.Params["request_id"]
	if found {
		return RequestId.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetPromoName(PromoName string) {
	req.Request.Params["promo_name"] = PromoName
}

func (req *SellerPromotionFullCreateRequest) GetPromoName() string {
	PromoName, found := req.Request.Params["promo_name"]
	if found {
		return PromoName.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetBeginTime(BeginTime string) {
	req.Request.Params["begin_time"] = BeginTime
}

func (req *SellerPromotionFullCreateRequest) GetBeginTime() string {
	BeginTime, found := req.Request.Params["begin_time"]
	if found {
		return BeginTime.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetEndTime(EndTime string) {
	req.Request.Params["end_time"] = EndTime
}

func (req *SellerPromotionFullCreateRequest) GetEndTime() string {
	EndTime, found := req.Request.Params["end_time"]
	if found {
		return EndTime.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetSlogan(Slogan string) {
	req.Request.Params["slogan"] = Slogan
}

func (req *SellerPromotionFullCreateRequest) GetSlogan() string {
	Slogan, found := req.Request.Params["slogan"]
	if found {
		return Slogan.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetComment(Comment string) {
	req.Request.Params["comment"] = Comment
}

func (req *SellerPromotionFullCreateRequest) GetComment() string {
	Comment, found := req.Request.Params["comment"]
	if found {
		return Comment.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetLink(Link string) {
	req.Request.Params["link"] = Link
}

func (req *SellerPromotionFullCreateRequest) GetLink() string {
	Link, found := req.Request.Params["link"]
	if found {
		return Link.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetPlusMember(PlusMember int64) {
	req.Request.Params["plusMember"] = PlusMember
}

func (req *SellerPromotionFullCreateRequest) GetPlusMember() int64 {
	PlusMember, found := req.Request.Params["plusMember"]
	if found {
		return PlusMember.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetAllowOthersOperate(AllowOthersOperate bool) {
	req.Request.Params["allow_others_operate"] = AllowOthersOperate
}

func (req *SellerPromotionFullCreateRequest) GetAllowOthersOperate() bool {
	AllowOthersOperate, found := req.Request.Params["allow_others_operate"]
	if found {
		return AllowOthersOperate.(bool)
	}
	return false
}

func (req *SellerPromotionFullCreateRequest) SetAllowOthersCheck(AllowOthersCheck bool) {
	req.Request.Params["allow_others_check"] = AllowOthersCheck
}

func (req *SellerPromotionFullCreateRequest) GetAllowOthersCheck() bool {
	AllowOthersCheck, found := req.Request.Params["allow_others_check"]
	if found {
		return AllowOthersCheck.(bool)
	}
	return false
}

func (req *SellerPromotionFullCreateRequest) SetAllowOtherUserOperate(AllowOtherUserOperate bool) {
	req.Request.Params["allow_other_user_operate"] = AllowOtherUserOperate
}

func (req *SellerPromotionFullCreateRequest) GetAllowOtherUserOperate() bool {
	AllowOtherUserOperate, found := req.Request.Params["allow_other_user_operate"]
	if found {
		return AllowOtherUserOperate.(bool)
	}
	return false
}

func (req *SellerPromotionFullCreateRequest) SetAllowOtherUserCheck(AllowOtherUserCheck bool) {
	req.Request.Params["allow_other_user_check"] = AllowOtherUserCheck
}

func (req *SellerPromotionFullCreateRequest) GetAllowOtherUserCheck() bool {
	AllowOtherUserCheck, found := req.Request.Params["allow_other_user_check"]
	if found {
		return AllowOtherUserCheck.(bool)
	}
	return false
}

func (req *SellerPromotionFullCreateRequest) SetNeedManualCheck(NeedManualCheck bool) {
	req.Request.Params["need_manual_check"] = NeedManualCheck
}

func (req *SellerPromotionFullCreateRequest) GetNeedManualCheck() bool {
	NeedManualCheck, found := req.Request.Params["need_manual_check"]
	if found {
		return NeedManualCheck.(bool)
	}
	return false
}

func (req *SellerPromotionFullCreateRequest) SetFreqBound(FreqBound int64) {
	req.Request.Params["freq_bound"] = FreqBound
}

func (req *SellerPromotionFullCreateRequest) GetFreqBound() int64 {
	FreqBound, found := req.Request.Params["freq_bound"]
	if found {
		return FreqBound.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetPerMaxNum(PerMaxNum int64) {
	req.Request.Params["per_max_num"] = PerMaxNum
}

func (req *SellerPromotionFullCreateRequest) GetPerMaxNum() int64 {
	PerMaxNum, found := req.Request.Params["per_max_num"]
	if found {
		return PerMaxNum.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetPerMinNum(PerMinNum int64) {
	req.Request.Params["per_min_num"] = PerMinNum
}

func (req *SellerPromotionFullCreateRequest) GetPerMinNum() int64 {
	PerMinNum, found := req.Request.Params["per_min_num"]
	if found {
		return PerMinNum.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetPropType(PropType int64) {
	req.Request.Params["prop_type"] = PropType
}

func (req *SellerPromotionFullCreateRequest) GetPropType() int64 {
	PropType, found := req.Request.Params["prop_type"]
	if found {
		return PropType.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetPropNum(PropNum int64) {
	req.Request.Params["prop_num"] = PropNum
}

func (req *SellerPromotionFullCreateRequest) GetPropNum() int64 {
	PropNum, found := req.Request.Params["prop_num"]
	if found {
		return PropNum.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetPropUsedWay(PropUsedWay int64) {
	req.Request.Params["prop_used_way"] = PropUsedWay
}

func (req *SellerPromotionFullCreateRequest) GetPropUsedWay() int64 {
	PropUsedWay, found := req.Request.Params["prop_used_way"]
	if found {
		return PropUsedWay.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetCouponValidDays(CouponValidDays int64) {
	req.Request.Params["coupon_valid_days"] = CouponValidDays
}

func (req *SellerPromotionFullCreateRequest) GetCouponValidDays() int64 {
	CouponValidDays, found := req.Request.Params["coupon_valid_days"]
	if found {
		return CouponValidDays.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetTokenUseNum(TokenUseNum int64) {
	req.Request.Params["token_use_num"] = TokenUseNum
}

func (req *SellerPromotionFullCreateRequest) GetTokenUseNum() int64 {
	TokenUseNum, found := req.Request.Params["token_use_num"]
	if found {
		return TokenUseNum.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetUserPins(UserPins string) {
	req.Request.Params["user_pins"] = UserPins
}

func (req *SellerPromotionFullCreateRequest) GetUserPins() string {
	UserPins, found := req.Request.Params["user_pins"]
	if found {
		return UserPins.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetPromoAreaType(PromoAreaType int64) {
	req.Request.Params["promo_area_type"] = PromoAreaType
}

func (req *SellerPromotionFullCreateRequest) GetPromoAreaType() int64 {
	PromoAreaType, found := req.Request.Params["promo_area_type"]
	if found {
		return PromoAreaType.(int64)
	}
	return 0
}

func (req *SellerPromotionFullCreateRequest) SetPromoAreas(PromoAreas string) {
	req.Request.Params["promo_areas"] = PromoAreas
}

func (req *SellerPromotionFullCreateRequest) GetPromoAreas() string {
	PromoAreas, found := req.Request.Params["promo_areas"]
	if found {
		return PromoAreas.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetSkuId(SkuId string) {
	req.Request.Params["sku_id"] = SkuId
}

func (req *SellerPromotionFullCreateRequest) GetSkuId() string {
	SkuId, found := req.Request.Params["sku_id"]
	if found {
		return SkuId.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetPromoPrice(PromoPrice string) {
	req.Request.Params["promo_price"] = PromoPrice
}

func (req *SellerPromotionFullCreateRequest) GetPromoPrice() string {
	PromoPrice, found := req.Request.Params["promo_price"]
	if found {
		return PromoPrice.(string)
	}
	return ""
}

func (req *SellerPromotionFullCreateRequest) SetLimitNum(LimitNum string) {
	req.Request.Params["limit_num"] = LimitNum
}

func (req *SellerPromotionFullCreateRequest) GetLimitNum() string {
	LimitNum, found := req.Request.Params["limit_num"]
	if found {
		return LimitNum.(string)
	}
	return ""
}
