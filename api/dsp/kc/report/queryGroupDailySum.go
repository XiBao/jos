package report

import (
	"encoding/json"
	"errors"
	"sort"
	"strconv"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/report"
)

type QueryGroupDailySumRequest struct {
	api.BaseRequest
	IsDaily             bool   `json:"is_daily,omitempty" codec:"is_daily,omitempty"`                           // 是否分日
	CampaignId          int    `json:"campaign_id,omitempty" codec:"campaign_id,omitempty"`                     // 计划标识
	StartDay            string `json:"start_day,omitempty" codec:"start_day,omitempty"`                         // 起始时间
	EndDay              string `json:"end_day,omitempty" codec:"end_day,omitempty"`                             // 结束时间
	IsOrderOrClick      bool   `json:"is_order_or_click,omitempty" codec:"is_order_or_click,omitempty"`         // 下单点击口径(true:下单口径;flase:点击口径)
	IsTodayOr15Days     bool   `json:"is_today_or_15_days,omitempty" codec:"is_today_or_15_days,omitempty"`     // 是当天15天口径(true:15天;flase:当天)
	OrderStatusCategory int    `json:"order_status_category,omitempty" codec:"order_status_category,omitempty"` // GMV订单类型(空:全部;1:成交订单)
	Platform            string `json:"platform,omitempty" codec:"platform,omitempty"`                           // 推广设备(全部:all;PC:pc;无线:mobile)
	PageIndex           int    `json:"page_index,omitempty" codec:"page_index,omitempty"`                       // 当前页码
	PageSize            int    `json:"page_size,omitempty" codec:"page_size,omitempty"`                         // 每页数量(最大值100)
}

type QueryGroupDailySumResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryGroupDailySumData `json:"jingdong_dsp_report_queryGroupDailySum_responce,omitempty" codec:"jingdong_dsp_report_queryGroupDailySum_responce,omitempty"`
}

type QueryGroupDailySumData struct {
	Result *QueryGroupDailySumResult `json:"querycampdailysum_result,omitempty" codec:"querycampdailysum_result,omitempty"`
}

