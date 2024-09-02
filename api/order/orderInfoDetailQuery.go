package order

import (
	"context"

	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/order"
)

type OrderInfoDetailQueryRequest struct {
	BaseRequest

	ActivityId uint64 `json:"activityId,omitempty" codec:"activityId,omitempty"`
	VenderId   uint64 `json:"venderId,omitempty" codec:"venderId,omitempty"`
	IsvSign    string `json:"isvSign,omitempty" codec:"isvSign,omitempty"`
	StartRow   int    `json:"startRow,omitempty" codec:"startRow,omitempty"`
	EndRow     int    `json:"endRow,omitempty" codec:"endRow,omitempty"`
	SearchDate string `json:"searchDate,omitempty" codec:"searchDate,omitempty"`
}

type OrderInfoDetailQueryResponse struct {
	ErrorResp *ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *OrderInfoDetailQueryData `json:"jingdong_orderInfoDetailQuery_responce,omitempty" codec:"jingdong_orderInfoDetailQuery_responce,omitempty"`
}

func (r OrderInfoDetailQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r OrderInfoDetailQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type OrderInfoDetailQueryData struct {
	Code      string                      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *OrderInfoDetailQueryResult `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r OrderInfoDetailQueryData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r OrderInfoDetailQueryData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type OrderInfoDetailQueryResult struct {
	Message string                        `json:"message,omitempty" codec:"message,omitempty"`
	Content []OrderInfoDetailQueryContent `json:"content,omitempty" codec:"content,omitempty"`
}

func (r OrderInfoDetailQueryResult) IsError() bool {
	return r.Message != "SUCCESS"
}

func (r OrderInfoDetailQueryResult) Error() string {
	return r.Message
}

type OrderInfoDetailQueryContent struct {
	SkuString  string  `json:"skuString"`
	Pin        string  `json:"pin"`
	CreateTime string  `json:"create_time"`
	OrderId    string  `json:"orderId"`
	SaleOrdDt  string  `json:"saleOrdDt"`
	DealAmount float64 `json:"dealAmount"`
	VenderId   string  `json:"venderId"`
	RankId     uint    `json:"rendId"`
	TotalRows  uint    `json:"totalRows"`
}

func OrderInfoDetailQuery(ctx context.Context, req *OrderInfoDetailQueryRequest) ([]OrderInfoDetailQueryContent, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewOrderInfoDetailQueryRequest()
	r.SetVenderId(req.VenderId)
	r.SetActivityId(req.ActivityId)
	r.SetStartRow(req.StartRow)
	r.SetEndRow(req.EndRow)

	if req.IsvSign != "" {
		r.SetIsvSign(req.IsvSign)
	}
	if req.SearchDate != "" {
		r.SetSearchDate(req.SearchDate)
	}

	var response OrderInfoDetailQueryResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Content, nil
}
