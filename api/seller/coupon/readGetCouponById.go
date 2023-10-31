package coupon

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
)

type CouponReadGetCouponByIdRequest struct {
	api.BaseRequest
	Ip       string `json:"ip" codec:"ip"`
	Port     string `json:"port" codec:"port"`
	CouponId uint64 `json:"couponId" codec:"couponId"` // 优惠券ID
}

type CouponReadGetCouponByIdResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ResponseData       `json:"jingdong_seller_coupon_read_getCouponById_responce,omitempty" codec:"jingdong_seller_coupon_read_getCouponById_responce,omitempty"`
}

func (r CouponReadGetCouponByIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CouponReadGetCouponByIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ResponseData struct {
	Code      string  `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string  `json:"error_description,omitempty" codec:"error_description,omitempty"`
	JosCoupon *Coupon `json:"josCoupon,omitempty" codec:"josCoupon,omitempty"`
}

func (r ResponseData) IsError() bool {
	return r.Code != "0" || r.JosCoupon == nil
}

func (r ResponseData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

func CouponReadGetCouponById(req *CouponReadGetCouponByIdRequest) (*Coupon, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponReadGetCouponByIdRequest()
	r.SetCouponId(req.CouponId)

	if len(req.Ip) > 0 {
		r.SetIp(req.Ip)
	}

	if len(req.Port) > 0 {
		r.SetPort(req.Port)
	}

	var response CouponReadGetCouponByIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.JosCoupon, nil

}
