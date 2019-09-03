package adkcunit

import "github.com/XiBao/jos/sdk"

type AdkcunitAdgroupListRequest struct {
	Request *sdk.Request
}

func NewAdkcunitAdgroupListRequest() (req *AdkcunitAdgroupListRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkcunit.adgroup.list", Params: make(map[string]interface{}, 3)}
	req = &AdkcunitAdgroupListRequest{
		Request: &request,
	}
	return
}

func (req *AdkcunitAdgroupListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *AdkcunitAdgroupListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *AdkcunitAdgroupListRequest) SetPageNum(pageNum int) {
	req.Request.Params["pageNum"] = pageNum
}

func (req *AdkcunitAdgroupListRequest) GetPageNum() int {
	pageNum, found := req.Request.Params["pageNum"]
	if found {
		return pageNum.(int)
	}
	return 0
}

func (req *AdkcunitAdgroupListRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaignId"] = campaignId
}

func (req *AdkcunitAdgroupListRequest) GetCampaignId() uint64 {
	campaignId, found := req.Request.Params["campaignId"]
	if found {
		return campaignId.(uint64)
	}
	return 0
}
