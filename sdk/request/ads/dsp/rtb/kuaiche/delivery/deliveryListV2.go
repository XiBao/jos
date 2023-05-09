package delivery

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheDeliveryListV2Request struct {
	Request *sdk.Request
}

type KuaicheDeliveryListV2RequestData struct {
	CampaignType        uint   `json:"campaignType,omitempty"`        // 计划类型，2:商品计划,41:门店商品计划，空为全部商品计划
	ClickOrOrderDay     uint   `json:"clickOrOrderDay"`               // 转化周期,0:当天，1:1天，3:3天，7:7天，15:15天,30:30天
	ClickOrOrderCaliber uint   `json:"clickOrOrderCaliber"`           // 点击口径：0 订单口径：1
	GiftFlag            uint   `json:"giftFlag"`                      // 包含赠品 0不含赠品 1含赠品
	OrderStatusCategory uint   `json:"orderStatusCategory,omitempty"` // 订单类型 空：全部订单 1：成交订单
	StartDay            string `json:"startDay"`                      // 开始时间
	EndDay              string `json:"endDay"`                        // 结束时间
	Page                int    `json:"page"`                          // 页码
	PageSize            int    `json:"pageSize,omitempty"`            // 每页条数
	Platform            string `json:"platform,omitempty"`            // 设备类型 空：全部 0 pc 1 无线
	CampaignId          uint64 `json:"campaignId,omitempty"`          // 计划id
	GroupId             uint64 `json:"groupId,omitempty"`             // 单元id
}

// create new request
func NewKuaicheDeliveryListV2Request() (req *KuaicheDeliveryListV2Request) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kc.deliveryList", Params: make(map[string]interface{}, 2)}
	req = &KuaicheDeliveryListV2Request{
		Request: &request,
	}
	return
}

func (req *KuaicheDeliveryListV2Request) SetData(data *KuaicheDeliveryListV2RequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheDeliveryListV2Request) GetData() *KuaicheDeliveryListV2RequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheDeliveryListV2RequestData)
	}
	return nil
}

func (req *KuaicheDeliveryListV2Request) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheDeliveryListV2Request) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
