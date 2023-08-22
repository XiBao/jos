package coupon

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/market/coupon"
)

type MarketCreateCouponRequest struct {
	api.BaseRequest
	WareGrade         uint8    `json:"wareGrade,omitempty" codec:"wareGrade,omitempty"`           // 参加商品级别，sku维度  1:SKU级，2:SPU级
	Num               uint64   `json:"num,omitempty" codec:"num,omitempty"`                       // 优惠券发放数量
	Discount          uint     `json:"discount,omitempty" codec:"discount,omitempty"`             // 优惠券面额，减金额，style=1时，必填
	StrategyParam     string   `json:"strategyParam,omitempty" codec:"strategyParam,omitempty"`   // 策略参数
	Type              uint8    `json:"type" codec:"type"`                                         // 优惠券类型 0：无门槛券(京券)；1：东券
	SkuIdList         string   `json:"skuIdList,omitempty" codec:"skuIdList,omitempty"`           // skuId
	CouponId          uint64   `json:"couponId,omitempty" codec:"couponId,omitempty"`             // 优惠券ID
	ShareType         uint8    `json:"shareType,omitempty" codec:"shareType,omitempty"`           // 是否可赠送(分享)：1：可赠送；2：不可赠送
	StoreId           []string `json:"storeId,omitempty" codec:"storeId,omitempty"`               // 门店列表
	TakeEndTime       string   `json:"takeEndTime,omitempty" codec:"takeBeginTime,omitempty"`     // 领券结束时间，示例：2022-11-03 00:00:00
	High              uint     `json:"high,omitempty" codec:"high,omitempty"`                     // 最高折扣 style=1时 high默认值等于discount；style=3时，high需要用法设置，并且大于最高折扣力度
	TakeNum           uint     `json:"takeNum,omitempty" codec:"takeNum,omitempty"`               // 领取数量
	Quota             float64  `json:"quota,omitempty" codec:"quota,omitempty"`                   // 优惠限额，满金额，style=1时，必填
	OfficialType      uint8    `json:"officialType,omitempty" codec:"officialType,omitempty"`     // 推广方式下二级官方营销推广子类型：1：互动营销；2：咚咚群聊；3：有价优惠券；4：cps联盟渠道 5：直播专享券
	BeginTime         string   `json:"beginTime,omitempty" codec:"beginTime,omitempty"`           // 使用开始时间，validityType=5时必填，示例：2022-11-03 00:00:00
	PromoteChannel    uint8    `json:"promoteChannel,omitempty" codec:"promoteChannel,omitempty"` // 推广方式一级：全网自动渠道：1 ；官方营销推广：2
	RemainNum         uint     `json:"remainNum,omitempty" codec:"remainNum,omitempty"`           // 剩余数量
	StoreType         uint     `json:"storeType,omitempty" codec:"storeType,omitempty"`           // 门店类型
	Display           uint8    `json:"display,omitempty" codec:"display,omitempty"`               // 是否公开 1：不公开（隐藏）；3：公开（显示）
	BusiCode          string   `json:"busiCode,omitempty" codec:"busiCode,omitempty"`             // 场景code，后端提供
	WareChoseType     uint8    `json:"wareChoseType,omitempty" codec:"wareChoseType,omitempty"`   // 选择商品类型；1:店铺券(全店商品参与)，2：部分商品参与，3：部分商品不参与
	UserClass         int      `json:"userClass,omitempty" codec:"userClass,omitempty"`           // 会员类别 20000-普通会员，30000-付费会员，60000-新付费会员
	UserLevel         int      `json:"userLevel,omitempty" codec:"userLevel,omitempty"`           // 2级人群标志
	TakeBeginTime     string   `json:"takeBeginTime,omitempty" codec:"takeBeginTime,omitempty"`   // 领券开始时间，时间格式，示例：2022-11-03 00:00:00
	ValidityType      uint8    `json:"validityType,omitempty" codec:"validityType,omitempty"`     // 有效期类型 5、绝对时间，1、相对时间
	TakeRule          uint8    `json:"takeRule" codec:"takeRule"`                                 // 领券规则 0不限制 1每人限领1张 2每人每天限领1张 3每人限领几张
	HourCoupon        bool     `json:"hourCoupon" codec:"hourCoupon"`                             // 是否是小时券，使用初始化的值
	Name              string   `json:"name,omitempty" codec:"name,omitempty"`                     // 优惠券名称
	ActivityLink      string   `json:"activityLink,omitempty" codec:"activityLink,omitempty"`     // 活动返回链接
	Days              uint     `json:"days,omitempty" codec:"days,omitempty"`                     // 有效期(领券后N天)，validityType=1时必填
	Style             uint8    `json:"style,omitempty" codec:"style,omitempty"`                   // 优惠方式 1：满减券； 3：满折券
	EndTime           string   `json:"endTime,omitempty" codec:"endTime,omitempty"`               // 使用结束时间，validityType=5时必填，示例：2022-11-03 00:00:00
	AdWord            string   `json:"adWord,omitempty" codec:"adWord,omitempty"`                 // 数据标签，专属优惠券使用
	SpuIdList         []uint64 `json:"spuIdList,omitempty" codec:"spuIdList,omitempty"`           // spuId
	Channels          string   `json:"channels,omitempty" codec:"channels,omitempty"`             // 渠道列表
	ChannelSelectType uint8    `json:"channelSelectType" codec:"channelSelectType"`               // 0全部1限选(非全部渠道) 传1
	SelectType        uint8    `json:"selectType" codec:"selectType"`                             // 一级推广平台的选择方式：0全部1限选(非全部渠道)
	Platforms         string   `json:"platforms,omitempty" codec:"platforms,omitempty"`           // 子渠道
	AppName           string   `json:"appName,omitempty" codec:"appName,omitempty"`               // appName客户端调用来源
	Ip                string   `json:"ip,omitempty" codec:"ip,omitempty"`                         // 调用方IP
	AppId             string   `json:"appId,omitempty" codec:"appId,omitempty"`                   // appId
}

type MarketCreateCouponResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *MarketCreateCouponData `json:"jingdong_market_tool_coupon_api_service_CouponWriteOuterService_createCoupon_responce,omitempty" codec:"jingdong_market_tool_coupon_api_service_CouponWriteOuterService_createCoupon_responce,omitempty"`
}

type MarketCreateCouponData struct {
	Code       string                            `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc  string                            `json:"error_description,omitempty" codec:"error_description,omitempty"`
	ReturnType *MarketCreateCouponDataReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type MarketCreateCouponDataReturnType struct {
	Msg     string `json:"msg,omitempty" codec:"msg,omitempty"`
	Code    uint   `json:"code,omitempty" codec:"code,omitempty"`
	Success bool   `json:"success,omitempty" codec:"success,omitempty"`
	Data    uint64 `json:"data,omitempty" codec:"data,omitempty"`
}

func MarketCreateCoupon(req *MarketCreateCouponRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := coupon.NewMarketCreateCouponRequest()
	r.SetNum(req.Num)
	r.SetDiscount(req.Discount)
	r.SetType(req.Type)
	r.SetSkuIdList(req.SkuIdList)
	r.SetShareType(req.ShareType)
	r.SetTakeEndTime(req.TakeEndTime)
	r.SetQuota(req.Quota)
	r.SetPromoteChannel(req.PromoteChannel)
	r.SetDisplay(req.Display)
	r.SetBusiCode(req.BusiCode)
	r.SetWareChoseType(req.WareChoseType)
	r.SetUserClass(req.UserClass)
	r.SetUserLevel(req.UserLevel)
	r.SetTakeBeginTime(req.TakeBeginTime)
	r.SetValidityType(req.ValidityType)
	r.SetTakeRule(req.TakeRule)
	r.SetHourCoupon(req.HourCoupon)
	r.SetName(req.Name)
	r.SetActivityLink(req.ActivityLink)
	r.SetStyle(req.Style)
	r.SetChannelSelectType(req.ChannelSelectType)
	r.SetSelectType(req.SelectType)
	r.SetAppName(req.AppName)
	r.SetAppId(req.AppId)
	if req.Platforms != "" {
		r.SetPlatforms(req.Platforms)
	}
	if req.Channels != "" {
		r.SetChannels(req.Channels)
	}
	if req.BeginTime != "" {
		r.SetBeginTime(req.BeginTime)
	}
	if req.EndTime != "" {
		r.SetEndTime(req.EndTime)
	}
	if req.WareGrade > 0 {
		r.SetWareGrade(req.WareGrade)
	}
	if req.StrategyParam != "" {
		r.SetStrategyParam(req.StrategyParam)
	}
	if req.CouponId > 0 {
		r.SetCouponId(req.CouponId)
	}
	if req.StoreId != nil {
		r.SetStoreId(req.StoreId)
	}
	if req.High > 0 {
		r.SetHigh(req.High)
	}
	if req.TakeNum > 0 {
		r.SetTakeNum(req.TakeNum)
	}
	if req.OfficialType > 0 {
		r.SetOfficialType(req.OfficialType)
	}
	if req.RemainNum > 0 {
		r.SetRemainNum(req.RemainNum)
	}
	if req.StoreType > 0 {
		r.SetStoreType(req.StoreType)
	}
	if req.Days > 0 {
		r.SetDays(req.Days)
	}
	if req.AdWord != "" {
		r.SetAdWord(req.AdWord)
	}
	if req.SpuIdList != nil {
		r.SetSpuIdList(req.SpuIdList)
	}
	if req.Ip != "" {
		r.SetIp(req.Ip)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	result = util.RemoveJsonSpace(result)

	var response MarketCreateCouponResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" && response.Data.ErrorDesc != "" {
		return 0, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.ReturnType == nil {
		return 0, errors.New("No return type.")
	}
	if response.Data.ReturnType.Msg != "" {
		return 0, errors.New(response.Data.ReturnType.Msg)
	}
	if response.Data.ReturnType.Data == 0 {
		return 0, errors.New("No coupon id.")
	}

	return response.Data.ReturnType.Data, nil
}
