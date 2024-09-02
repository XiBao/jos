package order

import (
	"context"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/order"
)

type PopOrderEnGetRequest struct {
	api.BaseRequest
	OrderState     []string `json:"order_state,omitempty" codec:"order_state,omitempty"`
	OptionalFields []string `json:"optional_fields,omitempty" codec:"optional_fields,omitempty"`
	OrderId        uint64   `json:"order_id,omitempty" codec:"order_id,omitempty"`
}

type PopOrderEnGetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *PopOrderEnGetData  `json:"jingdong_pop_order_enGet_responce,omitempty" codec:"jingdong_pop_order_enGet_responce,omitempty"`
}

func (r PopOrderEnGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r PopOrderEnGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type PopOrderEnGetData struct {
	Code            string           `json:"code,omitempty" codec:"code,omitempty"`
	OrderDetailInfo *OrderDetailInfo `json:"orderDetailInfo,omitempty" codec:"orderDetailInfo,omitempty"`
}

func (r PopOrderEnGetData) IsError() bool {
	return r.OrderDetailInfo == nil || r.OrderDetailInfo.IsError()
}

func (r PopOrderEnGetData) Error() string {
	if r.OrderDetailInfo != nil {
		return r.OrderDetailInfo.Error()
	}
	return "no result data"
}

type OrderDetailInfo struct {
	OrderInfo *OrderInfo     `json:"orderInfo,omitempty" codec:"orderInfo,omitempty"`
	ApiResult *api.ApiResult `json:"apiResult,omitempty" codec:"apiResult,omitempty"`
}

func (r OrderDetailInfo) IsError() bool {
	return r.ApiResult == nil || r.ApiResult.IsError()
}

func (r OrderDetailInfo) Error() string {
	if r.ApiResult != nil {
		return r.ApiResult.Error()
	}
	return "no result data"
}

// 输入单个订单id，得到所有相关订单信息
func PopOrderEnGet(ctx context.Context, req *PopOrderEnGetRequest) (*OrderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewPopOrderEnGetRequest()
	r.SetOrderState(strings.Join(req.OrderState, ","))
	r.SetOptionalFields(strings.Join(req.OptionalFields, ","))
	r.SetOrderId(req.OrderId)

	var response PopOrderEnGetResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.OrderDetailInfo.OrderInfo, nil
}
