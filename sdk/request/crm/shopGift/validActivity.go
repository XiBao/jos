package crm

import (
	"github.com/XiBao/jos/sdk"
)

type ShopGiftValidActivityRequest struct {
	Request *sdk.Request
}

// create new request
func NewShopGiftValidActivityRequest() (req *ShopGiftValidActivityRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.shopGift.getValidActivity", Params: make(map[string]interface{}, 1)}
	req = &ShopGiftValidActivityRequest{
		Request: &request,
	}
	return
}

func (req *ShopGiftValidActivityRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *ShopGiftValidActivityRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}
