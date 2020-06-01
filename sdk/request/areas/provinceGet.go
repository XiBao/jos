package areas

import (
	"github.com/XiBao/jos/sdk"
)

type AreasProvinceGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewAreasProvinceGetRequest() (req *AreasProvinceGetRequest) {
	request := sdk.Request{MethodName: "jingdong.areas.province.get", Params: nil}
	req = &AreasProvinceGetRequest{
		Request: &request,
	}
	return
}
