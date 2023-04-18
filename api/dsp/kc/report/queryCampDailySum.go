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

type QueryCampDailySumRequest struct {
	api.BaseRequest
	IsDaily             bool   `json:"is_daily,omitempty" codec:"is_daily,omitempty"`                           //是否分日
	CampaignId          int    `json:"campaign_id,omitempty" codec:"campaign_id,omitempty"`                     //计划标识
	StartDay            string `json:"start_day,omitempty" codec:"start_day,omitempty"`                         //起始时间
	EndDay              string `json:"end_day,omitempty" codec:"end_day,omitempty"`                             //结束时间
	IsOrderOrClick      bool   `json:"is_order_or_click,omitempty" codec:"is_order_or_click,omitempty"`         //下单点击口径(true:下单口径;flase:点击口径)
	IsTodayOr15Days     bool   `json:"is_today_or_15_days,omitempty" codec:"is_today_or_15_days,omitempty"`     //是当天15天口径(true:15天;flase:当天)
	OrderStatusCategory int    `json:"order_status_category,omitempty" codec:"order_status_category,omitempty"` //GMV订单类型(空:全部;1:成交订单)
	Platform            string `json:"platform,omitempty" codec:"platform,omitempty"`                           //推广设备(全部:all;PC:pc;无线:mobile)
	PageIndex           int    `json:"page_index,omitempty" codec:"page_index,omitempty"`                       //当前页码
	PageSize            int    `json:"page_size,omitempty" codec:"page_size,omitempty"`                         //每页数量(最大值100)
}

type QueryCampDailySumResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryCampDailySumData `json:"jingdong_dsp_report_queryCampDailySum_responce,omitempty" codec:"jingdong_dsp_report_queryCampDailySum_responce,omitempty"`
}

type QueryCampDailySumData struct {
	Result *QueryCampDailySumResult `json:"querycampdailysum_result,omitempty" codec:"querycampdailysum_result,omitempty"`
}

type QueryCampDailySumResult struct {
	Success    bool          `json:"success,omitempty" codec:"success,omitempty"`
	ResultCode string        `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string        `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Total      int           `json:"total,omitempty" codec:"total,omitempty"`
	Page       int           `json:"page,omitempty" codec:"page,omitempty"`
	Value      []*ReportInfo `json:"value,omitempty" codec:"value,omitempty"`
}

type ReportInfo struct {
	Id         int               `json:"id,omitempty" codec:"id,omitempty"`
	Pin        string            `json:"pin,omitempty" codec:"pin,omitempty"`
	Dimension  string            `json:"dimension,omitempty" codec:"dimension,omitempty"`
	Date       string            `json:"date,omitempty" codec:"date,omitempty"`
	FigureData map[string]string `json:"figureData,omitempty" codec:"figureData,omitempty"`
}

