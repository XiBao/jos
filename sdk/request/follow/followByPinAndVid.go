package follow

import (
	"github.com/XiBao/jos/sdk"
)

type FollowByPinAndVidRequest struct {
	Request *sdk.Request
}

// create new request
func NewFollowByPinAndVidRequest() (req *FollowByPinAndVidRequest) {
	request := sdk.Request{MethodName: "jingdong.follow.vender.write.followByPinAndVid", Params: make(map[string]interface{}, 2)}
	req = &FollowByPinAndVidRequest{
		Request: &request,
	}
	return
}

func (req *FollowByPinAndVidRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *FollowByPinAndVidRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *FollowByPinAndVidRequest) SetShopId(shopId uint64) {
	req.Request.Params["shopId"] = shopId
}

func (req *FollowByPinAndVidRequest) GetShopId() uint64 {
	shopId, found := req.Request.Params["shopId"]
	if found {
		return shopId.(uint64)
	}
	return 0
}
