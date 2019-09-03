package order

import (
	"github.com/XiBao/jos/sdk"
)

type PopOrderEnSearchRequest struct {
	Request *sdk.Request
}

// create new request
func NewPopOrderEnSearchRequest() (req *PopOrderEnSearchRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.order.enSearch", Params: make(map[string]interface{}, 8)}
	req = &PopOrderEnSearchRequest{
		Request: &request,
	}
	return
}

func (req *PopOrderEnSearchRequest) SetStartDate(startDate string) {
	req.Request.Params["start_date"] = startDate
}

func (req *PopOrderEnSearchRequest) GetStartDate() string {
	startDate, found := req.Request.Params["start_date"]
	if found {
		return startDate.(string)
	}
	return ""
}

func (req *PopOrderEnSearchRequest) SetEndDate(endDate string) {
	req.Request.Params["end_date"] = endDate
}

func (req *PopOrderEnSearchRequest) GetEndDate() string {
	endDate, found := req.Request.Params["end_date"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *PopOrderEnSearchRequest) SetOrderState(orderState string) {
	req.Request.Params["order_state"] = orderState
}

func (req *PopOrderEnSearchRequest) GetOrderState() string {
	orderState, found := req.Request.Params["order_state"]
	if found {
		return orderState.(string)
	}
	return ""
}

func (req *PopOrderEnSearchRequest) SetOptionalFields(optionalFields string) {
	req.Request.Params["optional_fields"] = optionalFields
}

func (req *PopOrderEnSearchRequest) GetOptionalFields() string {
	optionalFields, found := req.Request.Params["optional_fields"]
	if found {
		return optionalFields.(string)
	}
	return ""
}

func (req *PopOrderEnSearchRequest) SetPage(page string) {
	req.Request.Params["page"] = page
}

func (req *PopOrderEnSearchRequest) GetPage() string {
	page, found := req.Request.Params["page"]
	if found {
		return page.(string)
	}
	return ""
}

func (req *PopOrderEnSearchRequest) SetPageSize(pageSize string) {
	req.Request.Params["page_size"] = pageSize
}

func (req *PopOrderEnSearchRequest) GetPageSize() string {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(string)
	}
	return ""
}

func (req *PopOrderEnSearchRequest) SetSortType(sortType int) {
	req.Request.Params["sortType"] = sortType
}

func (req *PopOrderEnSearchRequest) GetSortType() int {
	sortType, found := req.Request.Params["sortType"]
	if found {
		return sortType.(int)
	}
	return 0
}

func (req *PopOrderEnSearchRequest) SetDateType(dateType int) {
	req.Request.Params["dateType"] = dateType
}

func (req *PopOrderEnSearchRequest) GetDateType() int {
	dateType, found := req.Request.Params["dateType"]
	if found {
		return dateType.(int)
	}
	return 0
}
