package coupon

import (
	"github.com/XiBao/jos/sdk"
)

type SellerPromoCouponCreateRequest struct {
	Request *sdk.Request
}

type SellerPromoCouponCreateClientInfoExtentsion struct {
	Value1 string `json:"value1"`
}

type SellerPromoCouponCreateClientInfo struct {
	AppName          string                                       `json:"appName"`
	BusinessIdentity string                                       `json:"businessIdentity"`
	AppId            string                                       `json:"appId,omitempty"`
	UserAgent        string                                       `json:"userAgent,omitempty"`
	Uuid             string                                       `json:"uuid,omitempty"`
	Ip               string                                       `json:"ip,omitempty"`
	Site             string                                       `json:"site,omitempty"`
	Language         string                                       `json:"language,omitempty"`
	Extension        *SellerPromoCouponCreateClientInfoExtentsion `json:"extension,omitempty"`
}

type SellerPromoCouponCreateParamStair struct {
	Quota    uint    `json:"quota"`    // 阶梯限额，阶梯序号递增，限额递增
	Discount float64 `json:"discount"` // 阶梯折扣1、折扣券8折，discount 0.8 * 2、阶梯序号递增，折扣递减
}

type SellerPromoCouponCreateParamSku struct {
	SkuId uint64 `json:"skuId"`
}

type SellerPromoCouponCreateParamSpu struct {
	SkuId uint64 `json:"skuId"`
}

type SellerPromoCouponCreateParamPlatform struct {
	SelectType        uint8  `json:"selectType"`         // 0全部1限选，对应
	Channel           int8   `json:"channel,omitempty"`  // 推广渠道：QQ平台、1号店、开普勒渠道
	ChannelSelectType uint8  `json:"channelSelectType"`  // 广平台下渠道选择方式 0全部 1限选，如 0：QQ平台或开普勒渠道下所有渠道
	Platform          string `json:"platform,omitempty"` // 推广平台，多个之间用逗号隔开
}

type SellerPromoCouponCreateParamMember struct {
	UserClass uint `json:"userClass,omitempty"` // 会员类别
	UserLevel uint `json:"userLevel,omitempty"` // 普通会员等级
}

type SellerPromoCouponCreateParam struct {
	WareChoseType     uint8                                 `json:"wareChoseType"`               // 选择商品类型；1:店铺券，2：商品券
	ValidityType      uint8                                 `json:"validityType"`                // 有效期类型 5、绝对时间，1、相对时间
	TakeRule          uint8                                 `json:"takeRule"`                    // 领券规则 0不限制 1活动期间每人限领1张 2每人每天限令1张 3活动期间每人限领几张
	Type              uint8                                 `json:"type"`                        // 优惠券类型 0：无门槛券(京券)；1：东券
	EndTime           string                                `json:"endTime,omitempty"`           // 有效期结束时间，validityType=5时必填
	MainQuota         uint                                  `json:"mainQuota,omitempty"`         // 优惠限额，style=1时，必填
	Style             uint8                                 `json:"style"`                       // 优惠方式 1：满减券； 3：满折券
	Name              string                                `json:"name"`                        // 优惠券名称
	TakeEndTime       string                                `json:"takeEndTime"`                 // 领券结束时间
	BeginTime         string                                `json:"beginTime,omitempty"`         // 有效期开始时间，validityType=5时必填
	Display           uint8                                 `json:"display"`                     // 是否公开 1：不公开（隐藏）；3：公开（显示）
	ActivityLink      string                                `json:"activityLink,omitempty"`      // 活动链接
	OfficialType      uint8                                 `json:"officialType,omitempty"`      // 官方渠道下的1：互动营销；2：咚咚群聊；3：有价优惠券；4：cps联盟渠道 5：直播专享券
	MainDiscount      uint                                  `json:"mainDiscount,omitempty"`      // 优惠券面额，style=1时，必填
	TakeNum           uint                                  `json:"takeNum,omitempty"`           // 领取数量
	Num               uint64                                `json:"num,omitempty"`               // 优惠券数量
	TakeBeginTime     string                                `json:"takeBeginTime"`               // 领券开始时间
	Days              uint                                  `json:"days,omitempty"`              // 有效期(领券后N天)，validityType=1时必填
	High              uint                                  `json:"high,omitempty"`              // 有效期类型 5、绝对时间，1、最高折扣 style=1时 high默认值等于discount；style=3时，high需要用法设置，并且大于最高折扣力度
	Spus              []SellerPromoCouponCreateParamSpu     `json:"spus,omitempty"`              // 参与活动spu集合
	ShareType         uint8                                 `json:"shareType"`                   // 是否可赠送(分享)：1：可赠送；2：不可赠送
	PromoteChannel    uint8                                 `json:"promoteChannel"`              // 全网渠道：1 ；官方渠道：2
	Stairs            []SellerPromoCouponCreateParamStair   `json:"stairs,omitempty"`            // 折扣阶梯，规则:style=3时 非空；并且满足折扣递减、满额递增
	Skus              []SellerPromoCouponCreateParamSku     `json:"skus,omitempty"`              // skuId集合
	BusiPlatformParam *SellerPromoCouponCreateParamPlatform `json:"busiPlatformParam,omitempty"` // 渠道
	MemberParam       *SellerPromoCouponCreateParamMember   `json:"memberParam,omitempty"`       // 人群
	BusiCode          string                                `json:"busiCode"`                    // 场景code
}

// create new request
func NewSellerPromoCouponCreateRequest() (req *SellerPromoCouponCreateRequest) {
	request := sdk.Request{MethodName: "jingdong.pop.promo.coupon.createCoupon", Params: make(map[string]interface{}, 2)}
	req = &SellerPromoCouponCreateRequest{
		Request: &request,
	}
	return
}

func (req *SellerPromoCouponCreateRequest) SetClientInfo(clientInfo *SellerPromoCouponCreateClientInfo) {
	req.Request.Params["clientInfo"] = clientInfo
}

func (req *SellerPromoCouponCreateRequest) GetClientInfo() *SellerPromoCouponCreateClientInfo {
	client, found := req.Request.Params["clientInfo"]
	if found {
		return client.(*SellerPromoCouponCreateClientInfo)
	}
	return nil
}

func (req *SellerPromoCouponCreateRequest) SetParam(param *SellerPromoCouponCreateParam) {
	req.Request.Params["couponOuterParam"] = param
}

func (req *SellerPromoCouponCreateRequest) GetParam() *SellerPromoCouponCreateParam {
	param, found := req.Request.Params["couponOuterParam"]
	if found {
		return param.(*SellerPromoCouponCreateParam)
	}
	return nil
}
