package campaign

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/campaign"
)

// 修改京东快车计划投放时间
type KuaicheCampaignUpdateDateV2Request struct {
	api.BaseRequest
	Data   *campaign.KuaicheCampaignUpdateDateV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt    `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheCampaignUpdateDateV2Response struct {
	Responce  *KuaicheCampaignUpdateDateV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_campaign_UpdateDate_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_campaign_UpdateDate_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                  `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheCampaignUpdateDateV2Responce struct {
	Data *KuaicheCampaignUpdateDateV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                   `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheCampaignUpdateDateV2ResponseData struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheCampaignUpdateDateV2(req *KuaicheCampaignUpdateDateV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := campaign.NewKuaicheCampaignUpdateDateV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	result, err := client.Execute(r.Request, req.Session)
	fmt.Println(string(result))
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}

	var response KuaicheCampaignUpdateDateV2Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.Data == nil {
		return false, errors.New("no result data.")
	}
	if !response.Responce.Data.Success {
		return false, errors.New(response.Responce.Data.Msg)
	}

	return response.Responce.Data.Success, nil
}
