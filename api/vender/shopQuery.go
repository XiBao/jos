package vender

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
	"github.com/daviddengcn/ljson"
)

type ShopQueryRequest struct {
	api.BaseRequest
}

type ShopQueryResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopQuerySubResponse `json:"jingdong_vender_shop_query_responce,omitempty" codec:"jingdong_vender_shop_query_responce,omitempty"`
}

type ShopQuerySubResponse struct {
	Code          string    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	ShopJosResult *ShopInfo `json:"shop_jos_result,omitempty" codec:"shop_jos_result,omitempty"`
}

// 店铺信息查询
func ShopQuery(req *ShopQueryRequest) (*ShopInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewVenderShopQueryRequest()
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}
	var response ShopQueryResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.ShopJosResult, nil
}
