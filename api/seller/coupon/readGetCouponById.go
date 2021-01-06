package coupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
	"github.com/daviddengcn/ljson"
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

type ResponseData struct {
	Code      string  `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string  `json:"error_description,omitempty" codec:"error_description,omitempty"`
	JosCoupon *Coupon `json:"josCoupon,omitempty" codec:"josCoupon,omitempty"`
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response CouponReadGetCouponByIdResponse
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

	return response.Data.JosCoupon, nil

}
