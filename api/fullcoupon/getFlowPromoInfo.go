package fullcoupon

import (
	"encoding/json"
	"errors"

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

type FullCouponGetFlowPromoInfoResponseResult struct {
	Result *FullCouponGetFlowPromoInfoResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

type FullCouponGetFlowPromoInfoResponseData struct {
	Msg     string                                      `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string                                      `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool                                        `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *FullCouponGetPromoListInfoResponseFlowList `json:"data,omitempty" codec:"data,omitempty"`
}

type FullCouponGetPromoListInfoResponseFlowList struct {
	Total          int         `json:"total" codec:"total"`
	TotalPageCount int         `json:"totalPageCount" codec:"totalPageCount"`
	PageIndex      int         `json:"pageIndex" codec:"pageIndex"`
	PageSize       int         `json:"pageSize" codec:"pageSize"`
	FlowList       []PromoFlow `json:"dataList" codec:"dataList"`
}

func GetFlowPromoInfo(req *FullCouponGetFlowPromoInfoRequest) ([]PromoFlow, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetFlowPromoInfoRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response FullCouponGetFlowPromoInfoResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data == nil || response.Data.Result == nil || response.Data.Result.Data == nil || response.Data.Result.Data.FlowList == nil {
		return nil, errors.New("no flow list.")
	}

	return response.Data.Result.Data.FlowList, nil
}
