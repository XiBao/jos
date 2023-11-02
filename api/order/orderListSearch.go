package order

import (
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

func (r OrderListSearchResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r OrderListSearchResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type OrderListSearchData struct {
	Code       string      `json:"code,omitempty" codec:"code,omitempty"`
	ReturnType *ReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r OrderListSearchData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r OrderListSearchData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
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

func (r ReturnType) IsError() bool {
	return !r.Success
}

func (r ReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
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

	var response OrderListSearchResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, 0, err
	}

	return response.Data.ReturnType.Orders, response.Data.ReturnType.RecordCount, nil
}
