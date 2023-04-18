package dsp

type JdDspPlatformGatewayApiVoParamExt struct {
	AuthType             string `json:"authType"`             // 授权模式
	AccessPin            string `json:"accessPin"`            // 被免密访问的pin
	PlatformBusinessType string `json:"platformBusinessType"` // 平台业务类型，DST_JZT：京准通
}
