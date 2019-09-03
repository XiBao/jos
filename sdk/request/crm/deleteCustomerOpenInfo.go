package crm

import (
	"github.com/XiBao/jos/sdk"
)

type DeleteCustomerOpenInfoRequest struct {
	Request *sdk.Request
}

// create new request
func NewDeleteCustomerOpenInfoRequest() (req *DeleteCustomerOpenInfoRequest) {
	request := sdk.Request{MethodName: "jingdong.crm.deleteCustomerOpenInfo", Params: nil}
	req = &DeleteCustomerOpenInfoRequest{
		Request: &request,
	}
	return
}
