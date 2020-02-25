package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
)

type CouponWriteCreateRequest struct {
	api.BaseRequest
	Ip       string `json:"ip,omitempty" codec:"ip,omitempty"`
	Port     string `json:"port,omitempty" codec:"port,omitempty"`
	CouponId uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"` // 优惠券编号
}

type CouponWriteCloseResponse struct {
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponWriteCloseResponseData `json:"jingdong_seller_coupon_write_close_response,omitempty" codec:"jingdong_seller_coupon_write_close_response,omitempty"`
}

type CouponWriteCloseResponseData struct {
	Result string `json:"close_result,omitempty" codec:"close_result,omitempty"`
}

func CouponWriteClose(req *CouponWriteCloseRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCouponWriteCloseRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetCouponId(req.CouponId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response CouponWriteCreateResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	return response.Data.Result, nil
}
