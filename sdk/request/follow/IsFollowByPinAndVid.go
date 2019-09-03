package follow

import (
	"github.com/XiBao/jos/sdk"
)

type IsFollowByPinAndVidRequest struct {
	Request *sdk.Request
}

// create new request
func NewIsFollowByPinAndVidRequest() (req *IsFollowByPinAndVidRequest) {
	request := sdk.Request{MethodName: "jingdong.follow.vender.read.isFollowByPinAndVid", Params: make(map[string]interface{}, 2)}
	req = &IsFollowByPinAndVidRequest{
		Request: &request,
	}
	return
}

func (req *IsFollowByPinAndVidRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *IsFollowByPinAndVidRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *IsFollowByPinAndVidRequest) SetShopId(shopId uint64) {
	req.Request.Params["shopId"] = shopId
}

func (req *IsFollowByPinAndVidRequest) GetShopId() uint64 {
	shopId, found := req.Request.Params["shopId"]
	if found {
		return shopId.(uint64)
	}
	return 0
}
