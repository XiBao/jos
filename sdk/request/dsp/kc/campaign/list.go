package campaign

import (
	"github.com/XiBao/jos/sdk"
)

type CampainListRequest struct {
	Request *sdk.Request
}

func NewCampainListRequest() (req *CampainListRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.kc.campain.list", Params: make(map[string]interface{}, 2)}
	req = &CampainListRequest{
		Request: &request,
	}
	return
}

func (req *CampainListRequest) SetPageNum(pageNum int) {
	req.Request.Params["pageNum"] = pageNum
}

func (req *CampainListRequest) GetPageNum() int {
	pageNum, found := req.Request.Params["pageNum"]
	if found {
		return pageNum.(int)
	}
	return 0
}

func (req *CampainListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *CampainListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}
