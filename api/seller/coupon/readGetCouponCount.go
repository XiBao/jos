package coupon

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
)

type CouponReadGetCouponCountRequest struct {
	api.BaseRequest
	Ip          string `json:"ip" codec:"ip"`
	Port        string `json:"port" codec:"port"`
	CouponId    uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`       // 优惠券ID
	Type        string `json:"type,omitempty" codec:"type,omitempty"`               // 优惠券类型 0京券 1东券
	GrantType   uint8  `json:"grantType,omitempty" codec:"grantType,omitempty"`     // 发放类型 1促销绑定 2推送 3免费领取 4京豆换券 5互动平台
	BindType    uint8  `json:"bindType,omitempty" codec:"bindType,omitempty"`       // 绑定类型 1全店参加 2指定sku参加
	GrantWay    uint8  `json:"grantWay,omitempty" codec:"grantWay,omitempty"`       // 推广方式 1卖家发放 2买家领取
	Name        string `json:"name,omitempty" codec:"name,omitempty"`               // 促销名称
	CreateMonth string `json:"createMonth,omitempty" codec:"createMonth,omitempty"` // 优惠券创建月份,格式（YYYY-MM）
	CreatorType string `json:"creatorType,omitempty" codec:"creatorType,omitempty"` // 优惠券创建者 0优惠券shop端 2促销发券等
	Closed      string `json:"closed,omitempty" codec:"closed,omitempty"`           // 店铺券状态 0未关闭 1关闭
}

type CouponReadGetCouponCountResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CountResponseData  `json:"jingdong_seller_coupon_read_getCouponCount_responce,omitempty" codec:"jingdong_seller_coupon_read_getCouponCount_responce,omitempty"`
}

func (r CouponReadGetCouponCountResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CouponReadGetCouponCountResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CountResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Count     uint64 `json:"getcouponcount_result,omitempty" codec:"getcouponcount_result,omitempty"`
}

func (r CountResponseData) IsError() bool {
	return r.Code != "0"
}

func (r CountResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func CouponReadGetCouponCount(ctx context.Context, req *CouponReadGetCouponCountRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponReadGetCouponCountRequest()

	if len(req.Ip) > 0 {
		r.SetIp(req.Ip)
	}

	if len(req.Port) > 0 {
		r.SetPort(req.Port)
	}

	if req.CouponId > 0 {
		r.SetCouponId(req.CouponId)
	}

	if len(req.Name) > 0 {
		r.SetName(req.Name)
	}

	if len(req.Type) > 0 {
		r.SetType(req.Type)
	}

	if req.GrantType > 0 {
		r.SetGrantType(req.GrantType)
	}

	if req.GrantWay > 0 {
		r.SetGrantWay(req.GrantWay)
	}

	if len(req.CreateMonth) > 0 {
		r.SetCreateMonth(req.CreateMonth)
	}

	if len(req.CreatorType) > 0 {
		r.SetCreatorType(req.CreatorType)
	}

	if len(req.Closed) > 0 {
		r.SetClosed(req.Closed)
	}

	if req.BindType > 0 {
		r.SetBindType(req.BindType)
	}

	var response CouponReadGetCouponCountResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Count, nil
}
