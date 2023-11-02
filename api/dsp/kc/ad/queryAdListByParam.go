package ad

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/dsp"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/kc/ad"
)

type AdQueryAdListByParamRequest struct {
	api.BaseRequest
	AdGroupId uint64 `json:"ad_group_id,omitempty" codec:"ad_group_id,omitempty"` // 单元id
	PageNum   int    `json:"page_num,omitempty" codec:"page_num,omitempty"`       // 页数
	PageSize  int    `json:"page_size,omitempty" codec:"page_size,omitempty"`     // 列数
}

type AdQueryAdListByParamResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AdQueryAdListByParamData `json:"jingdong_dsp_kc_ad_queryAdListByParam_responce,omitempty" codec:"jingdong_dsp_kc_ad_queryAdListByParam_responce,omitempty"`
}

func (r AdQueryAdListByParamResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AdQueryAdListByParamResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AdQueryAdListByParamData struct {
	Result *AdQueryAdListByParamResult `json:"querylistbyparam_result,omitempty" codec:"querylistbyparam_result,omitempty"`
}

func (r AdQueryAdListByParamData) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r AdQueryAdListByParamData) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type AdQueryAdListByParamResult struct {
	Value      *AdQueryAdListByParamValues `json:"value,omitempty" codec:"value,omitempty"`
	ResultCode string                      `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	ErrorMsg   string                      `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Success    bool                        `json:"success,omitempty" codec:"success,omitempty"`
}

func (r AdQueryAdListByParamResult) IsError() bool {
	return !r.Success || r.Value == nil
}

func (r AdQueryAdListByParamResult) Error() string {
	if !r.Success {
		return sdk.ErrorString(r.ResultCode, r.ErrorMsg)
	}
	return "no result data"
}

type AdQueryAdListByParamValues struct {
	Paginator *dsp.Paginator `json:"paginator,omitempty" codec:"paginator,omitempty"`
	Datas     []DspADQuery   `json:"datas,omitempty" codec:"datas,omitempty"`
}

type DspADQuery struct {
	ImgUrl        string      `json:"imgUrl,omitempty" codec:"imgUrl,omitempty"`               // 图片地址
	Id            uint64      `json:"id,omitempty" codec:"id,omitempty"`                       // 创意id
	Status        uint8       `json:"status,omitempty" codec:"status,omitempty"`               // 状态
	Name          string      `json:"name,omitempty" codec:"name,omitempty"`                   // 创意名称
	SkuId         string      `json:"skuId,omitempty" codec:"skuId,omitempty"`                 // SkuId
	AuditInfoList []AuditInfo `json:"auditInfoList,omitempty" codec:"auditInfoList,omitempty"` // 审核List
}

type AuditInfo struct {
	AuditInfo string `json:"auditInfo,omitempty" codec:"auditInfo,omitempty"`
}

// 查询.快车.指定单元下创意基本信息
func AdQueryAdListByParam(req *AdQueryAdListByParamRequest) ([]DspADQuery, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ad.NewAdQueryAdListByParamRequest()

	r.SetAdGroupId(req.AdGroupId)
	r.SetPageSize(req.PageSize)
	r.SetPageNum(req.PageNum)

	var response AdQueryAdListByParamResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, 0, err
	}
	return response.Data.Result.Value.Datas, response.Data.Result.Value.Paginator.Items, nil

}
