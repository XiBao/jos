package coupon

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/coupon"
)

type CouponWriteCreateRequest struct {
	api.BaseRequest
	Ip            string   `json:"ip,omitempty" codec:"ip,omitempty"`                       // 调用方IP
	Port          string   `json:"port,omitempty" codec:"port,omitempty"`                   // 调用方端口
	Name          string   `json:"name,omitempty" codec:"name,omitempty"`                   // 优惠券名称（长度小于30）
	Type          uint8    `json:"type,omitempty" codec:"type,omitempty"`                   // 优惠券类型 0京券 1东券
	BindType      uint8    `json:"bindType,omitempty" codec:"bindType,omitempty"`           // 绑定类型 1全店参加 2指定sku参加
	GrantType     uint8    `json:"grantType,omitempty" codec:"grantType,omitempty"`         // 发放类型 3免费领取 5互动平台 【仅允许这两种】
	Num           uint64   `json:"num,omitempty" codec:"num,omitempty"`                     // 优惠券数量（限制：(0,1000000000]）
	Discount      uint     `json:"discount,omitempty" codec:"discount,omitempty"`           // 优惠券面额（限制：[1,5000)）
	Quota         uint     `json:"quota,omitempty" codec:"quota,omitempty"`                 // 优惠限额（限制：[1,100000)）,京券无门槛传0
	ValidityType  uint8    `json:"validityType,omitempty" codec:"validityType,omitempty"`   // 有效期类型 1相对时间 5绝对时间
	Days          uint     `json:"days,omitempty" codec:"days,omitempty"`                   // 有效期(validityType为1时必填)
	BeginTime     uint64   `json:"beginTime,omitempty" codec:"beginTime,omitempty"`         // 有效期开始时间（validityType为5时必填）
	EndTime       uint64   `json:"endTime,omitempty" codec:"endTime,omitempty"`             // 有效期结束时间（validityType为5时必填）,有效期开始结束时间是券使用时间，必须包含在活动时间内，且小于 90 天
	Password      string   `json:"password,omitempty" codec:"password,omitempty"`           // 发放密码,(不用赋值)
	BatchKey      string   `json:"batchKey,omitempty" codec:"batchKey,omitempty"`           // 批次key(不赋值)
	Member        uint     `json:"member,omitempty" codec:"member,omitempty"`               // 会员等级 50注册会员 56铜牌 61银牌 62金牌 105钻石 UserClass=20000时有效
	PaidMembers   string   `json:"paidMembers,omitempty" codec:"paidMembers,omitempty"`     // 101,201 - 正式+试用，101-试用会员，201-正式会员
	TakeBeginTime uint64   `json:"takeBeginTime,omitempty" codec:"takeBeginTime,omitempty"` // 领券开始时间（晚于当前）
	TakeEndTime   uint64   `json:"takeEndTime,omitempty" codec:"takeBeginTime,omitempty"`   // 领券结束时间（活动时间最长90天）
	TakeRule      uint8    `json:"takeRule,omitempty" codec:"takeRule,omitempty"`           // 领券规则 5活动期间限领一张 4活动内每天限领一张 3自定义
	TakeNum       uint     `json:"takeNum,omitempty" codec:"takeNum,omitempty"`             // 限制条件内可以领取张数
	Display       uint8    `json:"display,omitempty" codec:"display,omitempty"`             // 是否公开 1不公开 3公开(grantType如设值5此参数必须为3)
	PlatformType  int8     `json:"platformType,omitempty" codec:"platformType,omitempty"`   // -1：全平台（体系：不限，应用：全部） 1：京东全平台（体系：京东商城，应用：京东商城） 3：京东限平台（体系：京东商城，应用：京东商城，平台：限平台） 4：京东到家（体系：京东到家，应用：全部，平台：京东客户端） 5：开普勒（体系：开普勒体系，应用：实际应用） 7：一号店（体系：一号店体系，应用：不限） 9：拍拍二手（体系：拍拍二手7，应用：不限） 15：特价街（体系：特价街体系，应用：不限） 18：药京采（体系：药京采，应用：不限） 24：医药店送（体系：医药店送体系，应用：不限） 26：区区购（体系：区区购体系，应用：不限）
	Platform      string   `json:"platform,omitempty" codec:"platform,omitempty"`           // platformType = 3 优惠券使用平台 0主站专享 1手机专享 3M版京东 4手Q专享 5微信专享 7京致衣橱（此参数需根据platformType设值，如限平台必填）
	ImgUrl        string   `json:"imgUrl,omitempty" codec:"imgUrl,omitempty"`               // 京豆换券，图片地址(不赋值)
	BoundStatus   uint     `json:"boundStatus,omitempty" codec:"boundStatus,omitempty"`     // 京豆换券(不赋值)
	JdNum         uint     `json:"jdNum,omitempty" codec:"jdNum,omitempty"`                 // 京豆数(不赋值)
	ItemId        uint64   `json:"itemId,omitempty" codec:"itemId,omitempty"`               // 京豆换券项目ID(不赋值)
	ShareType     uint8    `json:"shareType,omitempty" codec:"shareType,omitempty"`         // 分享类型 1分享 2不分享（如设置京券type=0,此参数必填2不分享）
	SkuId         []uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`                 // 商品sku编号(如设置bindType为2，此参数必填,需是有效sku)
	UserClass     int      `json:"userClass,omitempty" codec:"userClass,omitempty"`         // 会员类别 20000-普通会员，30000-付费会员，60000-新付费会员,70000-店铺粉丝会员
	ActivityLink  string   `json:"activityLink,omitempty" codec:"activityLink,omitempty"`   // 活动返回链接
	NumPerSending string   `json:"numPerSending,omitempty" codec:"numPerSending,omitempty"` // 单次发放张数，值是String，范围[1,10]；必须小于等于券数量num
}

type CouponWriteCreateResponse struct {
	ErrorResp *api.ErrorResponnse    `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponWriteCreateData `json:"jingdong_seller_coupon_write_create_responce,omitempty" codec:"jingdong_seller_coupon_write_create_responce,omitempty"`
}

