package imgzone

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/imgzone"
)

type CategoryQueryRequest struct {
	api.BaseRequest
	CateId       int64  `json:"cate_id,omitempty" codec:"cate_id,omitempty"`               // 分类ID
	CateName     string `json:"cate_name,omitempty" codec:"cate_name,omitempty"`           // 分类名称，不支持模糊查询
	ParentCateId int64  `json:"parent_cate_id,omitempty" codec:"parent_cate_id,omitempty"` // 父分类ID，查询二级分类时为对应父分类id，查询一级分类时为0，查询全部分类的时候为空
}

type CategoryQueryResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CategoryQueryData  `json:"jingdong_imgzone_category_query_responce,omitempty" codec:"jingdong_imgzone_category_query_responce,omitempty"`
}

func (c CategoryQueryResponse) IsError() bool {
	return c.ErrorResp != nil || c.Data == nil || c.Data.IsError()
}

func (c CategoryQueryResponse) Error() string {
	if c.ErrorResp != nil {
		return c.ErrorResp.Error()
	}
	if c.Data != nil {
		return c.Data.Error()
	}
	return "no result data"
}

type CategoryQueryData struct {
	CateList   []Categroy `json:"cateList,omitempty" codec:"cateList,omitempty"`
	Code       string     `json:"code,omitempty" codec:"code,omitempty"`
	Desc       string     `json:"desc1,omitempty" codec:"desc1,omitempty"`
	ReturnCode int        `json:"return_code,omitempty" codec:"return_code,omitempty"`
}

func (c CategoryQueryData) IsError() bool {
	return c.ReturnCode != 1
}

func (c CategoryQueryData) Error() string {
	return sdk.ErrorString(c.ReturnCode, c.Desc)
}

// 根据分类id、分类名称、父分类等查询分类信息
func CategoryQuery(ctx context.Context, req *CategoryQueryRequest) ([]Categroy, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := imgzone.NewCategoryQueryRequest()
	if req.CateId > 0 {
		r.SetCateId(req.CateId)
	}
	if req.CateName != "" {
		r.SetCateName(req.CateName)
	}
	if req.ParentCateId > 0 {
		r.SetParentCateId(req.ParentCateId)
	}

	var response CategoryQueryResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.CateList, nil
}
