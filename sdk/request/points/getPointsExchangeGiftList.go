package points

import "github.com/XiBao/jos/sdk"

type GetPointsExchangeGiftListRequest struct {
	Request *sdk.Request
}

func NewGetPointsExchangeGiftListRequest() (req *GetPointsExchangeGiftListRequest) {
	request := sdk.Request{MethodName: "jingdong.points.jos.getPointsExchangeGiftList", Params: nil}
	req = &GetPointsExchangeGiftListRequest{
		Request: &request,
	}
	return
}
