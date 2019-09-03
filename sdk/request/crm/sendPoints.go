package crm

import (
	"github.com/XiBao/jos/sdk"
)

type SendPointsRequest struct {
	Request *sdk.Request
}

// create new request
func NewSendPointsRequest() (req *SendPointsRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.sendPoints", Params: make(map[string]interface{}, 5)}
	req = &SendPointsRequest{
		Request: &request,
	}
	return
}

func (req *SendPointsRequest) SetSysName(sysName string) {
	req.Request.Params["sysName"] = sysName
}

func (req *SendPointsRequest) GetSysName() string {
	sysName, found := req.Request.Params["sysName"]
	if found {
		return sysName.(string)
	}
	return ""
}

func (req *SendPointsRequest) SetCustomerPin(customerPin string) {
	req.Request.Params["customerPin"] = customerPin
}

func (req *SendPointsRequest) GetCustomerPin() string {
	pin, found := req.Request.Params["customerPin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SendPointsRequest) SetResId(resId string) {
	req.Request.Params["resId"] = resId
}

func (req *SendPointsRequest) GetResId() string {
	resId, found := req.Request.Params["resId"]
	if found {
		return resId.(string)
	}
	return ""
}

func (req *SendPointsRequest) SetSourceType(sourceType string) {
	req.Request.Params["sourceType"] = sourceType
}

func (req *SendPointsRequest) GetSourceType() string {
	sourceType, found := req.Request.Params["sourceType"]
	if found {
		return sourceType.(string)
	}
	return ""
}

func (req *SendPointsRequest) SetPoints(points int64) {
	req.Request.Params["points"] = points
}

func (req *SendPointsRequest) GetPoints() int64 {
	points, found := req.Request.Params["points"]
	if found {
		return points.(int64)
	}
	return 0
}
