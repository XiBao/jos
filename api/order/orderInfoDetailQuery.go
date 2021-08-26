package order

import (
	"fmt"

	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/order"
	"github.com/bububa/ljson"

	"github.com/XiBao/jos/api/util"
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

type OrderInfoDetailQueryData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	Result *OrderInfoDetailQueryResult `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type OrderInfoDetailQueryResult struct {
	Message string                        `json:"message,omitempty" codec:"message,omitempty"`
	Content []OrderInfoDetailQueryContent `json:"content,omitempty" codec:"content,omitempty"`
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

func OrderInfoDetailQuery(req *OrderInfoDetailQueryRequest) ([]OrderInfoDetailQueryContent, error) {
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response OrderInfoDetailQueryResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, fmt.Errorf("%v", response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, fmt.Errorf("No result.")
	}
	if response.Data.Result.Message != "SUCCESS" {
		return nil, fmt.Errorf(response.Data.Result.Message)
	}

	return response.Data.Result.Content, nil
}
