package sku

import (
	"encoding/json"
	"errors"

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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No sku info.")
	}
	var response FindSkuByIdResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	return response.Data.Sku, nil
}
