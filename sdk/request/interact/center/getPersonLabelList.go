package center

import (
	"github.com/XiBao/jos/sdk"
)

type GetPersonLabelListRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetPersonLabelListRequest() (req *FindCollectInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.interact.center.api.service.read.getPersonLabelList", Params: make(map[string]interface{}, 2)}
	req = &FindCollectInfoRequest{
		Request: &request,
	}
	return
}

func (req *GetPersonLabelListRequest) SetAppName(appName string) {
	req.Request.Params["appName"] = appName
}

func (req *GetPersonLabelListRequest) GetAppName() string {
	appName, found := req.Request.Params["appName"]
	if found {
		return appName.(string)
	}
	return ""
}

func (req *GetPersonLabelListRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *GetPersonLabelListRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}
