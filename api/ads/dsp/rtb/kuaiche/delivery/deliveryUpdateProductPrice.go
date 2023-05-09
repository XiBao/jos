package delivery

import (
	"encoding/json"
	"errors"

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

type KuaicheDeliveryUpdateProductPriceResponce struct {
	ReturnType *KuaicheDeliveryUpdateProductPriceResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                                               `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheDeliveryUpdateProductPriceResponseReturnType struct {
	Data bool `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheDeliveryUpdateProductPrice(req *KuaicheDeliveryUpdateProductPriceRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := delivery.NewKuaicheDeliveryUpdateProductPriceRequest()
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

	var response KuaicheDeliveryUpdateProductPriceResponse
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

	return response.Responce.ReturnType.Data, nil
}
