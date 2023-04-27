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

// 计划详情
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

// 单元详情
type Adgroup struct {
	Id                   uint64   `json:"id"`                   // 单元id
	CampaignType         uint     `json:"campaignType"`         // 计划类型，2:普通快车，18:腰带店铺
	PutType              uint     `json:"putType"`              // 投放类型，3:商品，4:活动
	CampaignId           uint64   `json:"campaignId"`           // 计划id
	Name                 string   `json:"name"`                 // 单元名称
	Pin                  string   `json:"pin"`                  // 店铺PIN
	NewAreaIds           []string `json:"newAreaIds"`           // 地域ids
	Status               uint     `json:"status"`               // 状态,1:暂停，2:有效
	MobilePriceCoef      float64  `json:"mobilePriceCoef"`      // 单元无线出价系数
	InSearchFee          float64  `json:"inSearchFee"`          // 单元搜索出价
	RecommendFee         float64  `json:"recommendFee"`         // 单元推荐出价
	ShopId               uint64   `json:"shopId"`               // 店铺id
	AdOptimize           uint     `json:"adOptimize"`           // 创意优选设置创意优选开关 1打开;0关闭
	AutomatedBiddingType uint     `json:"automatedBiddingType"` // 出价控制方式，2：系统智能调价（投放目标为自定义的系统智能托管） 256：预算控制(投方目标可选择点击，下单，成交) 若使用智能出价必填，腰带店铺填写0：手动出价
	DeliveryTarget       uint     `json:"deliveryTarget"`       // 投放目标： 1, 自定义, 2, 点击, 3,下单, 4, 成交，选择智能出价控制必须填写
	DayCost              uint     `json:"dayCost"`              // 统一日消耗（范围50-计划设置的日预算值）,出价控制方式为预算控制必须填写
	TcpaMaxClickBid      uint     `json:"tcpaMaxClickBid"`      // CPC上限(范围 2-9999)，选择预算控制非系统托管需要填写
	PremiumCoef          uint     `json:"premiumCoef"`          // 溢价系数，投放目标为自定义，出价控制方式为系统智能调价选择自定义上限必填 30-300
	BiddingControlType   uint     `json:"biddingControlType"`   // 控制方式，1系统托管，3指定上限，若指定上限需和tcpaMaxClickBid配合使用
	OrientationRange     uint     `json:"orientationRange"`     // 智能出价定向范围，投放目标为自定义，出价控制方式为智能调价必填，1：关键词定向，2：商品定向，3：关键词+商品定向
}

type RetrievalData struct {
	CPA              string `json:"CPA"`              // CPA
	CTR              string `json:"CTR"`              // CTR
	CPM              string `json:"CPM"`              // CPM
	TcpaStatus       uint   `json:"tcpaStatus"`       // tcpa状态
	IndirectOrderCnt uint   `json:"indirectOrderCnt"` // 间接订单数
	DirectOrderCnt   uint   `json:"directOrderCnt"`   // 直接订单数
	IndirectCartCnt  uint   `json:"indirectCartCnt"`  // 间接加购数
	DirectCartCnt    uint   `json:"directCartCnt"`    // 直接加购数
	Cost             string `json:"cost"`             // cost
	TotalOrderSum    string `json:"totalOrderSum"`    // 总订单金额
	TotalCartCnt     uint   `json:"totalCartCnt"`     // 总加购数
	TotalOrderROI    string `json:"totalOrderROI"`    // ROI
	Impressions      uint   `json:"impressions"`      // 展现数
	IndirectOrderSum string `json:"indirectOrderSum"` // 间接订单金额
	DirectOrderSum   string `json:"directOrderSum"`   // 直接订单金额
	TotalOrderCVS    string `json:"totalOrderCVS"`    // 转化率
	CPC              string `json:"CPC"`              // CPC
	Clicks           uint   `json:"clicks"`           // clicks
	TotalOrderCnt    uint   `json:"totalOrderCnt"`    // 总订单数
}

