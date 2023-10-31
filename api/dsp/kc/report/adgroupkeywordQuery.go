package report

import (
	"fmt"
	"strconv"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/report"
)

type AdgroupkeywordQueryRequest struct {
	api.BaseRequest

	StartDate           string `json:"start_date,omitempty" codec:"start_date,omitempty"`                       // 起始时间
	EndDate             string `json:"end_date,omitempty" codec:"end_date,omitempty"`                           // 结束时间
	GroupId             uint64 `json:"group_id,omitempty" codec:"group_id,omitempty"`                           // 单元ID
	IsOrderOrClick      bool   `json:"is_order_or_click,omitempty" codec:"is_order_or_click,omitempty"`         // 下单点击口径(true:下单口径;flase:点击口径)
	OrderStatusCategory int    `json:"order_status_category,omitempty" codec:"order_status_category,omitempty"` // GMV订单类型(空:全部;1:成交订单)
	Platform            string `json:"platform,omitempty" codec:"platform,omitempty"`                           // 推广设备(全部:all;PC:pc;无线:mobile)
	ValType             string `json:"val_type,omitempty" codec:"val_type,omitempty"`                           // 购买类型exact精确匹配,term短语匹配,segment切词匹配
	Val                 string `json:"val,omitempty" codec:"val,omitempty"`                                     // 关键字
	ClickOrOrderDay     uint8  `json:"click_or_order_day,omitempty" codec:"click_or_order_day,omitempty"`       // 当天/1天/15天 当天：0,1天：1， 15天：15
	PageIndex           int    `json:"page_index,omitempty" codec:"page_index,omitempty"`                       // 当前页码
	PageSize            int    `json:"page_size,omitempty" codec:"page_size,omitempty"`                         // 每页数量(最大值100)
}

type AdgroupkeywordQueryResponse struct {
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AdgroupkeywordQueryData `json:"jingdong_dsp_report_adgroupkeyword_query_responce,omitempty" codec:"jingdong_dsp_report_adgroupkeyword_query_responce,omitempty"`
}

