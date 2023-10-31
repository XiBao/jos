package delivery

import (
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

func (r KuaicheDeliveryListV2Response) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDeliveryListV2Response) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDeliveryListV2Responce struct {
	ReturnType *KuaicheDeliveryListV2ResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                                   `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDeliveryListV2Responce) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r KuaicheDeliveryListV2Responce) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
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

	var response KuaicheDeliveryListV2Response
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.ReturnType.Data, nil
}
