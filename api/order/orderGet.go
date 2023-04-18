package order

import (
	"encoding/json"
	"errors"
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

type OrderGetData struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`

	OrderDetailResult *OrderDetailResult `json:"orderDetailResult,omitempty" codec:"orderDetailResult,omitempty"`
}

type OrderDetailResult struct {
	Code    int               `json:"code,omitempty" codec:"code,omitempty"`
	Message string            `json:"message,omitempty" codec:"message,omitempty"`
	ExtMap  map[string]string `json:"extMap,omitempty" codec:"extMap,omitempty"`
	Success bool              `json:"success,omitempty" codec:"success,omitempty"`
	Orders  []JdOrderInfo     `json:"orders,omitempty" codec:"orders,omitempty"`
}

// 输入单个自营订单id，得到所有相关订单信息
func OrderGet(req *OrderGetRequest) ([]JdOrderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewOrderGetRequest()
	r.SetOrders(req.Orders)
	r.SetOptionalFields(strings.Join(req.OptionalFields, ","))

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No order info.")
	}

	var response OrderGetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if !response.Data.OrderDetailResult.Success {
		return nil, errors.New(response.Data.OrderDetailResult.Message)
	}

	return response.Data.OrderDetailResult.Orders, nil
}
