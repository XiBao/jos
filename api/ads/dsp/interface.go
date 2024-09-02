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
	Status         int                     `json:"status"`         // 计划状态
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
	SkuId          string                  `json:"skuId"`          // skuId
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
	Type             int    `json:"type"`
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
	KeywordPCPrice               string  `json:"keywordPCPrice"`               // 关键词pc出价
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
	NewCurrentPcShowq            string  `json:"newCurrentPcShowq"`            // pc竞争力指数
	NewCurrentWlShowq            string  `json:"newCurrentWlShowq"`            // 无线竞争力指数
	GroupId                      uint64  `json:"groupId"`                      // 单元id
	GroupName                    string  `json:"groupName"`                    // 单元名称
	KeyWordType                  uint    `json:"keyWordType"`                  // 词类型：1，关键词；2，意图词，注意：若是绑定了意图词，查询时返回的数据量会大于请求分页的数量，ex:查询10条，返回数据12条，其中2条是意图词信息
}

// 定向数据
type DeliveryData struct {
	CPC                         string  `json:"CPC"`                         // CPC
	CPM                         string  `json:"CPM"`                         // CPM
	CTR                         string  `json:"CTR"`                         // CTR
	CampaignId                  uint64  `json:"campaignId"`                  // 计划id
	CampaignName                string  `json:"campaignName"`                // 计划名称
	CampaignType                uint    `json:"campaignType"`                // 计划类型
	ClickDate                   string  `json:"clickDate"`                   // 点击时间
	Clicks                      uint    `json:"clicks"`                      // clicks
	Cost                        string  `json:"cost"`                        // cost
	DeliveryType                uint    `json:"deliveryType"`                // 资源位0:汇总 1:推荐 ，2:搜索， 3：商品
	DirectCartCnt               uint    `json:"directCartCnt"`               // 直接加购数
	DirectOrderCnt              uint    `json:"directOrderCnt"`              // 直接订单数
	DirectOrderSum              string  `json:"directOrderSum"`              // 直接订单金额
	EffectCartCnt               uint    `json:"effectCartCnt"`               // 影响加购数
	EffectOrderCnt              uint    `json:"effectOrderCnt"`              // 影响订单数
	EffectOrderSum              string  `json:"effectOrderSum"`              // 影响订单金额
	GroupId                     uint64  `json:"groupId"`                     // 单元id
	GroupName                   string  `json:"groupName"`                   // 推广单元名称
	Impressions                 uint    `json:"impressions"`                 // 展现数
	IndirectCartCnt             uint    `json:"indirectCartCnt"`             // 间接加购数
	IndirectOrderCnt            uint    `json:"indirectOrderCnt"`            // 间接订单数
	IndirectOrderSum            string  `json:"indirectOrderSum"`            // 间接订单金额
	IsDefaultPrice              bool    `json:"isDefaultPrice"`              // 否独立出价，true:系数出价，false:独立出价
	PcPrice                     float64 `json:"pcPrice"`                     // PC出价
	ProductDeliveryMatchingId   uint64  `json:"productDeliveryMatchingId"`   // 定向条件id
	ProductDeliveryMatchingType uint    `json:"productDeliveryMatchingType"` // 定向条件匹配类型：1：sku定向，2：spu定向，3：类目定向
	SkuBrandId                  uint64  `json:"skuBrandId"`                  // 品牌id
	SkuBrandNameCn              string  `json:"skuBrandNameCn"`              // 品牌中文名
	SkuCName3                   string  `json:"skuCName3"`                   // 三级类目中文名
	SkuCid3                     uint64  `json:"skuCid3"`                     // 三级类目ID
	SkuId                       uint64  `json:"skuId"`                       // skuId
	TotalCartCnt                uint    `json:"totalCartCnt"`                // 总加购数
	TotalOrderCVS               string  `json:"totalOrderCVS"`               // 转化率
	TotalOrderCnt               uint    `json:"totalOrderCnt"`               // 总订单数
	TotalOrderROI               string  `json:"totalOrderROI"`               // ROI
	TotalOrderSum               string  `json:"totalOrderSum"`               // 总订单金额
	WlPrice                     float64 `json:"wlPrice"`                     // 无线出价
	WlPriceSource               string  `json:"wlPriceSource"`               // 无线出价系数
}

