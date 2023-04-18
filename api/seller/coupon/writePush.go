package coupon

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
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

type CouponWritePushData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	// PushResult bool `json:"msg,omitempty" codec:"msg,omitempty"` // 调用成功无返回
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	result = util.RemoveJsonSpace(result)

	var response CouponWritePushResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return false, errors.New(response.Data.ErrorDesc)
	}

	return true, nil
}
