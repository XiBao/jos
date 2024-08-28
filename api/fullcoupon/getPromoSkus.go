package fullcoupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 查询活动SPU信息 每页只支持查10条
type FullCouponGetPromoSkusRequest struct {
	api.BaseRequest
	WareId  uint64 `json:"wareId" codec:"wareId"`   // 商品ID
	AppKey  string `json:"appKey" codec:"appKey"`   // ISV渠道key
	PromoId uint64 `json:"promoId" codec:"promoId"` // 促销ID
}

type FullCouponGetPromoSkusResponse struct {
	ErrorResp *api.ErrorResponnse                   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCouponGetPromoSkusResponseResult `json:"jingdong_fullCoupon_getPromoSkus_responce,omitempty" codec:"jingdong_fullCoupon_getPromoSkus_responce,omitempty"`
}

func (r FullCouponGetPromoSkusResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FullCouponGetPromoSkusResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FullCouponGetPromoSkusResponseResult struct {
	Result *FullCouponGetPromoSkusResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

func (r FullCouponGetPromoSkusResponseResult) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FullCouponGetPromoSkusResponseResult) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FullCouponGetPromoSkusResponseData struct {
	Msg     string     `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string     `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool       `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    []PromoSku `json:"data,omitempty" codec:"data,omitempty"`
}

func (r FullCouponGetPromoSkusResponseData) IsError() bool {
	return !r.Success
}

func (r FullCouponGetPromoSkusResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

func GetPromoSkus(ctx context.Context, req *FullCouponGetPromoSkusRequest) ([]PromoSku, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetPromoSkusRequest()
	r.SetWareId(req.WareId)
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)

	var response FullCouponGetPromoSkusResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data, nil
}
