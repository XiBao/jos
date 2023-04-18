package campaign

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheCampaignListV2Request struct {
	Request *sdk.Request
}

type KuaicheCampaignListV2RequestData struct {
	PutType             uint   `json:"putType,omitempty"`             // 投放类型 3商品 4活动
	PageSize            int    `json:"pageSize,omitempty"`            // 每页条数
	ClickOrOrderDay     uint   `json:"clickOrOrderDay"`               // 当天:0;昨天：1;最近15天：15
	CampaignType        uint   `json:"campaignType,omitempty"`        // 计划类型 2普通快车计划 18快车腰带店铺计划
	GiftFlag            uint   `json:"giftFlag"`                      // 包含赠品 0不含赠品 1含赠品
	CampaignId          uint64 `json:"campaignId,omitempty"`          // 计划id
	OrderStatusCategory uint   `json:"orderStatusCategory,omitempty"` // 订单类型 空：全部订单 1：成交订单
	Status              uint   `json:"status,omitempty"`              // 计划状态 1暂停2启动3预算用完4不在投放时间段
	EndDay              string `json:"endDay"`                        // 结束时间
	NameLike            string `json:"nameLike,omitempty"`            // 计划名称
	ActivityId          uint64 `json:"activityId,omitempty"`          // 联合活动id
	StartDay            string `json:"startDay"`                      // 开始时间
	Page                int    `json:"page"`                          // 页码
	Platform            string `json:"platform,omitempty"`            // 设备类型 空：全部 0 pc 1 无线
	ClickOrOrderCaliber uint   `json:"clickOrOrderCaliber"`           // 点击口径：0 订单口径：1
}

// create new request
func NewKuaicheCampaignListV2Request() (req *KuaicheCampaignListV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kuaiche.campaign.list.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheCampaignListV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheCampaignListV2Request) SetData(data *KuaicheCampaignListV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheCampaignListV2Request) GetData() *KuaicheCampaignListV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheCampaignListV2RequestData)
	}
	return nil
}

func (req *KuaicheCampaignListV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheCampaignListV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
