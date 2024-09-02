package report

import (
	"github.com/XiBao/jos/sdk"
)

type AdgroupkeywordQueryRequest struct {
	Request *sdk.Request
}

func NewAdgroupkeywordQueryRequest() (req *AdgroupkeywordQueryRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.report.adgroupkeyword.query", Params: make(map[string]interface{}, 11)}
	req = &AdgroupkeywordQueryRequest{
		Request: &request,
	}
	return
}

func (req *AdgroupkeywordQueryRequest) SetStartDate(startDate string) {
	req.Request.Params["startDate"] = startDate
}

func (req *AdgroupkeywordQueryRequest) GetStartDate() string {
	startDate, found := req.Request.Params["startDate"]
	if found {
		return startDate.(string)
	}
	return ""
}

func (req *AdgroupkeywordQueryRequest) SetEndDate(endDate string) {
	req.Request.Params["endDate"] = endDate
}

func (req *AdgroupkeywordQueryRequest) GetEndDate() string {
	endDate, found := req.Request.Params["endDate"]
	if found {
		return endDate.(string)
	}
	return ""
}

func (req *AdgroupkeywordQueryRequest) SetGroupId(groupId uint64) {
	req.Request.Params["groupId"] = groupId
}

func (req *AdgroupkeywordQueryRequest) GetGroupId() uint64 {
	groupId, found := req.Request.Params["groupId"]
	if found {
		return groupId.(uint64)
	}
	return 0
}

func (req *AdgroupkeywordQueryRequest) SetPlatform(platform string) {
	req.Request.Params["platform"] = platform
}

func (req *AdgroupkeywordQueryRequest) GetPlatform() string {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(string)
	}
	return ""
}

func (req *AdgroupkeywordQueryRequest) SetValType(valType string) {
	req.Request.Params["valType"] = valType
}

func (req *AdgroupkeywordQueryRequest) GetValType() string {
	valType, found := req.Request.Params["valType"]
	if found {
		return valType.(string)
	}
	return ""
}

func (req *AdgroupkeywordQueryRequest) SetVal(val string) {
	req.Request.Params["val"] = val
}

func (req *AdgroupkeywordQueryRequest) GetVal() string {
	val, found := req.Request.Params["val"]
	if found {
		return val.(string)
	}
	return ""
}

func (req *AdgroupkeywordQueryRequest) SetClickOrOrderDay(clickOrOrderDay uint8) {
	req.Request.Params["clickOrOrderDay"] = clickOrOrderDay
}

func (req *AdgroupkeywordQueryRequest) GetClickOrOrderDay() uint8 {
	clickOrOrderDay, found := req.Request.Params["clickOrOrderDay"]
	if found {
		return clickOrOrderDay.(uint8)
	}
	return 0
}

func (req *AdgroupkeywordQueryRequest) SetIsOrderOrClick(isOrderOrClick bool) {
	req.Request.Params["isOrderOrClick"] = isOrderOrClick
}

func (req *AdgroupkeywordQueryRequest) GetIsOrderOrClick() bool {
	isOrderOrClick, found := req.Request.Params["isOrderOrClick"]
	if found {
		return isOrderOrClick.(bool)
	}
	return false
}

func (req *AdgroupkeywordQueryRequest) SetOrderStatusCategory(orderStatusCategory int) {
	req.Request.Params["orderStatusCategory"] = orderStatusCategory
}

func (req *AdgroupkeywordQueryRequest) GetOrderStatusCategory() int {
	orderStatusCategory, found := req.Request.Params["orderStatusCategory"]
	if found {
		return orderStatusCategory.(int)
	}
	return 0
}

func (req *AdgroupkeywordQueryRequest) SetPageIndex(pageIndex int) {
	req.Request.Params["pageIndex"] = pageIndex
}

func (req *AdgroupkeywordQueryRequest) GetPageIndex() int {
	pageIndex, found := req.Request.Params["pageIndex"]
	if found {
		return pageIndex.(int)
	}
	return 0
}

func (req *AdgroupkeywordQueryRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *AdgroupkeywordQueryRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}
