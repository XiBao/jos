package ware

import (
	"github.com/XiBao/jos/sdk"
)

type SearchWare4ValidRequest struct {
	Request *sdk.Request
}

// create new request
func NewSearchWare4ValidRequest() (req *SearchWare4ValidRequest) {
	request := sdk.Request{MethodName: "jingdong.ware.read.searchWare4Valid", Params: make(map[string]interface{}, 9)}
	req = &SearchWare4ValidRequest{
		Request: &request,
	}
	return
}

func (req *SearchWare4ValidRequest) SetWareStatusValue(wareStatusValue []int) {
	req.Request.Params["wareStatusValue"] = wareStatusValue
}

func (req *SearchWare4ValidRequest) GetWareStatusValue() []int {
	wareStatusValue, found := req.Request.Params["wareStatusValue"]
	if found {
		return wareStatusValue.([]int)
	}
	return nil
}

func (req *SearchWare4ValidRequest) SetPageNo(pageNo int) {
	req.Request.Params["pageNo"] = pageNo
}

func (req *SearchWare4ValidRequest) GetPageNo() int {
	pageNo, found := req.Request.Params["pageNo"]
	if found {
		return pageNo.(int)
	}
	return 0
}

func (req *SearchWare4ValidRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *SearchWare4ValidRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *SearchWare4ValidRequest) SetWareId(wareId []uint64) {
	req.Request.Params["wareId"] = wareId
}

func (req *SearchWare4ValidRequest) GetWareId() []uint64 {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.([]uint64)
	}
	return nil
}

func (req *SearchWare4ValidRequest) SetField(field string) {
	req.Request.Params["field"] = field
}

func (req *SearchWare4ValidRequest) GetField() string {
	field, found := req.Request.Params["field"]
	if found {
		return field.(string)
	}
	return ""
}

func (req *SearchWare4ValidRequest) SetSearchKey(searchKey string) {
	req.Request.Params["searchKey"] = searchKey
}

func (req *SearchWare4ValidRequest) GetSearchKey() string {
	searchKey, found := req.Request.Params["searchKey"]
	if found {
		return searchKey.(string)
	}
	return ""
}

func (req *SearchWare4ValidRequest) SetSearchField(searchField []string) {
	req.Request.Params["searchField"] = searchField
}

func (req *SearchWare4ValidRequest) GetSearchField() []string {
	searchField, found := req.Request.Params["searchField"]
	if found {
		return searchField.([]string)
	}
	return nil
}

func (req *SearchWare4ValidRequest) SetOrderField(orderField []string) {
	req.Request.Params["orderField"] = orderField
}

func (req *SearchWare4ValidRequest) GetOrderField() []string {
	orderField, found := req.Request.Params["orderField"]
	if found {
		return orderField.([]string)
	}
	return nil
}

func (req *SearchWare4ValidRequest) SetOrderType(orderType []string) {
	req.Request.Params["orderType"] = orderType
}

func (req *SearchWare4ValidRequest) GetOrderType() []string {
	orderType, found := req.Request.Params["orderType"]
	if found {
		return orderType.([]string)
	}
	return nil
}
