package asc

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/asc"
)

type ServiceAndRefundViewRequest struct {
	api.BaseRequest
	OrderId          uint64 `json:"order_id,omitempty" codec:"order_id,omitempty"`                 //订单号
	ApplyTimeBegin   string `json:"applyTimeBegin,omitempty" codec:"applyTimeBegin,omitempty"`     //申请时间开始（与结束两两不为空）
	ApplyTimeEnd     string `json:"applyTimeEnd,omitempty" codec:"applyTimeEnd,omitempty"`         //申请时间结束（与开始两两不为空）
	ApproveTimeBegin string `json:"approveTimeBegin,omitempty" codec:"approveTimeBegin,omitempty"` //审核时间开始（与结束两两不为空）
	ApproveTimeEnd   string `json:"approveTimeEnd,omitempty" codec:"approveTimeEnd,omitempty"`     //审核时间结束（与开始两两不为空）
	PageNumber       uint64 `json:"pageNumber,omitempty" codec:"pageNumber,omitempty"`             //页码(从1开始)
	PageSize         uint8  `json:"pageSize,omitempty" codec:"pageSize,omitempty"`                 //每页大小（1-50，默认10）
	ExtJsonStr       string `json:"extJsonStr,omitempty" codec:"extJsonStr,omitempty"`             //扩展条件（JSON格式）
}

type ServiceAndRefundViewResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ServiceAndRefundViewData `json:"jingdong_asc_serviceAndRefund_view_responce,omitempty" codec:"jingdong_asc_serviceAndRefund_view_responce,omitempty"`
}

func (r ServiceAndRefundViewResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ServiceAndRefundViewResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ServiceAndRefundViewData struct {
	PageResult *ServiceAndRefundViewPageResult `json:"pageResult,omitempty" codec:"pageResult,omitempty"`
}

func (r ServiceAndRefundViewData) IsError() bool {
	return r.PageResult == nil || r.PageResult.IsError()
}

func (r ServiceAndRefundViewData) Error() string {
	if r.PageResult != nil {
		return r.PageResult.Error()
	}
	return "no result data"
}

type ServiceAndRefundViewPageResult struct {
	Success    bool                `json:"success,omitempty" codec:"success,omitempty"`       //接口调用成功标识
	Code       string              `json:"code,omitempty" codec:"code,omitempty"`             //结果编码
	Msg        string              `json:"msg,omitempty" codec:"msg,omitempty"`               //结果描述
	PageSize   string              `json:"pageSize,omitempty" codec:"pageSize,omitempty"`     //分页大小
	PageNumber string              `json:"pageNumber,omitempty" codec:"pageNumber,omitempty"` //页码
	TotalCount string              `json:"totalCount,omitempty" codec:"totalCount,omitempty"` //总数
	Data       []OrderAfsAndRefund `json:"data,omitempty" codec:"data,omitempty"`             //售后和退款列表
}

func (r ServiceAndRefundViewPageResult) IsError() bool {
	return !r.Success
}

func (r ServiceAndRefundViewPageResult) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

type OrderAfsAndRefund struct {
	AfsRefundId          uint64                `json:"afsRefundId,omitempty" codec:"afsRefundId,omitempty"`                   //退款信息ID
	RefoundAmount        uint64                `json:"refoundAmount,omitempty" codec:"refoundAmount,omitempty"`               //实际退款金额
	CompleteTime         int64                 `json:"completeTime,omitempty" codec:"completeTime,omitempty"`                 //退款完成时间
	Status               uint8                 `json:"status,omitempty" codec:"status,omitempty"`                             //退款状态	13,成功,14,失败,others,其他
	SameOrderServiceBill *SameOrderServiceBill `json:"sameOrderServiceBill,omitempty" codec:"sameOrderServiceBill,omitempty"` //售后数据
}

type SameOrderServiceBill struct {
	ServiceId    uint64 `json:"serviceId,omitempty" codec:"serviceId,omitempty"`       //服务单号
	AfsApplyId   uint64 `json:"afsApplyId,omitempty" codec:"afsApplyId,omitempty"`     //申请单号
	OrderId      uint64 `json:"orderId,omitempty" codec:"orderId,omitempty"`           //订单单号
	WareId       uint64 `json:"wareId,omitempty" codec:"wareId,omitempty"`             //商品编号
	WareName     string `json:"wareName,omitempty" codec:"wareName,omitempty"`         //商品名称
	ApproveName  string `json:"approveName,omitempty" codec:"approveName,omitempty"`   //审核人姓名
	AfsApplyTime int64  `json:"afsApplyTime,omitempty" codec:"afsApplyTime,omitempty"` //申请时间
	ApproveDate  int64  `json:"approveDate,omitempty" codec:"approveDate,omitempty"`   //审核时间
	QuestionDesc string `json:"questionDesc,omitempty" codec:"questionDesc,omitempty"` //问题描述
	CustomerPin  string `json:"customerPin,omitempty" codec:"customerPin,omitempty"`   //客户账号
	ApproveNotes string `json:"approveNotes,omitempty" codec:"approveNotes,omitempty"` //审核意见
	ApplyReason  string `json:"applyReason,omitempty" codec:"applyReason,omitempty"`   //申请原因
}

// TODO 查询店铺退款明细
func ServiceAndRefundView(req *ServiceAndRefundViewRequest) (*ServiceAndRefundViewPageResult, error) {

	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := asc.NewServiceAndRefundViewRequest()

	if req.OrderId > 0 {
		r.SetOrderId(req.OrderId)
	}
	if len(req.ApplyTimeBegin) > 0 {
		r.SetApplyTimeBegin(req.ApplyTimeBegin)
	}
	if len(req.ApplyTimeEnd) > 0 {
		r.SetApplyTimeEnd(req.ApplyTimeEnd)
	}
	if len(req.ApproveTimeBegin) > 0 {
		r.SetApproveTimeBegin(req.ApproveTimeBegin)
	}
	if len(req.ApproveTimeEnd) > 0 {
		r.SetApproveTimeEnd(req.ApproveTimeEnd)
	}
	if req.PageNumber > 0 {
		r.SetPageNumber(req.PageNumber)
	}
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	}
	if len(req.ExtJsonStr) > 0 {
		r.SetExtJsonStr(req.ExtJsonStr)
	}

	var response ServiceAndRefundViewResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.PageResult, nil

}
