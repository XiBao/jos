package fullcoupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 满额返券促销分日效果数据
type FullCouponGetLastDataRequest struct {
	api.BaseRequest
	AppKey  string `json:"appKey" codec:"appKey"`   // ISV渠道key
	PromoId uint64 `json:"promoId" codec:"promoId"` // 促销ID
	ShopId  uint64 `json:"shopId" codec:"shopId"`   // 店铺ID
	Date    string `json:"Date" codec:"Date"`       // 查询的日期 yyyy-MM-dd
}

type FullCouponGetLastDataResponse struct {
	ErrorResp *api.ErrorResponnse                  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCouponGetLastDataResponseResult `json:"jingdong_fullCoupon_getLastData_responce,omitempty" codec:"jingdong_fullCoupon_getLastData_responce,omitempty"`
}

func (r FullCouponGetLastDataResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FullCouponGetLastDataResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FullCouponGetLastDataResponseResult struct {
	Result *FullCouponGetLastDataResponseResultData `json:"result,omitempty" codec:"result,omitempty"`
}

func (r FullCouponGetLastDataResponseResult) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FullCouponGetLastDataResponseResult) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FullCouponGetLastDataResponseResultData struct {
	Msg     string          `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string          `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool            `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *PromoTrendData `json:"data,omitempty" codec:"data,omitempty"`
}

func (r FullCouponGetLastDataResponseResultData) IsError() bool {
	return r.Data == nil
}

func (r FullCouponGetLastDataResponseResultData) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

func GetLastData(ctx context.Context, req *FullCouponGetLastDataRequest) (*PromoTrendData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetLastDataRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)
	r.SetShopId(req.ShopId)
	r.SetDate(req.Date)

	var response FullCouponGetLastDataResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data, nil
}
