package category

import (
	"encoding/json"
	"errors"
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

type FindCateByPidData struct {
	Code       string      `json:"code,omitempty" codec:"code,omitempty"`
	Categories []*Category `json:"categories,omitempty" codec:"categories,omitempty"`
}

// 获取单个SKU
func FindCateByPid(req *FindCateByPidRequest) ([]*Category, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := category.NewFindCateByPidRequest()
	if req.Fields != "" {
		r.SetFields(strings.Split(req.Fields, ","))
	}
	r.SetParentCid(req.ParentCid)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No cates info.")
	}
	var response FindCateByPidResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	return response.Data.Categories, nil
}
