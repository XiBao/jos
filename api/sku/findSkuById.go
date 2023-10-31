package sku

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ware"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/sku"
)

type FindSkuByIdRequest struct {
	api.BaseRequest
	Fields string `json:"fields,omitempty" codec:"fields,omitempty"` //
	SkuId  uint64 `json:"sku_id,omitempty" codec:"sku_id,omitempty"` // 自定义返回字段
}

type FindSkuByIdResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FindSkuByIdData    `json:"jingdong_sku_read_findSkuById_responce,omitempty" codec:"jingdong_sku_read_findSkuById_responce,omitempty"`
}

func (r *FindSkuByIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Sku == nil
}

func (r *FindSkuByIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type FindSkuByIdData struct {
	Code string    `json:"code,omitempty" codec:"code,omitempty"`
	Sku  *ware.Sku `json:"sku,omitempty" codec:"sku,omitempty"`
}

// 获取单个SKU
func FindSkuById(req *FindSkuByIdRequest) (*ware.Sku, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := sku.NewFindSkuByIdRequest()
	if req.Fields != "" {
		r.SetFields(req.Fields)
	}
	r.SetSkuId(req.SkuId)

	var response FindSkuByIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Sku, nil
}
