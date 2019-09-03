package vender

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
	"github.com/daviddengcn/ljson"
)

type CommonQueryRequest struct {
	api.BaseRequest
	Method    string `json:"method,omitempty" codec:"method,omitempty"`
	InputPara string `json:"input_para,omitempty" codec:"input_para,omitempty"`
}

type CommonQueryResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CommonQuerySubResponse `json:"jingdong_data_vender_common_query_responce,omitempty" codec:"jingdong_data_vender_common_query_responce,omitempty"`
}

type CommonQuerySubResponse struct {
	Code      string                  `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                  `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Response  *CommonQuerySecResponse `json:"response,omitempty" codec:"response,omitempty"`
}

type CommonQuerySecResponse struct {
	Code   int    `json:"code,omitempty" codec:"code,omitempty"`
	Msg    string `json:"msg,omitempty" codec:"msg,omitempty"`
	Result string `json:"result,omitempty" codec:"result,omitempty"`
}

type CommonQueryResult struct {
	Wjunionid        string `json:"wjunionid,omitempty" codec:"wjunionid,omitempty"`
	SaleOrdId        string `json:"sale_ord_id,omitempty" codec:"sale_ord_id,omitempty"`
	ItemId           string `json:"item_id,omitempty" codec:"item_id,omitempty"`
	SkuType          string `json:"sku_type,omitempty" codec:"sku_type,omitempty"`
	OpTime           string `json:"op_time,omitempty" codec:"op_time,omitempty"`
	NewbuyPinFlag    string `json:"newbuy_pin_flag,omitempty" codec:"newbuy_pin_flag,omitempty"`
	VenderId         string `json:"vender_id,omitempty" codec:"vender_id,omitempty"`
	ShopId           string `json:"shop_id,omitempty" codec:"shop_id,omitempty"`
	ServiceType      string `json:"service_type,omitempty" codec:"service_type,omitempty"`
	OpTimeStamp      uint64 `json:"op_time_stamp,omitempty" codec:"op_time_stamp,omitempty"`
	AfterPrefrAmount string `json:"after_prefr_amount,omitempty" codec:"after_prefr_amount,omitempty"`
	ItemSkuId        string `json:"item_sku_id,omitempty" codec:"item_sku_id,omitempty"`
	BucketNum500     string `json:"bucket_num_500,omitempty" codec:"bucket_num_500,omitempty"`
	Appkey           string `json:"appkey,omitempty" codec:"appkey,omitempty"`
	SaleOrdTm        string `json:"sale_ord_tm,omitempty" codec:"sale_ord_tm,omitempty"`
	OrderState       string `json:"order_state,omitempty" codec:"order_state,omitempty"`
	ParentSaleOrdId  string `json:"parent_sale_ord_id,omitempty" codec:"parent_sale_ord_id,omitempty"`
}

//通过组件化的方式，提供相关统一的查询方式
func CommonQuery(req *CommonQueryRequest) ([]*CommonQueryResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewVenderCommonQueryRequest()
	if req.Method != "" {
		r.SetMethod(req.Method)
	}
	if req.InputPara != "" {
		r.SetInputPara(req.InputPara)
	}
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No query info.")
	}

	var response CommonQueryResponse
	err = ljson.Unmarshal([]byte(result), &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Response.Code != 0 {
		return nil, errors.New(response.Data.Response.Msg)
	}

	var res []*CommonQueryResult
	err = ljson.Unmarshal([]byte(response.Data.Response.Result), &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
