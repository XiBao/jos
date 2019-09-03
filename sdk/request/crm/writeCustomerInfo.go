package crm

import (
	"github.com/XiBao/jos/sdk"
)

type WriteCustomerInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewWriteCustomerInfoRequest() (req *WriteCustomerInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.crm.writeCustomerInfo", Params: make(map[string]interface{}, 1)}
	req = &WriteCustomerInfoRequest{
		Request: &request,
	}
	return
}

func (req *WriteCustomerInfoRequest) SetLinkUrl(linkUrl string) {
	req.Request.Params["linkUrl"] = linkUrl
}

func (req *WriteCustomerInfoRequest) GetLinkUrl() string {
	linkUrl, found := req.Request.Params["linkUrl"]
	if found {
		return linkUrl.(string)
	}
	return ""
}
