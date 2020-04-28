package coupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
	"github.com/daviddengcn/ljson"
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

type CouponWriteLockData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	// LockResult string `json:"msg,omitempty" codec:"msg,omitempty"` // 调用成功无返回
}

func CouponWriteLock(req *CouponWriteLockRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponWriteLockRequest()
	r.SetPort(req.Port)
	r.SetRequestId(req.RequestId)
	r.SetTime(req.Time)
	r.SetPurpose(req.Purpose)
	r.SetOperateTime(req.OperateTime)
	r.SetCouponId(req.CouponId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	result = util.RemoveJsonSpace(result)

	var response CouponWriteLockResponse
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

	return true, nil
}
