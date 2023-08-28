package order

import (
	"github.com/XiBao/jos/sdk"
)

type OrderRowReq struct {
	PageIndex    uint   `json:"pageIndex"`
	PageSize     uint   `json:"pageSize"`               // 每页包含条数，上限为500
	Type         uint   `json:"type"`                   // 订单时间查询类型(1：下单时间，2：完成时间（购买用户确认收货时间），3：更新时间
	StartTime    string `json:"startTime"`              // 开始时间 格式yyyy-MM-dd HH:mm:ss，与endTime间隔不超过1小时
	EndTime      string `json:"endTime"`                // 结束时间 格式yyyy-MM-dd HH:mm:ss，与startTime间隔不超过1小时
	ChildUnionId string `json:"childUnionId,omitempty"` // 子推客unionID，传入该值可查询子推客的订单，注意不可和key同时传入。（需联系运营开通PID权限才能拿到数据）
	Key          string `json:"key,omitempty"`          // 工具商传入推客的授权key，可帮助该推客查询订单，注意不可和childUnionid同时传入。（需联系运营开通工具商权限才能拿到数据）
	Fields       string `json:"fields,omitempty"`       // 支持出参数据筛选，逗号','分隔，目前可用：goodsInfo（商品信息）,categoryInfo(类目信息）
}

type UnionOrderQueryRowRequest struct {
	Request *sdk.Request
}

// create new request
func NewUnionOrderQueryRowRequest() (req *UnionOrderQueryRowRequest) {
	request := sdk.Request{MethodName: "jd.union.open.order.row.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &UnionOrderQueryRowRequest{
		Request: &request,
	}
	return
}

func (req *UnionOrderQueryRowRequest) SetOrderRowReq(orderRowReq *OrderRowReq) {
	req.Request.Params["orderReq"] = orderRowReq
}

func (req *UnionOrderQueryRowRequest) GetOrderRowReq() *OrderRowReq {
	orderRowReq, found := req.Request.Params["orderReq"]
	if found {
		return orderRowReq.(*OrderRowReq)
	}
	return nil
}
