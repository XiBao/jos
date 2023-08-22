package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type MarketCreateCouponRequest struct {
	Request *sdk.Request
}

// create new request
func NewMarketCreateCouponRequest() (req *MarketCreateCouponRequest) {
	request := sdk.Request{MethodName: "jingdong.market.tool.coupon.api.service.CouponWriteOuterService.createCoupon", Params: make(map[string]interface{}, 41)}
	req = &MarketCreateCouponRequest{
		Request: &request,
	}
	return
}

func (req *MarketCreateCouponRequest) SetWareGrade(wareGrade uint8) {
	req.Request.Params["wareGrade"] = wareGrade
}

func (req *MarketCreateCouponRequest) GetWareGrade() uint8 {
	wareGrade, found := req.Request.Params["wareGrade"]
	if found {
		return wareGrade.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetNum(num uint64) {
	req.Request.Params["num"] = num
}

func (req *MarketCreateCouponRequest) GetNum() uint64 {
	num, found := req.Request.Params["num"]
	if found {
		return num.(uint64)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetDiscount(discount uint) {
	req.Request.Params["discount"] = discount
}

func (req *MarketCreateCouponRequest) GetDiscount() uint {
	discount, found := req.Request.Params["discount"]
	if found {
		return discount.(uint)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetStrategyParam(strategyParam string) {
	req.Request.Params["strategyParam"] = strategyParam
}

func (req *MarketCreateCouponRequest) GetStrategyParam() string {
	strategyParam, found := req.Request.Params["strategyParam"]
	if found {
		return strategyParam.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetType(cType uint8) {
	req.Request.Params["type"] = cType
}

func (req *MarketCreateCouponRequest) GetType() uint8 {
	cType, found := req.Request.Params["type"]
	if found {
		return cType.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetSkuIdList(skuIdList string) {
	req.Request.Params["skuIdList"] = skuIdList
}

func (req *MarketCreateCouponRequest) GetSkuIdList() string {
	skuIdList, found := req.Request.Params["skuIdList"]
	if found {
		return skuIdList.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *MarketCreateCouponRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetShareType(shareType uint8) {
	req.Request.Params["shareType"] = shareType
}

func (req *MarketCreateCouponRequest) GetShareType() uint8 {
	shareType, found := req.Request.Params["shareType"]
	if found {
		return shareType.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetStoreId(storeId []string) {
	req.Request.Params["storeId"] = storeId
}

func (req *MarketCreateCouponRequest) GetStoreId() []string {
	storeId, found := req.Request.Params["storeId"]
	if found {
		return storeId.([]string)
	}
	return nil
}

func (req *MarketCreateCouponRequest) SetTakeEndTime(takeEndTime string) {
	req.Request.Params["takeEndTime"] = takeEndTime
}

func (req *MarketCreateCouponRequest) GetTakeEndTime() string {
	takeEndTime, found := req.Request.Params["takeEndTime"]
	if found {
		return takeEndTime.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetHigh(high uint) {
	req.Request.Params["high"] = high
}

func (req *MarketCreateCouponRequest) GetHigh() uint {
	high, found := req.Request.Params["high"]
	if found {
		return high.(uint)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetTakeNum(takeNum uint) {
	req.Request.Params["takeNum"] = takeNum
}

func (req *MarketCreateCouponRequest) GetTakeNum() uint {
	takeNum, found := req.Request.Params["takeNum"]
	if found {
		return takeNum.(uint)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetQuota(quota float64) {
	req.Request.Params["quota"] = quota
}

func (req *MarketCreateCouponRequest) GetQuota() float64 {
	quota, found := req.Request.Params["quota"]
	if found {
		return quota.(float64)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetOfficialType(officialType uint8) {
	req.Request.Params["officialType"] = officialType
}

func (req *MarketCreateCouponRequest) GetOfficialType() uint8 {
	officialType, found := req.Request.Params["officialType"]
	if found {
		return officialType.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetBeginTime(beginTime string) {
	req.Request.Params["beginTime"] = beginTime
}

func (req *MarketCreateCouponRequest) GetBeginTime() string {
	beginTime, found := req.Request.Params["beginTime"]
	if found {
		return beginTime.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetPromoteChannel(promoteChannel uint8) {
	req.Request.Params["promoteChannel"] = promoteChannel
}

func (req *MarketCreateCouponRequest) GetPromoteChannel() uint8 {
	promoteChannel, found := req.Request.Params["promoteChannel"]
	if found {
		return promoteChannel.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetRemainNum(remainNum uint) {
	req.Request.Params["remainNum"] = remainNum
}

func (req *MarketCreateCouponRequest) GetRemainNum() uint {
	remainNum, found := req.Request.Params["remainNum"]
	if found {
		return remainNum.(uint)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetStoreType(storeType uint) {
	req.Request.Params["storeType"] = storeType
}

func (req *MarketCreateCouponRequest) GetStoreType() uint {
	storeType, found := req.Request.Params["storeType"]
	if found {
		return storeType.(uint)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetDisplay(display uint8) {
	req.Request.Params["display"] = display
}

func (req *MarketCreateCouponRequest) GetDisplay() uint8 {
	display, found := req.Request.Params["display"]
	if found {
		return display.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetBusiCode(busiCode string) {
	req.Request.Params["busiCode"] = busiCode
}

func (req *MarketCreateCouponRequest) GetBusiCode() string {
	busiCode, found := req.Request.Params["busiCode"]
	if found {
		return busiCode.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetWareChoseType(wareChoseType uint8) {
	req.Request.Params["wareChoseType"] = wareChoseType
}

func (req *MarketCreateCouponRequest) GetWareChoseType() uint8 {
	wareChoseType, found := req.Request.Params["wareChoseType"]
	if found {
		return wareChoseType.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetUserClass(userClass int) {
	req.Request.Params["userClass"] = userClass
}

func (req *MarketCreateCouponRequest) GetUserClass() int {
	userClass, found := req.Request.Params["userClass"]
	if found {
		return userClass.(int)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetUserLevel(userLevel int) {
	req.Request.Params["userLevel"] = userLevel
}

func (req *MarketCreateCouponRequest) GetUserLevel() int {
	userLevel, found := req.Request.Params["userLevel"]
	if found {
		return userLevel.(int)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetTakeBeginTime(takeBeginTime string) {
	req.Request.Params["takeBeginTime"] = takeBeginTime
}

func (req *MarketCreateCouponRequest) GetTakeBeginTime() string {
	takeBeginTime, found := req.Request.Params["takeBeginTime"]
	if found {
		return takeBeginTime.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetValidityType(validityType uint8) {
	req.Request.Params["validityType"] = validityType
}

func (req *MarketCreateCouponRequest) GetValidityType() uint8 {
	validityType, found := req.Request.Params["validityType"]
	if found {
		return validityType.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetTakeRule(takeRule uint8) {
	req.Request.Params["takeRule"] = takeRule
}

func (req *MarketCreateCouponRequest) GetTakeRule() uint8 {
	takeRule, found := req.Request.Params["takeRule"]
	if found {
		return takeRule.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetHourCoupon(hourCoupon bool) {
	req.Request.Params["hourCoupon"] = hourCoupon
}

func (req *MarketCreateCouponRequest) GetHourCoupon() bool {
	hourCoupon, found := req.Request.Params["hourCoupon"]
	if found {
		return hourCoupon.(bool)
	}
	return false
}

func (req *MarketCreateCouponRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *MarketCreateCouponRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetActivityLink(activityLink string) {
	req.Request.Params["activityLink"] = activityLink
}

func (req *MarketCreateCouponRequest) GetActivityLink() string {
	activityLink, found := req.Request.Params["activityLink"]
	if found {
		return activityLink.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetDays(days uint) {
	req.Request.Params["days"] = days
}

func (req *MarketCreateCouponRequest) GetDays() uint {
	days, found := req.Request.Params["days"]
	if found {
		return days.(uint)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetStyle(style uint8) {
	req.Request.Params["style"] = style
}

func (req *MarketCreateCouponRequest) GetStyle() uint8 {
	style, found := req.Request.Params["style"]
	if found {
		return style.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *MarketCreateCouponRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetAdWord(adWord string) {
	req.Request.Params["adWord"] = adWord
}

func (req *MarketCreateCouponRequest) GetAdWord() string {
	adWord, found := req.Request.Params["adWord"]
	if found {
		return adWord.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetSpuIdList(spuIdList []uint64) {
	req.Request.Params["spuIdList"] = spuIdList
}

func (req *MarketCreateCouponRequest) GetSpuIdList() []uint64 {
	spuIdList, found := req.Request.Params["spuIdList"]
	if found {
		return spuIdList.([]uint64)
	}
	return nil
}

func (req *MarketCreateCouponRequest) SetChannels(channels string) {
	req.Request.Params["channels"] = channels
}

func (req *MarketCreateCouponRequest) GetChannels() string {
	channels, found := req.Request.Params["channels"]
	if found {
		return channels.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetChannelSelectType(channelSelectType uint8) {
	req.Request.Params["channelSelectType"] = channelSelectType
}

func (req *MarketCreateCouponRequest) GetChannelSelectType() uint8 {
	channelSelectType, found := req.Request.Params["channelSelectType"]
	if found {
		return channelSelectType.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetSelectType(selectType uint8) {
	req.Request.Params["selectType"] = selectType
}

func (req *MarketCreateCouponRequest) GetSelectType() uint8 {
	selectType, found := req.Request.Params["selectType"]
	if found {
		return selectType.(uint8)
	}
	return 0
}

func (req *MarketCreateCouponRequest) SetPlatforms(platforms string) {
	req.Request.Params["platforms"] = platforms
}

func (req *MarketCreateCouponRequest) GetPlatforms() string {
	platforms, found := req.Request.Params["platforms"]
	if found {
		return platforms.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *MarketCreateCouponRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *MarketCreateCouponRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *MarketCreateCouponRequest) SetAppId(appId string) {
	req.Request.Params["appId"] = appId
}

func (req *MarketCreateCouponRequest) GetAppId() string {
	appId, found := req.Request.Params["appId"]
	if found {
		return appId.(string)
	}
	return ""
}
