package coupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
	"github.com/daviddengcn/ljson"
)

type CouponReadGetCouponListRequest struct {
	api.BaseRequest
	Ip          string `json:"ip,omitempty" codec:"ip,omitempty"`
	Port        string `json:"port,omitempty" codec:"port,omitempty"`
	CouponId    uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`       // 优惠券编号
	Name        string `json:"name,omitempty" codec:"name,omitempty"`               // 优惠券名称（长度小于30）
	Type        string `json:"type,omitempty" codec:"type,omitempty"`               // 优惠券类型 0京券 1东券
	BindType    uint   `json:"bindType,omitempty" codec:"bindType,omitempty"`       // 绑定类型 1全店参加 2指定sku参加
	GrantType   uint   `json:"grantType,omitempty" codec:"grantType,omitempty"`     // 发放类型 3免费领取 5互动平台 【仅允许这两种】
	GrantWay    uint   `json:"grantWay,omitempty" codec:"grantWay,omitempty"`       // 推广方式 1卖家发放 2买家领取
	CreateMonth string `json:"createMonth,omitempty" codec:"createMonth,omitempty"` // 优惠券创建月份
	CreatorType string `json:"creatorType,omitempty" codec:"creatorType,omitempty"` // 优惠券创建者 0优惠券shop端 2促销发券等，实际调用100为忽略
	Closed      string `json:"closed,omitempty" codec:"closed,omitempty"`           // 店铺券状态 0未关闭 1关闭，实际调用100为忽略
	Page        uint   `json:"page,omitempty" codec:"page,omitempty"`
	PageSize    uint   `json:"pageSize,omitempty" codec:"pageSize,omitempty"` // 页面大小 取值范围[1,20]
}

type CouponReadGetCouponListResponse struct {
	ErrorResp *api.ErrorResponnse                  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponReadGetCouponListResponseData `json:"jingdong_seller_coupon_read_getCouponList_responce,omitempty" codec:"jingdong_seller_coupon_read_getCouponList_responce,omitempty"`
}

type CouponReadGetCouponListResponseData struct {
	List []*Coupon `json:"couponList,omitempty" codec:"couponList,omitempty"`
}

func CouponReadGetCouponList(req *CouponReadGetCouponListRequest) ([]*Coupon, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponReadGetCouponListRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	if req.CouponId > 0 {
		r.SetCouponId(req.CouponId)
	}

	if req.Type == "0" || req.Type == "1" {
		r.SetType(req.Type)
	}

	if req.BindType == 1 || req.BindType == 2 {
		r.SetBindType(req.BindType)
	}

	if req.GrantType >= 1 && req.GrantType <= 5 {
		r.SetGrantType(req.GrantType)
	}

	if req.GrantWay == 1 || req.GrantWay == 2 {
		r.SetGrantWay(req.GrantWay)
	}

	if req.Name != "" {
		r.SetName(req.Name)
	}

	if req.CreateMonth != "" {
		r.SetCreateMonth(req.CreateMonth)
	}

	if req.CreatorType == "0" || req.CreatorType == "2" {
		r.SetCreatorType(req.CreatorType)
	}

	if req.Closed == "0" || req.Closed == "1" {
		r.SetClosed(req.Closed)
	}

	r.SetPage(req.Page)
	r.SetPageSize(req.PageSize)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response CouponReadGetCouponListResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	return response.Data.List, nil
}
