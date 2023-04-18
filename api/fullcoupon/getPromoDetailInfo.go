package fullcoupon

import (
	"encoding/json"
	"errors"

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

type FullCouponGetPromoDetailInfoResponseResult struct {
	Result *FullCouponGetPromoDetailInfoResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

type FullCouponGetPromoDetailInfoResponseData struct {
	Msg          string            `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code         string            `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success      bool              `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	PromoDetails *PromoDetailsInfo `json:"data,omitempty" codec:"data,omitempty"`
}

func GetPromoDetailInfo(req *FullCouponGetPromoDetailInfoRequest) (*PromoDetailsInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetPromoDetailInfoRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response FullCouponGetPromoDetailInfoResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data == nil || response.Data.Result == nil || response.Data.Result.PromoDetails == nil {
		return nil, errors.New("no promo details.")
	}

	return response.Data.Result.PromoDetails, nil
}
