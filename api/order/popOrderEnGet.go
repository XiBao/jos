package order

import (
	"encoding/json"
	"errors"
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

type PopOrderEnGetData struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`

	OrderDetailInfo *OrderDetailInfo `json:"orderDetailInfo,omitempty" codec:"orderDetailInfo,omitempty"`
}

type OrderDetailInfo struct {
	OrderInfo *OrderInfo     `json:"orderInfo,omitempty" codec:"orderInfo,omitempty"`
	ApiResult *api.ApiResult `json:"apiResult,omitempty" codec:"apiResult,omitempty"`
}

// 输入单个订单id，得到所有相关订单信息
func PopOrderEnGet(req *PopOrderEnGetRequest) (*OrderInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewPopOrderEnGetRequest()
	r.SetOrderState(strings.Join(req.OrderState, ","))
	r.SetOptionalFields(strings.Join(req.OptionalFields, ","))
	r.SetOrderId(req.OrderId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No order info.")
	}

	var response PopOrderEnGetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if !response.Data.OrderDetailInfo.ApiResult.Success {
		return nil, response.Data.OrderDetailInfo.ApiResult
	}

	return response.Data.OrderDetailInfo.OrderInfo, nil
}
