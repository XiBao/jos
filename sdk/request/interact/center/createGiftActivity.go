package center

import (
	"github.com/XiBao/jos/sdk"
)

type CreateGiftActivityRequest struct {
	Request *sdk.Request
}

// create new request
func NewCreateGiftActivityRequest() (req *CreateGiftActivityRequest) {
	request := sdk.Request{MethodName: "jingdong.interact.center.api.service.write.createGiftActivity", Params: make(map[string]interface{}, 18)}
	req = &CreateGiftActivityRequest{
		Request: &request,
	}
	return
}

func (req *CreateGiftActivityRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *CreateGiftActivityRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetAppId(appId uint64) {
	req.Request.Params["appId"] = appId
}

func (req *CreateGiftActivityRequest) GetAppId() uint64 {
	appId, found := req.Request.Params["appId"]
	if found {
		return appId.(uint64)
	}
	return 0
}

func (req *CreateGiftActivityRequest) SetCouponId(couponId uint64) {
	req.Request.Params["couponId"] = couponId
}

func (req *CreateGiftActivityRequest) GetCouponId() uint64 {
	couponId, found := req.Request.Params["couponId"]
	if found {
		return couponId.(uint64)
	}
	return 0
}

func (req *CreateGiftActivityRequest) SetQuota(quota string) {
	req.Request.Params["quota"] = quota
}

func (req *CreateGiftActivityRequest) GetQuota() string {
	quota, found := req.Request.Params["quota"]
	if found {
		return quota.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetDiscount(discount string) {
	req.Request.Params["discount"] = discount
}

func (req *CreateGiftActivityRequest) GetDiscount() string {
	discount, found := req.Request.Params["discount"]
	if found {
		return discount.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetSendCount(sendCount string) {
	req.Request.Params["sendCount"] = sendCount
}

func (req *CreateGiftActivityRequest) GetSendCount() string {
	discount, found := req.Request.Params["sendCount"]
	if found {
		return discount.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetPrizeEndTime(prizeEndTime string) {
	req.Request.Params["prizeEndTime"] = prizeEndTime
}

func (req *CreateGiftActivityRequest) GetPrizeEndTime() string {
	prizeEndTime, found := req.Request.Params["prizeEndTime"]
	if found {
		return prizeEndTime.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetValidateDay(validateDay string) {
	req.Request.Params["validateDay"] = validateDay
}

func (req *CreateGiftActivityRequest) GetValidateDay() string {
	validateDay, found := req.Request.Params["validateDay"]
	if found {
		return validateDay.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *CreateGiftActivityRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *CreateGiftActivityRequest) SetIsPrize(isPrize uint8) {
	req.Request.Params["isPrize"] = isPrize
}

func (req *CreateGiftActivityRequest) GetIsPrize() uint8 {
	isPrize, found := req.Request.Params["isPrize"]
	if found {
		return isPrize.(uint8)
	}
	return 0
}

func (req *CreateGiftActivityRequest) SetPrizeType(prizeType uint8) {
	req.Request.Params["prizeType"] = prizeType
}

func (req *CreateGiftActivityRequest) GetPrizeType() uint8 {
	prizeType, found := req.Request.Params["prizeType"]
	if found {
		return prizeType.(uint8)
	}
	return 0
}

func (req *CreateGiftActivityRequest) SetSource(source uint8) {
	req.Request.Params["source"] = source
}

func (req *CreateGiftActivityRequest) GetSource() uint8 {
	source, found := req.Request.Params["source"]
	if found {
		return source.(uint8)
	}
	return 0
}

func (req *CreateGiftActivityRequest) SetType(aType uint8) {
	req.Request.Params["type"] = aType
}

func (req *CreateGiftActivityRequest) GetType() uint8 {
	aType, found := req.Request.Params["type"]
	if found {
		return aType.(uint8)
	}
	return 0
}

func (req *CreateGiftActivityRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *CreateGiftActivityRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *CreateGiftActivityRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetCreator(creator string) {
	req.Request.Params["creator"] = creator
}

func (req *CreateGiftActivityRequest) GetCreator() string {
	creator, found := req.Request.Params["creator"]
	if found {
		return creator.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *CreateGiftActivityRequest) GetName() string {
	name, found := req.Request.Params["name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetSourceName(sourceName string) {
	req.Request.Params["sourceName"] = sourceName
}

func (req *CreateGiftActivityRequest) GetSourceName() string {
	sourceName, found := req.Request.Params["sourceName"]
	if found {
		return sourceName.(string)
	}
	return ""
}

func (req *CreateGiftActivityRequest) SetExt(ext string) {
	req.Request.Params["ext"] = ext
}

func (req *CreateGiftActivityRequest) GetExt() string {
	ext, found := req.Request.Params["ext"]
	if found {
		return ext.(string)
	}
	return ""
}
