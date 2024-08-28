package order

import (
	"context"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/order"
)

type OrderGetRequest struct {
	api.BaseRequest
	Orders         []uint64 `json:"orders" codec:"orders"`
	OptionalFields []string `json:"optional_fields" codec:"optional_fields"`
}

type OrderGetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *OrderGetData       `json:"jingdong_shop_order_get_responce,omitempty" codec:"jingdong_shop_order_get_responce,omitempty"`
}

func (r OrderGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r OrderGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type OrderGetData struct {
	Code              string             `json:"code,omitempty" codec:"code,omitempty"`
	OrderDetailResult *OrderDetailResult `json:"orderDetailResult,omitempty" codec:"orderDetailResult,omitempty"`
}

func (r OrderGetData) IsError() bool {
	return r.OrderDetailResult == nil || r.OrderDetailResult.IsError()
}

func (r OrderGetData) Error() string {
	if r.OrderDetailResult != nil {
		return r.OrderDetailResult.Error()
	}
	return "no result data"
}

type OrderDetailResult struct {
	Code    int               `json:"code,omitempty" codec:"code,omitempty"`
	Message string            `json:"message,omitempty" codec:"message,omitempty"`
	ExtMap  map[string]string `json:"extMap,omitempty" codec:"extMap,omitempty"`
	Success bool              `json:"success,omitempty" codec:"success,omitempty"`
	Orders  []JdOrderInfo     `json:"orders,omitempty" codec:"orders,omitempty"`
}

func (r OrderDetailResult) IsError() bool {
	return !r.Success
}

func (r OrderDetailResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

// 输入单个自营订单id，得到所有相关订单信息
func OrderGet(ctx context.Context, req *OrderGetRequest) ([]JdOrderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewOrderGetRequest()
	r.SetOrders(req.Orders)
	r.SetOptionalFields(strings.Join(req.OptionalFields, ","))

	var response OrderGetResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.OrderDetailResult.Orders, nil
}
