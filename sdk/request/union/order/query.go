package order

import (
	"github.com/XiBao/jos/sdk"
)

type OrderReq struct {
	PageNo       uint   `json:"pageNo"`
	PageSize     uint   `json:"pageSize"`
	Type         uint   `json:"type"`
	Time         string `json:"time"`
	ChildUnionId string `json:"childUnionId,omitempty"`
	Key          string `json:"key,omitempty"`
}

type UnionOrderQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewUnionOrderQueryRequest() (req *UnionOrderQueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.order.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &UnionOrderQueryRequest{
		Request: &request,
	}
	return
}

func (req *UnionOrderQueryRequest) SetOrderReq(orderReq *OrderReq) {
	req.Request.Params["orderReq"] = orderReq
}

func (req *UnionOrderQueryRequest) GetOrderReq() *OrderReq {
	orderReq, found := req.Request.Params["orderReq"]
	if found {
		return orderReq.(*OrderReq)
	}
	return nil
}
