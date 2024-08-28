package points

import (
	"github.com/XiBao/jos/sdk"
)

type QueryGiftSkuByExchangeGiftIdRequest struct {
	Request *sdk.Request
}

func NewQueryGiftSkuByExchangeGiftIdRequest() (req *QueryGiftSkuByExchangeGiftIdRequest) {
	request := sdk.Request{MethodName: "jingdong.points.jos.queryGiftSkuByExchangeGiftId", Params: make(map[string]interface{}, 1)}
	req = &QueryGiftSkuByExchangeGiftIdRequest{
		Request: &request,
	}
	return
}

func (req *QueryGiftSkuByExchangeGiftIdRequest) SetGiftId(giftId uint64) {
	req.Request.Params["giftId"] = giftId
}

func (req *QueryGiftSkuByExchangeGiftIdRequest) GetGiftId() uint64 {
	multiple, found := req.Request.Params["giftId"]
	if found {
		return multiple.(uint64)
	}
	return 0
}