type CampaignDailyRpt struct {
	Date                time.Time `json:"date,omitempty" codec:"date,omitempty"`
	CampaignId          uint64    `json:"campaign_id,omitempty" codec:"campaign_id,omitempty"`                 // 计划id
	CampaignName        string    `json:"campaign_name,omitempty" codec:"campaign_name,omitempty"`             // 计划名称
	CampaignStatus      int       `json:"campaign_status,omitempty" codec:"campaign_status,omitempty"`         //计划状态
	CPC                 float64   `json:"cpc,omitempty" codec:"cpc,omitempty"`                                 //平均点击成本
	Clicks              uint64    `json:"clicks,omitempty" codec:"clicks,omitempty"`                           //点击量
	Impressions         uint64    `json:"impressions,omitempty" codec:"impressions,omitempty"`                 //展现数
	Cost                float64   `json:"cost,omitempty" codec:"cost,omitempty"`                               //总费用
	CPM                 float64   `json:"cpm,omitempty" codec:"cpm,omitempty"`                                 //千次展现成本
	CTR                 float64   `json:"ctr,omitempty" codec:"ctr,omitempty"`                                 //点击率
	CouponCnt           uint64    `json:"coupon_cnt,omitempty" codec:"coupon_cnt,omitempty"`                   //领劵数
	DirectOrderSum      float64   `json:"direct_order_sum,omitempty" codec:"direct_order_sum,omitempty"`       //	直接订单金额
	IndirectOrderSum    float64   `json:"indirect_order_sum,omitempty" codec:"indirect_order_sum,omitempty"`   //间接订单金额
	TotalOrderSum       float64   `json:"total_order_sum,omitempty" codec:"total_order_sum,omitempty"`         //总订单金额
	DirectOrderCnt      uint64    `json:"direct_order_cnt,omitempty" codec:"direct_order_cnt,omitempty"`       //直接订单行
	IndirectOrderCnt    uint64    `json:"indirect_order_cnt,omitempty" codec:"indirect_order_cnt,omitempty"`   //	间接订单行
	TotalOrderCnt       uint64    `json:"total_order_cnt,omitempty" codec:"total_order_cnt,omitempty"`         //总订单行
	TotalOrderROI       float64   `json:"total_order_roi,omitempty" codec:"total_order_roi,omitempty"`         //
	OrderROI            float64   `json:"order_roi,omitempty" codec:"order_roi,omitempty"`                     //ROI
	OrderCVS            float64   `json:"order_cvs,omitempty" codec:"order_cvs,omitempty"`                     //转换率
	PreorderCnt         uint64    `json:"preorder_cnt,omitempty" codec:"preorder_cnt,omitempty"`               //预约数
	OrderDate           time.Time `json:"order_date,omitempty" codec:"order_date,omitempty"`                   //
	DirectCartCnt       uint64    `json:"direct_cart_cnt,omitempty" codec:"direct_cart_cnt,omitempty"`         //直接加购数
	IndirectCartCnt     uint64    `json:"indirect_cart_cnt,omitempty" codec:"indirect_cart_cnt,omitempty"`     //间接加购数
	TotalCartCnt        uint64    `json:"total_cart_cnt,omitempty" codec:"total_cart_cnt,omitempty"`           //总加购数
	TotalCartCntNCX     uint64    `json:"total_cart_cnt_ncx,omitempty" codec:"total_cart_cnt_ncx,omitempty"`   //
	PlatformCnt         uint64    `json:"platform_cnt,omitempty" codec:"platform_cnt,omitempty"`               //
	PlatformGmv         float64   `json:"platform_gmv,omitempty" codec:"platform_gmv,omitempty"`               //
	DepthPassengerCnt   uint64    `json:"depth_passenger_cnt,omitempty" codec:"depth_passenger_cnt,omitempty"` //深度进店数
	DepartmentCnt       uint64    `json:"department_cnt,omitempty" codec:"department_cnt,omitempty"`           //
	DepartmentGmv       float64   `json:"department_gmv,omitempty" codec:"department_gmv,omitempty"`           //
	MobileType          string    `json:"mobile_type,omitempty" codec:"mobile_type,omitempty"`                 //
	ChannelROI          float64   `json:"channel_roi,omitempty" codec:"channel_roi,omitempty"`                 //
	VisitTimeRange      float64   `json:"visit_time_range,omitempty" codec:"visit_time_range,omitempty"`       //访问时长
	VisitPageCnt        uint64    `json:"visit_page_cnt,omitempty" codec:"visit_page_cnt,omitempty"`           //访问页面数
	ShopAttentionCnt    uint64    `json:"shop_attention_cnt,omitempty" codec:"shop_attention_cnt,omitempty"`   //店铺关注数
	IsOrderOrClick      string    `json:"is_order_or_click,omitempty" codec:"is_order_or_click,omitempty"`
	IsTodayOr15Days     string    `json:"is_today_or_15days,omitempty" codec:"is_today_or_15days,omitempty"`
	PutType             string    `json:"put_type,omitempty" codec:"put_type,omitempty"`
	NewCustomersCnt     uint64    `json:"new_customers_cnt,omitempty" codec:"new_customers_cnt,omitempty"`       //下单新客数
	CampaignVisitorCnt  uint64    `json:"campaign_visitor_cnt,omitempty" codec:"campaign_visitor_cnt,omitempty"` //访问数
	AdType              int       `json:"ad_type,omitempty" codec:"ad_type,omitempty"`
	OrderStatusCategory string    `json:"order_status_category,omitempty" codec:"order_status_category,omitempty"` //
	GoodsAttentionCnt   uint64    `json:"goods_attention_cnt,omitempty" codec:"goods_attention_cnt,omitempty"`     //商品关注数
}

