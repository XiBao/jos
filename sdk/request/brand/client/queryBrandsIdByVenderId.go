package client

import "github.com/XiBao/jos/sdk"

type QueryBrandsIdByVenderIdRequest struct {
	Request *sdk.Request
}

// create new request
func NewQueryBrandsIdByVenderIdRequest() (req *QueryBrandsIdByVenderIdRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.brand.client.queryBrandsIdByVenderId", Params: nil}
	req = &QueryBrandsIdByVenderIdRequest{
		Request: &request,
	}
	return
}
