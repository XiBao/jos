package order

import (
	"fmt"

	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/order"
	"github.com/bububa/ljson"
)

type OrderInfoDetailQueryRequest struct {
	BaseRequest

	ActivityId uint64 `json:"activityId,omitempty" codec:"activityId,omitempty"`
	IsvSign    string `json:"isvSign,omitempty" codec:"isvSign,omitempty"`
	PageNo     int    `json:"pageNo,omitempty" codec:"pageNo,omitempty"`
	RowNo      int    `json:"rowNo,omitempty" codec:"rowNo,omitempty"`
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
	SkuString  string  `json:"sku_string"`
	Pin        string  `json:"pin"`
	CreateTime string  `json:"create_time"`
	OrderId    string  `json:"order_id"`
	DealAmount float64 `json:"deal_amount"`
	VenderId   string  `json:"vender_id"`
}

func OrderInfoDetailQuery(req *OrderInfoDetailQueryRequest) ([]OrderInfoDetailQueryContent, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewOrderInfoDetailQueryRequest()
	r.SetActivityId(req.ActivityId)
	if req.IsvSign != "" {
		r.SetIsvSign(req.IsvSign)
	}
	if req.PageNo > 0 {
		r.SetPageNo(req.PageNo)
	}
	if req.RowNo > 0 {
		r.SetRowNo(req.RowNo)
	}
	if req.SearchDate != "" {
		r.SetSearchDate(req.SearchDate)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = RemoveJsonSpace(result)

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
