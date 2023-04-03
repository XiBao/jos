package fullcoupon

import (
	"github.com/XiBao/jos/sdk"
)

type CreateFullCouponRequest struct {
	Request *sdk.Request
}

type CreateFullCouponParamBusiPlatform struct {
	SelectType        int      `json:"selectType"`            // 一级推广平台的选择方式：0全部1限选(非全部渠道)
	ChannelList       []int    `json:"channelList,omitempty"` // 一级推广平台，对应的code值：可传多个渠道
	ChannelSelectType int      `json:"channelSelectType"`     // 二级推广平台选择方式：0全部 1限选，目前只有限选1
	Platform          []string `json:"platform,omitempty"`    // 推广平台：细分平台的code值
}

type CreateFullCouponParamMember struct {
	UserClass       int    `json:"userClass,omitempty"`       // 1级人群标识
	UserLevel       int    `json:"userLevel,omitempty"`       // 2级人群标志
	MemberLevelName string `json:"memberLevelName,omitempty"` // 品牌会员名称
	BrandId         int    `json:"brandId,omitempty"`         // 品牌ID
}

type CreateFullCouponParamOrderModesCouponInfo struct {
	Quota           uint   `json:"quota"`                    // 满减额度
	Discount        uint   `json:"discount"`                 // 减免金额
	ValidateDays    uint   `json:"validateDays,omitempty"`   // 有效天数
	CouponType      int    `json:"couponType"`               // 0-商品券 1-店铺券
	OrderDoneTime   string `json:"orderDoneTime"`            // 完成订单时间
	LimitFirstOrder int    `json:"limitFirstOrder"`          // 是否首单返券 1-是 2-否
	OrderDayLimit   uint   `json:"orderDayLimit"`            // 订单完成多少天内有效
	ValidityType    int    `json:"validityType"`             // 有效期类型，必传，传5为绝对时间类型，传1为相对时间类型，兼容目前不传时则默认赋值为1相对时间类型
	ValidBeginTime  string `json:"validBeginTime,omitempty"` // 券开始时间
	ValidEndTime    string `json:"validEndTime,omitempty"`   // 券结束时间
	StoreNum        uint   `json:"storeNum"`                 // 库存数量
}

type CreateFullCouponParamOrderModes struct {
	Quota      uint                                       `json:"quota"`
	CouponInfo *CreateFullCouponParamOrderModesCouponInfo `json:"couponInfo"`
}

type CreateFullCouponParam struct {
	AppKey       string                             `json:"appKey"`                 // ISV渠道Key
	Name         string                             `json:"name"`                   // 促销名称
	StartTime    string                             `json:"startTime"`              // 开始时间
	EndTime      string                             `json:"endTime"`                // 结束时间
	Link         string                             `json:"link,omitempty"`         // 推广链接
	SkuIds       []uint64                           `json:"skuIds"`                 // 主商品ID列表
	BusiCode     string                             `json:"busiCode,omitempty"`     // busiCode
	BusiPlatform *CreateFullCouponParamBusiPlatform `json:"busiPlatform,omitempty"` // 推广平台
	Member       *CreateFullCouponParamMember       `json:"member,omitempty"`       // 会员模型 1-普通会员 2-PLUS 3-PLUS+普通 6-企业会员 30-品牌会员 31-店铺粉丝
	OrderModes   []CreateFullCouponParamOrderModes  `json:"orderModes"`             // 促销规则信息
}

// create new request
func NewCreateFullCouponRequest() (req *CreateFullCouponRequest) {
	request := sdk.Request{MethodName: "jingdong.fullCoupon.createFullCoupon", Params: make(map[string]interface{}, 1)}
	req = &CreateFullCouponRequest{
		Request: &request,
	}
	return
}

func (req *CreateFullCouponRequest) SetParam(param *CreateFullCouponParam) {
	req.Request.Params["param"] = param
}

func (req *CreateFullCouponRequest) GetParam() *CreateFullCouponParam {
	param, found := req.Request.Params["param"]
	if found {
		return param.(*CreateFullCouponParam)
	}
	return nil
}
