package areas

import (
	"github.com/XiBao/jos/sdk"
)

type AreasCountyGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewAreasCountyGetRequest() (req *AreasCountyGetRequest) {
	request := sdk.Request{MethodName: "jingdong.areas.county.get", Params: make(map[string]interface{}, 1)}
	req = &AreasCountyGetRequest{
		Request: &request,
	}
	return
}

func (req *AreasCountyGetRequest) SetParentId(parentId uint64) {
	req.Request.Params["parent_id"] = parentId
}

func (req *AreasCountyGetRequest) GetParentId() uint64 {
	parentId, found := req.Request.Params["parent_id"]
	if found {
		return parentId.(uint64)
	}
	return 0
}
