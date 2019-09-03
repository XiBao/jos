package order

import (
	"github.com/XiBao/jos/sdk"
)

type PopOrderEnGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewPopOrderEnGetRequest() (req *PopOrderEnGetRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.order.enGet", Params: make(map[string]interface{}, 3)}
	req = &PopOrderEnGetRequest{
		Request: &request,
	}
	return
}

func (req *PopOrderEnGetRequest) SetOrderState(orderState string) {
	req.Request.Params["order_state"] = orderState
}

func (req *PopOrderEnGetRequest) GetOrderState() string {
	orderState, found := req.Request.Params["order_state"]
	if found {
		return orderState.(string)
	}
	return ""
}

func (req *PopOrderEnGetRequest) SetOptionalFields(optionalFields string) {
	req.Request.Params["optional_fields"] = optionalFields
}

func (req *PopOrderEnGetRequest) GetOptionalFields() string {
	optionalFields, found := req.Request.Params["optional_fields"]
	if found {
		return optionalFields.(string)
	}
	return ""
}

func (req *PopOrderEnGetRequest) SetOrderId(orderId uint64) {
	req.Request.Params["order_id"] = orderId
}

func (req *PopOrderEnSearchRequest) GetOrderId() uint64 {
	orderId, found := req.Request.Params["order_id"]
	if found {
		return orderId.(uint64)
	}
	return 0
}