type AdgroupData struct {
	CampaignId       uint64         `json:"campaignId"`       // 计划id
	CampaignName     string         `json:"campaignName"`     // 计划名称
	CampaignType     uint           `json:"campaignType"`     // 计划类型，2:普通快车，18:腰带店铺
	ActivityId       uint64         `json:"activityId"`       // 联合活动id
	NewAreaId        string         `json:"newAreaId"`        // 地域id
	PutType          uint           `json:"putType"`          // 投放类型，3:商品，4:活动
	Status           uint           `json:"status"`           // 状态,1:暂停，2:有效
	GroupId          uint64         `json:"groupId"`          // 单元id
	GroupName        string         `json:"groupName"`        // 单元名称
	Device           uint           `json:"device"`           // 设备类型，1:pc,2:无线
	RecommendFee     float64        `json:"recommendFee"`     // 单元pc推荐出价
	SearchFee        float64        `json:"searchFee"`        // 单元pc搜索出价
	RetrievalType0   *RetrievalData `json:"retrievalType0"`   // 汇总展点消
	RetrievalType1   *RetrievalData `json:"retrievalType1"`   // 人群展点消
	RetrievalType2   *RetrievalData `json:"retrievalType2"`   // 关键词展点消
	RetrievalType3   *RetrievalData `json:"retrievalType3"`   // 商品展点消
	CrowdValidStatus uint           `json:"crowdValidStatus"` // 人群有效状态 0:未过期 1：已过期 2:即将过期
}

type CreativeAuditInfoData struct {
	MediaName string `json:"mediaName"` // 审核媒体
	Status    int    `json:"status"`    // 审核状态2有效 -2驳回
	AuditInfo string `json:"auditInfo"` // 审核信息
	FailUrl   string `json:"failUrl"`   // 驳回信息图片
	AuditTime string `json:"auditTime"` // 审核时间
}

type CreativeData struct {
	CampaignId     uint64                  `json:"campaignId"`     // 计划id
	CampaignName   string                  `json:"campaignName"`   // 计划名称
	CampaignType   uint                    `json:"campaignType"`   // 计划类型
	ActivityId     uint64                  `json:"activityId"`     // 联合活动id
	NewAreaId      string                  `json:"newAreaId"`      // 地域id
	PutType        uint                    `json:"putType"`        // 投放类型
	Status         uint                    `json:"status"`         // 计划状态
	GrouId         uint64                  `json:"groupId"`        // 投放类型
	GroupName      string                  `json:"groupName"`      // 单元名称
	Device         uint                    `json:"device"`         // 设备类型
	RecommendFee   float64                 `json:"recommendFee"`   // 单元pc推荐出价
	SearchFee      float64                 `json:"searchFee"`      // 单元pc搜索出价
	RetrievalType0 *RetrievalData          `json:"retrievalType0"` // 汇总展点消
	RetrievalType1 *RetrievalData          `json:"retrievalType1"` // 人群展点消
	RetrievalType2 *RetrievalData          `json:"retrievalType2"` // 关键词展点消
	RetrievalType3 *RetrievalData          `json:"retrievalType3"` // 商品展点消
	AdId           uint64                  `json:"adId"`           // 创意id
	AdName         string                  `json:"adName"`         // 创意名称
	CustomTitle    string                  `json:"customTitle"`    // 创意标题
	ImgUrl         string                  `json:"imgUrl"`         // 创意图片地址
	SizeStr        string                  `json:"sizeStr"`        // 图片尺寸
	SkuId          uint64                  `json:"skuId"`          // skuId
	Url            string                  `json:"url"`            // 落地页
	AuditInfoList  []CreativeAuditInfoData `json:"auditInfoList"`  // 创意审核信息列表
	SkuState       uint                    `json:"skuState"`       // 商品上架状态，非1为正常上架状态，1为不在架状态
	AdCreativeType uint                    `json:"adCreativeType"` // 0:无创意类型,11：智能创意，18：自定义创意，19：默认创意
}

type KeywordQueryData struct {
	KeywordName             string  `json:"keywordName"`             // 关键词名称
	KeywordMobilePrice      float64 `json:"keywordMobilePrice"`      // 关键词无线出价
	KeywordPrice            float64 `json:"keywordPrice"`            // 关键词出价
	Type                    uint    `json:"type"`                    // 关键词类型
	KeywordId               uint64  `json:"keywordId"`               // 关键词id
	SearchPromoteRankEnable uint    `json:"searchPromoteRankEnable"` // 抢位助手是否开启 0不开启 1开启
}

type KeywordRecommend struct {
	KeyWord             string  `json:"keyWord"`              // 关键词名称
	SourceType          uint    `json:"sourceType,emitempty"` // 关键词标签，热词：12或者20，潜力词：72或者80，热+潜：76或者84，无标签：8或者16
	Pv                  uint    `json:"pv"`                   // 搜索量
	AvgBigPrice         float64 `json:"avgBigPrice"`          // 平均出价
	Ctr                 float64 `json:"ctr"`                  // 点击率
	Cvr                 float64 `json:"cvr"`                  // 点击转化率，-1代表不置信，对应pc页面的“-”
	StarCount           uint    `json:"starCount"`            // 推荐买词热度
	PurchasedKeyWordNum uint    `json:"purchasedKeyWordNum"`  // 关键词购买量
}

