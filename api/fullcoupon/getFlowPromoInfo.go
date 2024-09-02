package fullcoupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
)

// 满额返券促销详情查询审核信息
type FullCouponGetFlowPromoInfoRequest struct {
	api.BaseRequest
	AppKey  string `json:"appKey" codec:"appKey"`   // ISV渠道key
	PromoId uint64 `json:"promoId" codec:"promoId"` // 促销ID
}

type FullCouponGetFlowPromoInfoResponse struct {
	ErrorResp *api.ErrorResponnse                       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCouponGetFlowPromoInfoResponseResult `json:"jingdong_fullCoupon_getFlowPromoInfo_responce,omitempty" codec:"jingdong_fullCoupon_getFlowPromoInfo_responce,omitempty"`
}

func (r FullCouponGetFlowPromoInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FullCouponGetFlowPromoInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FullCouponGetFlowPromoInfoResponseResult struct {
	Result *FullCouponGetFlowPromoInfoResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

func (r FullCouponGetFlowPromoInfoResponseResult) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r FullCouponGetFlowPromoInfoResponseResult) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type FullCouponGetFlowPromoInfoResponseData struct {
	Msg     string                                      `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string                                      `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool                                        `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *FullCouponGetPromoListInfoResponseFlowList `json:"data,omitempty" codec:"data,omitempty"`
}

func (r FullCouponGetFlowPromoInfoResponseData) IsError() bool {
	return r.Data == nil
}

func (r FullCouponGetFlowPromoInfoResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

type FullCouponGetPromoListInfoResponseFlowList struct {
	Total          int         `json:"total" codec:"total"`
	TotalPageCount int         `json:"totalPageCount" codec:"totalPageCount"`
	PageIndex      int         `json:"pageIndex" codec:"pageIndex"`
	PageSize       int         `json:"pageSize" codec:"pageSize"`
	FlowList       []PromoFlow `json:"dataList" codec:"dataList"`
}

func GetFlowPromoInfo(ctx context.Context, req *FullCouponGetFlowPromoInfoRequest) ([]PromoFlow, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetFlowPromoInfoRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)

	var response FullCouponGetFlowPromoInfoResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data.FlowList, nil
}
