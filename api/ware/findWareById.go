package ware

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type FindWareByIdRequest struct {
	api.BaseRequest

	Fields string `json:"fields,omitempty" codec:"fields,omitempty"`   //
	WareId uint64 `json:"ware_id,omitempty" codec:"ware_id,omitempty"` // 自定义返回字段
}

type FindWareByIdResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FindWareByIdData   `json:"jingdong_ware_read_findWareById_responce,omitempty" codec:"jingdong_ware_read_findWareById_responce,omitempty"`
}

type FindWareByIdData struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Ware *Ware  `json:"ware,omitempty" codec:"ware,omitempty"`
}

// 获取单个SKU
func FindWareById(req *FindWareByIdRequest) (*Ware, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewFindWareByIdRequest()
	if req.Fields != "" {
		r.SetFields(req.Fields)
	}
	r.SetWareId(req.WareId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No ware info.")
	}
	var response FindWareByIdResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	return response.Data.Ware, nil
}
