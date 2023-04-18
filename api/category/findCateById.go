package category

import (
	"encoding/json"
	"errors"

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
	Data      *FindCateByIdData   `json:"jingdong_category_read_findById_response,omitempty" codec:"jingdong_category_read_findById_response,omitempty"`
}

type FindCateByIdData struct {
	Code     string    `json:"code,omitempty" codec:"code,omitempty"`
	Category *Category `json:"category,omitempty" codec:"category,omitempty"`
}

// 获取单个SKU
func FindCateById(req *FindCateByIdRequest) (*Category, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := category.NewFindCateByIdRequest()
	if req.Fields != "" {
		r.SetFields(req.Fields)
	}
	r.SetCid(req.Cid)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No sku info.")
	}
	var response FindCateByIdResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	return response.Data.Category, nil
}
