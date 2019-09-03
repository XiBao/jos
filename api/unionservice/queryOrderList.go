package unionservice

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/ware"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/unionservice"
	"github.com/daviddengcn/ljson"
)

type QueryOrderListRequest struct {
	api.BaseRequest
	UnionId  uint64 `json:"union_id,omitempty" codec:"union_id,omitempty"`   // 站长Id
	Time     string `json:"time,omitempty" codec:"time,omitempty"`           // 查询时间, 格式yyyyMMddHH:2018012316 (按数据更新时间查询)
	Page     int    `json:"page,omitempty" codec:"page,omitempty"`           // 页数, 从1开始
	PageSize int    `json:"page_size,omitempty" codec:"page_size,omitempty"` // 每页条数, 上限500
}

type QueryOrderListResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryOrderListData `json:"jingdong_sku_read_findSkuById_responce,omitempty" codec:"jingdong_sku_read_findSkuById_responce,omitempty"`
}

type QueryOrderListData struct {
	Code string    `json:"code,omitempty" codec:"code,omitempty"`
	Sku  *ware.Sku `json:"sku,omitempty" codec:"sku,omitempty"`
}

// 获取单个SKU
func QueryOrderList(req *QueryOrderListRequest) (*ware.Sku, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := unionservice.NewUnionQueryOrderListRequest()
	r.SetUnionId(req.UnionId)
	r.SetTime(req.Time)
	r.SetPage(req.Page)
	r.SetPageSize(req.PageSize)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No sku info.")
	}
	var response QueryOrderListResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	return response.Data.Sku, nil
}