type DmpExtData struct {
	CTR              string `json:"CTR"`
	CPM              string `json:"CPM"`
	Cost             string `json:"cost"`
	TotalOrderSum    string `json:"totalOrderSum"`
	TotalCartCnt     uint   `json:"totalCartCnt"`
	TotalOrderROI    string `json:"totalOrderROI"`
	Impressions      uint   `json:"impressions"`
	IndirectOrderSum string `json:"indirectOrderSum"`
	DirectOrderSum   string `json:"directOrderSum"`
	IndirectOrderCnt uint   `json:"indirectOrderCnt"`
	DirectOrderCnt   uint   `json:"directOrderCnt"`
	IndirectCartCnt  uint   `json:"indirectCartCnt"`
	TotalOrderCVS    string `json:"totalOrderCVS"`
	CPC              string `json:"CPC"`
	Clicks           uint   `json:"clicks"`
	TotalOrderCnt    uint   `json:"totalOrderCnt"`
	DirectCartCnt    uint   `json:"directCartCnt"`
}

type DmpData struct {
	Impressions      uint   `json:"impressions"`      // 展现数
	Clicks           uint   `json:"clicks"`           // clicks
	CTR              string `json:"CTR"`              // CTR
	Cost             string `json:"cost"`             // cost
	CPM              string `json:"CPM"`              // CPM
	CPC              string `json:"CPC"`              // CPC
	DirectOrderCnt   uint   `json:"directOrderCnt"`   // 直接订单数
	DirectOrderSum   string `json:"directOrderSum"`   // 直接订单金额
	IndirectOrderCnt uint   `json:"indirectOrderCnt"` // 间接订单数
	IndirectOrderSum string `json:"indirectOrderSum"` // 间接订单金额
	TotalOrderCnt    uint   `json:"totalOrderCnt"`    // 总订单数
	TotalOrderSum    string `json:"totalOrderSum"`    // 总订单金额
	DirectCartCnt    uint   `json:"directCartCnt"`    // 直接加购数
	IndirectCartCnt  uint   `json:"indirectCartCnt"`  // 间接加购数
	TotalCartCnt     uint   `json:"totalCartCnt"`     // 总加购数
	TotalOrderCVS    string `json:"totalOrderCVS"`    // 转化率
	TotalOrderROI    string `json:"totalOrderROI"`    // ROI
	DmpId            uint64 `json:"dmpId"`            // 人群id
	DmpName          string `json:"dmpName"`          // 人群名称
	DmpStatus        uint   `json:"dmpStatus"`        // 人群状态 0未启用 1 启用
	DmpFactor        uint   `json:"dmpFactor"`        // 人群溢价系数
	CampaignId       uint64 `json:"campaignId"`       // 计划id
	GroupId          uint64 `json:"groupId"`          // 单元id
}

type DmpCommonCrowdQuery struct {
	CrowdId            uint64   `json:"crowdId"`            // 人群id
	CrowdName          string   `json:"crowdName"`          // 人群名称
	CrowdType          uint     `json:"crowdType"`          // 人群类型:1标签人群,2种子包人群,3扩展人群，4场景人群，5只能定向人群，6动态规则人群,7推荐促点击人群,8兴趣定向人群,9营销策略人群
	SenceFirstCategory string   `json:"senceFirstCategory"` // 京选人群一级分类
	EstimateUv         uint     `json:"estimateUv"`         // 曝光人数(默认人群为null)
	EstimatePv         uint     `json:"estimatePv"`         // 曝光量(默认人群为-1)
	SenceCrowdDesc     string   `json:"senceCrowdDesc"`     // 人群的描述信息
	IsNewTagCompose    uint     `json:"isNewTagCompose"`    // 是否是新标题体系人群:0不是;1是
	CrowdTabType       uint     `json:"crowdTabType"`       // 按照人群快车投放端tab的分类枚举:1.DMP，2定向，3京选,4智能定向人群
	ExpiredTime        uint64   `json:"expiredTime"`        // 过期时间
	CrowdKeyWordDesc   []string `json:"crowdKeyWordDesc"`   // 人群tag描述：对应应用情景及人群特点列
}

type TagVO struct {
	TagId                uint64 `json:"tagId"`                // 标签ID
	TagName              string `json:"tagName"`              // 标签名称
	TagDesc              string `json:"tagDesc"`              // 标签描述
	FilterDemo           string `json:"filterDemo"`           // 筛选项示例
	IndustryHot          string `json:"industryHot"`          // 行业热度
	CoverageRate         string `json:"coverageRate"`         // 覆盖率
	IsFavorite           uint   `json:"isFavorite"`           // 是否收藏，1收藏
	NewTag               bool   `json:"newTag"`               // 是否为新上标签
	ExtInfo              string `json:"extInfo"`              // 扩展信息，Json串
	TagSourceId          uint   `json:"tagSourceId"`          // 标签来源ID
	ImpUrl               string `json:"impUrl"`               // 曝光url
	ViewDetailsClickUrl  string `json:"viewDetailsClickUrl"`  // 点击url
	AddIntoBoardClickUrl string `json:"addIntoBoardClickUrl"` // 添加看板点击url
	OfflineTag           bool   `json:"offlineTag"`           // 是否是即将下线标签
	ValidNum             uint   `json:"validNum"`             // 有效种子包数量
}

