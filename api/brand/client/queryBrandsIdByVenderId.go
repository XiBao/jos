package client

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	clt "github.com/XiBao/jos/sdk/request/brand/client"
)

type QueryBrandsIdByVenderIdRequest struct {
	api.BaseRequest
}

type QueryBrandsIdByVenderIdResponse struct {
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryBrandsIdByVenderIdData `json:"jingdong_pop_brand_client_queryBrandsIdByVenderId_response,omitempty" codec:"jingdong_pop_brand_client_queryBrandsIdByVenderId_response,omitempty"`
}

func (r QueryBrandsIdByVenderIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == nil
}

func (r QueryBrandsIdByVenderIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type QueryBrandsIdByVenderIdData struct {
	Result *QueryBrandsIdByVenderIdResult `json:"result,omitempty" codec:"result,omitempty"`
}

type QueryBrandsIdByVenderIdResult struct {
	BrandsId uint64 `json:"brandsId,omitempty" codec:"brandsId,omitempty"`
}

// TODO 根据商家id 查询品牌id
func QueryBrandsIdByVenderId(req *QueryBrandsIdByVenderIdRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := clt.NewQueryBrandsIdByVenderIdRequest()

	var response QueryBrandsIdByVenderIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result.BrandsId, nil

}
