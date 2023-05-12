package dmp

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 批量修改不同人群不同溢价
type KuaicheDmpUpdatePriceRequest struct {
	api.BaseRequest
	DmpVOS   []dmp.KuaicheDmpUpdatePriceRequestDmpVOS      `json:"dmpVOS,omitempty" codec:"dmpVOS,omitempty"`     // 业务参数
	ParamExt *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"paramExt,omitempty" codec:"paramExt,omitempty"` // 系统参数
}

type KuaicheDmpUpdatePriceResponse struct {
	Responce  *KuaicheDmpUpdatePriceResponce `json:"jingdong_ads_dsp_rtb_kuaiche_batchUpdateDifferentDmpPrice_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_batchUpdateDifferentDmpPrice_responce,omitempty"`
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheDmpUpdatePriceResponce struct {
	ReturnType *dsp.DataCommonResponse `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func KuaicheDmpUpdatePrice(req *KuaicheDmpUpdatePriceRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpUpdatePriceRequest()
	r.SetDmpVOS(req.DmpVOS)
	if req.ParamExt != nil {
		r.SetParamExt(req.ParamExt)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}

	var response KuaicheDmpUpdatePriceResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.ReturnType == nil {
		return false, errors.New("no result data.")
	}
	if !response.Responce.ReturnType.Success {
		return false, errors.New(response.Responce.ReturnType.Msg)
	}

	return response.Responce.ReturnType.Success, nil
}
