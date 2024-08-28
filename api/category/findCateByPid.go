package category

import (
	"context"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/category"
)

type FindCateByPidRequest struct {
	api.BaseRequest

	Fields    string `json:"fields,omitempty" codec:"fields,omitempty"`         //
	ParentCid uint64 `json:"parent_cid,omitempty" codec:"parent_cid,omitempty"` // 自定义返回字段
}

type FindCateByPidResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FindCateByPidData  `json:"jingdong_category_read_findByPId_responce,omitempty" codec:"jingdong_category_read_findByPId_responce,omitempty"`
}

func (r FindCateByPidResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r FindCateByPidResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type FindCateByPidData struct {
	Code       string     `json:"code,omitempty" codec:"code,omitempty"`
	Categories []Category `json:"categories,omitempty" codec:"categories,omitempty"`
}

// 获取单个SKU
func FindCateByPid(ctx context.Context, req *FindCateByPidRequest) ([]Category, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := category.NewFindCateByPidRequest()
	if req.Fields != "" {
		r.SetFields(strings.Split(req.Fields, ","))
	}
	r.SetParentCid(req.ParentCid)

	var response FindCateByPidResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Categories, nil
}
