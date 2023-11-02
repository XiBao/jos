package promotion

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type LimitRequest struct {
	api.BaseRequest
	Ip         string `json:"ip,omitempty" codec:"ip,omitempty"`                   // 调用方IP
	Port       string `json:"port,omitempty" codec:"port,omitempty"`               // 调用方端口
	CategoryId uint64 `json:"category_id,omitempty" codec:"category_id,omitempty"` // 二级类目编号
	StartTime  string `json:"start_time,omitempty" codec:"start_time,omitempty"`   // 活动开始时间
	EndTime    string `json:"end_time,omitempty" codec:"end_time,omitempty"`       // 活动结束时间
}

type LimitResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *LimitData          `json:"jingdong_seller_promotion_v2_getPromoLimit_responce,omitempty" codec:"jingdong_seller_promotion_v2_getPromoLimit_responce,omitempty"`
}

func (r LimitResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r LimitResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type LimitData struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PLimit    *PromoLimit `json:"jos_promo_limit,omitempty" codec:"jos_promo_limit,omitempty"`
}

func (r LimitData) IsError() bool {
	return r.Code != "0" || r.PLimit == nil
}

func (r LimitData) Error() string {
	if r.Code != "0" {
		return sdk.ErrorString(r.Code, r.ErrorDesc)
	}
	return "no result data"
}

type PromoLimit struct {
	DiscountLimit float64 `json:"discount_limit,omitempty" codec:"discount_limit,omitempty"`
}

func Limit(req *LimitRequest) (*PromoLimit, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionLimitRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetCategoryId(req.CategoryId)
	r.SetStartTime(req.StartTime)
	r.SetEndTime(req.EndTime)

	var response LimitResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	return response.Data.PLimit, nil
}
