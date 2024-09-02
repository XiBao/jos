package center

import (
	"github.com/XiBao/jos/sdk"
)

type FindGiftInfoByIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewFindGiftInfoByIdRequest() (req *FindGiftInfoByIdRequest) {
	request := sdk.Request{MethodName: "jingdong.interact.center.api.service.read.GiftActivityResultReadService.findGiftInfoById", Params: make(map[string]interface{}, 5)}
	req = &FindGiftInfoByIdRequest{
		Request: &request,
	}
	return
}

func (req *FindGiftInfoByIdRequest) SetAppName(AppName string) {
	req.Request.Params["appName"] = AppName
}

func (req *FindGiftInfoByIdRequest) GetAppName() string {
	AppName, found := req.Request.Params["appName"]
	if found {
		return AppName.(string)
	}
	return ""
}

func (req *FindGiftInfoByIdRequest) SetAppId(appId uint64) {
	req.Request.Params["appId"] = appId
}

func (req *FindGiftInfoByIdRequest) GetAppId() uint64 {
	appId, found := req.Request.Params["appId"]
	if found {
		return appId.(uint64)
	}
	return 0
}

func (req *FindGiftInfoByIdRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *FindGiftInfoByIdRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *FindGiftInfoByIdRequest) SetType(Type uint8) {
	req.Request.Params["type"] = Type
}

func (req *FindGiftInfoByIdRequest) GetType() uint8 {
	Type, found := req.Request.Params["type"]
	if found {
		return Type.(uint8)
	}
	return 0
}

func (req *FindGiftInfoByIdRequest) SetActivityId(ActivityId uint64) {
	req.Request.Params["activityId"] = ActivityId
}

func (req *FindGiftInfoByIdRequest) GetActivityId() uint64 {
	ActivityId, found := req.Request.Params["activityId"]
	if found {
		return ActivityId.(uint64)
	}
	return 0
}
