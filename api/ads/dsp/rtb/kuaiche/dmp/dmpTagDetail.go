package dmp

import (
	"encoding/json"
	"errors"

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

type KuaicheDmpNewTagDetailResponce struct {
	Data *KuaicheDmpNewTagDetailResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                              `json:"code,omitempty" codec:"code,omitempty"`
}

type KuaicheDmpNewTagDetailResponseData struct {
	Data *dsp.TagDetail `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

func KuaicheDmpNewTagDetail(req *KuaicheDmpNewTagDetailRequest) (*dsp.TagDetail, error) {
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response KuaicheDmpNewTagDetailResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, errors.New(response.ErrorResp.ZhDesc)
	}
	if response.Responce == nil || response.Responce.Data == nil {
		return nil, errors.New("no result data.")
	}
	if !response.Responce.Data.Success {
		return nil, errors.New(response.Responce.Data.Msg)
	}

	return response.Responce.Data.Data, nil
}
