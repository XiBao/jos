package adgroup

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheAdgroupListV2Request struct {
	Request *sdk.Request
}

type KuaicheAdgroupListV2RequestData struct {
	PutType             uint     `json:"putType,omitempty"`             // 投放类型 3商品 4活动
	PageSize            int      `json:"pageSize,omitempty"`            // 每页条数
	ClickOrOrderDay     uint     `json:"clickOrOrderDay"`               // 转化周期 当天:0 1天:1 3天:3 7天:7 15天: 15
	CampaignType        uint     `json:"campaignType,omitempty"`        // 计划类型 2普通快车 18 快车腰带店铺
	GiftFlag            uint     `json:"giftFlag"`                      // 包含赠品 0不含赠品 1含赠品
	CampaignId          uint64   `json:"campaignId,omitempty"`          // 计划id
	OrderStatusCategory uint     `json:"orderStatusCategory,omitempty"` // 订单类型 空：全部订单 1：成交订单
	Status              uint     `json:"status,omitempty"`              // 单元状态 1 暂停 2 启动
	EndDay              string   `json:"endDay"`                        // 结束时间
	NameLike            string   `json:"nameLike,omitempty"`            // 单元名称
	StartDay            string   `json:"startDay"`                      // 开始时间
	Page                int      `json:"page"`                          // 页码
	Platform            string   `json:"platform,omitempty"`            // 设备类型 空：全部 0 pc 1 无线
	ClickOrOrderCaliber uint     `json:"clickOrOrderCaliber"`           // 口径，0：点击，1：下单
	Filters             []string `json:"filters,omitempty"`             // 高级查询，格式为 属性|条件|属性值，条件为：gt大于，lt小于，如impressions|gt|10，表示展现数大于10
}

// create new request
func NewKuaicheAdgroupListV2Request() (req *KuaicheAdgroupListV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.group.list.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheAdgroupListV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheAdgroupListV2Request) SetData(data *KuaicheAdgroupListV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheAdgroupListV2Request) GetData() *KuaicheAdgroupListV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheAdgroupListV2RequestData)
	}
	return nil
}

func (req *KuaicheAdgroupListV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheAdgroupListV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
