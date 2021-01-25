package crm

import (
	"github.com/XiBao/jos/sdk"
)

type ShopGiftDrawRequest struct {
	Request *sdk.Request
}

// create new request
func NewShopGiftDrawRequest() (req *ShopGiftDrawRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.shopGift.drawShopGift", Params: make(map[string]interface{}, 6)}
	req = &ShopGiftDrawRequest{
		Request: &request,
	}
	return
}

func (req *ShopGiftDrawRequest) SetUserPin(userPin string) {
	req.Request.Params["userPin"] = userPin
}

func (req *ShopGiftDrawRequest) GetUserPin() string {
	userPin, found := req.Request.Params["userPin"]
	if found {
		return userPin.(string)
	}
	return ""
}

func (req *ShopGiftDrawRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *ShopGiftDrawRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}

func (req *ShopGiftDrawRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *ShopGiftDrawRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *ShopGiftDrawRequest) SetBussinessId(bussinessId string) {
	req.Request.Params["bussinessId"] = bussinessId
}

func (req *ShopGiftDrawRequest) GetBussinessId() string {
	bussinessId, found := req.Request.Params["bussinessId"]
	if found {
		return bussinessId.(string)
	}
	return ""
}

func (req *ShopGiftDrawRequest) SetChannel(channel uint8) {
	req.Request.Params["channel"] = channel
}

func (req *ShopGiftDrawRequest) GetChannel() uint8 {
	channel, found := req.Request.Params["channel"]
	if found {
		return channel.(uint8)
	}
	return 0
}

func (req *ShopGiftDrawRequest) SetOpenIdBuyer(openIdBuyer string) {
	req.Request.Params["open_id_buyer"] = openIdBuyer
}

func (req *ShopGiftDrawRequest) GetOpenIdBuyer() string {
	openIdBuyer, found := req.Request.Params["open_id_buyer"]
	if found {
		return openIdBuyer.(string)
	}
	return ""
}
