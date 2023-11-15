package order

import (
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

func (r PopOrderEnSearchResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r PopOrderEnSearchResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type PopOrderEnSearchData struct {
	Code            string           `json:"code,omitempty" codec:"code,omitempty"`
	SearchOrderInfo *SearchOrderInfo `json:"searchorderinfo_result,omitempty" codec:"searchorderinfo_result,omitempty"`
}

func (r PopOrderEnSearchData) IsError() bool {
	return r.SearchOrderInfo == nil || r.SearchOrderInfo.IsError()
}

func (r PopOrderEnSearchData) Error() string {
	if r.SearchOrderInfo != nil {
		return r.SearchOrderInfo.Error()
	}
	return "no result data"
}

type SearchOrderInfo struct {
	OrderInfoList []OrderInfo    `json:"orderInfoList,omitempty" codec:"orderInfoList,omitempty"`
	OrderTotal    int            `json:"orderTotal,omitempty" codec:"orderTotal,omitempty"`
	ApiResult     *api.ApiResult `json:"apiResult,omitempty" codec:"apiResult,omitempty"`
}

func (r SearchOrderInfo) IsError() bool {
	return r.ApiResult == nil || r.ApiResult.IsError()
}

func (r SearchOrderInfo) Error() string {
	if r.ApiResult != nil {
		return r.ApiResult.Error()
	}
	return "no result data"
}

// 根据条件检索订单信息 （仅适用于SOP、LBP，SOPL类型，FBP类型请调取FBP订单检索 ）
func PopOrderEnSearch(req *PopOrderEnSearchRequest) ([]OrderInfo, int, error) {
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

	var response PopOrderEnSearchResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, 0, err
	}
	return response.Data.SearchOrderInfo.OrderInfoList, response.Data.SearchOrderInfo.OrderTotal, nil
}
