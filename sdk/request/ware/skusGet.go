package ware

import (
	"github.com/XiBao/jos/sdk"
)

type WareSkusGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewWareSkusGetRequest() (req *WareSkusGetRequest) {
	request := sdk.Request{MethodName: "360buy.ware.skus.get", Params: make(map[string]interface{}, 2)}
	req = &WareSkusGetRequest{
		Request: &request,
	}
	return
}

func (req *WareSkusGetRequest) SetWareIds(wareIds string) {
	req.Request.Params["ware_ids"] = wareIds
}

func (req *WareSkusGetRequest) GetWareIds() string {
	wareIds, found := req.Request.Params["ware_ids"]
	if found {
		return wareIds.(string)
	}
	return ""
}

func (req *WareSkusGetRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *WareSkusGetRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}
