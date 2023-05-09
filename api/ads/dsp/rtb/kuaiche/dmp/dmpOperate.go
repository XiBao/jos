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

type KuaicheDmpOperateResponce struct {
	ReturnType *dsp.DataCommonResponse `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func KuaicheDmpOperate(req *KuaicheDmpOperateRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpOperateRequest()
	r.SetDmpVO(req.DmpVO)
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

	var response KuaicheDmpOperateResponse
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
