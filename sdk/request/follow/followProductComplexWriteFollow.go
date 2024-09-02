package follow

import (
	"github.com/XiBao/jos/sdk"
)

type FollowProductComplexWriteFollow struct {
	Request *sdk.Request
}

func NewFollowProductComplexWriteFollow() (req *FollowProductComplexWriteFollow) {
	request := sdk.Request{MethodName: "jingdong.follow.product.complex.write.follow", Params: make(map[string]interface{}, 2)}
	req = &FollowProductComplexWriteFollow{
		Request: &request,
	}
	return
}

func (req *FollowProductComplexWriteFollow) SetPin(pin string) {
	req.Request.Params["pin"] = pin
}

func (req *FollowProductComplexWriteFollow) GetPin() string {
	pin, found := req.Request.Params["pin"]
	if found {
		return pin.(string)
	}
	return ``
}

func (req *FollowProductComplexWriteFollow) SetProductId(productId uint64) {
	req.Request.Params["productId"] = productId
}

func (req *FollowProductComplexWriteFollow) GetProductId() uint64 {
	productId, found := req.Request.Params["productId"]
	if found {
		return productId.(uint64)
	}
	return 0
}
