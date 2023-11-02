package coupon

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
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

func (r PromoCouponCheckDiscountAuthResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r PromoCouponCheckDiscountAuthResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type PromoCouponCheckDiscountAuthData struct {
	Code      string                              `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                              `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *PromoCouponCheckDiscountAuthResult `json:"returnType,omitempty" codec:"AuthResult,omitempty"`
}

func (r PromoCouponCheckDiscountAuthData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r PromoCouponCheckDiscountAuthData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type PromoCouponCheckDiscountAuthResult struct {
	Code    string `json:"code,omitempty" codec:"code,omitempty"`
	Success bool   `json:"success,omitempty" codec:"success,omitempty"`
	Data    bool   `json:"data,omitempty" codec:"data,omitempty"`
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`
}

func (r PromoCouponCheckDiscountAuthResult) IsError() bool {
	return r.Code != "200" || !r.Success
}

func (r PromoCouponCheckDiscountAuthResult) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
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

	var response PromoCouponCheckDiscountAuthResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
