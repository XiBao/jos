package crm

import (
	"github.com/XiBao/jos/sdk"
)

type GetGradesRequest struct {
	Request *sdk.Request
}

// create new request
func NewGetGradesRequest() (req *GetGradesRequest) {
	request := sdk.Request{MethodName: "jingdong.crm.grade.get", Params: nil}
	req = &GetGradesRequest{
		Request: &request,
	}
	return
}
