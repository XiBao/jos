package adgroup

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/adgroup"
)

// 修改快车单元出价，pc搜索出价，无线出价系数
type KuaicheAdgroupUpdateFeeV2Request struct {
	api.BaseRequest
	Data   *adgroup.KuaicheAdgroupUpdateFeeV2RequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheAdgroupUpdateFeeV2Response struct {
	Responce  *KuaicheAdgroupUpdateFeeV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_group_updatefee_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_group_updatefee_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheAdgroupUpdateFeeV2Responce struct {
	Data *KuaicheAdgroupUpdateFeeV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                 `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheAdgroupUpdateFeeV2ResponseData struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheAdgroupUpdateFeeV2(req *KuaicheAdgroupUpdateFeeV2Request) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adgroup.NewKuaicheAdgroupUpdateFeeV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}

	var response KuaicheAdgroupUpdateFeeV2Response
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
