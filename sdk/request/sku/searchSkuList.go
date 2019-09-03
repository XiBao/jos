package sku

import (
	"github.com/XiBao/jos/sdk"
)

type SearchSkuListRequest struct {
	Request *sdk.Request
}

// create new request
func NewSearchSkuListRequest() (req *SearchSkuListRequest) {
	request := sdk.Request{MethodName: "jingdong.sku.read.searchSkuList", Params: make(map[string]interface{}, 7)}
	req = &SearchSkuListRequest{
		Request: &request,
	}
	return
}

func (req *SearchSkuListRequest) SetSkuStatusValue(skuStatuValue []int) {
	req.Request.Params["skuStatuValue"] = skuStatuValue
}

func (req *SearchSkuListRequest) GetSkuStatusValue() []int {
	skuStatuValue, found := req.Request.Params["skuStatuValue"]
	if found {
		return skuStatuValue.([]int)
	}
	return nil
}

func (req *SearchSkuListRequest) SetPageNo(pageNo int) {
	req.Request.Params["pageNo"] = pageNo
}

func (req *SearchSkuListRequest) GetPageNo() int {
	pageNo, found := req.Request.Params["pageNo"]
	if found {
		return pageNo.(int)
	}
	return 0
}

func (req *SearchSkuListRequest) SetPageSize(pageSize int) {
	req.Request.Params["page_size"] = pageSize
}

func (req *SearchSkuListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *SearchSkuListRequest) SetWareId(wareId []uint64) {
	req.Request.Params["wareId"] = wareId
}

func (req *SearchSkuListRequest) GetWareId() []uint64 {
	wareId, found := req.Request.Params["wareId"]
	if found {
		return wareId.([]uint64)
	}
	return nil
}

func (req *SearchSkuListRequest) SetField(field string) {
	req.Request.Params["field"] = field
}

func (req *SearchSkuListRequest) GetField() string {
	field, found := req.Request.Params["field"]
	if found {
		return field.(string)
	}
	return ""
}

func (req *SearchSkuListRequest) SetOrderField(orderField []string) {
	req.Request.Params["orderField"] = orderField
}

func (req *SearchSkuListRequest) GetOrderField() []string {
	orderField, found := req.Request.Params["orderField"]
	if found {
		return orderField.([]string)
	}
	return nil
}

func (req *SearchSkuListRequest) SetOrderType(orderType []string) {
	req.Request.Params["orderType"] = orderType
}

func (req *SearchSkuListRequest) GetOrderType() []string {
	orderType, found := req.Request.Params["orderType"]
	if found {
		return orderType.([]string)
	}
	return nil
}
