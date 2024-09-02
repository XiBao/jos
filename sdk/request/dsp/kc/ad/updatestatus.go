package ad

import (
	"github.com/XiBao/jos/sdk"
)

type AdUpdateStatusRequest struct {
	Request *sdk.Request
}

func NewAdUpdateStatusRequest() (req *AdUpdateStatusRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.ad.updatestatus", Params: make(map[string]interface{}, 2)}
	req = &AdUpdateStatusRequest{
		Request: &request,
	}
	return
}

func (req *AdUpdateStatusRequest) SetStatus(status int) {
	req.Request.Params["status"] = status
}

func (req *AdUpdateStatusRequest) GetStatus() int {
	status, found := req.Request.Params["status"]
	if found {
		return status.(int)
	}
	return 0
}

func (req *AdUpdateStatusRequest) SetId(id string) {
	req.Request.Params["id"] = id
}

func (req *AdUpdateStatusRequest) GetId() string {
	id, found := req.Request.Params["id"]
	if found {
		return id.(string)
	}
	return ""
}
