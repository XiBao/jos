package dmp

import (
	"context"

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

func (r KuaicheDmpUpdatePriceResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpUpdatePriceResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpUpdatePriceResponce struct {
	ReturnType *dsp.DataCommonResponse `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r KuaicheDmpUpdatePriceResponce) IsError() bool {
	return r.ReturnType != nil && !r.ReturnType.Success
}

func (r KuaicheDmpUpdatePriceResponce) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Msg
	}
	return "no result data"
}

func KuaicheDmpUpdatePrice(ctx context.Context, req *KuaicheDmpUpdatePriceRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpUpdatePriceRequest()
	r.SetDmpVOS(req.DmpVOS)
	if req.ParamExt != nil {
		r.SetParamExt(req.ParamExt)
	}

	var response KuaicheDmpUpdatePriceResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.ReturnType.Success, nil
}
