package coupon

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
)

type CouponWriteCloseRequest struct {
	api.BaseRequest
	Ip       string `json:"ip,omitempty" codec:"ip,omitempty"`             // 调用方IP
	Port     string `json:"port,omitempty" codec:"port,omitempty"`         // 调用方端口
	CouponId uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"` // 优惠券编号
}

type CouponWriteCloseResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponWriteCloseData `json:"jingdong_seller_coupon_write_close_responce,omitempty" codec:"jingdong_seller_coupon_write_close_responce,omitempty"`
}

func (r CouponWriteCloseResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CouponWriteCloseResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CouponWriteCloseData struct {
	Code        string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc   string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	CloseResult bool   `json:"close_result,omitempty" codec:"close_result,omitempty"`
}

func (r CouponWriteCloseData) IsError() bool {
	return r.Code != "0"
}

func (r CouponWriteCloseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CouponWriteClose(req *CouponWriteCloseRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponWriteCloseRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetCouponId(req.CouponId)

	var response CouponWriteCloseResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}

	return response.Data.CloseResult, nil
}
