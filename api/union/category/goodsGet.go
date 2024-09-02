package category

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/union/category"
)

type GoodsGetRequest struct {
	api.BaseRequest
	ParentId uint64 `json:"parentId"` // 父类目id(一级父类目为0)
	Grade    uint   `json:"grade"`    // 类目级别(类目级别 0，1，2 代表一、二、三级类目)
}

type GoodsGetResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty"`
	Data      *GoodsGetResponseData `json:"jd_union_open_category_goods_get_response,omitempty"`
}

func (r GoodsGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.Result == ""
}

func (r GoodsGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return "no result data"
}

type GoodsGetResponseData struct {
	Result string `json:"result,omitempty"`
}

type GoodsGetResult struct {
	Code    int64          `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    []CategoryResp `json:"data,omitempty"`
}

func (r GoodsGetResult) IsError() bool {
	return r.Code != 200
}

func (r GoodsGetResult) Error() string {
	return sdk.ErrorString(r.Code, r.Message)
}

type CategoryResp struct {
	Id       uint64 `json:"id,omitempty"`       // 类目Id
	Name     string `json:"name,omitempty"`     // 类目名称
	Grade    uint   `json:"grade,omitempty"`    // 类目级别(类目级别 0，1，2 代表一、二、三级类目)
	ParentId uint64 `json:"parentId,omitempty"` // 父类目Id
}

// 商品类目查询
func GoodsGet(ctx context.Context, req *GoodsGetRequest) ([]CategoryResp, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := category.NewGoodsGetRequest()
	goodsReq := &category.GoodsGetReq{
		ParentId: req.ParentId,
		Grade:    req.Grade,
	}
	r.SetReq(goodsReq)

	var response GoodsGetResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	var resp GoodsGetResult
	if err := client.Logger().DecodeJSON([]byte(response.Data.Result), &resp); err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, resp
	}

	return resp.Data, nil
}
