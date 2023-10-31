package report

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/report"
)

type QueryAccountReportRequest struct {
	api.BaseRequest

	StartDay            string `json:"startDay"`                      // 起始日期
	EndDay              string `json:"endDay"`                        // 结束日期
	IsDaily             string `json:"isDaily"`                       // 是否分日
	Platform            string `json:"platform"`                      // 设备(all: 全部，pc: PC，mobile:无线)
	ClickOrOrderDay     uint8  `json:"clickOrOrderDay"`               // 口径(0:当天，1:1天，3:3天，7:7天，15:15天)
	ClickOrOrderCaliber uint8  `json:"clickOrOrderCaliber"`           // 点击口径/下单口径(0-点击口径，1-下单口径)
	OrderStatusCategory uint8  `json:"orderStatusCategory,omitempty"` // 全部订单/成交订单(null:全部订单，1: 成交订单)
	PageIndex           uint16 `json:"pageIndex"`                     // 当前页码
	PageSize            uint8  `json:"pageSize"`                      // 每页数量(最大值100)
}

type QueryAccountReportResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryAccountReportData `json:"jingdong_dsp_report_queryAccountReport_responce,omitempty" codec:"jingdong_dsp_report_queryAccountReport_responce,omitempty"`
}

func (r QueryAccountReportResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r QueryAccountReportResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type QueryAccountReportData struct {
	Code      string                    `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                    `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *QueryAccountReportResult `json:"querycampdailysum_result,omitempty" codec:"querycampdailysum_result,omitempty"`
}

func (r QueryAccountReportData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r QueryAccountReportData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	if r.Code != "0" {
		return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type QueryAccountReportResult struct {
	Value      *QueryAccountReportResultValue `json:"value,omitempty" codec:"value,omitempty"`
	ResultCode string                         `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string                         `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool                           `json:"success" codec:"success"`
}

func (r QueryAccountReportResult) IsError() bool {
	return !r.Success || r.Value == nil
}

func (r QueryAccountReportResult) Error() string {
	if !r.Success {
		return fmt.Sprintf("code:%s, msg:%s", r.ResultCode, r.ErrorMsg)
	}
	return "no result data"
}

type QueryAccountReportResultValue struct {
	Datas []*AccountReport `json:"datas" codec:"datas"`
	// Ext          *DspAccountReport    `json:"ext" codec:"ext"`
	// Paginator    string               `json:"paginator" codec:"paginator"`
}

type AccountReport struct {
	Date              string  `json:"date"`
	Pin               string  `json:"pin"`
	MobileType        string  `json:"mobileType"`
	Impressions       uint64  `json:"impressions"`
	Clicks            uint64  `json:"clicks"`
	Ctr               string  `json:"CTR"`
	Cost              string  `json:"cost"`
	Cpm               string  `json:"CPM"`
	Cpc               string  `json:"CPC"`
	DirectOrderCnt    uint64  `json:"directOrderCnt"`
	DirectOrderSum    string  `json:"directOrderSum"`
	IndirectOrderCnt  uint64  `json:"indirectOrderCnt"`
	IndirectOrderSum  string  `json:"indirectorderSum"`
	TotalOrderCnt     uint64  `json:"totalOrderCnt"`
	TotalOrderSum     string  `json:"totalOrderSum"`
	DirectCartCnt     uint64  `json:"directCartCnt"`
	IndirectCartCnt   uint64  `json:"indirectCartCnt"`
	TotalCartCnt      uint64  `json:"totalCartCnt"`
	TotalOrderCvs     string  `json:"totalOrderCvs"`
	TotalOrderRoi     string  `json:"totalOrderRoi"`
	Cpa               string  `json:"CPA"`
	DepartmentCnt     uint64  `json:"departmentCnt"`
	DepartmentGmv     float64 `json:"departmentGmv"`
	PlatformCnt       uint64  `json:"platformCnt"`
	PlatformGmv       float64 `json:"platformGmv"`
	ChannelRoi        float64 `json:"channelRoi"`
	NewCustomersCnt   uint64  `json:"newCustomersCnt"`
	UserVisitorCnt    uint64  `json:"userVisitorCnt"`
	VisitPageCnt      uint64  `json:"visitPageCnt"`
	VisitTimeAverage  float64 `json:"visitTimeAverage"`
	DepthPassengerCnt uint64  `json:"depthPassengerCnt"`
	GoodsAttentionCnt uint64  `json:"goodsAttentionCnt"`
	ShopAttentionCnt  uint64  `json:"shopAttentionCnt"`
	PreorderCnt       uint64  `json:"preorderCnt"`
	CouponCnt         uint64  `json:"couponCnt"`
}

func QueryAccountReport(req *QueryAccountReportRequest) ([]*AccountReport, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := report.NewDspReportQueryAccountReportRequest()
	r.SetStartDay(req.StartDay)
	r.SetEndDay(req.EndDay)
	r.SetIsDaily(req.IsDaily)
	r.SetPlatform(req.Platform)
	r.SetClickOrOrderDay(req.ClickOrOrderDay)
	r.SetClickOrOrderCaliber(req.ClickOrOrderCaliber)
	r.SetPageSize(req.PageSize)
	r.SetPageIndex(req.PageIndex)
	if req.OrderStatusCategory > 0 {
		r.SetOrderStatusCategory(req.OrderStatusCategory)
	}

	var response QueryAccountReportResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Value.Datas, nil
}
