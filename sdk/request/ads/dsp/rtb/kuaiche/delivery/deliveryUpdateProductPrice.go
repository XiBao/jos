package delivery

import (
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp"
)

type KuaicheDeliveryUpdateProductPriceRequest struct {
	Request *sdk.Request
}

type KuaicheDeliveryUpdateProductPriceRequestPrice struct {
	DevType uint    `json:"devType"` // 出价类型 1:pc出价 2:无线出价
	Number  float64 `json:"number"`  // 具体值 出价范围0.1-9999，若是百分比则为0-500
	Change  uint    `json:"change"`  // 出价幅度 0:对应type自定义类型 1:提高 2:降低
	Type    uint    `json:"type"`    // 改价类型，1:自定义类型 2:提高/降低出价幅度 3:提供/降低百分比
}

type KuaicheDeliveryUpdateProductPriceRequestData struct {
	Ids       []uint64                                        `json:"ids"`       // 商品定向id
	PrictList []KuaicheDeliveryUpdateProductPriceRequestPrice `json:"priceList"` // 调价信息（列表数量上限为2），包含无线和pc出价，若不填写无线出价则根据无线出价系数进行换算（pc无线出价已做合并）
}

// create new request
func NewKuaicheDeliveryUpdateProductPriceRequest() (req *KuaicheDeliveryUpdateProductPriceRequest) {
	request := sdk.Request{MethodName: "jingdong.ads.dsp.rtb.kc.batchUpdateProductDeliveryPrice", Params: make(map[string]interface{}, 2)}
	req = &KuaicheDeliveryUpdateProductPriceRequest{
		Request: &request,
	}
	return
}

func (req *KuaicheDeliveryUpdateProductPriceRequest) SetData(data *KuaicheDeliveryUpdateProductPriceRequestData) {
	req.Request.Params["data"] = data
}

func (req *KuaicheDeliveryUpdateProductPriceRequest) GetData() *KuaicheDeliveryUpdateProductPriceRequestData {
	data, found := req.Request.Params["data"]
	if found {
		return data.(*KuaicheDeliveryUpdateProductPriceRequestData)
	}
	return nil
}

func (req *KuaicheDeliveryUpdateProductPriceRequest) SetSystem(system *dsp.JdDspPlatformGatewayApiVoParamExt) {
	req.Request.Params["system"] = system
}

func (req *KuaicheDeliveryUpdateProductPriceRequest) GetSystem() *dsp.JdDspPlatformGatewayApiVoParamExt {
	system, found := req.Request.Params["system"]
	if found {
		return system.(*dsp.JdDspPlatformGatewayApiVoParamExt)
	}
	return nil
}
