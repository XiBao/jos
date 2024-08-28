package jzone

import (
	"github.com/XiBao/jos/sdk"
)

type AddCartItemByPinRequest struct {
	Request *sdk.Request
}

func NewAddCartItemByPinRequest() (req *AddCartItemByPinRequest) {
	request := sdk.Request{MethodName: "jingdong.jzone.addCartItemByPin", Params: make(map[string]interface{}, 3)}
	req = &AddCartItemByPinRequest{
		Request: &request,
	}
	return
}

func (req *AddCartItemByPinRequest) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *AddCartItemByPinRequest) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ``
}

func (req *AddCartItemByPinRequest) SetItemId(itemId uint64) {
	req.Request.Params["itemId"] = itemId
}

func (req *AddCartItemByPinRequest) GetItemId() uint64 {
	itemId, found := req.Request.Params["itemId"]
	if found {
		return itemId.(uint64)
	}
	return 0
}

func (req *AddCartItemByPinRequest) SetNum(num uint64) {
	req.Request.Params["num"] = num
}

func (req *AddCartItemByPinRequest) GetNum() uint64 {
	num, found := req.Request.Params["num"]
	if found {
		return num.(uint64)
	}
	return 0
}
