package ware

import (
	"context"
	"strconv"
	"strings"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/ware"
)

type WareBaseProductGetRequest struct {
	api.BaseRequest
	Ids        []uint64 `json:"ids,omitempty" codec:"ids,omitempty"`                 //
	BaseFields string   `json:"base_fields,omitempty" codec:"base_fields,omitempty"` // 自定义返回字段
}

type WareBaseProductGetResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *WareBaseProductGetData `json:"jingdong_new_ware_baseproduct_get_responce,omitempty" codec:"jingdong_new_ware_baseproduct_get_responce,omitempty"`
}

func (r WareBaseProductGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil
}

func (r WareBaseProductGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type WareBaseProductGetData struct {
	Code   string         `json:"code,omitempty" codec:"code,omitempty"`
	Result []ProductsBase `json:"listproductbase_result,omitempty" codec:"listproductbase_result,omitempty"`
}

// 获取单个SKU
func WareBaseProductGet(ctx context.Context, req *WareBaseProductGetRequest) ([]ProductsBase, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := ware.NewWareBaseproductGetRequest()
	var ids []string
	for _, v := range req.Ids {
		ids = append(ids, strconv.FormatUint(v, 10))
	}
	r.SetIds(strings.Join(ids, ","))
	r.SetBaseFields(req.BaseFields)

	var response WareBaseProductGetResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
