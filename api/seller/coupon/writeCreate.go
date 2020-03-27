package coupon

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
	"github.com/daviddengcn/ljson"
)

type CouponWriteCreateRequest struct {
	api.BaseRequest
	Ip            string `json:"ip,omitempty" codec:"ip,omitempty"`                       // 调用方IP
	Port          string `json:"string,omitempty" codec:"string,omitempty"`               // 调用方端口
	Name          string `json:"name,omitempty" codec:"name,omitempty"`                   // 优惠券名称（长度小于30）
	Type          uint8  `json:"type,omitempty" codec:"type,omitempty"`                   // 优惠券类型 0京券 1东券
	BindType      uint8  `json:"bindType,omitempty" codec:"bindType,omitempty"`           // 绑定类型 1全店参加 2指定sku参加
	GrantType     uint8  `json:"grantType,omitempty" codec:"grantType,omitempty"`         // 发放类型 3免费领取 5互动平台 【仅允许这两种】
	Num           uint64 `json:"num,omitempty" codec:"num,omitempty"`                     // 优惠券数量（限制：(0,1000000000]）
	Discount      uint   `json:"discount,omitempty" codec:"discount,omitempty"`           // 优惠券面额（限制：[1,5000)）
	Quota         uint   `json:"quota,omitempty" codec:"quota,omitempty"`                 // 优惠限额（限制：[1,100000)）,京券无门槛传0
	ValidityType  uint8  `json:"validityType,omitempty" codec:"validityType,omitempty"`   // 有效期类型 1相对时间 5绝对时间
	Days          uint   `json:"days,omitempty" codec:"days,omitempty"`                   // 有效期(validityType为1时必填)
	BeginTime     uint64 `json:"beginTime,omitempty" codec:"beginTime,omitempty"`         // 有效期开始时间（validityType为5时必填）
	EndTime       uint64 `json:"endTime,omitempty" codec:"endTime,omitempty"`             // 有效期结束时间（validityType为5时必填）,有效期开始结束时间是券使用时间，必须包含在活动时间内，且小于 90 天
	Password      string `json:"password,omitempty" codec:"password,omitempty"`           // 发放密码,(不用赋值)
	BatchKey      string `json:"batchKey,omitempty" codec:"batchKey,omitempty"`           // 批次key(不赋值)
	Member        uint   `json:"member,omitempty" codec:"member,omitempty"`               // 会员等级 50注册会员 56铜牌 61银牌 62金牌 105钻石 110VIP 90企业会员
	TakeBeginTime uint64 `json:"takeBeginTime,omitempty" codec:"takeBeginTime,omitempty"` // 领券开始时间（晚于当前）
	TakeEndTime   uint64 `json:"takeBeginTime,omitempty" codec:"takeBeginTime,omitempty"` // 领券结束时间（活动时间最长90天）
	TakeRule      uint8  `json:"takeRule,omitempty" codec:"takeRule,omitempty"`           // 领券规则 5活动期间限领一张 4活动内每天限领一张
	TakeNum       uint   `json:"takeNum,omitempty" codec:"takeNum,omitempty"`             // 限制条件内可以领取张数
	Display       uint8  `json:"display,omitempty" codec:"display,omitempty"`             // 是否公开 1不公开 3公开(grantType如设值5此参数必须为3)
	PlatformType  uint8  `json:"platformType,omitempty" codec:"platformType,omitempty"`   // 使用平台 1全平台 3限平台
	Platform      uint8  `json:"platform,omitempty" codec:"platform,omitempty"`           // 优惠券使用平台 0主站专享 1手机专享 3M版京东 4手Q专享 5微信专享 7京致衣橱（此参数需根据platformType设值，如限平台必填）
	ImgUrl        string `json:"imgUrl,omitempty" codec:"imgUrl,omitempty"`               // 京豆换券，图片地址(不赋值)
	BoundStatus   uint   `json:"boundStatus,omitempty" codec:"boundStatus,omitempty"`     // 京豆换券(不赋值)
	JdNum         uint   `json:"jdNum,omitempty" codec:"jdNum,omitempty"`                 // 京豆数(不赋值)
	ItemId        uint64 `json:"itemId,omitempty" codec:"itemId,omitempty"`               // 京豆换券项目ID(不赋值)
	ShareType     uint8  `json:"shareType,omitempty" codec:"shareType,omitempty"`         // 分享类型 1分享 2不分享（如设置京券type=0,此参数必填2不分享）
	SkuId         string `json:"skuId,omitempty" codec:"skuId,omitempty"`                 // 商品sku编号(如设置bindType为2，此参数必填,需是有效sku)
}

type CouponWriteCreateResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponWriteCreateData `json:"jingdong_seller_coupon_write_create_responce,omitempty" codec:"jingdong_seller_coupon_write_create_responce,omitempty"`
}

type CouponWriteCreateData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	CouponId uint64 `json:"coupon_id,omitempty" codec:"coupon_id,omitempty"`
}

func CouponWriteCreate(req *CouponWriteCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewSellerCouponWriteCreateRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetName(req.Name)
	r.SetType(req.Type)
	r.SetBindType(req.BindType)
	r.SetGrantType(req.GrantType)
	r.SetNum(req.Num)
	r.SetDiscount(req.Discount)
	r.SetQuota(req.Quota)
	r.SetValidityType(req.ValidityType)
	r.SetMember(req.Member)
	r.SetTakeBeginTime(req.TakeBeginTime)
	r.SetTakeEndTime(req.TakeEndTime)
	r.SetTakeRule(req.TakeRule)
	r.SetTakeNum(req.TakeNum)
	r.SetDisplay(req.Display)
	r.SetPlatformType(req.PlatformType)
	r.SetShareType(req.ShareType)

	if req.Days > 0 {
		r.SetDays(req.Days)
	}

	if req.BeginTime > 0 {
		r.SetBeginTime(req.BeginTime)
	}

	if req.EndTime > 0 {
		r.SetEndTime(req.EndTime)
	}

	if req.Platform > 0 {
		r.SetPlatform(req.Platform)
	}

	if len(req.SkuId) > 0 {
		r.SetSkuId(req.SkuId)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	result = util.RemoveJsonSpace(result)

	var response CouponWriteCreateResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.CouponId == 0 {
		return 0, errors.New("No coupon id.")
	}

	return response.Data.CouponId, nil
}
