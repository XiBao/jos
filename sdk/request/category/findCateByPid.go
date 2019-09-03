package category

import (
	"github.com/XiBao/jos/sdk"
)

type FindCateByPidRequest struct {
	Request *sdk.Request
}

// create new request
func NewFindCateByPidRequest() (req *FindCateByPidRequest) {
	request := sdk.Request{MethodName: "jingdong.category.read.findByPId", Params: make(map[string]interface{}, 2)}
	req = &FindCateByPidRequest{
		Request: &request,
	}
	return
}

func (req *FindCateByPidRequest) SetFields(fields []string) {
	req.Request.Params["field"] = fields
}

func (req *FindCateByPidRequest) GetFields() []string {
	fields, found := req.Request.Params["field"]
	if found {
		return fields.([]string)
	}
	return nil
}

func (req *FindCateByPidRequest) SetParentCid(parentCid uint64) {
	req.Request.Params["parentCid"] = parentCid
}

func (req *FindCateByPidRequest) GetParentCid() uint64 {
	parentCid, found := req.Request.Params["parentCid"]
	if found {
		return parentCid.(uint64)
	}
	return 0
}
