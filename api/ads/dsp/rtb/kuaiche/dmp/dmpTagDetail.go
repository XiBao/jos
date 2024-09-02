package dmp

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 获取标签详情
type KuaicheDmpNewTagDetailRequest struct {
	api.BaseRequest
	AccessPin    string `json:"accessPin,omitempty"` // 被免密访问的pin
	AuthType     string `json:"authType,omitempty"`  // 授权模式:0: 普通登录模式(无授权关系) 1:普通授权(不同商家pin互相授权) 2:主子pin关系授权 3:代理授权
	TagId        uint64 `json:"tagId"`               // 标签id
	CrowdId      int64  `json:"crowdId"`             // 人群ID，新建人群是-1，否则传已经创建的人群id
	IndustryHot  string `json:"industryHot"`         // 使用热度
	CoverageRate string `json:"coverageRate"`        // 覆盖分
}

type KuaicheDmpNewTagDetailResponse struct {
	Responce  *KuaicheDmpNewTagDetailResponce `json:"jingdong_dmp_new_tag_detail_responce,omitempty" codec:"jingdong_dmp_new_tag_detail_responce,omitempty"`
	ErrorResp *api.ErrorResponnse             `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDmpNewTagDetailResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpNewTagDetailResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpNewTagDetailResponce struct {
	Data *KuaicheDmpNewTagDetailResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                              `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDmpNewTagDetailResponce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheDmpNewTagDetailResponce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheDmpNewTagDetailResponseData struct {
	Data *dsp.TagDetail `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheDmpNewTagDetail(ctx context.Context, req *KuaicheDmpNewTagDetailRequest) (*dsp.TagDetail, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpNewTagDetailRequest()
	if req.AccessPin != "" {
		r.SetAccessPin(req.AccessPin)
	}
	if req.AuthType != "" {
		r.SetAuthType(req.AuthType)
	}
	r.SetTagId(req.TagId)
	r.SetCrowdId(req.CrowdId)
	r.SetIndustryHot(req.IndustryHot)
	r.SetCoverageRate(req.CoverageRate)

	var response KuaicheDmpNewTagDetailResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
