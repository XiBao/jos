package item

import (
	"github.com/XiBao/jos/sdk"
)

type FindJoinActivitiesRequest struct {
	Request *sdk.Request
}

// create new request
func NewFindJoinActivitiesRequest() (req *FindJoinActivitiesRequest) {
	request := sdk.Request{MethodName: "jd.kpl.open.item.findjoinactives", Params: make(map[string]interface{}, 2)}
	req = &FindJoinActivitiesRequest{
		Request: &request,
	}
	return
}

func (req *FindJoinActivitiesRequest) SetUid(uid string) {
	req.Request.Params["uid"] = uid
}

func (req *FindJoinActivitiesRequest) GetUid() string {
	uid, found := req.Request.Params["uid"]
	if found {
		return uid.(string)
	}
	return ""
}

func (req *FindJoinActivitiesRequest) SetSku(sku uint64) {
	req.Request.Params["sku"] = sku
}

func (req *FindJoinActivitiesRequest) GetSku() uint64 {
	sku, found := req.Request.Params["sku"]
	if found {
		return sku.(uint64)
	}
	return 0
}
