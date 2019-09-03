package crm

import (
	"github.com/XiBao/jos/sdk"
)

type GetPointsDetailRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetPointsDetailRequest() (req *GetPointsDetailRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.getPointsDetail", Params: make(map[string]interface{}, 8)}
	req = &GetPointsDetailRequest{
		Request: &request,
	}
	return
}

func (req *GetPointsDetailRequest) SetCustomerPin(customerPin string) {
	req.Request.Params["customerPin"] = customerPin
}

func (req *GetPointsDetailRequest) GetPin() string {
	pin, found := req.Request.Params["customerPin"]
	if found {
		return pin.(string)
	}
	return ""
}

func (req *GetPointsDetailRequest) SetStartTime(startTime string) {
	req.Request.Params["startTime"] = startTime
}

func (req *GetPointsDetailRequest) GetStartTime() string {
	startTime, found := req.Request.Params["startTime"]
	if found {
		return startTime.(string)
	}
	return ""
}

func (req *GetPointsDetailRequest) SetEndTime(endTime string) {
	req.Request.Params["endTime"] = endTime
}

func (req *GetPointsDetailRequest) GetEndTime() string {
	endTime, found := req.Request.Params["endTime"]
	if found {
		return endTime.(string)
	}
	return ""
}

func (req *GetPointsDetailRequest) SetStartPage(startPage int) {
	req.Request.Params["startPage"] = startPage
}

func (req *GetPointsDetailRequest) GetStartPage() int {
	startPage, found := req.Request.Params["startPage"]
	if found {
		return startPage.(int)
	}
	return 0
}

func (req *GetPointsDetailRequest) SetEndPage(endPage int) {
	req.Request.Params["endPage"] = endPage
}

func (req *GetPointsDetailRequest) GetEndPage() int {
	endPage, found := req.Request.Params["endPage"]
	if found {
		return endPage.(int)
	}
	return 0
}

func (req *GetPointsDetailRequest) SetPage(page int) {
	req.Request.Params["page"] = page
}

func (req *GetPointsDetailRequest) GetPage() int {
	page, found := req.Request.Params["page"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *GetPointsDetailRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *GetPointsDetailRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *GetPointsDetailRequest) SetStartRowKey(startRowKey string) {
	req.Request.Params["startRowKey"] = startRowKey
}

func (req *GetPointsDetailRequest) GetStartRowKey() string {
	startRowKey, found := req.Request.Params["startRowKey"]
	if found {
		return startRowKey.(string)
	}
	return ""
}
