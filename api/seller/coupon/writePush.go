package coupon

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
)

type CouponWritePushRequest struct {
	api.BaseRequest
	Port      string `json:"port,omitempty" codec:"port,omitempty"`           // 调用方端口
	RequestId string `json:"requestId,omitempty" codec:"requestId,omitempty"` // 参数描述
	Pin       string `json:"pin,omitempty" codec:"pin,omitempty"`             // 用户pin(密文）
	DistrTime string `json:"distrTime,omitempty" codec:"distrTime,omitempty"` // 发券时间（yyyy-MM-dd hh:mm:ss）
	CouponId  uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`   // 优惠券ID
	Uuid      string `json:"uuid,omitempty" codec:"uuid,omitempty"`           // 发券唯一标识
}

type CouponWritePushResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponWritePushData `json:"jingdong_seller_coupon_write_pushCoupon_responce,omitempty" codec:"jingdong_seller_coupon_write_pushCoupon_responce,omitempty"`
}

func (r CouponWritePushResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CouponWritePushResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CouponWritePushData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	// PushResult bool `json:"msg,omitempty" codec:"msg,omitempty"` // 调用成功无返回
}

func (r CouponWritePushData) IsError() bool {
	return r.Code != "0"
}

func (r CouponWritePushData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CouponWritePush(req *CouponWritePushRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponWritePushRequest()
	r.SetPort(req.Port)
	r.SetRequestId(req.RequestId)
	r.SetPin(req.Pin)
	r.SetDistrTime(req.DistrTime)
	r.SetUuid(req.Uuid)
	r.SetCouponId(req.CouponId)

	var response CouponWritePushResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
