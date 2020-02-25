package promotion

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromotionCouponWriteCreateRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerPromotionCouponWriteCreateRequest() (req *SellerPromotionCouponWriteCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.write.create", Params: make(map[string]interface{}, 29)}
	req = &SellerPromotionCouponWriteCreateRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromotionCouponWriteCreateRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerPromotionCouponWriteCreateRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCreateRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerPromotionCouponWriteCreateRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCreateRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *SellerPromotionCouponWriteCreateRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCreateRequest) SetType(cType uint) {
	req.Request.Params["type"] = cType
}

func (req *SellerPromotionCouponWriteCreateRequest) GetType() uint {
	cType, found := req.Request.Params["type"]
	if found {
		return cType.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetBindType(bindType uint) {
	req.Request.Params["bindType"] = bindType
}

func (req *SellerPromotionCouponWriteCreateRequest) GetBindType() uint {
	bindType, found := req.Request.Params["bindType"]
	if found {
		return bindType.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetGrantType(grantType uint) {
	req.Request.Params["grantType"] = grantType
}

func (req *SellerPromotionCouponWriteCreateRequest) GetGrantType() uint {
	grantType, found := req.Request.Params["grantType"]
	if found {
		return grantType.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetNum(num uint) {
	req.Request.Params["num"] = num
}

func (req *SellerPromotionCouponWriteCreateRequest) GetNum() uint {
	num, found := req.Request.Params["num"]
	if found {
		return num.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetDiscount(discount uint) {
	req.Request.Params["discount"] = discount
}

func (req *SellerPromotionCouponWriteCreateRequest) GetDiscount() uint {
	discount, found := req.Request.Params["discount"]
	if found {
		return discount.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetQuota(quota uint) {
	req.Request.Params["quota"] = quota
}

func (req *SellerPromotionCouponWriteCreateRequest) GetQuota() uint {
	quota, found := req.Request.Params["quota"]
	if found {
		return quota.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetValidityType(validityType uint) {
	req.Request.Params["validityType"] = validityType
}

func (req *SellerPromotionCouponWriteCreateRequest) GetValidityType() uint {
	validityType, found := req.Request.Params["validityType"]
	if found {
		return validityType.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetDays(days uint) {
	req.Request.Params["days"] = days
}

func (req *SellerPromotionCouponWriteCreateRequest) GetDays() uint {
	days, found := req.Request.Params["days"]
	if found {
		return days.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetBeginTime(beginTime int64) {
	req.Request.Params["beginTime"] = beginTime
}

func (req *SellerPromotionCouponWriteCreateRequest) GetBeginTime() int64 {
	beginTime, found := req.Request.Params["beginTime"]
	if found {
		return beginTime.(int64)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetEndTime(endTime int64) {
	req.Request.Params["endTime"] = endTime
}

func (req *SellerPromotionCouponWriteCreateRequest) GetEndTime() int64 {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(int64)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetPassword(password string) {
	req.Request.Params["password"] = password
}

func (req *SellerPromotionCouponWriteCreateRequest) GetPassword() string {
	password, found := req.Request.Params["password"]
	if found {
		return password.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCreateRequest) SetBatchKey(batchKey string) {
	req.Request.Params["batchKey"] = batchKey
}

func (req *SellerPromotionCouponWriteCreateRequest) GetBatchKey() string {
	batchKey, found := req.Request.Params["batchKey"]
	if found {
		return batchKey.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCreateRequest) SetMember(member uint) {
	req.Request.Params["member"] = member
}

func (req *SellerPromotionCouponWriteCreateRequest) GetMember() uint {
	member, found := req.Request.Params["member"]
	if found {
		return member.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetTakeBeginTime(takeBeginTime int64) {
	req.Request.Params["takeBeginTime"] = takeBeginTime
}

func (req *SellerPromotionCouponWriteCreateRequest) GetTakeBeginTime() int64 {
	takeBeginTime, found := req.Request.Params["takeBeginTime"]
	if found {
		return takeBeginTime.(int64)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetTakeEndTime(takeEndTime int64) {
	req.Request.Params["takeEndTime"] = takeEndTime
}

func (req *SellerPromotionCouponWriteCreateRequest) GetTakeEndTime() int64 {
	takeEndTime, found := req.Request.Params["takeEndTime"]
	if found {
		return takeEndTime.(int64)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetTakeRule(takeRule uint) {
	req.Request.Params["takeRule"] = takeRule
}

func (req *SellerPromotionCouponWriteCreateRequest) GetTakeRule() uint {
	takeRule, found := req.Request.Params["takeRule"]
	if found {
		return takeRule.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetTakeNum(takeNum uint) {
	req.Request.Params["takeNum"] = takeNum
}

func (req *SellerPromotionCouponWriteCreateRequest) GetTakeNum() uint {
	takeNum, found := req.Request.Params["takeNum"]
	if found {
		return takeNum.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetDisplay(display uint) {
	req.Request.Params["display"] = display
}

func (req *SellerPromotionCouponWriteCreateRequest) GetDisplay() uint {
	display, found := req.Request.Params["display"]
	if found {
		return display.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetPlatformType(platformType uint) {
	req.Request.Params["platformType"] = platformType
}

func (req *SellerPromotionCouponWriteCreateRequest) GetPlatformType() uint {
	platformType, found := req.Request.Params["platformType"]
	if found {
		return display.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetPlatform(platform uint) {
	req.Request.Params["platform"] = platform
}

func (req *SellerPromotionCouponWriteCreateRequest) GetPlatform() uint {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetImgUrl(imgUrl string) {
	req.Request.Params["imgUrl"] = imgUrl
}

func (req *SellerPromotionCouponWriteCreateRequest) GetImgUrl() string {
	imgUrl, found := req.Request.Params["imgUrl"]
	if found {
		return imgUrl.(string)
	}
	return ""
}

func (req *SellerPromotionCouponWriteCreateRequest) SetBoundStatus(boundStatus uint) {
	req.Request.Params["boundStatus"] = boundStatus
}

func (req *SellerPromotionCouponWriteCreateRequest) GetBoundStatus() uint {
	boundStatus, found := req.Request.Params["boundStatus"]
	if found {
		return boundStatus.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetJdNum(jdNum uint) {
	req.Request.Params["jdNum"] = jdNum
}

func (req *SellerPromotionCouponWriteCreateRequest) GetJdNum() uint {
	jdNum, found := req.Request.Params["jdNum"]
	if found {
		return jdNum.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetItemId(itemId uint64) {
	req.Request.Params["itemId"] = itemId
}

func (req *SellerPromotionCouponWriteCreateRequest) GetItemId() uint64 {
	itemId, found := req.Request.Params["itemId"]
	if found {
		return itemId.(uint64)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetShareType(shareType uint) {
	req.Request.Params["shareType"] = shareType
}

func (req *SellerPromotionCouponWriteCreateRequest) GetShareType() uint {
	shareType, found := req.Request.Params["shareType"]
	if found {
		return shareType.(uint)
	}
	return 0
}

func (req *SellerPromotionCouponWriteCreateRequest) SetSkuId(skuId string) {
	req.Request.Params["skuId"] = skuId
}

func (req *SellerPromotionCouponWriteCreateRequest) GetSkuId() string {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(string)
	}
	return ""
}
