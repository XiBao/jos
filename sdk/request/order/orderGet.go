package order

import (
	"github.com/XiBao/jos/sdk"
)

type OrderGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewOrderGetRequest() (req *OrderGetRequest) {
	request := sdk.Request{MethodName: "jingdong.shop.order.get", Params: make(map[string]interface{}, 2)}
	req = &OrderGetRequest{
		Request: &request,
	}
	return
}

func (req *OrderGetRequest) SetOrders(orders []uint64) {
	req.Request.Params["orders"] = orders
}

func (req *OrderGetRequest) GetOrders() []uint64 {
	orders, found := req.Request.Params["orders"]
	if found {
		return orders.([]uint64)
	}
	return nil
}

func (req *OrderGetRequest) SetOptionalFields(optionalFields string) {
	req.Request.Params["optional_fields"] = optionalFields
}

func (req *OrderGetRequest) GetOptionalFields() string {
	optionalFields, found := req.Request.Params["optional_fields"]
	if found {
		return optionalFields.(string)
	}
	return ""
}
