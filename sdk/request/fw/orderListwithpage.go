package fw

import (
	"github.com/XiBao/jos/sdk"
)

type OrderListwithpageRequest struct {
	Request *sdk.Request
}

// create new request
func NewOrderListwithpageRequest() (req *OrderListwithpageRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.fw.order.listwithpage", Params: make(map[string]interface{}, 4)}
	req = &OrderListwithpageRequest{
		Request: &request,
	}
	return
}

func (req *OrderListwithpageRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *OrderListwithpageRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *OrderListwithpageRequest) SetFwsPin(fwsPin string) {
	req.Request.Params["fwsPin"] = fwsPin
}

func (req *OrderListwithpageRequest) GetFwsPin() string {
	fwsPin, found := req.Request.Params["fwsPin"]
	if found {
		return fwsPin.(string)
	}
	return ""
}

func (req *OrderListwithpageRequest) SetCurrentPage(currentPage int) {
	req.Request.Params["currentPage"] = currentPage
}

func (req *OrderListwithpageRequest) GetCurrentPage() int {
	currentPage, found := req.Request.Params["currentPage"]
	if found {
		return currentPage.(int)
	}
	return 0
}

func (req *OrderListwithpageRequest) SetServiceCode(serviceCode string) {
	req.Request.Params["serviceCode"] = serviceCode
}

func (req *OrderListwithpageRequest) GetServiceCode() string {
	serviceCode, found := req.Request.Params["serviceCode"]
	if found {
		return serviceCode.(string)
	}
	return ""
}
