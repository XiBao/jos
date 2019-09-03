package link

import (
	"github.com/XiBao/jos/sdk"
)

type GetOpenLinkRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetOpenLinkRequest() (req *GetOpenLinkRequest) {
	request := sdk.Request{MethodName: "jingdong.wxsq.mjgj.link.GetOpenLink", Params: make(map[string]interface{}, 2)}
	req = &GetOpenLinkRequest{
		Request: &request,
	}
	return
}

func (req *GetOpenLinkRequest) SetRUrl(rUrl string) {
	req.Request.Params["rurl"] = rUrl
}

func (req *GetOpenLinkRequest) GetRUrl() string {
	rUrl, found := req.Request.Params["rurl"]
	if found {
		return rUrl.(string)
	}
	return ""
}

func (req *GetOpenLinkRequest) SetJump(jump uint8) {
	req.Request.Params["jump"] = jump
}

func (req *GetOpenLinkRequest) GetJump() uint8 {
	jump, found := req.Request.Params["jump"]
	if found {
		return jump.(uint8)
	}
	return 0
}
