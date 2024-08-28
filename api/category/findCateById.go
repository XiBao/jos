package category

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/category"
)

type FindCateByIdRequest struct {
	api.BaseRequest

	Fields string `json:"fields,omitempty" codec:"fields,omitempty"` //
	Cid    uint64 `json:"cid,omitempty" codec:"cid,omitempty"`       // 自定义返回字段
}

type FindCateByIdResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FindCateByIdData   `json:"jingdong_category_read_findById_responce,omitempty" codec:"jingdong_category_read_findById_responce,omitempty"`
}

func (r FindCateByIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r FindCateByIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type FindCateByIdData struct {
	Code     string    `json:"code,omitempty" codec:"code,omitempty"`
	Category *Category `json:"category,omitempty" codec:"category,omitempty"`
}

// 获取单个SKU
func FindCateById(ctx context.Context, req *FindCateByIdRequest) (*Category, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := category.NewFindCateByIdRequest()
	if req.Fields != "" {
		r.SetFields(req.Fields)
	}
	r.SetCid(req.Cid)

	var response FindCateByIdResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Category, nil
}
