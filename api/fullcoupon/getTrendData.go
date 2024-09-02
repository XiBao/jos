package fullcoupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 满额送券促销趋势效果数据
type FullCouponGetTrendDataRequest struct {
	api.BaseRequest
	AppKey    string `json:"appKey" codec:"appKey"`       // ISV渠道key
	PromoId   uint64 `json:"promoId" codec:"promoId"`     // 促销ID
	ShopId    uint64 `json:"shopId" codec:"shopId"`       // 店铺ID
	StartDate string `json:"startDate" codec:"startDate"` // 开始日期 yyyy-MM-dd
	EndDate   string `json:"endDate" codec:"endDate"`     // 参数描述结束日期 yyyy-MM-dd
}

type FullCouponGetTrendDataResponse struct {
	ErrorResp *api.ErrorResponnse                   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCouponGetTrendDataResponseResult `json:"jingdong_fullCoupon_getTrendData_responce,omitempty" codec:"jingdong_fullCoupon_getTrendData_responce,omitempty"`
}

func (r FullCouponGetTrendDataResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FullCouponGetTrendDataResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FullCouponGetTrendDataResponseResult struct {
	Result *FullCouponGetTrendDataResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

func (r FullCouponGetTrendDataResponseResult) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FullCouponGetTrendDataResponseResult) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FullCouponGetTrendDataResponseData struct {
	Msg     string           `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string           `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool             `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    []PromoTrendData `json:"data,omitempty" codec:"data,omitempty"`
}

func (r FullCouponGetTrendDataResponseData) IsError() bool {
	return !r.Success
}

func (r FullCouponGetTrendDataResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

func GetTrendData(ctx context.Context, req *FullCouponGetTrendDataRequest) ([]PromoTrendData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetTrendDataRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)
	r.SetShopId(req.ShopId)
	r.SetStartDate(req.StartDate)
	r.SetEndDate(req.EndDate)

	var response FullCouponGetTrendDataResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data, nil
}
