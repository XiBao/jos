package order

import "github.com/XiBao/jos/sdk"

type OrderBonusQueryReq struct {
	// OptType 时间类型（1.下单时间拉取、2.更新时间拉取）
	OptType int `json:"optType,omitempty"`
	// StartTime 订单开始时间，时间戳（毫秒），与endTime间隔不超过10分钟
	StartTime int64 `json:"startTime,omitempty"`
	// EndTime 订单结束时间，时间戳（毫秒），与startTime间隔不超过10分钟
	EndTime int64 `json:"endTime,omitempty"`
	// PageNo 页码，默认值为1
	PageNo int `json:"pageNo,omitempty"`
	// PageSize 每页数量，上限100
	PageSize int `json:"pageSize,omitempty"`
	// SortValue 与pageNo、pageSize组合使用。获取当前页最后一条记录的sortValue，下一页请求传入该值
	SortValue string `json:"sortValue,omitempty"`
	// ActivityID 奖励活动ID
	ActivityID uint64 `json:"activityId,omitempty"`
}

type UnionOrderBonusQueryRequest struct {
	Request *sdk.Request
}

// create new request
func NewUnionOrderBonusQueryRequest() (req *UnionOrderBonusQueryRequest) {
	request := sdk.Request{MethodName: "jd.union.open.order.bonus.query", IsUnionGW: true, Params: make(map[string]interface{}, 1)}
	req = &UnionOrderBonusQueryRequest{
		Request: &request,
	}
	return
}

func (req *UnionOrderBonusQueryRequest) SetOrderReq(orderReq *OrderBonusQueryReq) {
	req.Request.Params["orderReq"] = orderReq
}

func (req *UnionOrderBonusQueryRequest) GetOrderReq() *OrderBonusQueryReq {
	orderReq, found := req.Request.Params["orderReq"]
	if found {
		return orderReq.(*OrderBonusQueryReq)
	}
	return nil
}
