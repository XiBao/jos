package vender

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
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

func (r CommonQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CommonQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CommonQuerySubResponse struct {
	Code      string                  `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                  `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Response  *CommonQuerySecResponse `json:"response,omitempty" codec:"response,omitempty"`
}

func (r CommonQuerySubResponse) IsError() bool {
	return r.Code != "" || r.Response == nil || r.Response.IsError()
}

func (r CommonQuerySubResponse) Error() string {
	if r.Response != nil && r.Response.IsError() {
		return r.Response.Error()
	}
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type CommonQuerySecResponse struct {
	Code   int    `json:"code,omitempty" codec:"code,omitempty"`
	Msg    string `json:"msg,omitempty" codec:"msg,omitempty"`
	Result string `json:"result,omitempty" codec:"result,omitempty"`
}

func (r CommonQuerySecResponse) IsError() bool {
	return r.Code != 0 || r.Result == ""
}

func (r CommonQuerySecResponse) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
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

// 通过组件化的方式，提供相关统一的查询方式
func CommonQuery(req *CommonQueryRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewVenderCommonQueryRequest()
	if req.Method != "" {
		r.SetMethod(req.Method)
	}
	if req.InputPara != "" {
		r.SetInputPara(req.InputPara)
	}

	var response CommonQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Data.Response.Result, nil
}
