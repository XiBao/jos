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

// 商品定向列表
type KuaicheDeliveryListV2Request struct {
	api.BaseRequest
	Data   *delivery.KuaicheDeliveryListV2RequestData    `json:"data,omitempty" codec:"data,omitempty"`     // 业务参数
	System *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"system,omitempty" codec:"system,omitempty"` // 系统参数
}

type KuaicheDeliveryListV2Response struct {
	Responce  *KuaicheDeliveryListV2Responce `json:"jingdong_ads_dsp_rtb_kc_deliveryList_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kc_deliveryList_responce,omitempty"`
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

type KuaicheDeliveryListV2Responce struct {
	ReturnType *KuaicheDeliveryListV2ResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                                   `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheDeliveryListV2ResponseReturnType struct {
	Data *KuaicheDeliveryListV2ResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheDeliveryListV2ResponseDataData struct {
	Deliveries []dsp.DeliveryData `json:"datas,omitempty" codec:"datas,omitempty"`
	Paginator  *dsp.Paginator     `json:"paginator,omitempty" codec:"paginator,omitempty"`
}

func KuaicheDeliveryListV2(req *KuaicheDeliveryListV2Request) (*KuaicheDeliveryListV2ResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := delivery.NewKuaicheDeliveryListV2Request()
	r.SetData(req.Data)
	if req.System != nil {
		r.SetSystem(req.System)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response KuaicheDeliveryListV2Response
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.ReturnType == nil {
		return nil, errors.New("no result data.")
	}
	if !response.Responce.ReturnType.Success {
		return nil, errors.New(response.Responce.ReturnType.Msg)
	}

	return response.Responce.ReturnType.Data, nil
}
