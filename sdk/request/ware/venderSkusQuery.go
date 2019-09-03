package ware

import (
	"github.com/XiBao/jos/sdk"
)

type WareVenderSkusQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewWareVenderSkusQueryRequest() (req *WareVenderSkusQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.new.ware.vender.skus.query", Params: make(map[string]interface{}, 1)}
	req = &WareVenderSkusQueryRequest{
		Request: &request,
	}
	return
}

func (req *WareVenderSkusQueryRequest) SetIndex(index int) {
	req.Request.Params["index"] = index
}

func (req *WareVenderSkusQueryRequest) GetIndex() int {
	index, found := req.Request.Params["index"]
	if found {
		return index.(int)
	}
	return 0
}
