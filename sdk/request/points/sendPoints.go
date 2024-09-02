package points

import (
	"github.com/XiBao/jos/sdk"
)

type SendPointsRequest struct {
	Request *sdk.Request
}

func NewSendPointsRequest() (req *SendPointsRequest) {
	request := sdk.Request{MethodName: "jingdong.points.jos.sendPoints", Params: make(map[string]interface{}, 4)}
	req = &SendPointsRequest{
		Request: &request,
	}
	return
}

func (req *SendPointsRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *SendPointsRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *SendPointsRequest) SetBusinessId(businessId string) {
	req.Request.Params["businessId"] = businessId
}

func (req *SendPointsRequest) GetBusinessId() string {
	businessId, found := req.Request.Params["businessId"]
	if found {
		return businessId.(string)
	}
	return ""
}

func (req *SendPointsRequest) SetSourceType(sourceType uint8) {
	req.Request.Params["sourceType"] = sourceType
}

func (req *SendPointsRequest) GetSourceType() uint8 {
	sourceType, found := req.Request.Params["sourceType"]
	if found {
		return sourceType.(uint8)
	}
	return 0
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