func (r CouponWriteCreateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CouponWriteCreateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CouponWriteCreateData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	CouponId  uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`
}

func (r CouponWriteCreateData) IsError() bool {
	return r.Code != "0"
}

func (r CouponWriteCreateData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
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
	r.SetTakeBeginTime(req.TakeBeginTime)
	r.SetTakeEndTime(req.TakeEndTime)
	r.SetTakeRule(req.TakeRule)
	r.SetTakeNum(req.TakeNum)
	r.SetDisplay(req.Display)
	r.SetPlatformType(req.PlatformType)
	r.SetShareType(req.ShareType)

	if req.Member > 0 {
		r.SetMember(req.Member)
	}

	if len(req.PaidMembers) > 0 {
		r.SetPaidMembers(req.PaidMembers)
	}

	if req.Days > 0 {
		r.SetDays(req.Days)
	}

	if req.BeginTime > 0 && req.EndTime > 0 {
		r.SetBeginTime(req.BeginTime)
		r.SetEndTime(req.EndTime)
	}

	if req.PlatformType != 1 && len(req.Platform) > 0 {
		r.SetPlatform(req.Platform)
	}

	if len(req.SkuId) > 0 {
		r.SetSkuId(req.SkuId)
	}

	if req.UserClass > 0 {
		r.SetUserClass(req.UserClass)
	}

	if len(req.ActivityLink) > 0 {
		r.SetActivityLink(req.ActivityLink)
	}

	if req.NumPerSending != "" {
		r.SetNumPerSending(req.NumPerSending)
	}

	var response CouponWriteCreateResponse
	if err := client.PostExecute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.CouponId, nil
}
