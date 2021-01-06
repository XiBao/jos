package user

import (
	"github.com/XiBao/jos/sdk"
)

type GetSocialInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetSocialInfoRequest() (req *GetSocialInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.user.getUserSocialInfo", Params: make(map[string]interface{}, 1)}
	req = &GetSocialInfoRequest{
		Request: &request,
	}
	return
}

func (req *GetSocialInfoRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *GetSocialInfoRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ""
}
