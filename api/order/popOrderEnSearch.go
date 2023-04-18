package order

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/order"
)

type PopOrderEnSearchRequest struct {
	api.BaseRequest

	StartDate      string   `json:"start_date,omitempty" codec:"start_date,omitempty"`
	EndDate        string   `json:"end_date,omitempty" codec:"end_date,omitempty"`
	OrderState     []string `json:"order_state,omitempty" codec:"order_state,omitempty"`
	OptionalFields []string `json:"optional_fields,omitempty" codec:"optional_fields,omitempty"`
	Page           int      `json:"page,omitempty" codec:"page,omitempty"`
	PageSize       int      `json:"page_size,omitempty" codec:"page_size,omitempty"`
	SortType       uint8    `json:"sort_type,omitempty" codec:"sort_type,omitempty"`
	DateType       uint8    `json:"date_type,omitempty" codec:"date_type,omitempty"`
}

type PopOrderEnSearchResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *PopOrderEnSearchData `json:"jingdong_pop_order_enSearch_responce,omitempty" codec:"jingdong_pop_order_enSearch_responce,omitempty"`
}

type PopOrderEnSearchData struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`

	SearchOrderInfo *SearchOrderInfo `json:"searchorderinfo_result,omitempty" codec:"searchorderinfo_result,omitempty"`
}

type SearchOrderInfo struct {
	OrderInfoList []*OrderInfo   `json:"orderInfoList,omitempty" codec:"orderInfoList,omitempty"`
	OrderTotal    int            `json:"orderTotal,omitempty" codec:"orderTotal,omitempty"`
	ApiResult     *api.ApiResult `json:"apiResult,omitempty" codec:"apiResult,omitempty"`
}

// 根据条件检索订单信息 （仅适用于SOP、LBP，SOPL类型，FBP类型请调取FBP订单检索 ）
func PopOrderEnSearch(req *PopOrderEnSearchRequest) ([]*OrderInfo, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewPopOrderEnSearchRequest()
	if req.StartDate != "" {
		r.SetStartDate(req.StartDate)
	}
	if req.EndDate != "" {
		r.SetEndDate(req.EndDate)
	}
	r.SetOrderState(strings.Join(req.OrderState, ","))
	r.SetOptionalFields(strings.Join(req.OptionalFields, ","))
	r.SetPage(strconv.Itoa(req.Page))
	r.SetPageSize(strconv.Itoa(req.PageSize))
	r.SetSortType(int(req.SortType))
	r.SetDateType(int(req.DateType))

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, 0, err
	}
	var response PopOrderEnSearchResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, 0, err
	}
	if response.ErrorResp != nil {
		return nil, 0, response.ErrorResp
	}
	if !response.Data.SearchOrderInfo.ApiResult.Success {
		return nil, 0, response.Data.SearchOrderInfo.ApiResult
	}

	return response.Data.SearchOrderInfo.OrderInfoList, response.Data.SearchOrderInfo.OrderTotal, nil
}