func (r AdgroupkeywordQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AdgroupkeywordQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AdgroupkeywordQueryData struct {
	Result *AdgroupkeywordQueryResult `json:"keywordreportquery_result,omitempty" codec:"keywordreportquery_result,omitempty"`
}

func (r AdgroupkeywordQueryData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r AdgroupkeywordQueryData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type AdgroupkeywordQueryResult struct {
	Status     int    `json:"status,omitempty" codec:"status,omitempty"`
	ResultCode string `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool   `json:"success,omitempty" codec:"success,omitempty"`
	Value      *AdgroupkeywordQueryValue
}

func (r AdgroupkeywordQueryResult) IsError() bool {
	return !r.Success || r.Value == nil
}

func (r AdgroupkeywordQueryResult) Error() string {
	if !r.Success {
		return fmt.Sprintf("code:%s, msg:%s", r.ResultCode, r.ErrorMsg)
	}
	return "no result data"
}

type AdgroupkeywordQueryValue struct {
	Datas []*JosKeywordTmp `json:"datas,omitempty" codec:"datas,omitempty"`
}

type JosKeywordTmp struct {
	Id                      uint64  `json:"id,omitempty" codec:"id,omitempty"`
	Keyword                 string  `json:"Keyword,omitempty" codec:"Keyword,omitempty"`
	KeywordType             string  `json:"type,omitempty" codec:"type,omitempty"`
	Price                   float64 `json:"keywordPrice,omitempty" codec:"keywordPrice,omitempty"`
	PriceWireless           float64 `json:"keywordPriceWireless,omitempty" codec:"keywordPriceWireless,omitempty"`
	Clicks                  string  `json:"Clicks,omitempty" codec:"Clicks,omitempty"`
	Cost                    string  `json:"Cost,omitempty" codec:"Cost,omitempty"`
	CTR                     string  `json:"CTR,omitempty" codec:"CTR,omitempty"`
	CPC                     string  `json:"CPC,omitempty" codec:"CPC,omitempty"`
	CPM                     string  `json:"CPM,omitempty" codec:"CPM,omitempty"`
	TotalOrderROI           string  `json:"TotalOrderROI,omitempty" codec:"TotalOrderROI,omitempty"`
	Impressions             string  `json:"Impressions,omitempty" codec:"Impressions,omitempty"`
	TotalOrderCnt           string  `json:"TotalOrderCnt,omitempty" codec:"TotalOrderCnt,omitempty"`
	TotalOrderSum           string  `json:"TotalOrderSum,omitempty" codec:"TotalOrderSum,omitempty"`
	TotalOrderCVS           string  `json:"TotalOrderCVS,omitempty" codec:"TotalOrderCVS,omitempty"`
	IndirectOrderCnt        string  `json:"IndirectOrderCnt,omitempty" codec:"IndirectOrderCnt,omitempty"`
	DirectOrderCnt          string  `json:"DirectOrderCnt,omitempty" codec:"DirectOrderCnt,omitempty"`
	IndirectOrderSum        string  `json:"IndirectOrderSum,omitempty" codec:"IndirectOrderSum,omitempty"`
	DirectOrderSum          string  `json:"DirectOrderSum,omitempty" codec:"DirectOrderSum,omitempty"`
	IndirectCartCnt         string  `json:"IndirectCartCnt,omitempty" codec:"IndirectCartCnt,omitempty"`
	DirectCartCnt           string  `json:"DirectCartCnt,omitempty" codec:"DirectCartCnt,omitempty"`
	TotalCartCnt            string  `json:"TotalCartCnt,omitempty" codec:"TotalCartCnt,omitempty"`
	CurrentWlShowq          string  `json:"CurrentWlShowq,omitempty" codec:"CurrentWlShowq,omitempty"`
	CurrentPcShowq          string  `json:"CurrentPcShowq,omitempty" codec:"CurrentPcShowq,omitempty"`
	CurrentHourPcRank       string  `json:"CurrentHourPcRank,omitempty" codec:"CurrentHourPcRank,omitempty"`
	AverageCurrentPcRank    string  `json:"AverageCurrentPcRank,omitempty" codec:"AverageCurrentPcRank,omitempty"`
	AverageHistoryWlRank    string  `json:"AverageHistoryWlRank,omitempty" codec:"AverageHistoryWlRank,omitempty"`
	AverageHistoryPcRank    string  `json:"AverageHistoryPcRank,omitempty" codec:"AverageHistoryPcRank,omitempty"`
	AverageHistoryTotalRank string  `json:"AverageHistoryTotalRank,omitempty" codec:"AverageHistoryTotalRank,omitempty"`
	EffectOrderSum          string  `json:"EffectOrderSum,omitempty" codec:"EffectOrderSum,omitempty"`
	EffectOrderCnt          string  `json:"EffectOrderCnt,omitempty" codec:"EffectOrderCnt,omitempty"`
	EffectCartCnt           string  `json:"EffectCartCnt,omitempty" codec:"EffectCartCnt,omitempty"`
	IsDefaultPrice          bool    `json:"isDefaultPrice,omitempty" codec:"isDefaultPrice,omitempty"`
	OrderCVS                string  `json:"OrderCVS,omitempty" codec:"OrderCVS,omitempty"`
	NewTotalOrderCnt        string  `json:"NewTotalOrderCnt,omitempty" codec:"NewTotalOrderCnt,omitempty"`
	NewTotalOrderSum        string  `json:"NewTotalOrderSum,omitempty" codec:"NewTotalOrderCnt,omitempty"`
	NewTotalOrderCVS        string  `json:"NewTotalOrderCVS,omitempty" codec:"NewTotalOrderCVS,omitempty"`
	KeywordFlag             string  `json:"KeywordFlag,omitempty" codec:"KeywordFlag,omitempty"`
	MobileType              string  `json:"MobileType,omitempty" codec:"MobileType,omitempty"`
}

type JosKeyword struct {
	Id                      uint64    `json:"id,omitempty" codec:"id,omitempty"`
	Keyword                 string    `json:"keyword,omitempty" codec:"keyword,omitempty"`
	KeywordType             string    `json:"type,omitempty" codec:"type,omitempty"`
	Price                   float64   `json:"keyword_price,omitempty" codec:"keyword_price,omitempty"`
	PriceWireless           float64   `json:"keyword_price_wireless,omitempty" codec:"keyword_price_wireless,omitempty"`
	Clicks                  uint64    `json:"clicks,omitempty" codec:"clicks,omitempty"`
	Cost                    float64   `json:"cost,omitempty" codec:"cost,omitempty"`
	CTR                     float64   `json:"ctr,omitempty" codec:"ctr,omitempty"`
	CPC                     float64   `json:"cpc,omitempty" codec:"cpc,omitempty"`
	CPM                     float64   `json:"cpm,omitempty" codec:"cpm,omitempty"`
	TotalOrderROI           float64   `json:"total_order_roi,omitempty" codec:"total_order_roi,omitempty"`
	Impressions             uint64    `json:"impressions,omitempty" codec:"impressions,omitempty"`
	TotalOrderCnt           uint64    `json:"total_order_cnt,omitempty" codec:"total_order_cnt,omitempty"`
	TotalOrderSum           float64   `json:"total_order_sum,omitempty" codec:"total_order_sum,omitempty"`
	TotalOrderCVS           float64   `json:"total_order_cvs,omitempty" codec:"total_order_cvs,omitempty"`
	IndirectOrderCnt        uint64    `json:"indirect_order_cnt,omitempty" codec:"indirect_order_cnt,omitempty"`
	DirectOrderCnt          uint64    `json:"direct_order_cnt,omitempty" codec:"direct_order_cnt,omitempty"`
	IndirectOrderSum        float64   `json:"indirect_order_sum,omitempty" codec:"indirect_order_sum,omitempty"`
	DirectOrderSum          float64   `json:"direct_order_sum,omitempty" codec:"direct_order_sum,omitempty"`
	IndirectCartCnt         uint64    `json:"indirect_cart_cnt,omitempty" codec:"indirect_cart_cnt,omitempty"`
	DirectCartCnt           uint64    `json:"direct_cart_cnt,omitempty" codec:"direct_cart_cnt,omitempty"`
	TotalCartCnt            uint64    `json:"total_cart_cnt,omitempty" codec:"total_cart_cnt,omitempty"`
	CurrentWlShowq          uint      `json:"current_wl_showq,omitempty" codec:"current_wl_showq,omitempty"`
	CurrentPcShowq          uint      `json:"current_pc_showq,omitempty" codec:"current_pc_showq,omitempty"`
	CurrentHourPcRank       uint      `json:"current_hour_pc_rank,omitempty" codec:"current_hour_pc_rank,omitempty"`
	AverageCurrentPcRank    uint      `json:"average_current_pc_rank,omitempty" codec:"average_current_pc_rank,omitempty"`
	AverageHistoryWlRank    uint      `json:"average_history_wl_rank,omitempty" codec:"average_history_wl_rank,omitempty"`
	AverageHistoryPcRank    uint      `json:"average_history_pc_rank,omitempty" codec:"average_history_pc_rank,omitempty"`
	AverageHistoryTotalRank uint      `json:"average_history_total_rank,omitempty" codec:"average_history_total_rank,omitempty"`
	EffectOrderSum          float64   `json:"effect_order_sum,omitempty" codec:"effect_order_sum,omitempty"`
	EffectOrderCnt          uint64    `json:"effect_order_cnt,omitempty" codec:"effect_order_cnt,omitempty"`
	EffectCartCnt           uint64    `json:"effect_cart_cnt,omitempty" codec:"effect_cart_cnt,omitempty"`
	IsDefaultPrice          bool      `json:"is_default_price,omitempty" codec:"is_default_price,omitempty"`
	OrderCVS                float64   `json:"order_cvs,omitempty" codec:"order_cvs,omitempty"`
	NewTotalOrderCnt        uint64    `json:"new_total_order_cnt,omitempty" codec:"new_total_order_cnt,omitempty"`
	NewTotalOrderSum        float64   `json:"new_total_order_sum,omitempty" codec:"new_total_order_sum,omitempty"`
	NewTotalOrderCVS        float64   `json:"new_total_order_cvs,omitempty" codec:"new_total_order_cvs,omitempty"`
	KeywordFlag             int       `json:"keyword_flag,omitempty" codec:"keyword_flag,omitempty"`
	MobileType              string    `json:"mobile_type,omitempty" codec:"mobile_type,omitempty"`
	RecordOn                time.Time `json:"-" codec:"-"`
}

func (k *JosKeywordTmp) ToJosKeyword() *JosKeyword {
	kw := &JosKeyword{
		Id:             k.Id,
		Keyword:        k.Keyword,
		KeywordType:    k.KeywordType,
		Price:          k.Price,
		PriceWireless:  k.PriceWireless,
		MobileType:     k.MobileType,
		IsDefaultPrice: k.IsDefaultPrice,
	}
	kw.Clicks, _ = strconv.ParseUint(k.Clicks, 10, 64)
	kw.Cost, _ = strconv.ParseFloat(k.Cost, 64)
	kw.CTR, _ = strconv.ParseFloat(k.CTR, 64)
	kw.CPC, _ = strconv.ParseFloat(k.CPC, 64)
	kw.CPM, _ = strconv.ParseFloat(k.CPM, 64)
	kw.TotalOrderROI, _ = strconv.ParseFloat(k.TotalOrderROI, 64)
	kw.Impressions, _ = strconv.ParseUint(k.Impressions, 10, 64)
	kw.TotalOrderCnt, _ = strconv.ParseUint(k.TotalOrderCnt, 10, 64)
	kw.TotalOrderSum, _ = strconv.ParseFloat(k.TotalOrderSum, 64)
	kw.TotalOrderCVS, _ = strconv.ParseFloat(k.TotalOrderCVS, 64)
	kw.IndirectOrderCnt, _ = strconv.ParseUint(k.IndirectOrderCnt, 10, 64)
	kw.DirectOrderCnt, _ = strconv.ParseUint(k.DirectOrderCnt, 10, 64)
	kw.IndirectOrderSum, _ = strconv.ParseFloat(k.IndirectOrderSum, 64)
	kw.DirectOrderSum, _ = strconv.ParseFloat(k.DirectOrderSum, 64)
	kw.IndirectCartCnt, _ = strconv.ParseUint(k.IndirectCartCnt, 10, 64)
	kw.DirectCartCnt, _ = strconv.ParseUint(k.DirectCartCnt, 10, 64)
	kw.TotalCartCnt, _ = strconv.ParseUint(k.TotalCartCnt, 10, 64)
	vv, err := strconv.ParseUint(k.CurrentWlShowq, 10, 64)
	if err == nil {
		kw.CurrentWlShowq = uint(vv)
	}
	vv, err = strconv.ParseUint(k.CurrentPcShowq, 10, 64)
	if err == nil {
		kw.CurrentPcShowq = uint(vv)
	}
	vv, err = strconv.ParseUint(k.CurrentHourPcRank, 10, 64)
	if err == nil {
		kw.CurrentHourPcRank = uint(vv)
	}
	vv, err = strconv.ParseUint(k.AverageCurrentPcRank, 10, 64)
	if err == nil {
		kw.AverageCurrentPcRank = uint(vv)
	}
	vv, err = strconv.ParseUint(k.AverageHistoryWlRank, 10, 64)
	if err == nil {
		kw.AverageHistoryWlRank = uint(vv)
	}
	vv, err = strconv.ParseUint(k.AverageHistoryPcRank, 10, 64)
	if err == nil {
		kw.AverageHistoryPcRank = uint(vv)
	}
	vv, err = strconv.ParseUint(k.AverageHistoryTotalRank, 10, 64)
	if err == nil {
		kw.AverageHistoryTotalRank = uint(vv)
	}
	kw.EffectOrderSum, _ = strconv.ParseFloat(k.EffectOrderSum, 64)
	kw.EffectOrderCnt, _ = strconv.ParseUint(k.EffectOrderCnt, 10, 64)
	kw.EffectCartCnt, _ = strconv.ParseUint(k.EffectCartCnt, 10, 64)
	kw.OrderCVS, _ = strconv.ParseFloat(k.OrderCVS, 64)
	kw.NewTotalOrderCnt, _ = strconv.ParseUint(k.NewTotalOrderCnt, 10, 64)
	kw.NewTotalOrderSum, _ = strconv.ParseFloat(k.NewTotalOrderSum, 64)
	kw.NewTotalOrderCVS, _ = strconv.ParseFloat(k.NewTotalOrderCVS, 64)
	vvv, err := strconv.ParseInt(k.KeywordFlag, 10, 64)
	if err == nil {
		kw.KeywordFlag = int(vvv)
	}
	return kw
}

// 获取关键词报表
func AdgroupkeywordQuery(req *AdgroupkeywordQueryRequest) ([]*JosKeyword, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := report.NewAdgroupkeywordQueryRequest()

	r.SetStartDate(req.StartDate)
	r.SetEndDate(req.EndDate)
	r.SetGroupId(req.GroupId)
	r.SetClickOrOrderDay(req.ClickOrOrderDay)
	r.SetIsOrderOrClick(req.IsOrderOrClick)
	r.SetPageSize(req.PageSize)
	r.SetPageIndex(req.PageIndex)
	if len(req.Platform) > 0 {
		r.SetPlatform(req.Platform)
	}
	if len(req.ValType) > 0 {
		r.SetValType(req.ValType)
	}
	if len(req.Val) > 0 {
		r.SetVal(req.Val)
	}
	if req.OrderStatusCategory > 0 {
		r.SetOrderStatusCategory(req.OrderStatusCategory)
	}

	var response AdgroupkeywordQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	record := make([]*JosKeyword, 0, len(response.Data.Result.Value.Datas))
	for _, v := range response.Data.Result.Value.Datas {
		record = append(record, v.ToJosKeyword())
	}

	return record, nil
}
