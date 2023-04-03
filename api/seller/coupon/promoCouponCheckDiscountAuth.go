package coupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
	"github.com/daviddengcn/ljson"
)

type PromoCouponCheckDiscountAuthRequest struct {
	api.BaseRequest
	AppName string `json:"appName"`
	AppId   string `json:"appId"`
	Uuid    string `json:"uuid"`
	Ip      string `json:"ip"`
}

type PromoCouponCheckDiscountAuthResponse struct {
	ErrorResp *api.ErrorResponnse               `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *PromoCouponCheckDiscountAuthData `json:"jingdong_pop_promo_coupon_checkDiscountAuth_responce,omitempty" codec:"jingdong_pop_promo_coupon_checkDiscountAuth_responce,omitempty"`
}

type PromoCouponCheckDiscountAuthData struct {
	Code      string                              `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                              `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *PromoCouponCheckDiscountAuthResult `json:"returnType,omitempty" codec:"AuthResult,omitempty"`
}

type PromoCouponCheckDiscountAuthResult struct {
	Code    string `json:"code,omitempty" codec:"code,omitempty"`
	Success bool   `json:"success,omitempty" codec:"success,omitempty"`
	Data    bool   `json:"data,omitempty" codec:"data,omitempty"`
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func PromoCouponCheckDiscountAuth(req *PromoCouponCheckDiscountAuthRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewPromoCouponCheckDiscountAuthRequest()
	r.SetAppName(req.AppName)
	if req.Ip != "" {
		r.SetIp(req.Ip)
	}
	if req.Uuid != "" {
		r.SetUuid(req.Uuid)
	}
	if req.AppId != "" {
		r.SetAppId(req.AppId)
	}
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}
	var response PromoCouponCheckDiscountAuthResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}
	if response.Data.Result.Code != "200" || !response.Data.Result.Success {
		if response.Data.Result.Msg == "" {
			return false, errors.New("未知错误")
		} else {
			return false, errors.New(response.Data.Result.Msg)
		}
	}

	if response.Data.Result == nil {
		return false, errors.New("no result.")
	}

	return response.Data.Result.Data, nil
}
