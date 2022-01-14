package vender

import (
	"github.com/XiBao/jos/sdk"
)

type GetBasicVenderInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetBasicVenderInfoRequest() (req *GetBasicVenderInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.vender.vbinfo.getBasicVenderInfoByVenderId", Params: make(map[string]interface{}, 2)}
	req = &GetBasicVenderInfoRequest{
		Request: &request,
	}
	return
}

func (req *GetBasicVenderInfoRequest) SetColNames(colNames []string) {
	req.Request.Params["colNames"] = colNames
}

func (req *GetBasicVenderInfoRequest) GetColNames() []string {
	colNames, found := req.Request.Params["colNames"]
	if found {
		return colNames.([]string)
	}
	return nil
}

func (req *GetBasicVenderInfoRequest) SetSource(source int) {
	req.Request.Params["source"] = source
}

func (req *GetBasicVenderInfoRequest) GetSource() int {
	source, found := req.Request.Params["source"]
	if found {
		return source.(int)
	}
	return 0
}
