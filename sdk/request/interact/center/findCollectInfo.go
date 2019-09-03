package center

import (
	"github.com/XiBao/jos/sdk"
)

type FindCollectInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewFindCollectInfoRequest() (req *FindCollectInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.interact.center.api.service.read.findCollectInfo", Params: make(map[string]interface{}, 6)}
	req = &FindCollectInfoRequest{
		Request: &request,
	}
	return
}

func (req *FindCollectInfoRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *FindCollectInfoRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *FindCollectInfoRequest) SetAppId(appId uint64) {
	req.Request.Params["appId"] = appId
}

func (req *FindCollectInfoRequest) GetAppId() uint64 {
	appId, found := req.Request.Params["appId"]
	if found {
		return appId.(uint64)
	}
	return 0
}

func (req *FindCollectInfoRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *FindCollectInfoRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *FindCollectInfoRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *FindCollectInfoRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *FindCollectInfoRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *FindCollectInfoRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}

func (req *FindCollectInfoRequest) SetType(aType uint8) {
	req.Request.Params["type"] = aType
}

func (req *FindCollectInfoRequest) GetType() uint8 {
	aType, found := req.Request.Params["type"]
	if found {
		return aType.(uint8)
	}
	return 0
}
