package ware

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type VenderSkusQueryRequest struct {
	api.BaseRequest
	Index int `json:"index,omitempty" codec:"index,omitempty"` //
}

type VenderSkusQueryResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *VenderSkusQueryData `json:"jingdong_new_ware_vender_skus_query_responce,omitempty" codec:"jingdong_new_ware_vender_skus_query_responce,omitempty"`
}

type VenderSkusQueryData struct {
	Code         string                 `json:"code,omitempty" codec:"code,omitempty"`
	SearchResult *VenderSkusQueryResult `json:"search_result,omitempty" codec:"search_result,omitempty"`
}

type VenderSkusQueryResult struct {
	Code    int      `json:"code,omitempty" codec:"code,omitempty"`
	Total   int      `json:"total,omitempty" codec:"total,omitempty"`
	SkuList []uint64 `json:"skuList,omitempty" codec:"skuList,omitempty"`
}

// 获取单个SKU
func VenderSkusQuery(req *VenderSkusQueryRequest) ([]uint64, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewWareVenderSkusQueryRequest()
	r.SetIndex(req.Index)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, 0, err
	}
	if len(result) == 0 {
		return nil, 0, errors.New("No result.")
	}
	var response VenderSkusQueryResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, 0, err
	}

	if response.ErrorResp != nil {
		return nil, 0, response.ErrorResp
	}

	searchResult := response.Data.SearchResult
	return searchResult.SkuList, searchResult.Total, nil
}
