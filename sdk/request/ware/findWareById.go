package ware

import (
	"github.com/XiBao/jos/sdk"
)

type FindWareByIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewFindWareByIdRequest() (req *FindWareByIdRequest) {
	request := sdk.Request{MethodName: "jingdong.ware.read.findWareById", Params: make(map[string]interface{}, 2)}
	req = &FindWareByIdRequest{
		Request: &request,
	}
	return
}

func (req *FindWareByIdRequest) SetFields(fields string) {
	req.Request.Params["field"] = fields
}

func (req *FindWareByIdRequest) GetFields() string {
	fields, found := req.Request.Params["field"]
	if found {
		return fields.(string)
	}
	return ""
}

func (req *FindWareByIdRequest) SetWareId(wareId uint64) {
	req.Request.Params["wareId"] = wareId
}

func (req *FindWareByIdRequest) GetWareId() uint64 {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.(uint64)
	}
	return 0
}
