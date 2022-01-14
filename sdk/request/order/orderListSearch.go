package order

import (
	"github.com/XiBao/jos/sdk"
)

type OrderListSearchRequest struct {
	Request *sdk.Request
}

// create new request
func NewOrderListSearchRequest() (req *OrderListSearchRequest) {
	request := sdk.Request{MethodName: "jingdong.shop.order.list.search", Params: make(map[string]interface{}, 5)}
	req = &OrderListSearchRequest{
		Request: &request,
	}
	return
}

func (req *OrderListSearchRequest) SetStartDate(startDate string) {
	req.Request.Params["startDate"] = startDate
}

func (req *OrderListSearchRequest) GetStartDate() string {
	startDate, found := req.Request.Params["startDate"]
	if found {
		return startDate.(string)
	}
	return ""
}

func (req *OrderListSearchRequest) SetEndDate(endDate string) {
	req.Request.Params["endDate"] = endDate
}

func (req *OrderListSearchRequest) GetEndDate() string {
	endDate, found := req.Request.Params["endDate"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *OrderListSearchRequest) SetOrderState(orderState string) {
	req.Request.Params["orderState"] = orderState
}

func (req *OrderListSearchRequest) GetOrderState() string {
	orderState, found := req.Request.Params["orderState"]
	if found {
		return orderState.(string)
	}
	return ""
}

func (req *OrderListSearchRequest) SetPage(page int) {
	req.Request.Params["page"] = page
}

func (req *OrderListSearchRequest) GetPage() int {
	page, found := req.Request.Params["page"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *OrderListSearchRequest) SetSize(size int) {
	req.Request.Params["size"] = size
}

func (req *OrderListSearchRequest) GetSzie() int {
	size, found := req.Request.Params["size"]
	if found {
		return size.(int)
	}
	return 0
}
