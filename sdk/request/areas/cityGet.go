package areas

import (
	"github.com/XiBao/jos/sdk"
)

type AreasCityGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewAreasCityGetRequest() (req *AreasCityGetRequest) {
	request := sdk.Request{MethodName: "jingdong.areas.city.get", Params: make(map[string]interface{}, 1)}
	req = &AreasCityGetRequest{
		Request: &request,
	}
	return
}

func (req *AreasCityGetRequest) SetParentId(parentId uint64) {
	req.Request.Params["parent_id"] = parentId
}

func (req *AreasCityGetRequest) GetParentId() uint64 {
	parentId, found := req.Request.Params["parent_id"]
	if found {
		return parentId.(uint64)
	}
	return 0
}
