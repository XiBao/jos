package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionCreateRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionCreateRequest() (req *SellerPromotionCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.promotion.create", Params: make(map[string]interface{}, 40)}
	req = &SellerPromotionCreateRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionCreateRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionCreateRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionCreateRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *SellerPromotionCreateRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetBeginTime(beginTime string) {
	req.Request.Params["beginTime"] = beginTime
}

func (req *SellerPromotionCreateRequest) GetBeginTime() string {
	beginTime, found := req.Request.Params["beginTime"]
	if found {
		return beginTime.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *SellerPromotionCreateRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *SellerPromotionCreateRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetBound(bound uint8) {
	req.Request.Params["bound"] = bound
}

func (req *SellerPromotionCreateRequest) GetBound() uint8 {
	bound, found := req.Request.Params["bound"]
	if found {
		return bound.(uint8)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetMember(member uint) {
	req.Request.Params["member"] = member
}

func (req *SellerPromotionCreateRequest) GetMember() uint {
	member, found := req.Request.Params["member"]
	if found {
		return member.(uint)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetSlogan(slogan string) {
	req.Request.Params["slogan"] = slogan
}

func (req *SellerPromotionCreateRequest) GetSlogan() string {
	slogan, found := req.Request.Params["slogan"]
	if found {
		return slogan.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetComment(comment string) {
	req.Request.Params["comment"] = comment
}

func (req *SellerPromotionCreateRequest) GetComment() string {
	comment, found := req.Request.Params["comment"]
	if found {
		return comment.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetPlatform(platform uint8) {
	req.Request.Params["platform"] = platform
}

func (req *SellerPromotionCreateRequest) GetPlatform() uint8 {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint8)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetFavorMode(favorMode uint) {
	req.Request.Params["favorMode"] = favorMode
}

func (req *SellerPromotionCreateRequest) GetFavorMode() uint {
	favorMode, found := req.Request.Params["favorMode"]
	if found {
		return favorMode.(uint)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetShopMember(shopMember uint) {
	req.Request.Params["ShopMember"] = shopMember
}

func (req *SellerPromotionCreateRequest) GetShopMember() uint {
	shopMember, found := req.Request.Params["shopMember"]
	if found {
		return shopMember.(uint)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetQqMember(qqMember uint8) {
	req.Request.Params["qqMember"] = qqMember
}

func (req *SellerPromotionCreateRequest) GetQqMember() uint8 {
	qqMember, found := req.Request.Params["qqMember"]
	if found {
		return qqMember.(uint8)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetPlusMember(plusMember uint8) {
	req.Request.Params["plusMember"] = plusMember
}

func (req *SellerPromotionCreateRequest) GetPlusMember() uint8 {
	plusMember, found := req.Request.Params["plusMember"]
	if found {
		return plusMember.(uint8)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetSamMember(samMember uint8) {
	req.Request.Params["samMember"] = samMember
}

func (req *SellerPromotionCreateRequest) GetSamMember() uint8 {
	samMember, found := req.Request.Params["samMember"]
	if found {
		return samMember.(uint8)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetTokenId(tokenId uint64) {
	req.Request.Params["tokenId"] = tokenId
}

func (req *SellerPromotionCreateRequest) GetTokenId() uint64 {
	tokenId, found := req.Request.Params["tokenId"]
	if found {
		return tokenId.(uint64)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetPromoChannel(promoChannel string) {
	req.Request.Params["promoChannel"] = promoChannel
}

func (req *SellerPromotionCreateRequest) GetPromoChannel() string {
	promoChannel, found := req.Request.Params["promoChannel"]
	if found {
		return promoChannel.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetMemberLevelOnly(memberLevelOnly bool) {
	req.Request.Params["memberLevelOnly"] = memberLevelOnly
}

func (req *SellerPromotionCreateRequest) GetMemberLevelOnly() bool {
	memberLevelOnly, found := req.Request.Params["memberLevelOnly"]
	if found {
		return memberLevelOnly.(bool)
	}
	return false
}

func (req *SellerPromotionCreateRequest) SetTokenUseNum(tokenUseNum uint) {
	req.Request.Params["tokenUseNum"] = tokenUseNum
}

func (req *SellerPromotionCreateRequest) GetTokenUseNum() uint {
	tokenUseNum, found := req.Request.Params["tokenUseNum"]
	if found {
		return tokenUseNum.(uint)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetAllowOthersOperate(allowOthersOperate bool) {
	req.Request.Params["allowOthersOperate"] = allowOthersOperate
}

func (req *SellerPromotionCreateRequest) GetAllowOthersOperate() bool {
	allowOthersOperate, found := req.Request.Params["allowOthersOperate"]
	if found {
		return allowOthersOperate.(bool)
	}
	return false
}

func (req *SellerPromotionCreateRequest) SetAllowOthersCheck(allowOthersCheck bool) {
	req.Request.Params["allowOthersCheck"] = allowOthersCheck
}

func (req *SellerPromotionCreateRequest) GetAllowOthersCheck() bool {
	allowOthersCheck, found := req.Request.Params["allowOthersCheck"]
	if found {
		return allowOthersCheck.(bool)
	}
	return false
}

func (req *SellerPromotionCreateRequest) SetAllowOtherUserOperate(allowOtherUserOperate bool) {
	req.Request.Params["allowOtherUserOperate"] = allowOtherUserOperate
}

func (req *SellerPromotionCreateRequest) GetAllowOtherUserOperate() bool {
	allowOtherUserOperate, found := req.Request.Params["allowOtherUserOperate"]
	if found {
		return allowOtherUserOperate.(bool)
	}
	return false
}

func (req *SellerPromotionCreateRequest) SetAllowOtherUserCheck(allowOtherUserCheck bool) {
	req.Request.Params["allowOtherUserCheck"] = allowOtherUserCheck
}

func (req *SellerPromotionCreateRequest) GetAllowOtherUserCheck() bool {
	allowOtherUserCheck, found := req.Request.Params["allowOtherUserCheck"]
	if found {
		return allowOtherUserCheck.(bool)
	}
	return false
}

func (req *SellerPromotionCreateRequest) SetNeedManualCheck(needManualCheck bool) {
	req.Request.Params["needManualCheck"] = needManualCheck
}

func (req *SellerPromotionCreateRequest) GetNeedManualCheck() bool {
	needManualCheck, found := req.Request.Params["needManualCheck"]
	if found {
		return needManualCheck.(bool)
	}
	return false
}

func (req *SellerPromotionCreateRequest) SetSkuId(skuId string) {
	req.Request.Params["skuId"] = skuId
}

func (req *SellerPromotionCreateRequest) GetSkuId() string {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetBindType(bindType string) {
	req.Request.Params["bindType"] = bindType
}

func (req *SellerPromotionCreateRequest) GetBindType() string {
	bindType, found := req.Request.Params["bindType"]
	if found {
		return bindType.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetPromoPrice(promoPrice string) {
	req.Request.Params["promoPrice"] = promoPrice
}

func (req *SellerPromotionCreateRequest) GetPromoPrice() string {
	promoPrice, found := req.Request.Params["promoPrice"]
	if found {
		return promoPrice.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetNum(num string) {
	req.Request.Params["num"] = num
}

func (req *SellerPromotionCreateRequest) GetNum() string {
	num, found := req.Request.Params["num"]
	if found {
		return num.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetWareId(wareId string) {
	req.Request.Params["wareId"] = wareId
}

func (req *SellerPromotionCreateRequest) GetWareId() string {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetSkuName(skuName string) {
	req.Request.Params["skuName"] = skuName
}

func (req *SellerPromotionCreateRequest) GetSkuName() string {
	skuName, found := req.Request.Params["skuName"]
	if found {
		return skuName.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetJdPrice(jdPrice string) {
	req.Request.Params["jdPrice"] = jdPrice
}

func (req *SellerPromotionCreateRequest) GetJdPrice() string {
	jdPrice, found := req.Request.Params["jdPrice"]
	if found {
		return jdPrice.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetItemNum(itemNum string) {
	req.Request.Params["itemNum"] = itemNum
}

func (req *SellerPromotionCreateRequest) GetItemNum() string {
	itemNum, found := req.Request.Params["itemNum"]
	if found {
		return itemNum.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetFreqBound(freqBound uint8) {
	req.Request.Params["freqBound"] = freqBound
}

func (req *SellerPromotionCreateRequest) GetFreqBound() uint8 {
	freqBound, found := req.Request.Params["freqBound"]
	if found {
		return freqBound.(uint8)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetPerMaxNum(perMaxNum uint) {
	req.Request.Params["perMaxNum"] = perMaxNum
}

func (req *SellerPromotionCreateRequest) GetPerMaxNum() uint {
	perMaxNum, found := req.Request.Params["perMaxNum"]
	if found {
		return perMaxNum.(uint)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetPerMinNum(perMinNum uint) {
	req.Request.Params["perMinNum"] = perMinNum
}

func (req *SellerPromotionCreateRequest) GetPerMinNum() uint {
	perMinNum, found := req.Request.Params["perMinNum"]
	if found {
		return perMinNum.(uint)
	}
	return 0
}

func (req *SellerPromotionCreateRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *SellerPromotionCreateRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetUseBeginTime(useBeginTime string) {
	req.Request.Params["useBeginTime"] = useBeginTime
}

func (req *SellerPromotionCreateRequest) GetUseBeginTime() string {
	useBeginTime, found := req.Request.Params["useBeginTime"]
	if found {
		return useBeginTime.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetUseEndTime(useEndTime string) {
	req.Request.Params["useEndTime"] = useEndTime
}

func (req *SellerPromotionCreateRequest) GetUseEndTime() string {
	useEndTime, found := req.Request.Params["useEndTime"]
	if found {
		return useEndTime.(string)
	}
	return ""
}

func (req *SellerPromotionCreateRequest) SetShowTokenPrice(showTokenPrice bool) {
	req.Request.Params["showTokenPrice"] = showTokenPrice
}

func (req *SellerPromotionCreateRequest) GetShowTokenPrice() bool {
	showTokenPrice, found := req.Request.Params["showTokenPrice"]
	if found {
		return showTokenPrice.(bool)
	}
	return false
}
