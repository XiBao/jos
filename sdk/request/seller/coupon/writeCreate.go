package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerCouponWriteCreateRequest struct {
	Request *sdk.Request
}

// create new request
func NewSellerCouponWriteCreateRequest() (req *SellerCouponWriteCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.seller.coupon.write.create", Params: make(map[string]interface{}, 18)}
	req = &SellerCouponWriteCreateRequest{
		Request: &request,
	}
	return
}

func (req *SellerCouponWriteCreateRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SellerCouponWriteCreateRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *SellerCouponWriteCreateRequest) SetPort(port string) {
	req.Request.Params["port"] = port
}

func (req *SellerCouponWriteCreateRequest) GetPort() string {
	port, found := req.Request.Params["port"]
	if found {
		return port.(string)
	}
	return ""
}

func (req *SellerCouponWriteCreateRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *SellerCouponWriteCreateRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *SellerCouponWriteCreateRequest) SetType(cType uint8) {
	req.Request.Params["type"] = cType
}

func (req *SellerCouponWriteCreateRequest) GetType() uint8 {
	cType, found := req.Request.Params["type"]
	if found {
		return cType.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetBindType(bindType uint8) {
	req.Request.Params["bindType"] = bindType
}

func (req *SellerCouponWriteCreateRequest) GetBindType() uint8 {
	bindType, found := req.Request.Params["bindType"]
	if found {
		return bindType.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetGrantType(grantType uint8) {
	req.Request.Params["grantType"] = grantType
}

func (req *SellerCouponWriteCreateRequest) GetGrantType() uint8 {
	grantType, found := req.Request.Params["grantType"]
	if found {
		return grantType.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetNum(num uint64) {
	req.Request.Params["num"] = num
}

func (req *SellerCouponWriteCreateRequest) GetNum() uint64 {
	num, found := req.Request.Params["num"]
	if found {
		return num.(uint64)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetDiscount(discount uint) {
	req.Request.Params["discount"] = discount
}

func (req *SellerCouponWriteCreateRequest) GetDiscount() uint {
	discount, found := req.Request.Params["discount"]
	if found {
		return discount.(uint)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetQuota(quota uint) {
	req.Request.Params["quota"] = quota
}

func (req *SellerCouponWriteCreateRequest) GetQuota() uint {
	quota, found := req.Request.Params["quota"]
	if found {
		return quota.(uint)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetValidityType(validityType uint8) {
	req.Request.Params["validityType"] = validityType
}

func (req *SellerCouponWriteCreateRequest) GetValidityType() uint8 {
	validityType, found := req.Request.Params["validityType"]
	if found {
		return validityType.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetDays(days uint) {
	req.Request.Params["days"] = days
}

func (req *SellerCouponWriteCreateRequest) GetDays() uint {
	days, found := req.Request.Params["days"]
	if found {
		return days.(uint)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetBeginTime(beginTime uint64) {
	req.Request.Params["beginTime"] = beginTime
}

func (req *SellerCouponWriteCreateRequest) GetBeginTime() uint64 {
	beginTime, found := req.Request.Params["beginTime"]
	if found {
		return beginTime.(uint64)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetEndTime(endTime uint64) {
	req.Request.Params["endTime"] = endTime
}

func (req *SellerCouponWriteCreateRequest) GetEndTime() uint64 {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(uint64)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetPassword(password string) {
	req.Request.Params["password"] = password
}

func (req *SellerCouponWriteCreateRequest) GetPassword() string {
	password, found := req.Request.Params["password"]
	if found {
		return password.(string)
	}
	return ""
}

func (req *SellerCouponWriteCreateRequest) SetBatchKey(batchKey string) {
	req.Request.Params["batchKey"] = batchKey
}

func (req *SellerCouponWriteCreateRequest) GetBatchKey() string {
	batchKey, found := req.Request.Params["batchKey"]
	if found {
		return batchKey.(string)
	}
	return ""
}

func (req *SellerCouponWriteCreateRequest) SetMember(member uint) {
	req.Request.Params["member"] = member
}

func (req *SellerCouponWriteCreateRequest) GetMember() uint {
	member, found := req.Request.Params["member"]
	if found {
		return member.(uint)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetTakeBeginTime(takeBeginTime uint64) {
	req.Request.Params["takeBeginTime"] = takeBeginTime
}

func (req *SellerCouponWriteCreateRequest) GetTakeBeginTime() uint64 {
	takeBeginTime, found := req.Request.Params["takeBeginTime"]
	if found {
		return takeBeginTime.(uint64)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetTakeEndTime(takeEndTime uint64) {
	req.Request.Params["takeEndTime"] = takeEndTime
}

func (req *SellerCouponWriteCreateRequest) GetTakeEndTime() uint64 {
	takeEndTime, found := req.Request.Params["takeEndTime"]
	if found {
		return takeEndTime.(uint64)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetTakeRule(takeRule uint8) {
	req.Request.Params["takeRule"] = takeRule
}

func (req *SellerCouponWriteCreateRequest) GetTakeRule() uint8 {
	takeRule, found := req.Request.Params["takeRule"]
	if found {
		return takeRule.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetTakeNum(takeNum uint) {
	req.Request.Params["takeNum"] = takeNum
}

func (req *SellerCouponWriteCreateRequest) GetTakeNum() uint {
	takeNum, found := req.Request.Params["takeNum"]
	if found {
		return takeNum.(uint)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetDisplay(display uint8) {
	req.Request.Params["display"] = display
}

func (req *SellerCouponWriteCreateRequest) GetDisplay() uint8 {
	display, found := req.Request.Params["display"]
	if found {
		return display.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetPlatformType(platformType uint8) {
	req.Request.Params["platformType"] = platformType
}

func (req *SellerCouponWriteCreateRequest) GetPlatformType() uint8 {
	platformType, found := req.Request.Params["platformType"]
	if found {
		return platformType.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetPlatform(platform uint8) {
	req.Request.Params["platform"] = platform
}

func (req *SellerCouponWriteCreateRequest) GetPlatform() uint8 {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetImgUrl(imgUrl string) {
	req.Request.Params["imgUrl"] = imgUrl
}

func (req *SellerCouponWriteCreateRequest) GetImgUrl() string {
	imgUrl, found := req.Request.Params["imgUrl"]
	if found {
		return imgUrl.(string)
	}
	return ""
}

func (req *SellerCouponWriteCreateRequest) SetBoundStatus(boundStatus uint) {
	req.Request.Params["boundStatus"] = boundStatus
}

func (req *SellerCouponWriteCreateRequest) GetBoundStatus() uint {
	boundStatus, found := req.Request.Params["boundStatus"]
	if found {
		return boundStatus.(uint)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetJdNum(jdNum uint) {
	req.Request.Params["jdNum"] = jdNum
}

func (req *SellerCouponWriteCreateRequest) GetJdNum() uint {
	jdNum, found := req.Request.Params["jdNum"]
	if found {
		return jdNum.(uint)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetItemId(itemId uint64) {
	req.Request.Params["itemId"] = itemId
}

func (req *SellerCouponWriteCreateRequest) GetItemId() uint64 {
	itemId, found := req.Request.Params["itemId"]
	if found {
		return itemId.(uint64)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetShareType(shareType uint8) {
	req.Request.Params["shareType"] = shareType
}

func (req *SellerCouponWriteCreateRequest) GetShareType() uint8 {
	shareType, found := req.Request.Params["shareType"]
	if found {
		return shareType.(uint8)
	}
	return 0
}

func (req *SellerCouponWriteCreateRequest) SetSkuId(skuId string) {
	req.Request.Params["skuId"] = skuId
}

func (req *SellerCouponWriteCreateRequest) GetSkuId() string {
	skuId, found := req.Request.Params["skuId"]
	if found {
		return skuId.(string)
	}
	return ""
}
