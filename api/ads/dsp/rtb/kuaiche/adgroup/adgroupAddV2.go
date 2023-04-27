package adgroup

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	adgroup "github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/adgroup"
)

// 新建快车商品，店铺类型单元
type KuaicheAdgroupAddV2Request struct {
	api.BaseRequest
	Data   *adgroup.KuaicheAdgroupAddV2RequestData       `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheAdgroupAddV2Response struct {
	Responce  *KuaicheAdgroupAddV2Responce `json:"jingdong_ads_dsp_rtb_kuaiche_productgroup_add_v2_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_productgroup_add_v2_responce,omitempty"`
	ErrorResp *api.ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheAdgroupAddV2Responce struct {
	Data *KuaicheAdgroupAddV2ResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                           `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheAdgroupAddV2ResponseData struct {
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheAdgroupAddV2(req *KuaicheAdgroupAddV2Request) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adgroup.NewKuaicheAdgroupAddV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response KuaicheAdgroupAddV2Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.Data == nil {
		return 0, errors.New("no result data.")
	}
	if !response.Responce.Data.Success {
		return 0, errors.New(response.Responce.Data.Msg)
	}

	return response.Responce.Data.Data, nil
}