type TagDetail struct {
	TemplateId         string                        `json:"templateId"`         // 模版id
	TemplateDesc       string                        `json:"templateDesc"`       // 模版描述
	IndustryHot        string                        `json:"industryHot"`        // 热度
	CoverageRate       string                        `json:"coverageRate"`       // 覆盖分
	TemplateSearchDesc string                        `json:"templateSearchDesc"` // 搜索文案
	DimensionType      uint                          `json:"dimensionType"`      // 标签类型，根据不同的标签类型从不同的属性中取筛选项
	TagId              uint64                        `json:"tagId"`              // 标签id
	TagName            string                        `json:"tagName"`            // 标签名称
	CommitAttribute    []string                      `json:"commitAttribute"`    // 设置标签时提交的属性名称
	ChoosedItems       []TagDetailChoosedItem        `json:"choosedItems"`       // 已经勾选的筛选项
	CategoryNodes      []TagDetailCategoryNode       `json:"categoryNodes"`      // 三级类目筛选项，仅当dimenType为2时有效
	Level2data         []TagDetailLevel2data         `json:"level2data"`         // level2data 筛选项数据
	Level2Type         uint                          `json:"level2Type"`         // level2Type筛选项数据类型
	PriceAttributes    []TagDetailPriceAttribute     `json:"priceAttributes"`    // 价格属性信息
	BaseDimenId        uint64                        `json:"baseDimenId"`        // 后台维度id
	DimenStatus        uint                          `json:"dimenStatus"`        // 标签状态
	FrequencyAttribute []TagDetailFrequencyAttribute `json:"frequencyAttribute"` // 频率相关属性配置
	CategoryLimit      uint                          `json:"categoryLimit"`      // 筛选项个数限制，通常为200
	IsActionComplexTag uint                          `json:"isActionComplexTag"` // 是否是行为复合标签：0不是;1是
	CompeteGoodsType   uint                          `json:"competeGoodsType"`   // 竞品圈选程度：高、中、低
	HasCompeteGoods    bool                          `json:"hasCompeteGoods"`    // 判断用户是否有竞品数据
	BoardId            int64                         `json:"boardId"`            // 看板id
	RelevantDegreeType uint                          `json:"relevantDegreeType"` // 关联度选择：高中低
	ResourceType       string                        `json:"resourceType"`       // 资源类型 0：全部 1：主站 2：京喜
	ExtendInfo         *TagDetailExtendInfo          `json:"extendInfo"`         // 扩展属性对象
	SelectedData       []TagDetailSelectedData       `json:"selectedData"`       // 扩展属性
}

type TagDetailSelectedData struct {
	Wid              uint64                 `json:"wid"`                   // 当前记录ID
	Pid              uint64                 `json:"pid"`                   // 父ID
	Name             string                 `json:"name"`                  // 名称
	Level            uint                   `json:"level"`                 // 层级
	CidLevel         uint                   `json:"cidLevel"`              // 类目层级
	CategoryName     string                 `json:"categoryName"`          // 类目名称
	Yn               uint                   `json:"yn"`                    // yn
	SonTreeNode      *TagDetailSelectedData `json:"sonTreeNode,omitempty"` // 嵌套子节点
	JointCampaignIds []uint64               `json:"jointCampaignIds"`      // 计划ID
}

type TagDetailExtendInfo struct {
	ResourceTypes  []TagDetailExtendInfoResourceType  `json:"resourceTypes"`  // 数据范围
	CategoryLevels []TagDetailExtendInfoCategoryLevel `json:"categoryLevels"` // 类目级别
	Attribute      []TagDetailExtendInfoAttribute     `json:"attribute"`      // 标签属性
	ComplexActions []TagDetailExtendInfoComplexAction `json:"complexActions"` // 复合行为标签里面的行为信息
}

