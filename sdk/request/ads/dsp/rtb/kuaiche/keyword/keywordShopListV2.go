package keyword

import (
	"github.com/XiBao/jos/sdk"
)

type KuaicheKeywordShopListRequest struct {
	Request *sdk.Request
}

type KuaicheKeywordShopListRequestData struct {
	PageSize             int    `json:"pageSize,omitempty"`             // 每页条数
	ClickOrOrderDay      uint   `json:"clickOrOrderDay"`                // 当天:0;昨天：1;最近15天：15
	GiftFlag             uint   `json:"giftFlag"`                       // 包含赠品 0不含赠品 1含赠品
	CampaignId           uint64 `json:"campaignId,omitempty"`           // 计划id
	PutType              uint   `json:"putType,omitempty"`              // 投放类型 3商品 4活动
	EndDay               string `json:"endDay"`                         // 结束时间
	GroupId              uint64 `json:"groupId,omitempty"`              // 单元id
	StartDay             string `json:"startDay"`                       // 开始时间
	Page                 int    `json:"page"`                           // 页码
	Platform             string `json:"platform,omitempty"`             // 设备类型 空：全部 0 pc 1 无线
	ClickOrOrderCaliber  uint   `json:"clickOrOrderCaliber"`            // 点击口径：0 订单口径：1
	OrderStatusCategory  uint   `json:"orderStatusCategory,omitempty"`  // 订单类型 空：全部订单 1：成交订单
	NameLike             string `json:"nameLike,omitempty"`             // 计划名称
	AccessPin            string `json:"accessPin,omitempty"`            // 被免密访问的pin
	AuthType             string `json:"authType,omitempty"`             // 授权模式
	PlatformBusinessType string `json:"platformBusinessType,omitempty"` // 平台业务类型，DST_JZT：京准通
}

// create new request
func NewKuaicheKeywordShopListRequest() (req *KuaicheKeywordShopListRequest) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.shopkeyword.list", Params: make(map[string]interface{}, 16)}
	req = &KuaicheKeywordShopListRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordShopListRequest) SetPageSize(pageSize int) {
	req.Request.Params["pageSize"] = pageSize
}

func (req *KuaicheKeywordShopListRequest) GetPageSize() int {
	pageSize, found := req.Request.Params["pageSize"]
	if found {
		return pageSize.(int)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetClickOrOrderDay(clickOrOrderDay uint) {
	req.Request.Params["clickOrOrderDay"] = clickOrOrderDay
}

func (req *KuaicheKeywordShopListRequest) GetClickOrOrderDay() uint {
	clickOrOrderDay, found := req.Request.Params["clickOrOrderDay"]
	if found {
		return clickOrOrderDay.(uint)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetGiftFlag(giftFlag uint) {
	req.Request.Params["giftFlag"] = giftFlag
}

func (req *KuaicheKeywordShopListRequest) GetGiftFlag() uint {
	giftFlag, found := req.Request.Params["giftFlag"]
	if found {
		return giftFlag.(uint)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaignId"] = campaignId
}

func (req *KuaicheKeywordShopListRequest) GetCampaignId() uint64 {
	campaignId, found := req.Request.Params["campaignId"]
	if found {
		return campaignId.(uint64)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetEndDay(endDay string) {
	req.Request.Params["endDay"] = endDay
}

func (req *KuaicheKeywordShopListRequest) GetEndDay() string {
	endDay, found := req.Request.Params["endDay"]
	if found {
		return endDay.(string)
	}
	return ""
}

func (req *KuaicheKeywordShopListRequest) SetGroupId(groupId uint64) {
	req.Request.Params["groupId"] = groupId
}

func (req *KuaicheKeywordShopListRequest) GetGroupId() uint64 {
	groupId, found := req.Request.Params["groupId"]
	if found {
		return groupId.(uint64)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetStartDay(startDay string) {
	req.Request.Params["startDay"] = startDay
}

func (req *KuaicheKeywordShopListRequest) GetStartDay() string {
	startDay, found := req.Request.Params["startDay"]
	if found {
		return startDay.(string)
	}
	return ""
}

func (req *KuaicheKeywordShopListRequest) SetPage(page int) {
	req.Request.Params["page"] = page
}

func (req *KuaicheKeywordShopListRequest) GetPage() int {
	page, found := req.Request.Params["page"]
	if found {
		return page.(int)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetPlatform(platform string) {
	req.Request.Params["platform"] = platform
}

func (req *KuaicheKeywordShopListRequest) GetPlatform() string {
	platform, found := req.Request.Params["platform"]
	if found {
		return platform.(string)
	}
	return ""
}

func (req *KuaicheKeywordShopListRequest) SetClickOrOrderCaliber(clickOrOrderCaliber uint) {
	req.Request.Params["clickOrOrderCaliber"] = clickOrOrderCaliber
}

func (req *KuaicheKeywordShopListRequest) GetClickOrOrderCaliber() uint {
	clickOrOrderCaliber, found := req.Request.Params["clickOrOrderCaliber"]
	if found {
		return clickOrOrderCaliber.(uint)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetOrderStatusCategory(orderStatusCategory uint) {
	req.Request.Params["orderStatusCategory"] = orderStatusCategory
}

func (req *KuaicheKeywordShopListRequest) GetOrderStatusCategory() uint {
	orderStatusCategory, found := req.Request.Params["orderStatusCategory"]
	if found {
		return orderStatusCategory.(uint)
	}
	return 0
}

func (req *KuaicheKeywordShopListRequest) SetNameLike(nameLike string) {
	req.Request.Params["nameLike"] = nameLike
}

func (req *KuaicheKeywordShopListRequest) GetNameLike() string {
	nameLike, found := req.Request.Params["nameLike"]
	if found {
		return nameLike.(string)
	}
	return ""
}

func (req *KuaicheKeywordShopListRequest) SetAccessPin(accessPin string) {
	req.Request.Params["accessPin"] = accessPin
}

func (req *KuaicheKeywordShopListRequest) GetAccessPin() string {
	accessPin, found := req.Request.Params["accessPin"]
	if found {
		return accessPin.(string)
	}
	return ""
}

func (req *KuaicheKeywordShopListRequest) SetAuthType(authType string) {
	req.Request.Params["authType"] = authType
}

func (req *KuaicheKeywordShopListRequest) GetAuthType() string {
	authType, found := req.Request.Params["authType"]
	if found {
		return authType.(string)
	}
	return ""
}

func (req *KuaicheKeywordShopListRequest) SetPlatformBusinessType(platformBusinessType string) {
	req.Request.Params["platformBusinessType"] = platformBusinessType
}

func (req *KuaicheKeywordShopListRequest) GetPlatformBusinessType() string {
	platformBusinessType, found := req.Request.Params["platformBusinessType"]
	if found {
		return platformBusinessType.(string)
	}
	return ""
}
