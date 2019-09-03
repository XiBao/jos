package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/crm"
	"github.com/daviddengcn/ljson"
)

type MemberSearchNewRequest struct {
	api.BaseRequest

	Pin              string `json:"pin,omitempty" codec:"pin,omitempty"`                                 // 会员名称,用户在京东的唯一标识
	Grade            string `json:"grade,omitempty" codec:"grade,omitempty"`                             // 会员等级(可支持多个，用逗号分隔，如：1,2,3)
	MinLastTradeTime string `json:"min_last_trade_time,omitempty" codec:"min_last_trade_time,omitempty"` // 最早上次交易时间，精确至年月日
	MaxLastTradeTime string `json:"max_last_trade_time,omitempty" codec:"max_last_trade_time,omitempty"` // 最迟上次交易时间，精确至年月日
	MinTradeCount    int    `json:"min_trade_count,omitempty" codec:"min_trade_count,omitempty"`         // 最小交易量
	MaxTradeCount    int    `json:"max_trade_count,omitempty" codec:"max_trade_count,omitempty"`         // 最大交易量
	AvgPrice         int    `json:"avg_price,omitempty" codec:"avg_price,omitempty"`                     // 最小平均客单价，单位为元
	MinTradeAmount   int    `json:"min_trade_amount,omitempty" codec:"min_trade_amount,omitempty"`       // 最小交易额,单位为元
	Page             int    `json:"page,omitempty" codec:"page,omitempty"`                               // 显示第几页的会员，如果输入的页码大于总共的页码数，例如总共10页，但是current_page的值为11，则返回空白页，最小页数为1，最大页数为100
	PageSize         int    `json:"page_size,omitempty" codec:"page_size,omitempty"`                     // 表示每页显示的会员数量,page_size的最大值不能超过100条,最小值不能低于1
}

type MemberSearchNewResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *MemberSearchNewData `json:"jingdong_crm_member_searchNew_responce",omitempty" codec:"jingdong_crm_member_searchNew_responce",omitempty"`
}

type MemberSearchNewData struct {
	Code   string                 `json:"code,omitempty" codec:"code,omitempty"`
	Result *MemberSearchNewResult `json:"crm_member_result,omitempty" codec:"crm_member_result,omitempty"`
}

type MemberSearchNewResult struct {
	TotalResult int       `json:"total_result,omitempty" codec:"total_result,omitempty"`
	Members     []*Member `json:"crm_members,omitempty" codec:"crm_members,omitempty"`
}

func MemberSearchNew(req *MemberSearchNewRequest) (*MemberSearchNewResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewMemberSearchNewRequest()
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	} else {
		r.SetPageSize(100)
	}
	if req.Page > 0 {
		r.SetPage(req.Page)
	} else {
		r.SetPage(1)
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result info.")
	}
	var response MemberSearchNewResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.Result == nil {
		return nil, nil
	}

	return response.Data.Result, nil
}
