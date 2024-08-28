package dmp

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	requestDsp "github.com/XiBao/jos/sdk/request/ads/dsp"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 快车增量绑定(不支持定向推荐人群)、解绑人群
type KuaicheDmpOperateRequest struct {
	api.BaseRequest
	DmpVO    *dmp.KuaicheDmpOperateRequestDmpVO            `json:"dmpVO,omitempty" codec:"dmpVO,omitempty"`       // 业务参数
	ParamExt *requestDsp.JdDspPlatformGatewayApiVoParamExt `json:"paramExt,omitempty" codec:"paramExt,omitempty"` // 系统参数
}

type KuaicheDmpOperateResponse struct {
	Responce  *KuaicheDmpOperateResponce `json:"jingdong_ads_dsp_rtb_kuaiche_dmp_operate_responce,omitempty" codec:"jingdong_ads_dsp_rtb_kuaiche_dmp_operate_responce,omitempty"`
	ErrorResp *api.ErrorResponnse        `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDmpOperateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpOperateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpOperateResponce struct {
	ReturnType *dsp.DataCommonResponse `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r KuaicheDmpOperateResponce) IsError() bool {
	return r.ReturnType != nil && !r.ReturnType.Success
}

func (r KuaicheDmpOperateResponce) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Msg
	}
	return "no result data"
}

func KuaicheDmpOperate(ctx context.Context, req *KuaicheDmpOperateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpOperateRequest()
	r.SetDmpVO(req.DmpVO)
	if req.ParamExt != nil {
		r.SetParamExt(req.ParamExt)
	}

	var response KuaicheDmpOperateResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Responce.ReturnType.Success, nil
}
