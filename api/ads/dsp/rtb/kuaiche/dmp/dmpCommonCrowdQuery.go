package dmp

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 查询人群列表(包含京选和自建人群)
type KuaicheDmpCommonCrowdQueryRequest struct {
	api.BaseRequest
	AdGroupAdType       uint   `json:"adGroupAdType,omitempty"`       // 广告单元的类型：1推荐；2搜索
	AdGroupBidPrice     uint   `json:"adGroupBidPrice,omitempty"`     // 广告出价，用户出价乘以100w倍
	AdGroupBillingType  uint   `json:"adGroupBillingType,omitempty"`  // 广告出价类型
	AdGroupId           uint64 `json:"adGroupId,omitempty"`           // 广告单元ID，新建为空
	BusinessType        uint   `json:"businessType,omitempty"`        // 通用产品线类型
	CrowdName           string `json:"crowdName,omitempty"`           // 人群名称关键字默认为空
	CrowdTabType        uint   `json:"crowdTabType"`                  // 人群类型枚举值:1京选；2自建。京选人群不允许编辑，不应该调用crowd.detail接口
	FirstSenceCategory  string `json:"firstSenceCategory,omitempty"`  // 场景人群一级类目
	PageIndex           int    `json:"pageIndex"`                     // 当前页数
	PageSize            int    `json:"pageSize,omitempty"`            // 每页显示的条数
	SecondSenceCategory string `json:"secondSenceCategory,omitempty"` // 场景人群二级类目
	ResourcesList       string `json:"resourcesList,omitempty"`       // 广告位ID，逗号分隔
	AccessPin           string `json:"accessPin,omitempty"`           // 被免密访问的pin
	AuthType            string `json:"authType,omitempty"`            // 授权模式:0: 普通登录模式(无授权关系) 1:普通授权(不同商家pin互相授权) 2:主子pin关系授权 3:代理授权
}

type KuaicheDmpCommonCrowdQueryResponse struct {
	Responce  *KuaicheDmpCommonCrowdQueryResponce `json:"jingdong_dmp_common_crowd_query_responce,omitempty" codec:"jingdong_dmp_common_crowd_query_responce,omitempty"`
	ErrorResp *api.ErrorResponnse                 `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDmpCommonCrowdQueryResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpCommonCrowdQueryResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpCommonCrowdQueryResponce struct {
	Data *KuaicheDmpCommonCrowdQueryResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                                  `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDmpCommonCrowdQueryResponce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheDmpCommonCrowdQueryResponce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheDmpCommonCrowdQueryResponseData struct {
	Data *KuaicheDmpCommonCrowdQueryResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheDmpCommonCrowdQueryResponseDataData struct {
	PageTotal uint                      `json:"pageTotal"` // 总页数
	PageSize  uint                      `json:"pageSize"`  // 数据库偏移量，页面最多记录条数
	Count     uint                      `json:"count"`     // 总记录数
	PageIndex uint                      `json:"pageIndex"` // 页码
	Data      []dsp.DmpCommonCrowdQuery `json:"data"`
}

func KuaicheDmpCommonCrowdQuery(req *KuaicheDmpCommonCrowdQueryRequest) (*KuaicheDmpCommonCrowdQueryResponseDataData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpCommonCrowdQueryRequest()
	r.SetCrowdTabType(req.CrowdTabType)
	r.SetPageIndex(req.PageIndex)
	r.SetPageSize(req.PageSize)
	if req.AdGroupAdType > 0 {
		r.SetAdGroupAdType(req.AdGroupAdType)
	}
	if req.AdGroupBidPrice > 0 {
		r.SetAdGroupBidPrice(req.AdGroupBidPrice)
	}
	if req.AdGroupBillingType > 0 {
		r.SetAdGroupBillingType(req.AdGroupBillingType)
	}
	if req.AdGroupId > 0 {
		r.SetAdGroupId(req.AdGroupId)
	}
	if req.BusinessType > 0 {
		r.SetBusinessType(req.BusinessType)
	}
	if req.CrowdName != "" {
		r.SetCrowdName(req.CrowdName)
	}
	if req.FirstSenceCategory != "" {
		r.SetFirstSenceCategory(req.FirstSenceCategory)
	}
	if req.SecondSenceCategory != "" {
		r.SetSecondSenceCategory(req.SecondSenceCategory)
	}
	if req.ResourcesList != "" {
		r.SetResourcesList(req.ResourcesList)
	}
	if req.AccessPin != "" {
		r.SetAccessPin(req.AccessPin)
	}
	if req.AuthType != "" {
		r.SetAuthType(req.AuthType)
	}

	var response KuaicheDmpCommonCrowdQueryResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data, nil
}
