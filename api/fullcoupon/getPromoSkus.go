package fullcoupon

import (
	"encoding/json"
	"errors"

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

type FullCouponGetPromoSkusResponseResult struct {
	Result *FullCouponGetPromoSkusResponseData `json:"result,omitempty" codec:"result,omitempty"`
}

type FullCouponGetPromoSkusResponseData struct {
	Msg     string     `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string     `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool       `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    []PromoSku `json:"data,omitempty" codec:"data,omitempty"`
}

func GetPromoSkus(req *FullCouponGetPromoSkusRequest) ([]PromoSku, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetPromoSkusRequest()
	r.SetWareId(req.WareId)
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response FullCouponGetPromoSkusResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data == nil || response.Data.Result == nil || response.Data.Result.Data == nil {
		return nil, errors.New("no sku list.")
	}

	return response.Data.Result.Data, nil
}
