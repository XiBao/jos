package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
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

type LimitData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	PLimit *PromoLimit `json:"jos_promo_limit,omitempty" codec:"jos_promo_limit,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	result = util.RemoveJsonSpace(result)

	var response LimitResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.PLimit, nil
}
