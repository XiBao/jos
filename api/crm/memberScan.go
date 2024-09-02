package crm

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
)

type MemberScanRequest struct {
	api.BaseRequest

	Pin              string `json:"pin,omitempty" codec:"pin,omitempty"`                                 // 会员名称,用户在京东的唯一标识
	Grade            string `json:"grade,omitempty" codec:"grade,omitempty"`                             // 会员等级(可支持多个，用逗号分隔，如：1,2,3)
	MinLastTradeTime string `json:"min_last_trade_time,omitempty" codec:"min_last_trade_time,omitempty"` // 最早上次交易时间，精确至年月日
	MaxLastTradeTime string `json:"max_last_trade_time,omitempty" codec:"max_last_trade_time,omitempty"` // 最迟上次交易时间，精确至年月日
	MinTradeCount    int    `json:"min_trade_count,omitempty" codec:"min_trade_count,omitempty"`         // 最小交易量
	MaxTradeCount    int    `json:"max_trade_count,omitempty" codec:"max_trade_count,omitempty"`         // 最大交易量
	AvgPrice         int    `json:"avg_price,omitempty" codec:"avg_price,omitempty"`                     // 最小平均客单价，单位为元
	MinTradeAmount   int    `json:"min_trade_amount,omitempty" codec:"min_trade_amount,omitempty"`       // 最小交易额,单位为元
	PageSize         int    `json:"page_size,omitempty" codec:"page_size,omitempty"`                     // 页大小，查询第一页时必填
	ScrollId         string `json:"scroll_id,omitempty" codec:"scroll_id,omitempty"`                     // 上一次查询的游标，查询第一页时不填。第二页及以后scrollID为必填项，其他参数均可不填，scrollID为唯一有效参数。
}

type MemberScanResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *MemberScanData     `json:"jingdong_crm_member_scan_responce,omitempty" codec:"jingdong_crm_member_scan_responce,omitempty"`
}

func (r MemberScanResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == nil
}

func (r MemberScanResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type MemberScanData struct {
	Code   string            `json:"code,omitempty" codec:"code,omitempty"`
	Result *MemberScanResult `json:"crm_member_scan_result,omitempty" codec:"crm_member_scan_result,omitempty"`
}

type MemberScanResult struct {
	TotalResult int      `json:"total_result,omitempty" codec:"total_result,omitempty"`
	Members     []Member `json:"crm_members,omitempty" codec:"crm_members,omitempty"`
	ScrollId    string   `json:"scroll_id,omitempty" codec:"scroll_id,omitempty"`
}

func MemberScan(ctx context.Context, req *MemberScanRequest) (*MemberScanResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewMemberScanRequest()
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	}
	if req.ScrollId != "" {
		r.SetScrollId(req.ScrollId)
	}
	if req.Pin != "" {
		r.SetPin(req.Pin)
	}
	if req.Grade != "" {
		r.SetGrade(req.Grade)
	}
	if req.MinLastTradeTime != "" {
		r.SetMinLastTradeTime(req.MinLastTradeTime)
	}
	if req.MaxLastTradeTime != "" {
		r.SetMaxLastTradeTime(req.MaxLastTradeTime)
	}
	if req.MinTradeCount > 0 {
		r.SetMinTradeCount(req.MinTradeCount)
	}
	if req.MaxTradeCount > 0 {
		r.SetMaxTradeCount(req.MaxTradeCount)
	}
	if req.AvgPrice > 0 {
		r.SetAvgPrice(req.AvgPrice)
	}
	if req.MinTradeAmount > 0 {
		r.SetMinTradeAmount(req.MinTradeAmount)
	}

	var response MemberScanResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