// 查询.快车.计划报表数据
func QueryCampDailySum(req *QueryCampDailySumRequest) ([]*CampaignDailyRpt, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := report.NewQueryCampDailySumRequest()

	r.SetIsDaily(req.IsDaily)
	if req.CampaignId > 0 {
		r.SetCampaignId(req.CampaignId)
	}
	//yyyy-MM-dd HH:mm:ss
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
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}

	var response QueryCampDailySumResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return nil, errors.New(response.Data.Result.ErrorMsg)
	}

	loc := time.Now().Location()
	rpts := []*CampaignDailyRpt{}
	for _, data := range response.Data.Result.Value {
		rpt := &CampaignDailyRpt{}
		for k, v := range data.FigureData {
			switch k {
			case "TotalCartCntNCX":
				rpt.TotalCartCntNCX, _ = strconv.ParseUint(v, 10, 64)
			case "CPM":
				rpt.CPM, _ = strconv.ParseFloat(v, 64)
			case "DepthPassengerCnt":
				rpt.DepthPassengerCnt, _ = strconv.ParseUint(v, 10, 64)
			case "DepartmentGmv":
				rpt.DepartmentGmv, _ = strconv.ParseFloat(v, 64)
			case "CouponCnt":
				rpt.CouponCnt, _ = strconv.ParseUint(v, 10, 64)
			case "IndirectOrderCnt":
				rpt.IndirectOrderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "OrderROI":
				rpt.OrderROI, _ = strconv.ParseFloat(v, 64)
			case "PreorderCnt":
				rpt.PreorderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "OrderCVS":
				rpt.OrderCVS, _ = strconv.ParseFloat(v, 64)
			case "date":
				rpt.Date, _ = time.ParseInLocation("20060102", v, loc)
			case "MobileType":
				rpt.MobileType = v
			case "ChannelROI":
				rpt.ChannelROI, _ = strconv.ParseFloat(v, 64)
			case "Clicks":
				rpt.Clicks, _ = strconv.ParseUint(v, 10, 64)
			case "TotalOrderCnt":
				rpt.TotalOrderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "VisitTimeRange":
				rpt.VisitTimeRange, _ = strconv.ParseFloat(v, 64)
			case "VisitPageCnt":
				rpt.VisitPageCnt, _ = strconv.ParseUint(v, 10, 64)
			case "ShopAttentionCnt":
				rpt.ShopAttentionCnt, _ = strconv.ParseUint(v, 10, 64)
			case "IndirectCartCnt":
				rpt.IndirectCartCnt, _ = strconv.ParseUint(v, 10, 64)
			case "isOrderOrClick":
				rpt.IsOrderOrClick = v
			case "DirectOrderSum":
				rpt.DirectOrderSum, _ = strconv.ParseFloat(v, 64)
			case "OrderDate":
				rpt.OrderDate, _ = time.ParseInLocation("20060102", v, loc)
			case "isTodayOr15Days":
				rpt.IsTodayOr15Days = v
			case "TotalOrderSum":
				rpt.TotalOrderSum, _ = strconv.ParseFloat(v, 64)
			case "putType":
				rpt.PutType = v
			case "NewCustomersCnt":
				rpt.NewCustomersCnt, _ = strconv.ParseUint(v, 10, 64)
			case "Impressions":
				rpt.Impressions, _ = strconv.ParseUint(v, 10, 64)
			case "IndirectOrderSum":
				rpt.IndirectOrderSum, _ = strconv.ParseFloat(v, 64)
			case "CampaignVisitorCnt":
				rpt.CampaignVisitorCnt, _ = strconv.ParseUint(v, 10, 64)
			case "AdType":
				val, err := strconv.ParseInt(v, 10, 64)
				if err == nil {
					rpt.AdType = int(val)
				}
			case "PlatformCnt":
				rpt.PlatformCnt, _ = strconv.ParseUint(v, 10, 64)
			case "DirectCartCnt":
				rpt.DirectCartCnt, _ = strconv.ParseUint(v, 10, 64)
			case "TotalOrderROI":
				rpt.TotalOrderROI, _ = strconv.ParseFloat(v, 64)
			case "CampaignId":
				rpt.CampaignId, _ = strconv.ParseUint(v, 10, 64)
			case "Cost":
				rpt.Cost, _ = strconv.ParseFloat(v, 64)
			case "PlatformGmv":
				rpt.PlatformGmv, _ = strconv.ParseFloat(v, 64)
			case "orderStatusCategory":
				rpt.OrderStatusCategory = v
			case "CampaignStatus":
				val, err := strconv.ParseInt(v, 10, 64)
				if err == nil {
					rpt.CampaignStatus = int(val)
				}
			case "CampaignName":
				rpt.CampaignName = v
			case "DepartmentCnt":
				rpt.DepartmentCnt, _ = strconv.ParseUint(v, 10, 64)
			case "TotalCartCnt":
				rpt.TotalCartCnt, _ = strconv.ParseUint(v, 10, 64)
			case "DirectOrderCnt":
				rpt.DirectOrderCnt, _ = strconv.ParseUint(v, 10, 64)
			case "CPC":
				rpt.CPC, _ = strconv.ParseFloat(v, 64)
			case "GoodsAttentionCnt":
				rpt.GoodsAttentionCnt, _ = strconv.ParseUint(v, 10, 64)
			case "CTR":
				rpt.CTR, _ = strconv.ParseFloat(v, 64)

			}
		}
		rpts = append(rpts, rpt)
	}
	if req.IsDaily {
		sort.Slice(rpts, func(i, j int) bool {
			return rpts[i].Date.Before(rpts[j].Date)
		})
	}

	return rpts, nil

}
