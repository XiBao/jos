package order

import (
	"github.com/XiBao/jos/sdk"
)

type OrderInfoDetailQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewOrderInfoDetailQueryRequest() (req *OrderInfoDetailQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.orderInfoDetailQuery", Params: make(map[string]interface{}, 5)}
	req = &OrderInfoDetailQueryRequest{
		Request: &request,
	}
	return
}

func (req *OrderInfoDetailQueryRequest) SetVenderId(venderId uint64) {
	req.Request.Params["venderId"] = venderId
}

func (req *OrderInfoDetailQueryRequest) GetVenderId() uint64 {
	venderId, found := req.Request.Params["venderId"]
	if found {
		return venderId.(uint64)
	}
	return 0
}

func (req *OrderInfoDetailQueryRequest) SetActivityId(activityId uint64) {
	req.Request.Params["activityId"] = activityId
}

func (req *OrderInfoDetailQueryRequest) GetActivityId() uint64 {
	activityId, found := req.Request.Params["activityId"]
	if found {
		return activityId.(uint64)
	}
	return 0
}

func (req *OrderInfoDetailQueryRequest) SetIsvSign(isvSign string) {
	req.Request.Params["isvSign"] = isvSign
}

func (req *OrderInfoDetailQueryRequest) GetIsvSign() string {
	isvSign, found := req.Request.Params["isvSign"]
	if found {
		return isvSign.(string)
	}
	return ""
}

func (req *OrderInfoDetailQueryRequest) SetStartRow(startRow int) {
	req.Request.Params["startRow"] = startRow
}

func (req *OrderInfoDetailQueryRequest) GetStartRow() int {
	startRow, found := req.Request.Params["startRow"]
	if found {
		return startRow.(int)
	}
	return 0
}

func (req *OrderInfoDetailQueryRequest) SetEndRow(endRow int) {
	req.Request.Params["endRow"] = endRow
}

func (req *OrderInfoDetailQueryRequest) GetEndRow() int {
	endRow, found := req.Request.Params["endRow"]
	if found {
		return endRow.(int)
	}
	return 0
}

func (req *OrderInfoDetailQueryRequest) SetSearchDate(searchDate string) {
	req.Request.Params["searchDate"] = searchDate
}

func (req *OrderInfoDetailQueryRequest) GetSearchDate() string {
	searchDate, found := req.Request.Params["searchDate"]
	if found {
		return searchDate.(string)
	}
	return ""
}
