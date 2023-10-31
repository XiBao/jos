package ware

import (
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

func (r FindWareByIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Ware == nil
}

func (r FindWareByIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
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

	var response FindWareByIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Ware, nil
}
