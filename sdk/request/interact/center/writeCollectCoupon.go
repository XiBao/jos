package center

import (
	"github.com/XiBao/jos/sdk"
)

type WriteCollectCouponRequest struct {
	Request *sdk.Request
}

// create new request
func NewWriteCollectCouponRequest() (req *WriteCollectCouponRequest) {
	request := sdk.Request{MethodName: "jingdong.interact.center.api.service.write.collectCoupon", Params: make(map[string]interface{}, 10)}
	req = &WriteCollectCouponRequest{
		Request: &request,
	}
	return
}

func (req *WriteCollectCouponRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *WriteCollectCouponRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *WriteCollectCouponRequest) SetAppId(appId uint64) {
	req.Request.Params["appId"] = appId
}

func (req *WriteCollectCouponRequest) GetAppId() uint64 {
	appId, found := req.Request.Params["appId"]
	if found {
		return appId.(uint64)
	}
	return 0
}

func (req *WriteCollectCouponRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *WriteCollectCouponRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *WriteCollectCouponRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *WriteCollectCouponRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *WriteCollectCouponRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *WriteCollectCouponRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}

func (req *WriteCollectCouponRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *WriteCollectCouponRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *WriteCollectCouponRequest) SetRuleId(ruleId uint64) {
	req.Request.Params["ruleId"] = ruleId
}

func (req *WriteCollectCouponRequest) GetRuldId() uint64 {
	ruleId, found := req.Request.Params["ruleId"]
	if found {
		return ruleId.(uint64)
	}
	return 0
}

func (req *WriteCollectCouponRequest) SetRfId(rfId string) {
	req.Request.Params["rfId"] = rfId
}

func (req *WriteCollectCouponRequest) GetRfId() string {
	rfId, found := req.Request.Params["rfId"]
	if found {
		return rfId.(string)
	}
	return ""
}

func (req *WriteCollectCouponRequest) SetSource(source uint8) {
	req.Request.Params["source"] = source
}

func (req *WriteCollectCouponRequest) GetSource() uint8 {
	source, found := req.Request.Params["source"]
	if found {
		return source.(uint8)
	}
	return 0
}

func (req *WriteCollectCouponRequest) SetType(aType uint8) {
	req.Request.Params["type"] = aType
}

func (req *WriteCollectCouponRequest) GetType() uint8 {
	aType, found := req.Request.Params["type"]
	if found {
		return aType.(uint8)
	}
	return 0
}
