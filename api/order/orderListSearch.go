package order

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/order"
)

type OrderListSearchRequest struct {
	api.BaseRequest
	StartDate  string `json:"startDate" codec:"startDate"`   // 开始时间 [时间间隔：最大一个月，不能跨年]
	EndDate    string `json:"endDate" codec:"endDate"`       // 结束时间
	OrderState string `json:"orderState" codec:"orderState"` // ALL 全状态 ；NOT_PAY 等待付款 ；SUSPEND 暂停； WAIT_DELIVERY 等待出库；WAIT_SHIPMENTS 等待发货；FINISH 完成 ；CANCEL 取消；LOCK 锁定;
	Page       int    `json:"page" codec:"page"`             // 页数 最大：50
	Size       int    `json:"size" codec:"size"`             // 每页数量 每页最大：100
}

type OrderListSearchResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *OrderListSearchData `json:"jingdong_shop_order_list_search_responce,omitempty" codec:"jingdong_shop_order_list_search_responce,omitempty"`
}

type OrderListSearchData struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`

	ReturnType *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type ReturnType struct {
	Code        int               `json:"code,omitempty" codec:"code,omitempty"`
	Message     string            `json:"message,omitempty" codec:"message,omitempty"`
	ExtMap      map[string]string `json:"extMap,omitempty" codec:"extMap,omitempty"`
	Success     bool              `json:"success,omitempty" codec:"success,omitempty"`
	Orders      []uint64          `json:"orders,omitempty" codec:"orders,omitempty"`
	CurPage     int               `json:"curPage,omitempty" codec:"curPage,omitempty"`
	RecordCount int               `json:"recordCount,omitempty" codec:"recordCount,omitempty"`
	PageSize    int               `json:"pageSize,omitempty" codec:"pageSize,omitempty"`
}

// 自营订单列表查询
func OrderListSearch(req *OrderListSearchRequest) ([]uint64, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewOrderListSearchRequest()
	r.SetStartDate(req.StartDate)
	r.SetEndDate(req.EndDate)
	r.SetOrderState(req.OrderState)
	r.SetPage(req.Page)
	r.SetSize(req.Size)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, 0, err
	}
	if len(result) == 0 {
		return nil, 0, errors.New("No order info.")
	}

	var response OrderListSearchResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, 0, err
	}
	if response.ErrorResp != nil {
		return nil, 0, response.ErrorResp
	}
	if !response.Data.ReturnType.Success {
		return nil, 0, errors.New(response.Data.ReturnType.Message)
	}

	return response.Data.ReturnType.Orders, response.Data.ReturnType.RecordCount, nil
}
