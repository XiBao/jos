package ad

import (
	"github.com/XiBao/jos/sdk"
)

type AdQueryAdListByParamRequest struct {
	Request *sdk.Request
}

func NewAdQueryAdListByParamRequest() (req *AdQueryAdListByParamRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.ad.queryAdListByParam", Params: make(map[string]interface{}, 3)}
	req = &AdQueryAdListByParamRequest{
		Request: &request,
	}
	return
}

func (req *AdQueryAdListByParamRequest) SetAdGroupId(adGroupId uint64) {
	req.Request.Params["adGroupId"] = adGroupId
}

func (req *AdQueryAdListByParamRequest) GetAdGroupId() uint64 {
	adGroupId, found := req.Request.Params["adGroupId"]
	if found {
		return adGroupId.(uint64)
	}
	return 0
}

func (req *AdQueryAdListByParamRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *AdQueryAdListByParamRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *AdQueryAdListByParamRequest) SetPageNum(pageNum int) {
	req.Request.Params["pageNum"] = pageNum
}

func (req *AdQueryAdListByParamRequest) GetPageNum() int {
	pageNum, found := req.Request.Params["pageNum"]
	if found {
		return pageNum.(int)
	}
	return 0
}
