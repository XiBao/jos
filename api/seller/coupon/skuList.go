package coupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	coupon "github.com/XiBao/jos/sdk/request/seller/coupon"
)

type SkuListRequest struct {
	api.BaseRequest
	Port         string `json:"port" codec:"port"`
	Ip           string `json:"ip" codec:"ip"`
	AppKey       string `json:"appKey" codec:"appKey"`
	Page         int    `json:"page" codec:"page"`
	SearchStatus int    `json:"searchStatus,omitempty" codec:"searchStatus,omitempty"`
	PageSize     int    `json:"page_size" codec:"page_size"`
	SkuId        uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`
	WareId       uint64 `json:"wareId,omitempty" codec:"wareId,omitempty"`
	CouponId     uint64 `json:"couponId" codec:"couponId"`
}

type SkuListResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SkuListResponseData `json:"jingdong_seller_coupon_read_getCouponSkuList_responce,omitempty" codec:"jingdong_seller_coupon_read_getCouponSkuList_responce,omitempty"`
}

func (r SkuListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r SkuListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type SkuListResponseData struct {
	Code          string          `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string          `json:"error_description,omitempty" codec:"error_description,omitempty"`
	CouponSkuList []CouponSkuList `json:"couponSkuList,omitempty" codec:"couponSkuList,omitempty"`
}

func (r SkuListResponseData) IsError() bool {
	return r.Code != "0"
}

func (r SkuListResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 优惠券商品查询
func SkuList(ctx context.Context, req *SkuListRequest) ([]CouponSkuList, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponSkuListRequest()
	r.SetCouponId(req.CouponId)
	r.SetPort(req.Port)
	r.SetIp(req.Ip)

	if req.SearchStatus != 0 {
		r.SetSearchStatus(req.SearchStatus)
	}

	if req.SkuId > 0 {
		r.SetSkuId(req.SkuId)
	}

	if req.WareId > 0 {
		r.SetWareId(req.WareId)
	}

	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	} else {
		r.SetPageSize(20)
	}

	if req.Page > 0 {
		r.SetPage(req.Page)
	} else {
		r.SetPage(1)
	}

	var response SkuListResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.CouponSkuList, nil
}
