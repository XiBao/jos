package category

import (
	"github.com/XiBao/jos/sdk"
)

type FindCateByIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewFindCateByIdRequest() (req *FindCateByIdRequest) {
	request := sdk.Request{MethodName: "jingdong.category.read.findById", Params: make(map[string]interface{}, 2)}
	req = &FindCateByIdRequest{
		Request: &request,
	}
	return
}

func (req *FindCateByIdRequest) SetFields(fields string) {
	req.Request.Params["field"] = fields
}

func (req *FindCateByIdRequest) GetFields() string {
	fields, found := req.Request.Params["field"]
	if found {
		return fields.(string)
	}
	return ""
}

func (req *FindCateByIdRequest) SetCid(cid uint64) {
	req.Request.Params["cid"] = cid
}

func (req *FindCateByIdRequest) GetCid() uint64 {
	cid, found := req.Request.Params["cid"]
	if found {
		return cid.(uint64)
	}
	return 0
}
