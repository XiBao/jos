package ware

import (
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

func (r VenderSkusQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r VenderSkusQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data.IsError() {
		return r.Data.Error()
	}
	return "no result data"
}

type VenderSkusQueryData struct {
	Code         string                 `json:"code,omitempty" codec:"code,omitempty"`
	SearchResult *VenderSkusQueryResult `json:"search_result,omitempty" codec:"search_result,omitempty"`
}

func (r VenderSkusQueryData) IsError() bool {
	return r.SearchResult == nil
}

func (r VenderSkusQueryData) Error() string {
	return "no result data"
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

	var response VenderSkusQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, 0, err
	}
	searchResult := response.Data.SearchResult
	return searchResult.SkuList, searchResult.Total, nil
}
