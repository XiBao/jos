package vender

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type ShopQueryRequest struct {
	api.BaseRequest
}

type ShopQueryResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopQuerySubResponse `json:"jingdong_vender_shop_query_responce,omitempty" codec:"jingdong_vender_shop_query_responce,omitempty"`
}

func (r ShopQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ShopQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ShopQuerySubResponse struct {
	Code          string    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	ShopJosResult *ShopInfo `json:"shop_jos_result,omitempty" codec:"shop_jos_result,omitempty"`
}

func (r ShopQuerySubResponse) IsError() bool {
	return r.Code != "0" || r.ShopJosResult == nil
}

func (r ShopQuerySubResponse) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 店铺信息查询
func ShopQuery(req *ShopQueryRequest) (*ShopInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewVenderShopQueryRequest()

	var response ShopQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	return response.Data.ShopJosResult, nil
}
