package crm

import (
	"github.com/XiBao/jos/sdk"
)

type PopShopGiftListRequest struct {
	Request *sdk.Request
}

// create new request
func NewPopShopGiftListRequest() (req *PopShopGiftListRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.shopGift.getPopShopGiftList", Params: make(map[string]interface{}, 3)}
	req = &PopShopGiftListRequest{
		Request: &request,
	}
	return
}

func (req *PopShopGiftListRequest) SetUserPin(userPin string) {
	req.Request.Params["userPin"] = userPin
}

func (req *PopShopGiftListRequest) GetUserPin() string {
	userPin, found := req.Request.Params["userPin"]
	if found {
		return userPin.(string)
	}
	return ""
}

func (req *PopShopGiftListRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *PopShopGiftListRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *PopShopGiftListRequest) SetOpenIdBuyer(openIdBuyer string) {
	req.Request.Params["open_id_buyer"] = openIdBuyer
}

func (req *PopShopGiftListRequest) GetOpenIdBuyer() string {
	openIdBuyer, found := req.Request.Params["open_id_buyer"]
	if found {
		return openIdBuyer.(string)
	}
	return ""
}
