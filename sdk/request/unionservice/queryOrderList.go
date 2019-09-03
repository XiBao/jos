package unionservice

import (
	"github.com/XiBao/jos/sdk"
)

type UnionQueryOrderListRequest struct {
	Request *sdk.Request
}

// create new request
func NewUnionQueryOrderListRequest() (req *UnionQueryOrderListRequest) {
	request := sdk.Request{MethodName: "jingdong.UnionService.queryOrderList", Params: make(map[string]interface{}, 4)}
	req = &UnionQueryOrderListRequest{
		Request: &request,
	}
	return
}

func (req *UnionQueryOrderListRequest) SetUnionId(unionId uint64) {
	req.Request.Params["unionId"] = unionId
}

func (req *UnionQueryOrderListRequest) GetUnionId() uint64 {
	unionId, found := req.Request.Params["unionId"]
	if found {
		return unionId.(uint64)
	}
	return 0
}

func (req *UnionQueryOrderListRequest) SetTime(time string) { // 格式yyyyMMddHH:2018012316 (按数据更新时间查询)
	req.Request.Params["time"] = time
}

func (req *UnionQueryOrderListRequest) GetTime() string {
	time, found := req.Request.Params["time"]
	if found {
		return time.(string)
	}
	return ""
}

func (req *UnionQueryOrderListRequest) SetPage(page int) {
	req.Request.Params["pageIndex"] = page
}

func (req *UnionQueryOrderListRequest) GetPage() int {
	page, found := req.Request.Params["pageIndex"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *UnionQueryOrderListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *UnionQueryOrderListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}
