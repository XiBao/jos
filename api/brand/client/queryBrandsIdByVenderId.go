package client

import (
	"encoding/json"
	"errors"

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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result info")
	}
	var response QueryBrandsIdByVenderIdResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}

	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	if response.Data.Result == nil {
		return 0, errors.New(`查询失败`)
	}

	return response.Data.Result.BrandsId, nil

}
