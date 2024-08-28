package crm

import (
	"github.com/XiBao/jos/sdk"
)

type GetMemeberDiscountRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetMemeberDiscountRequest() (req *GetMemeberDiscountRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.crm.getMemeberDiscount", Params: nil}
	req = &GetMemeberDiscountRequest{
		Request: &request,
	}
	return
}
