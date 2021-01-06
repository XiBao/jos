package crm

import (
	"github.com/XiBao/jos/sdk"
)

type ShopGiftGroupModelListRequest struct {
	Request *sdk.Request
}

// create new request
func NewShopGiftGroupModelListRequest() (req *ShopGiftGroupModelListRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.shopGift.getGroupModelList", Params: make(map[string]interface{}, 1)}
	req = &ShopGiftGroupModelListRequest{
		Request: &request,
	}
	return
}

func (req *ShopGiftGroupModelListRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *ShopGiftGroupModelListRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}
