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

type Coupon struct {
	Id            uint64            `json:"couponId,omitempty"`      // 优惠券ID
	VenderId      uint64            `json:"venderId,omitempty"`      // 商家ID
	LockType      uint              `json:"lockType,omitempty"`      // 优惠券锁定状态 4已锁定 1未锁定
	Name          string            `json:"name,omitempty"`          // 优惠券名称
	Type          uint              `json:"type"`                    // 优惠券类型 0京券 1东券
	BindType      uint              `json:"bindType,omitempty"`      // 绑定类型 1全店参加 2指定sku参加
	GrantType     uint              `json:"grantType,omitempty"`     // 发放类型 1促销绑定 2推送 3免费领取 4京豆换券 5互动平台
	Num           uint              `json:"num,omitempty"`           // 优惠券数量
	Discount      uint              `json:"discount,omitempty"`      // 优惠券面额
	Quota         uint              `json:"quota,omitempty"`         // 优惠限额
	ValidityType  uint              `json:"validityType,omitempty"`  // 有效期类型 1相对时间 5绝对时间
	Days          uint              `json:"days,omitempty"`          // 有效期
	BeginTime     int64             `json:"beginTime,omitempty"`     // 有效期开始时间
	EndTime       int64             `json:"endTime,omitempty"`       // 有效期结束时间
	Password      string            `json:"password,omitempty"`      // 发放密码
	RfId          uint64            `json:"rfId,omitempty"`          // 优惠券系统EVT_ID
	Member        uint              `json:"member,omitempty"`        // 会员等级 50注册会员 56铜牌 61银牌 62金牌 105钻石 110VIP 90企业会员
	TakeBeginTime int64             `json:"takeBeginTime,omitempty"` // 领券开始时间
	TakeEndTime   int64             `json:"takeEndTime,omitempty"`   // 领券结束时间
	TakeRule      uint              `json:"takeRule,omitempty"`      // 领券规则 5限领一张 4每天限领一张 3自定义每天限量数量
	TakeNum       uint              `json:"takeNum,omitempty"`       // 限制条件内可以领取张数
	Link          string            `json:"link,omitempty"`          // 领券链接
	ActivityRfId  uint64            `json:"activityRfId,omitempty"`  // ERP优惠券活动EVT_ID
	ActivityLink  string            `json:"activityLink,omitempty"`  // 活动链接
	UsedNum       uint              `json:"usedNum,omitempty"`       // 已使用数量
	SendNum       uint              `json:"sendNum,omitempty"`       // 已发放数量
	Deleted       bool              `json:"deleted"`                 // 关闭状态 0未关闭 1已关闭
	Display       uint              `json:"display,omitempty"`       // 是否公开 1不公开 3公开
	Created       int64             `json:"created,omitempty"`       // 优惠券创建时间
	PlatformType  int               `json:"platformType,omitempty"`  // 使用平台 1全平台 3限平台
	Platform      string            `json:"platform"`                // 优惠券使用平台 0主站专享 1手机专享 3M版京东 4手Q专享 5微信专享 7京致衣橱
	ImgUrl        string            `json:"imgUrl,omitempty"`        // 京豆换券，图片地址
	BoundStatus   uint              `json:"boundStatus,omitempty"`   // 京豆换券
	JdItem        uint              `json:"jdItem,omitempty"`        // 京豆数
	ItemId        uint64            `json:"itemId,omitempty"`        // 京豆换券项目ID
	ShareType     uint              `json:"shareType,omitempty"`     // 分享类型 1分享 2不分享
	Ext           map[string]string `json:"extMapInfo,omitempty"`    // 券扩展信息Map(当前key有：putKey，encLink，encMobileLink等,key与value均为String类型)
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
