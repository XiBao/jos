package delivery

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/delivery"
)

// 商品定向批量改价
type KuaicheDeliveryUpdateProductPriceRequest struct {
	api.BaseRequest
	Data   *delivery.KuaicheDeliveryUpdateProductPriceRequestData `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt          `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheDeliveryUpdateProductPriceResponse struct {
	Responce  *KuaicheDeliveryUpdateProductPriceResponce `json:"jingdong_ads_dsp_rtb_kc_batchUpdateProductDeliveryPrice_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kc_batchUpdateProductDeliveryPrice_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                        `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDeliveryUpdateProductPriceResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDeliveryUpdateProductPriceResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDeliveryUpdateProductPriceResponce struct {
	ReturnType *KuaicheDeliveryUpdateProductPriceResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                                               `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDeliveryUpdateProductPriceResponce) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r KuaicheDeliveryUpdateProductPriceResponce) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type KuaicheDeliveryUpdateProductPriceResponseReturnType struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheDeliveryUpdateProductPrice(ctx context.Context, req *KuaicheDeliveryUpdateProductPriceRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := delivery.NewKuaicheDeliveryUpdateProductPriceRequest()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	var response KuaicheDeliveryUpdateProductPriceResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.ReturnType.Data, nil
}
