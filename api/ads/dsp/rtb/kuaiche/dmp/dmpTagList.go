package dmp

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ads/dsp"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ads/dsp/rtb/kuaiche/dmp"
)

// 获取标签列表
type KuaicheDmpTagListRequest struct {
	api.BaseRequest
	AccessPin       string `json:"accessPin,emitempty"`       // 被免密访问的pin
	AuthType        string `json:"authType,emitempty"`        // 授权模式:0: 普通登录模式(无授权关系) 1:普通授权(不同商家pin互相授权) 2:主子pin关系授权 3:代理授权
	CategoryId      string `json:"categoryId,emitempty"`      // 维度分类ID
	PageIndex       int    `json:"pageIndex"`                 // 页码
	PageSize        int    `json:"pageSize,omitempty"`        // 每页条数
	TagName         string `json:"tagName,omitempty"`         // 标签名称关键字：界面上输入搜索框内的文本
	SortType        uint   `json:"sortType,omitempty"`        //8行业热度降序(点击页面图标);9行业热度升序(点击页面图标);10覆盖度降序(点击页面图标);11覆盖度升序(点击页面图标),12相关性降序（点击搜索按钮），1标签创建时间降序（点击维度分类按此顺序）
	IsFavorite      int    `json:"isFavorite,omitempty"`      // -1 查询全部，1查询收藏列表。收藏列表时排序默认为1
	Level           uint   `json:"level,omitempty"`           // 类目级别
	TagCategoryType uint   `json:"tagCategoryType,omitempty"` // 1、类目对应列表（含新上标签类目） 2收藏列表 3、最近使用 4、同行常用 5、可能感兴趣 6、标签组合 ，收藏和最近使用列表排序默认为1，不支持其他类型排序
}

type KuaicheDmpTagListResponse struct {
	Responce  *KuaicheDmpTagListResponce `json:"jingdong_dmp_new_tag_list_v1_responce,omitempty" codec:"jingdong_dmp_new_tag_list_v1_responce,omitempty"`
	ErrorResp *api.ErrorResponnse        `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r KuaicheDmpTagListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r KuaicheDmpTagListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type KuaicheDmpTagListResponce struct {
	Data *KuaicheDmpTagListResponseData `json:"data,omitempty" codec:"data,omitempty"`
	Code string                         `json:"code,omitempty" codec:"code,omitempty"`
}

func (r KuaicheDmpTagListResponce) IsError() bool {
	return r.Data == nil || r.Data.IsError()
}

func (r KuaicheDmpTagListResponce) Error() string {
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type KuaicheDmpTagListResponseData struct {
	Data *KuaicheDmpTagListResponseDataData `json:"data,omitempty" codec:"data,omitempty"`
	dsp.DataCommonResponse
}

type KuaicheDmpTagListResponseDataData struct {
	TagPageInfo *KuaicheDmpTagListResponseDataTagPageInfo `json:"tagPageInfo,omitempty" codec:"tagPageInfo,omitempty"`
}

type KuaicheDmpTagListResponseDataTagPageInfo struct {
	PageIndex uint        `json:"pageIndex" codec:"pageIndex"`
	Count     uint        `json:"count" codec:"count"`
	PageTotal uint        `json:"pageTotal" codec:"pageTotal"`
	PageSize  uint        `json:"pageSize" codec:"pageSize"`
	TagVOList []dsp.TagVO `json:"data,omitempty" codec:"data,omitempty"`
}

func KuaicheDmpTagList(req *KuaicheDmpTagListRequest) (*KuaicheDmpTagListResponseDataTagPageInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := dmp.NewKuaicheDmpTagListRequest()
	if req.AccessPin != "" {
		r.SetAccessPin(req.AccessPin)
	}
	if req.AuthType != "" {
		r.SetAuthType(req.AuthType)
	}
	if req.CategoryId != "" {
		r.SetCategoryId(req.CategoryId)
	}
	if req.PageIndex > 0 {
		r.SetPageIndex(req.PageIndex)
	}
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	}
	if req.TagName != "" {
		r.SetTagName(req.TagName)
	}
	if req.SortType > 0 {
		r.SetSortType(req.SortType)
	}
	if req.IsFavorite != 0 {
		r.SetIsFavorite(req.IsFavorite)
	}
	if req.Level > 0 {
		r.SetLevel(req.Level)
	}
	if req.TagCategoryType > 0 {
		r.SetTagCategoryType(req.TagCategoryType)
	}

	var response KuaicheDmpTagListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.Data.Data.TagPageInfo, nil
}
