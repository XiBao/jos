package keyword

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheKeywordListV2Request struct {
	Request *sdk.Request
}

type KuaicheKeywordListV2RequestData struct {
	PutType             uint   `json:"putType,omitempty"`             // 投放类型 3商品 4活动
	PageSize            int    `json:"pageSize,omitempty"`            // 每页条数
	ClickOrOrderDay     uint   `json:"clickOrOrderDay"`               // 当天:0;昨天：1;最近15天：15
	GiftFlag            uint   `json:"giftFlag"`                      // 包含赠品 0不含赠品 1含赠品
	CampaignId          uint64 `json:"campaignId,omitempty"`          // 计划id
	EndDay              string `json:"endDay"`                        // 结束时间
	GroupId             uint64 `json:"groupId,omitempty"`             // 单元id
	StartDay            string `json:"startDay"`                      // 开始时间
	Page                int    `json:"page"`                          // 页码
	Platform            string `json:"platform,omitempty"`            // 设备类型 空：全部 0 pc 1 无线
	ClickOrOrderCaliber uint   `json:"clickOrOrderCaliber"`           // 点击口径：0 订单口径：1
	OrderStatusCategory uint   `json:"orderStatusCategory,omitempty"` // 订单类型 空：全部订单 1：成交订单
	NameLike            string `json:"nameLike,omitempty"`            // 计划名称
}

// create new request
func NewKuaicheKeywordListV2Request() (req *KuaicheKeywordListV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.keyword.list.v2", Params: make(map[string]interface{}, 2)}
	req = &KuaicheKeywordListV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheKeywordListV2Request) SetData(data *KuaicheKeywordListV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheKeywordListV2Request) GetData() *KuaicheKeywordListV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheKeywordListV2RequestData)
	}
	return nil
}

func (req *KuaicheKeywordListV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheKeywordListV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
