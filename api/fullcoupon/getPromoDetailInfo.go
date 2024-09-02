package fullcoupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 满额返券活动详情查询
type FullCouponGetPromoDetailInfoRequest struct {
	api.BaseRequest
	AppKey  string `json:"appKey" codec:"appKey"`   // ISV渠道key
	PromoId uint64 `json:"promoId" codec:"promoId"` // 促销ID
}

type FullCouponGetPromoDetailInfoResponse struct {
	ErrorResp *api.ErrorResponnse                         `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCouponGetPromoDetailInfoResponseResult `json:"jingdong_fullCoupon_getPromoDetailInfo_responce,omitempty" codec:"jingdong_fullCoupon_getPromoDetailInfo_responce,omitempty"`
}

func (r FullCouponGetPromoDetailInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FullCouponGetPromoDetailInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FullCouponGetPromoDetailInfoResponseResult struct {
	Result *FullCouponGetPromoDetailInfoResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

func (r FullCouponGetPromoDetailInfoResponseResult) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FullCouponGetPromoDetailInfoResponseResult) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FullCouponGetPromoDetailInfoResponseData struct {
	Msg          string            `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code         string            `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success      bool              `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	PromoDetails *PromoDetailsInfo `json:"data,omitempty" codec:"data,omitempty"`
}

func (r FullCouponGetPromoDetailInfoResponseData) IsError() bool {
	return r.PromoDetails == nil
}

func (r FullCouponGetPromoDetailInfoResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

func GetPromoDetailInfo(ctx context.Context, req *FullCouponGetPromoDetailInfoRequest) (*PromoDetailsInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetPromoDetailInfoRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)

	var response FullCouponGetPromoDetailInfoResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.PromoDetails, nil
}
