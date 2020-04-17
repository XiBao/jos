package coupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
	"github.com/daviddengcn/ljson"
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

type CouponWriteCloseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	CloseResult bool `json:"close_result,omitempty" codec:"close_result,omitempty"`
}

func CouponWriteClose(req *CouponWriteCloseRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponWriteCloseRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetCouponId(req.CouponId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	result = util.RemoveJsonSpace(result)

	var response CouponWriteCloseResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return false, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.CloseResult, nil
}