type KeywordExtData struct {
	CTR              string `json:"CTR"`
	CPM              string `json:"CPM"`
	KeywordName      string `json:"keywordName"`
	Type             string `json:"type"`
	IndirectOrderCnt uint   `json:"indirectOrderCnt"`
	DirectOrderCnt   uint   `json:"directOrderCnt"`
	IndirectCartCnt  uint   `json:"indirectCartCnt"`
	Id               string `json:"id"`
	DirectCartCnt    uint   `json:"directCartCnt"`
	Cost             string `json:"cost"`
	TotalOrderSum    string `json:"totalOrderSum"`
	TotalCartCnt     uint   `json:"totalCartCnt"`
	TotalOrderROI    string `json:"totalOrderROI"`
	Impressions      uint   `json:"impressions"`
	IndirectOrderSum string `json:"indirectOrderSum"`
	DirectOrderSum   string `json:"directOrderSum"`
	TotalOrderCVS    string `json:"totalOrderCVS"`
	CPC              string `json:"CPC"`
	Clicks           uint   `json:"clicks"`
	TotalOrderCnt    uint   `json:"totalOrderCnt"`
}

type KeywordData struct {
	Impressions                  uint    `json:"impressions"`                  // 展现数
	Clicks                       uint    `json:"clicks"`                       // clicks
	CTR                          string  `json:"CTR"`                          // CTR
	Cost                         string  `json:"cost"`                         // cost
	CPM                          string  `json:"CPM"`                          // CPM
	CPC                          string  `json:"CPC"`                          // CPC
	DirectOrderCnt               uint    `json:"directOrderCnt"`               // 直接订单数
	DirectOrderSum               string  `json:"directOrderSum"`               // 直接订单金额
	IndirectOrderCnt             uint    `json:"indirectOrderCnt"`             // 间接订单数
	IndirectOrderSum             string  `json:"indirectOrderSum"`             // 间接订单金额
	TotalOrderCnt                uint    `json:"totalOrderCnt"`                // 总订单数
	TotalOrderSum                string  `json:"totalOrderSum"`                // 总订单金额
	DirectCartCnt                uint    `json:"directCartCnt"`                // 直接加购数
	IndirectCartCnt              uint    `json:"indirectCartCnt"`              // 间接加购数
	TotalCartCnt                 uint    `json:"totalCartCnt"`                 // 总加购数
	TotalOrderCVS                string  `json:"totalOrderCVS"`                // 转化率
	TotalOrderROI                string  `json:"totalOrderROI"`                // ROI
	Id                           uint64  `json:"id"`                           // 关键词id
	KeywordName                  string  `json:"keywordName"`                  // 关键词名称
	KeywordPCPrice               float64 `json:"keywordPCPrice"`               // 关键词pc出价
	KeywordWlPrice               float64 `json:"keywordWlPrice"`               // 关键词无线出价
	CampaignId                   uint64  `json:"campaignId"`                   // 计划id
	CampaignName                 string  `json:"campaignName"`                 // 计划名称
	NewPcRank                    uint    `json:"newPcRank"`                    // 近一小时pc排名
	NewWlRank                    uint    `json:"newWlRank"`                    // 近一小时无线排名
	SearchPromoteRankCoef        uint    `json:"searchPromoteRankCoef"`        // 关键词抢排位溢价
	SearchPromoteRankEnable      uint    `json:"searchPromoteRankEnable"`      // 是否开启关键词抢排位 0 关闭 1 开启
	SearchPromoteRankSuccessRate string  `json:"searchPromoteRankSuccessRate"` // 关键词抢排名成功率
	AverageHistoryRankExpand     string  `json:"averageHistoryRankExpand"`     // 展现排名
	Type                         uint    `json:"type"`                         // 关键词匹配类型,1:精确匹配 4:短语匹配 8:切词匹配
	KeywordFlag                  string  `json:"keywordFlag"`                  // 标签
	NewCurrentPcShowq            float64 `json:"newCurrentPcShowq"`            // pc竞争力指数
	NewCurrentWlShowq            float64 `json:"newCurrentWlShowq"`            // 无线竞争力指数
	GroupId                      uint64  `json:"groupId"`                      // 单元id
	GroupName                    string  `json:"groupName"`                    // 单元名称
	KeyWordType                  uint    `json:"keyWordType"`                  // 词类型：1，关键词；2，意图词，注意：若是绑定了意图词，查询时返回的数据量会大于请求分页的数量，ex:查询10条，返回数据12条，其中2条是意图词信息
}