type TagDetailExtendInfoComplexAction struct {
	Id             string `json:"id"`             // id
	Name           string `json:"name"`           // name
	Checked        uint   `json:"checked"`        // checked
	Type           uint   `json:"type"`           // type
	TagValueValid  uint   `json:"tagValueValid"`  // 标签值是否有效，默认有效
	IsCompeteGoods uint   `json:"isCompeteGoods"` // 标签是否为竞品，默认不是
}

type TagDetailExtendInfoAttribute struct {
	DefineId  uint64                               `json:"defineId"`  // 属性主键
	CommitKey string                               `json:"commitKey"` // 提交英文属性名称
	CnName    string                               `json:"cnName"`    // 属性中文名称
	Config    []TagDetailExtendInfoAttributeConfig `json:"config"`    // 选中的具体属性值
}

type TagDetailExtendInfoAttributeConfig struct {
	Name                string                                        `json:"name"`                // 前台属性中文名称
	Wid                 string                                        `json:"wid"`                 // 值
	Checked             bool                                          `json:"checked"`             // 是否勾选
	Configured          bool                                          `json:"configured"`          // 是否显示此属性
	PriceBeginValue     float64                                       `json:"priceBeginValue"`     // 价格区间
	PriceEndValue       float64                                       `json:"priceEndValue"`       // 价格区间
	FrequencyBeginValue uint                                          `json:"frequencyBeginValue"` // 频次区间
	FrequencyEndValue   uint                                          `json:"frequencyEndValue"`   // 频次区间
	CommonDays          []TagDetailExtendInfoAttributeConfigCommonDay `json:"commonDays"`          // 复合类目行为时间
}

type TagDetailExtendInfoAttributeConfigCommonDay struct {
	CommitKey string `json:"commitKey"` // 提交英文属性名称
}

type TagDetailExtendInfoCategoryLevel struct {
	Pid        uint64 `json:"pid"`        // 父ID
	Id         uint64 `json:"id"`         // id
	Name       string `json:"name"`       // 商品名称
	Wid        uint64 `json:"wid"`        // 当前ID
	IsParent   bool   `json:"isParent"`   // 是否为父节点
	Checked    bool   `json:"checked"`    // 是否选中
	Nocheck    bool   `json:"nocheck"`    // 是否显示复选框
	Configured bool   `json:"configured"` // 后台是否配置
}

type TagDetailExtendInfoResourceType struct {
	Pid        uint64 `json:"pid"`        // 父ID
	Id         uint64 `json:"id"`         // id
	Name       string `json:"name"`       // 商品名称
	Wid        uint64 `json:"wid"`        // 当前ID
	IsParent   bool   `json:"isParent"`   // 是否为父节点
	Checked    bool   `json:"checked"`    // 是否选中
	Nocheck    bool   `json:"nocheck"`    // 是否显示复选框
	Configured bool   `json:"configured"` // 后台是否配置
}

type TagDetailFrequencyAttribute struct {
	BeginValue uint `json:"beginValue"` // 频率的起始值
	EndValue   uint `json:"endValue"`   // 频率的结束值
}

type TagDetailPriceAttribute struct {
	AttributeId   string  `json:"attributeId"`   // attributeId
	AttributeName string  `json:"attributeName"` // attributeName
	Checked       uint    `json:"checked"`       // 是否选中，1选中，0不选中
	EndPrice      float64 `json:"endPrice"`      // 结束价格
	StartPrice    float64 `json:"startPrice"`    // 起始价格
}

type TagDetailLevel2data struct {
	Id               string                `json:"id"`                 // id
	Pid              string                `json:"pid"`                // pid
	Name             string                `json:"name"`               // name
	CommitId         uint64                `json:"commitId"`           // commitId
	Type             uint                  `json:"type"`               // 表示通天塔or联合活动，1：通天塔；2：联合活动
	JointCampaignIds []uint64              `json:"jointCampaignIds"`   // 如果type=2,此字段有值，表示联合活动对应的通天塔活动id
	SonNodes         []TagDetailLevel2data `json:"sonNodes,omitempty"` // sonNodes
}

type TagDetailCategoryNode struct {
	Wid          uint64 `json:"wid"`          // 当前记录id
	Pid          uint64 `json:"pid"`          // 父记录id
	Name         string `json:"name"`         // 名称
	Level        uint   `json:"level"`        // 层级
	CidLevel     uint   `json:"cidLevel"`     // 子节点层级
	CategoryName string `json:"categoryName"` // 类目名称
}

type TagDetailChoosedItem struct {
	Id      string `json:"id"`      // 筛选项id
	Name    string `json:"name"`    // 筛选项名称
	Checked uint   `json:"checked"` // 	是否被选中
	Type    uint   `json:"type"`    // 类型
}
