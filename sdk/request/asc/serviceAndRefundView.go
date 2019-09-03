package asc

import "github.com/XiBao/jos/sdk"

type ServiceAndRefundViewRequest struct {
	Request *sdk.Request
}

// create new request
func NewServiceAndRefundViewRequest() (req *ServiceAndRefundViewRequest) {
	request := sdk.Request{MethodName: "jingdong.asc.serviceAndRefund.view", Params: make(map[string]interface{}, 8)}
	req = &ServiceAndRefundViewRequest{
		Request: &request,
	}
	return
}

func (req *ServiceAndRefundViewRequest) SetOrderId(orderId uint64) {
	req.Request.Params["orderId"] = orderId
}

func (req *ServiceAndRefundViewRequest) GetOrderId() uint64 {
	orderId, found := req.Request.Params["customerPin"]
	if found {
		return orderId.(uint64)
	}
	return 0
}

func (req *ServiceAndRefundViewRequest) SetApplyTimeBegin(applyTimeBegin string) {
	req.Request.Params["applyTimeBegin"] = applyTimeBegin
}

func (req *ServiceAndRefundViewRequest) GetApplyTimeBegin() string {
	applyTimeBegin, found := req.Request.Params["applyTimeBegin"]
	if found {
		return applyTimeBegin.(string)
	}
	return ""
}

func (req *ServiceAndRefundViewRequest) SetApplyTimeEnd(applyTimeEnd string) {
	req.Request.Params["applyTimeEnd"] = applyTimeEnd
}

func (req *ServiceAndRefundViewRequest) GetApplyTimeEnd() string {
	applyTimeEnd, found := req.Request.Params["applyTimeEnd"]
	if found {
		return applyTimeEnd.(string)
	}
	return ""
}

func (req *ServiceAndRefundViewRequest) SetApproveTimeBegin(approveTimeBegin string) {
	req.Request.Params["approveTimeBegin"] = approveTimeBegin
}

func (req *ServiceAndRefundViewRequest) GetApproveTimeBegin() string {
	approveTimeBegin, found := req.Request.Params["approveTimeBegin"]
	if found {
		return approveTimeBegin.(string)
	}
	return ""
}

func (req *ServiceAndRefundViewRequest) SetApproveTimeEnd(approveTimeEnd string) {
	req.Request.Params["approveTimeEnd"] = approveTimeEnd
}

func (req *ServiceAndRefundViewRequest) GetApproveTimeEnd() string {
	approveTimeEnd, found := req.Request.Params["approveTimeEnd"]
	if found {
		return approveTimeEnd.(string)
	}
	return ""
}

func (req *ServiceAndRefundViewRequest) SetPageNumber(pageNumber uint64) {
	req.Request.Params["pageNumber"] = pageNumber
}

func (req *ServiceAndRefundViewRequest) GetPageNumber() uint64 {
	pageNumber, found := req.Request.Params["pageNumber"]
	if found {
		return pageNumber.(uint64)
	}
	return 0
}

func (req *ServiceAndRefundViewRequest) SetPageSize(pageSize uint8) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *ServiceAndRefundViewRequest) GetPageSize() uint8 {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(uint8)
	}
	return 0
}

func (req *ServiceAndRefundViewRequest) SetExtJsonStr(extJsonStr string) {
	req.Request.Params["extJsonStr"] = extJsonStr
}

func (req *ServiceAndRefundViewRequest) GetExtJsonStr() string {
	extJsonStr, found := req.Request.Params["extJsonStr"]
	if found {
		return extJsonStr.(string)
	}
	return ""
}
