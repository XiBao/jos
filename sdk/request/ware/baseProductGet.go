package ware

import (
	"github.com/XiBao/jos/sdk"
)

type WareBaseProductGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewWareBaseproductGetRequest() (req *WareBaseProductGetRequest) {
	request := sdk.Request{MethodName: "jingdong.new.ware.baseproduct.get", Params: make(map[string]interface{}, 2)}
	req = &WareBaseProductGetRequest{
		Request: &request,
	}
	return
}

func (req *WareBaseProductGetRequest) SetIds(ids string) {
	req.Request.Params["ids"] = ids
}

func (req *WareBaseProductGetRequest) GetIds() string {
	ids, found := req.Request.Params["ids"]
	if found {
		return ids.(string)
	}
	return ""
}

func (req *WareBaseProductGetRequest) SetBaseFields(baseFields string) {
	req.Request.Params["basefields"] = baseFields
}

func (req *WareBaseProductGetRequest) GetBaseFields() string {
	baseFields, found := req.Request.Params["basefields"]
	if found {
		return baseFields.(string)
	}
	return ""
}
