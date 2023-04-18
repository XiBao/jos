package coupon

import (
	"encoding/json"
	"errors"

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

type CountResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Count     uint64 `json:"getcouponcount_result,omitempty" codec:"getcouponcount_result,omitempty"`
}

func CouponReadGetCouponCount(req *CouponReadGetCouponCountRequest) (uint64, error) {
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response CouponReadGetCouponCountResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Count, nil

}
