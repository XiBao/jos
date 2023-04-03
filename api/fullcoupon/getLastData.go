package fullcoupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fullcoupon"
	"github.com/daviddengcn/ljson"
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

type FullCouponGetLastDataResponseResult struct {
	Result *FullCouponGetLastDataResponseResultData `json:"result,omitempty" codec:"result,omitempty"`
}

type FullCouponGetLastDataResponseResultData struct {
	Msg     string          `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string          `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool            `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Data    *PromoTrendData `json:"data,omitempty" codec:"data,omitempty"`
}

func GetLastData(req *FullCouponGetLastDataRequest) (*PromoTrendData, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fullcoupon.NewFullCouponGetLastDataRequest()
	r.SetAppKey(req.AppKey)
	r.SetPromoId(req.PromoId)
	r.SetShopId(req.ShopId)
	r.SetDate(req.Date)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response FullCouponGetLastDataResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data == nil || response.Data.Result == nil || response.Data.Result.Data == nil {
		return nil, errors.New("no data.")
	}

	return response.Data.Result.Data, nil
}