type QueryGroupDailySumResult struct {
	Success    bool          `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string        `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string        `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Total      int           `json:"total,omitempty" codec:"total,omitempty"`
	Page       int           `json:"page,omitempty" codec:"page,omitempty"`
	Value      []*ReportInfo `json:"value,omitempty" codec:"value,omitempty"`
}
type GroupDailyRpt struct {
	Date                              time.Time `json:"date,omitempty" codec:"date,omitempty"`
	GroupId                           uint64    `json:"group_id,omitempty" codec:"group_id,omitempty"`
	GroupName                         string    `json:"group_name,omitempty" codec:"group_name,omitempty"`
	CampaignId                        uint64    `json:"campaign_id,omitempty" codec:"campaign_id,omitempty"`
	CampaignName                      string    `json:"campaign_name,omitempty" codec:"campaign_name,omitempty"`
	CPC                               float64   `json:"cpc,omitempty" codec:"cpc,omitempty"`
	Clicks                            uint64    `json:"clicks,omitempty" codec:"clicks,omitempty"`
	Impressions                       uint64    `json:"impressions,omitempty" codec:"impressions,omitempty"`
	Cost                              float64   `json:"cost,omitempty" codec:"cost,omitempty"`
	CPM                               float64   `json:"cpm,omitempty" codec:"cpm,omitempty"`
	CTR                               float64   `json:"ctr,omitempty" codec:"ctr,omitempty"`
	CouponCnt                         uint64    `json:"coupon_cnt,omitempty" codec:"coupon_cnt,omitempty"`
	DirectOrderSum                    float64   `json:"direct_order_sum,omitempty" codec:"direct_order_sum,omitempty"`
	IndirectOrderSum                  float64   `json:"indirect_order_sum,omitempty" codec:"indirect_order_sum,omitempty"`
	TotalOrderSum                     float64   `json:"total_order_sum,omitempty" codec:"total_order_sum,omitempty"`
	DirectOrderCnt                    uint64    `json:"direct_order_cnt,omitempty" codec:"direct_order_cnt,omitempty"`
	IndirectOrderCnt                  uint64    `json:"indirect_order_cnt,omitempty" codec:"indirect_order_cnt,omitempty"`
	TotalOrderCnt                     uint64    `json:"total_order_cnt,omitempty" codec:"total_order_cnt,omitempty"`
	TotalOrderROI                     float64   `json:"total_order_roi,omitempty" codec:"total_order_roi,omitempty"`
	TotalOrderCVS                     float64   `json:"total_order_cvs,omitempty" codec:"total_order_cvs,omitempty"`
	OrderROI                          float64   `json:"order_roi,omitempty" codec:"order_roi,omitempty"`
	OrderCVS                          float64   `json:"order_cvs,omitempty" codec:"order_cvs,omitempty"`
	NewTotalOrderCnt                  uint64    `json:"new_total_order_cnt,omitempty" codec:"new_total_order_cnt,omitempty"`
	NewTotalOrderSum                  float64   `json:"new_total_order_sum,omitempty" codec:"new_total_order_sum,omitempty"`
	PreorderCnt                       uint64    `json:"preorder_cnt,omitempty" codec:"preorder_cnt,omitempty"`
	OrderDate                         time.Time `json:"order_date,omitempty" codec:"order_date,omitempty"`
	DirectCartCnt                     uint64    `json:"direct_cart_cnt,omitempty" codec:"direct_cart_cnt,omitempty"`
	IndirectCartCnt                   uint64    `json:"indirect_cart_cnt,omitempty" codec:"indirect_cart_cnt,omitempty"`
	TotalCartCnt                      uint64    `json:"total_cart_cnt,omitempty" codec:"total_cart_cnt,omitempty"`
	TotalCartCntNCX                   uint64    `json:"total_cart_cnt_ncx,omitempty" codec:"total_cart_cnt_ncx,omitempty"`
	PlatformCnt                       uint64    `json:"platform_cnt,omitempty" codec:"platform_cnt,omitempty"`
	PlatformGmv                       float64   `json:"platform_gmv,omitempty" codec:"platform_gmv,omitempty"`
	DepthPassengerCnt                 uint64    `json:"depth_passenger_cnt,omitempty" codec:"depth_passenger_cnt,omitempty"`
	DepartmentCnt                     uint64    `json:"department_cnt,omitempty" codec:"department_cnt,omitempty"`
	DepartmentGmv                     float64   `json:"department_gmv,omitempty" codec:"department_gmv,omitempty"`
	MobileType                        string    `json:"mobile_type,omitempty" codec:"mobile_type,omitempty"`
	ChannelROI                        float64   `json:"channel_roi,omitempty" codec:"channel_roi,omitempty"`
	VisitTimeRange                    float64   `json:"visit_time_range,omitempty" codec:"visit_time_range,omitempty"`
	VisitPageCnt                      uint64    `json:"visit_page_cnt,omitempty" codec:"visit_page_cnt,omitempty"`
	ShopAttentionCnt                  uint64    `json:"shop_attention_cnt,omitempty" codec:"shop_attention_cnt,omitempty"`
	NewCustomersCnt                   uint64    `json:"new_customers_cnt,omitempty" codec:"new_customers_cnt,omitempty"`
	OrderStatusCategory               string    `json:"order_status_category,omitempty" codec:"order_status_category,omitempty"`
	GoodsAttentionCnt                 uint64    `json:"goods_attention_cnt,omitempty" codec:"goods_attention_cnt,omitempty"`
	GroupVisitorCntForInternalSummary uint64    `json:"group_visitor_cnt_for_internal_summary,omitempty" codec:"group_visitor_cnt_for_internal_summary,omitempty"`
	RetrievalType                     int       `json:"retrieval_type,omitempty" codec:"retrieval_type,omitempty"`
}

// 查询.快车.单元报表数据
func QueryGroupDailySum(req *QueryGroupDailySumRequest) ([]*GroupDailyRpt, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := report.NewQueryGroupDailySumRequest()

	r.SetIsDaily(req.IsDaily)
	if req.CampaignId > 0 {
		r.SetCampaignId(req.CampaignId)
	}
	// yyyy-MM-dd HH:mm:ss
	r.SetStartDay(req.StartDay)
	r.SetEndDay(req.EndDay)
	r.SetIsOrderOrClick(req.IsOrderOrClick)
	r.SetIsTodayOr15Days(req.IsTodayOr15Days)
	if req.OrderStatusCategory > 0 {
		r.SetOrderStatusCategory(req.OrderStatusCategory)
	}
	if req.Platform == "" {
		r.SetPlatform("all")
	} else {
		r.SetPlatform(req.Platform)
	}
	r.SetPageIndex(req.PageIndex)
	r.SetPageSize(req.PageSize)

	result, err := client.Execute(r.Request, req.Session)

	if err != nil {
		return nil, 0, err
	}
	if len(result) == 0 {
		return nil, 0, errors.New("no result info")
	}

	var response QueryGroupDailySumResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, 0, err
	}
	if response.ErrorResp != nil {
		return nil, 0, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return nil, 0, errors.New(response.Data.Result.ErrorMsg)
	}

	var rpts []*GroupDailyRpt

	loc := time.Now().Location()
	for _, data := range response.Data.Result.Value {
		rpt := &GroupDailyRpt{
			GroupId: uint64(data.Id),
		}
		for k, v := range data.FigureData {
			switch k {
			case "GroupName":
				rpt.GroupName = v
			case "r0_GroupVisitorCntForInternalSummary":
				rpt.GroupVisitorCntForInternalSummary, _ = strconv.ParseUint(v, 10, 64)
			case "r0_RetrievalType":
				val, err := strconv.ParseInt(v, 10, 64)
				if err == nil {
					rpt.RetrievalType = int(val)
				}
			case "r0_TotalCartCntNCX":
				rpt.TotalCartCntNCX, _ = strconv.ParseUint(v, 10, 64)
			case "r0_CPM":
				rpt.CPM, _ = strconv.ParseFloat(v, 64)
			case "r0_DepartmentGmv":
				rpt.DepartmentGmv, _ = strconv.ParseFloat(v, 64)
			case "r0_DepartmentCnt":
				rpt.DepartmentCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_DepthPassengerCnt":
				rpt.DepthPassengerCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_CouponCnt":
				rpt.CouponCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_IndirectOrderCnt":
				rpt.IndirectOrderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_OrderROI":
				rpt.OrderROI, _ = strconv.ParseFloat(v, 64)
			case "r0_PreorderCnt":
				rpt.PreorderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_OrderCVS":
				rpt.OrderCVS, _ = strconv.ParseFloat(v, 64)
			case "date":
				rpt.Date, _ = time.ParseInLocation("20060102", v, loc)
			case "MobileType":
				rpt.MobileType = v
			case "r0_ChannelROI":
				rpt.ChannelROI, _ = strconv.ParseFloat(v, 64)
			case "r0_Clicks":
				rpt.Clicks, _ = strconv.ParseUint(v, 10, 64)
			case "r0_TotalOrderCnt":
				rpt.TotalOrderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_VisitTimeRange":
				rpt.VisitTimeRange, _ = strconv.ParseFloat(v, 64)
			case "r0_VisitPageCnt":
				rpt.VisitPageCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_NewTotalOrderCnt":
				rpt.NewTotalOrderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_ShopAttentionCnt":
				rpt.ShopAttentionCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_IndirectCartCnt":
				rpt.IndirectCartCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_DirectOrderSum":
				rpt.DirectOrderSum, _ = strconv.ParseFloat(v, 64)
			case "r0_NewTotalOrderSum":
				rpt.NewTotalOrderSum, _ = strconv.ParseFloat(v, 64)
			case "r0_OrderDate":
				rpt.OrderDate, _ = time.ParseInLocation("20060102", v, loc)
			case "r0_TotalOrderSum":
				rpt.TotalOrderSum, _ = strconv.ParseFloat(v, 64)
			case "r0_NewCustomersCnt":
				rpt.NewCustomersCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_Impressions":
				rpt.Impressions, _ = strconv.ParseUint(v, 10, 64)
			case "r0_IndirectOrderSum":
				rpt.IndirectOrderSum, _ = strconv.ParseFloat(v, 64)
			case "r0_PlatformCnt":
				rpt.PlatformCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_DirectCartCnt":
				rpt.DirectCartCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_TotalOrderROI":
				rpt.TotalOrderROI, _ = strconv.ParseFloat(v, 64)
			case "CampaignId":
				rpt.CampaignId, _ = strconv.ParseUint(v, 10, 64)
			case "r0_Cost":
				rpt.Cost, _ = strconv.ParseFloat(v, 64)
			case "r0_PlatformGmv":
				rpt.PlatformGmv, _ = strconv.ParseFloat(v, 64)
			case "orderStatusCategory":
				rpt.OrderStatusCategory = v
			case "CampaignName":
				rpt.CampaignName = v
			case "r0_TotalCartCnt":
				rpt.TotalCartCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_DirectOrderCnt":
				rpt.DirectOrderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_CPC":
				rpt.CPC, _ = strconv.ParseFloat(v, 64)
			case "r0_GoodsAttentionCnt":
				rpt.GoodsAttentionCnt, _ = strconv.ParseUint(v, 10, 64)
			case "r0_CTR":
				rpt.CTR, _ = strconv.ParseFloat(v, 64)
			case "r0_TotalOrderCVS":
				rpt.TotalOrderCVS, _ = strconv.ParseFloat(v, 64)
			}
		}
		rpts = append(rpts, rpt)
	}
	if req.IsDaily {
		sort.Slice(rpts, func(i, j int) bool {
			return rpts[i].Date.Before(rpts[j].Date)
		})

	}
	return rpts, response.Data.Result.Total, nil

}
