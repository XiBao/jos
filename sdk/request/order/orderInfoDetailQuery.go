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

func (req *OrderInfoDetailQueryRequest) SetPageNo(pageNo int) {
	req.Request.Params["pageNo"] = pageNo
}

func (req *OrderInfoDetailQueryRequest) GetPageNo() int {
	pageNo, found := req.Request.Params["pageNo"]
	if found {
		return pageNo.(int)
	}
	return 0
}

func (req *OrderInfoDetailQueryRequest) SetRowNo(rowNo int) {
	req.Request.Params["rowNo"] = rowNo
}

func (req *OrderInfoDetailQueryRequest) GetRowNo() int {
	rowNo, found := req.Request.Params["rowNo"]
	if found {
		return rowNo.(int)
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
