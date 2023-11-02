package order

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/order"
)

type UnionOrderBonusQueryRequest struct {
	api.BaseRequest
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

type UnionOrderBonusQueryResponse struct {
	ErrorResp *api.ErrorResponnse               `json:"error_response,omitempty"`
	Data      *UnionOrderBonusQueryResponseData `json:"jd_union_open_order_bonus_query_responce,omitempty"`
}

func (r UnionOrderBonusQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == ""
}

func (r UnionOrderBonusQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type UnionOrderBonusQueryResponseData struct {
	Result string `json:"queryResult,omitempty"`
}

type UnionOrderBonusQueryResult struct {
	Code    int          `json:"code,omitempty"`
	Message string       `json:"message,omitempty"`
	Data    []OrderBonus `json:"data,omitempty"`
}

func (r UnionOrderBonusQueryResult) IsError() bool {
	return r.Code != 200
}

func (r UnionOrderBonusQueryResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

type OrderBonus struct {
	// UnionId 联盟ID
	UnionId uint64 `json:"unionId,omitempty"`
	// BonusInvalidCode 无效状态码，-1:无效、2:无效-拆单、3:无效-取消、4:无效-京东帮帮主订单、5:无效-账号异常、6:无效-赠品类目不返佣 等
	BonusInvalidCode string `json:"bonusInvalidCode,omitempty"`
	// BonusState
	BonusState int `json:"bonusState,omitempty"`
	// BonusInvalidText 无效状态码对应的无效状态文案
	BonusInvalidText string `json:"bonusInvalidText,omitempty"`
	// EstimateBonusFee 预估奖励金额：查询时间范围内，已付款且奖励有效，满足奖励规则的奖励金额
	EstimateBonusFee float64 `json:"estimateBonusFee,omitempty"`
	// ActualBonusFee 实际奖励金额：查询时间范围内，已付款或已完成（视具体规则），奖励有效且满足奖励规则的奖励金额
	ActualBonusFee float64 `json:"actualBonusFee,omitempty"`
	// OrderState 奖励订单状态，1:已完成、2:已付款、3:待付款
	OrderState int `json:"orderState,omitempty"`
	// OrderText 奖励订单状态，待付款/已付款/已完成
	OrderText string `json:"orderText,omitempty"`
	// ActivityName  活动名称
	ActivityName string `json:"activityName,omitempty"`
	// ActivityId 奖励活动ID
	ActivityId uint64 `json:"activityId,omitempty"`
	FinishTime int64  `json:"finishTime"`
	OrderTime  int64  `json:"orderTime"`
	OrderId    uint64 `json:"orderId"`
	ParentId   uint64 `json:"ParentId"`
	Ext1       string `json:"ext1,omitempty"`
	// SortValue 排序值，按'下单时间'分页查询时使用
	SortValue string `json:"sortValue,omitempty"`
	SkuInfo
}

// UnionOrderBonusQuery
func UnionOrderBonusQuery(req *UnionOrderBonusQueryRequest) ([]OrderBonus, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := order.NewUnionOrderBonusQueryRequest()
	orderReq := &order.OrderBonusQueryReq{
		OptType:    req.OptType,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		PageNo:     req.PageNo,
		PageSize:   req.PageSize,
		SortValue:  req.SortValue,
		ActivityID: req.ActivityID,
	}
	r.SetOrderReq(orderReq)

	var response UnionOrderBonusQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	var ret UnionOrderBonusQueryResult
	if err := client.Logger().DecodeJSON([]byte(response.Data.Result), &ret); err != nil {
		return nil, err
	}
	if ret.IsError() {
		return nil, ret
	}

	return ret.Data, nil
}
