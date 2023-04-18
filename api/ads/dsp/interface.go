package dsp

type DataCommonResponse struct {
	Msg     string                             `json:"msg,omitempty"`
	Code    int                                `json:"code"`
	Success bool                               `json:"success"`
	System  *JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty"`
}

type Paginator struct {
	PageNum  int `json:"pageNum"`  // 当前页码
	PageSize int `json:"pageSize"` // 每页项数
	Items    int `json:"items"`    // 总共项数
}

type JdDspPlatformGatewayApiVoParamExt struct {
	Pin         string `json:"pin"`         // 用户pin
	VenderId    string `json:"venderId"`    // 商家ID
	RequestFrom string `json:"requestFrom"` // 业务来源
	AppKey      string `json:"appKey"`      // appKey
	TraceId     string `json:"traceId"`     // traceId
}

type Campaign struct {
	PutType            uint    `json:"putType"`            // 投放类型
	DateRange          string  `json:"dateRange"`          // 自定义日预算
	TimeRangePriceCoef string  `json:"timeRangePriceCoef"` // 投放折扣系数
	DayBudgetCustom    float64 `json:"dayBudgetCustom"`    // 统一预算
	StartTime          uint64  `json:"startTime"`          // 计划投放开始时间
	DayBudget          float64 `json:"dayBudget"`          // 计划预算
	Id                 uint64  `json:"id"`                 // 计划id
	CampaignType       uint    `json:"campaignType"`       // 计划类型
	Name               string  `json:"name"`               // 计划名称
	EndTime            uint64  `json:"endTime"`            // 计划投放结束时间
	Status             uint    `json:"status"`             // 计划状态1暂停2有效
}

type CampaignData struct {
	CampaignId         uint64 `json:"campaignId"`         // 计划id
	CampaignName       string `json:"campaignName"`       // 计划名称
	CampaignType       uint   `json:"campaignType"`       // 计划类型
	ClickDate          string `json:"clickDate"`          // 点击时间
	CPA                string `json:"CPA"`                // CPA
	CPC                string `json:"CPC"`                // CPC
	CPM                string `json:"CPM"`                // CPM
	CTR                string `json:"CTR"`                // CTR
	ActivityId         uint64 `json:"activityId"`         // 联合活动id
	Clicks             uint   `json:"clicks"`             // clicks
	Cost               string `json:"cost"`               // cost
	DirectCartCnt      uint   `json:"directCartCnt"`      // 直接加购数
	DirectOrderCnt     uint   `json:"directOrderCnt"`     // 直接订单数
	DirectOrderSum     string `json:"directOrderSum"`     // 直接订单金额
	EffectCartCnt      uint   `json:"effectCartCnt"`      // 影响加购数
	EffectOrderCnt     uint   `json:"effectOrderCnt"`     // 影响订单数
	EffectOrderSum     string `json:"effectOrderSum"`     // 影响订单金额
	Impressions        uint   `json:"impressions"`        // 展现数
	IndirectCartCnt    uint   `json:"indirectCartCnt"`    // 间接加购数
	IndirectOrderCnt   uint   `json:"indirectOrderCnt"`   // 间接订单数
	IndirectOrderSum   string `json:"indirectOrderSum"`   // 间接订单金额
	PutType            uint   `json:"putType"`            // 投放类型
	Status             uint   `json:"status"`             // 计划状态
	TimeRange          string `json:"timeRange"`          // 投放时段
	TimeRangePriceCoef string `json:"timeRangePriceCoef"` // 投放折扣系数
	TotalCartCnt       uint   `json:"totalCartCnt"`       // 总加购数
	TotalOrderCVS      string `json:"totalOrderCVS"`      // 转化率
	TotalOrderCnt      uint   `json:"totalOrderCnt"`      // 总订单数
	TotalOrderROI      string `json:"totalOrderROI"`      // ROI
	TotalOrderSum      string `json:"totalOrderSum"`      // 总订单金额
}
