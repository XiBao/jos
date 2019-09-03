package follow

import (
	"github.com/XiBao/jos/sdk"
)

type QueryForCountByVidRequest struct {
	Request *sdk.Request
}

func NewQueryForCountByVidRequest() (req *QueryForCountByVidRequest) {
	request := sdk.Request{MethodName: "jingdong.follow.vender.read.queryForCountByVid", Params: make(map[string]interface{}, 1)}
	req = &QueryForCountByVidRequest{
		Request: &request,
	}
	return
}

func (req *QueryForCountByVidRequest) SetShopId(shopId uint64) {
	req.Request.Params["shopId"] = shopId
}

func (req *QueryForCountByVidRequest) GetShopId() uint64 {
	shopId, found := req.Request.Params["shopId"]
	if found {
		return shopId.(uint64)
	}
	return 0
}
