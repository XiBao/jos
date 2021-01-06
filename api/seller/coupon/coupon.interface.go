package coupon

import (
	coupon "github.com/XiBao/jos/sdk/request/seller/coupon"
)

type Coupon struct {
	Id                     uint64                                      `json:"couponId,omitempty"`               // 优惠券ID
	VenderId               uint64                                      `json:"venderId,omitempty"`               // 商家ID
	LockType               uint                                        `json:"lockType,omitempty"`               // 优惠券锁定状态 4已锁定 1未锁定
	Name                   string                                      `json:"name,omitempty"`                   // 优惠券名称
	Type                   uint                                        `json:"type"`                             // 优惠券类型 0京券 1东券
	BindType               uint                                        `json:"bindType,omitempty"`               // 绑定类型 1全店参加 2指定sku参加
	GrantType              uint                                        `json:"grantType,omitempty"`              // 发放类型 1促销绑定 2推送 3免费领取 4京豆换券 5互动平台
	Num                    uint                                        `json:"num,omitempty"`                    // 优惠券数量
	Discount               float64                                     `json:"discount,omitempty"`               // 优惠券面额
	Quota                  uint                                        `json:"quota,omitempty"`                  // 优惠限额
	ValidityType           uint                                        `json:"validityType,omitempty"`           // 有效期类型 1相对时间 5绝对时间
	Days                   uint                                        `json:"days,omitempty"`                   // 有效期
	BeginTime              int64                                       `json:"beginTime,omitempty"`              // 有效期开始时间
	EndTime                int64                                       `json:"endTime,omitempty"`                // 有效期结束时间
	Password               string                                      `json:"password,omitempty"`               // 发放密码
	BatchKey               string                                      `json:"batchKey,omitempty"`               // 批次key
	RfId                   uint64                                      `json:"rfId,omitempty"`                   // 优惠券系统EVT_ID
	Member                 uint                                        `json:"member,omitempty"`                 // 会员等级 50注册会员 56铜牌 61银牌 62金牌 105钻石 110VIP 90企业会员
	TakeBeginTime          int64                                       `json:"takeBeginTime,omitempty"`          // 领券开始时间
	TakeEndTime            int64                                       `json:"takeEndTime,omitempty"`            // 领券结束时间
	TakeRule               uint                                        `json:"takeRule,omitempty"`               // 领券规则 5限领一张 4每天限领一张 3自定义每天限量数量
	TakeNum                uint                                        `json:"takeNum,omitempty"`                // 限制条件内可以领取张数
	Link                   string                                      `json:"link,omitempty"`                   // 领券链接
	ActivityRfId           uint64                                      `json:"activityRfId,omitempty"`           // ERP优惠券活动EVT_ID
	ActivityLink           string                                      `json:"activityLink,omitempty"`           // 活动链接
	UsedNum                uint                                        `json:"usedNum,omitempty"`                // 已使用数量
	SendNum                uint                                        `json:"sendNum,omitempty"`                // 已发放数量
	Deleted                bool                                        `json:"deleted"`                          // 关闭状态 0未关闭 1已关闭
	Display                uint                                        `json:"display,omitempty"`                // 是否公开 1不公开 3公开
	Created                int64                                       `json:"created,omitempty"`                // 优惠券创建时间
	PlatformType           int                                         `json:"platformType,omitempty"`           // 使用平台 1全平台 3限平台
	Platform               string                                      `json:"platform"`                         // 优惠券使用平台 0主站专享 1手机专享 3M版京东 4手Q专享 5微信专享 7京致衣橱
	ImgUrl                 string                                      `json:"imgUrl,omitempty"`                 // 京豆换券，图片地址
	BoundStatus            uint                                        `json:"boundStatus,omitempty"`            // 京豆换券
	JdNum                  uint                                        `json:"jdNum,omitempty"`                  // 京豆数
	ItemId                 uint64                                      `json:"itemId,omitempty"`                 // 京豆换券项目ID
	ShareType              uint                                        `json:"shareType,omitempty"`              // 分享类型 1分享 2不分享
	UserClass              int                                         `json:"userClass,omitempty"`              // 会员类别 20000-普通会员，30000-付费会员，60000-新付费会员
	PaidMembers            string                                      `json:"paidMembers,omitempty"`            // 101,201 - 正式+试用，101-试用会员，201-正式会员
	CouponStatus           int                                         `json:"couponStatus,omitempty"`           // 提交状态 -1未提交 1已提交
	MobileLink             string                                      `json:"mobileLink,omitempty"`             // 移动端领券链接
	SkuId                  []uint64                                    `json:"skuId,omitempty"`                  // 商品sku编号(如设置bindType为2，此参数必填,需是有效sku)
	ExtMapInfo             map[string]string                           `json:"extMapInfo,omitempty"`             // 券扩展信息Map(当前key有：putKey，encLink，encMobileLink等,key与value均为String类型)
	ChannelStock           map[string]interface{}                      `json:"channelStock,omitempty"`           // 其他营销渠道及渠道投放量，key渠道，value投放量
	Style                  uint8                                       `json:"style,omitempty"`                  // 优惠类型 1-满减，3-满折
	MaxDiscount            uint                                        `json:"maxDiscount,omitempty"`            // 最高折扣
	Stairs                 []*coupon.SellerPromoCouponCreateParamStair `json:"stairs,omitempty"`                 // 折扣阶梯
	RiskLevel              int                                         `json:"riskLevel,omitempty"`              // 风控等级
	ActivityUserGetTopType uint8                                       `json:"activityUserGetTopType,omitempty"` // 用户领取上限类型0-限制，1-不限
	UserGetTop             uint                                        `json:"userGetTop,omitempty"`             // 用户领取最大数
	IsSpecialChannel       bool                                        `json:"isSpecialChannel,omitempty"`       // 是否是独立渠道
	IsHourCoupon           bool                                        `json:"isHourCoupon,omitempty"`           // 是否是小时券
	ToLink                 string                                      `json:"toLink,omitempty"`                 // 立即使用跳转链接
	UserLabel              string                                      `json:"userLabel,omitempty"`              // 用户标签
	BusinessLabel          string                                      `json:"businessLabel,omitempty"`          // 业务类型标签
	ShopId                 uint64                                      `json:"shopId,omitempty"`                 // 店铺ID
	OffPage                uint                                        `json:"off_page,omitempty"`
}

type CouponSkuList struct {
	SkuId     uint64  `json:"skuId" codec:"skuId"`
	WareId    uint64  `json:"wareId" codec:"wareId"`
	BindType  uint8   `json:"bindType" codec:"bindType"`
	VenderId  uint64  `json:"venderId" codec:"venderId"`
	CouponId  uint64  `json:"couponId" codec:"couponId"`
	WareTitle string  `json:"wareTitle" codec:"wareTitle"`
	Logo      string  `json:"logo" codec:"logo"`
	SkuName   string  `json:"skuName" codec:"skuName"`
	StockNum  int     `json:"stockNum,omitempty" codec:"stockNum,omitempty"`
	JdPrice   float64 `json:"jdPrice" codec:"jdPrice"`
}
