package fw

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fw"
	"github.com/daviddengcn/ljson"
)

type OrderListWithPageRequest struct {
	api.BaseRequest
	PageSize    int    `json:"pageSize,omitempty" codec:"pageSize,omitempty"`       //页面中展示个数
	CurrentPage int    `json:"currentPage,omitempty" codec:"currentPage,omitempty"` //
	FwsPin      string `json:"fwsPin,omitempty" codec:"fwsPin,omitempty"`           //服务商PIN
	ServiceCode string `json:"serviceCode,omitempty" codec:"serviceCode,omitempty"` //服务编码
}

type OrderListWithPageResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *OrderListWithPageData `json:"jingdong_pop_fw_order_listwithpage_responce,omitempty" codec:"jingdong_pop_fw_order_listwithpage_responce,omitempty"`
}

type OrderListWithPageData struct {
	ReturnType *OrderListWithPageReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                       `json:"code"`
}

type OrderListWithPageReturnType struct {
	TotalPage   uint64     `json:"totalPage,omitempty" codec:"totalPage,omitempty"`     //实际页数
	PageSize    uint64     `json:"pageSize,omitempty" codec:"pageSize,omitempty"`       //每页个数
	ErrorCode   uint64     `json:"errorCode,omitempty" codec:"errorCode,omitempty"`     //错误码
	CurrentPage uint64     `json:"currentPage,omitempty" codec:"currentPage,omitempty"` //当前页
	TotalCount  uint64     `json:"totalCount,omitempty" codec:"totalCount,omitempty"`   //总个数
	IsSuccess   bool       `json:"isSuccess,omitempty" codec:"isSuccess,omitempty"`     //是否成功
	ErrorMsg    string     `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`       // 错误信息
	OrderList   []*FWOrder `json:"orderList,omitempty" codec:"orderList,omitempty"`     //订单列表
}

type FWOrder struct {
	OrderId           uint64 `json:"orderId,omitempty" codec:"orderId,omitempty"`                     //订单号
	EndDate           int64  `json:"endDate,omitempty" codec:"endDate,omitempty"`                     //结束时间
	ItemCode          string `json:"itemCode,omitempty" codec:"itemCode,omitempty"`                   //收费项目编码
	ItemName          string `json:"itemName,omitempty" codec:"itemName,omitempty"`                   //收费项目名称
	ArticleType       uint8  `json:"articleType,omitempty" codec:"articleType,omitempty"`             //订购类型（1：订购、2：续费、3：升级）
	ErpOrderId        uint64 `json:"erpOrderId,omitempty" codec:"erpOrderId,omitempty"`               // 结算订单编号
	TotalRealpayPrice uint64 `json:"totalRealpayPrice,omitempty" codec:"totalRealpayPrice,omitempty"` // 实付金额（单位：分）
	SkuId             uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`                         // 商品skuId
	FwsPin            string `json:"fwsPin,omitempty" codec:"fwsPin,omitempty"`                       //服务商pin
	TpOrderId         uint64 `json:"tpOrderId,omitempty" codec:"tpOrderId,omitempty"`                 // 钱包订单号
	NickName          string `json:"nickName,omitempty" codec:"nickName,omitempty"`                   // 购买人昵称
	Created           int64  `json:"created,omitempty" codec:"created,omitempty"`                     // 创建时间
	Buyer             string `json:"buyer,omitempty" codec:"buyer,omitempty"`                         //购买人pin
	OrderCycle        int64  `json:"orderCycle,omitempty" codec:"orderCycle,omitempty"`               // 订购周期
	OrderDate         int64  `json:"orderDate,omitempty" codec:"orderDate,omitempty"`                 // 下单时间
	StartDate         int64  `json:"startDate,omitempty" codec:"startDate,omitempty"`                 // 服务开始时间
	ServiceCode       string `json:"serviceCode,omitempty" codec:"serviceCode,omitempty"`             // 服务编码
	OrderStatus       uint8  `json:"orderStatus,omitempty" codec:"orderStatus,omitempty"`             // 订单状态（1：等待付款、2：等待付款确认、4：订单完成、8：取消完成、16：退款中、64、退款完成）
	OrderNum          uint64 `json:"orderNum,omitempty" codec:"orderNum,omitempty"`                   // 订单数量
	ChildBuyer        string `json:"childBuyer,omitempty" codec:"childBuyer,omitempty"`               // 购买人子账号
	ServiceName       string `json:"serviceName,omitempty" codec:"serviceName,omitempty"`             //服务名称
}

func OrderListWithPage(req *OrderListWithPageRequest) (*OrderListWithPageReturnType, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fw.NewOrderListwithpageRequest()

	if req.CurrentPage > 0 {
		r.SetCurrentPage(req.CurrentPage)
	} else {
		r.SetCurrentPage(1)
	}

	if len(req.FwsPin) > 0 {
		r.SetFwsPin(req.FwsPin)
	}

	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	} else {
		r.SetPageSize(20)
	}

	if len(req.ServiceCode) > 0 {
		r.SetServiceCode(req.ServiceCode)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result info.")
	}
	var response OrderListWithPageResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.ReturnType.ErrorCode != 0 {
		return nil, errors.New(response.Data.ReturnType.ErrorMsg)
	}
	return response.Data.ReturnType, nil

}
