package asset

import (
	"github.com/XiBao/jos/sdk"
)

type BenefitSendRequest struct {
	Request *sdk.Request
}

// create new request
func NewBenefitSendRequest() (req *BenefitSendRequest) {
	request := sdk.Request{MethodName: "jingdong.asset.benefit.send", Params: make(map[string]interface{}, 9)}
	req = &BenefitSendRequest{
		Request: &request,
	}
	return
}

func (req *BenefitSendRequest) SetTypeId(typeId int) {
	req.Request.Params["type_id"] = typeId
}

func (req *BenefitSendRequest) GetTypeId() int {
	typeId, found := req.Request.Params["type_id"]
	if found {
		return typeId.(int)
	}
	return 0
}

func (req *BenefitSendRequest) SetItemId(itemId int) {
	req.Request.Params["item_id"] = itemId
}

func (req *BenefitSendRequest) GetItemId() int {
	itemId, found := req.Request.Params["item_id"]
	if found {
		return itemId.(int)
	}
	return 0
}

func (req *BenefitSendRequest) SetQuantity(quantity int) {
	req.Request.Params["quantity"] = quantity
}

func (req *BenefitSendRequest) GetQuantity() int {
	quantityId, found := req.Request.Params["quantity_id"]
	if found {
		return quantityId.(int)
	}
	return 0
}

func (req *BenefitSendRequest) SetUserPin(userPin string) {
	req.Request.Params["user_pin"] = userPin
}

func (req *BenefitSendRequest) GetUserPin() string {
	userPin, found := req.Request.Params["user_pin"]
	if found {
		return userPin.(string)
	}
	return ""
}

func (req *BenefitSendRequest) SetToken(token string) {
	req.Request.Params["token"] = token
}

func (req *BenefitSendRequest) GetToken() string {
	token, found := req.Request.Params["token"]
	if found {
		return token.(string)
	}
	return ""
}

func (req *BenefitSendRequest) SetRequestId(requestId string) {
	req.Request.Params["request_id"] = requestId
}

func (req *BenefitSendRequest) GetRequestId() string {
	requestId, found := req.Request.Params["request_id"]
	if found {
		return requestId.(string)
	}
	return ""
}

func (req *BenefitSendRequest) SetRemark(remark string) {
	req.Request.Params["remark"] = remark
}

func (req *BenefitSendRequest) GetRemark() string {
	remark, found := req.Request.Params["remark"]
	if found {
		return remark.(string)
	}
	return ""
}

func (req *BenefitSendRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *BenefitSendRequest) GetIp() string {
	ip, found := req.Request.Params["ip"]
	if found {
		return ip.(string)
	}
	return ""
}

func (req *BenefitSendRequest) SetOpenIdBuyer(openIdBuyer string) {
	req.Request.Params["open_id_buyer"] = openIdBuyer
}

func (req *BenefitSendRequest) GetOpenIdBuyer() string {
	openIdBuyer, found := req.Request.Params["open_id_buyer"]
	if found {
		return openIdBuyer.(string)
	}
	return ""
}
