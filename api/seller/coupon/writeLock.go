package coupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
)

type CouponWriteLockRequest struct {
	api.BaseRequest
	Port        string `json:"port,omitempty" codec:"port,omitempty"`               // 调用方端口
	RequestId   string `json:"requestId,omitempty" codec:"requestId,omitempty"`     // 参数描述
	Time        uint64 `json:"time,omitempty" codec:"time,omitempty"`               // 请求时间，时间戳
	Purpose     string `json:"purpose,omitempty" codec:"purpose,omitempty"`         // 锁券原因
	OperateTime string `json:"operateTime,omitempty" codec:"operateTime,omitempty"` // 操作时间，格式yyyy-MM-dd HH:mm:ss
	CouponId    uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`       // 优惠券ID
}

type CouponWriteLockResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponWriteLockData `json:"jingdong_seller_coupon_write_lockCoupon_responce,omitempty" codec:"jingdong_seller_coupon_write_lockCoupon_responce,omitempty"`
}

func (r CouponWriteLockResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CouponWriteLockResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CouponWriteLockData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	// LockResult string `json:"msg,omitempty" codec:"msg,omitempty"` // 调用成功无返回
}

func (r CouponWriteLockData) IsError() bool {
	return r.Code != "0"
}

func (r CouponWriteLockData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CouponWriteLock(ctx context.Context, req *CouponWriteLockRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponWriteLockRequest()
	r.SetPort(req.Port)
	r.SetRequestId(req.RequestId)
	r.SetTime(req.Time)
	r.SetPurpose(req.Purpose)
	r.SetOperateTime(req.OperateTime)
	r.SetCouponId(req.CouponId)

	var response CouponWriteLockResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
