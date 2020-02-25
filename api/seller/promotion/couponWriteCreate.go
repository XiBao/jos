package promotion

import (
	"errors"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
)

type CouponWriteCreateRequest struct {
	api.BaseRequest
	Ip            string `json:"ip,omitempty" codec:"ip,omitempty"`
	Port          string `json:"port,omitempty" codec:"port,omitempty"`
	Name          string `json:"name,omitempty" codec:"name,omitempty"`           // 优惠券名称（长度小于30）
	Type          uint   `json:"type" codec:"type"`                               // 优惠券类型 0京券 1东券
	BindType      uint   `json:"bindType" codec:"bindType"`                       // 绑定类型 1全店参加 2指定sku参加
	GrantType     uint   `json:"grantType" codec:"grantType"`                     // 发放类型 3免费领取 5互动平台 【仅允许这两种】
	Num           uint   `json:"num" codec:"num"`                                 // 优惠券数量（限制：(0,1000000000]）
	Discount      uint   `json:"discount" codec:"discount"`                       // 优惠券面额（限制：[1,5000)）
	Quota         uint   `json:"quota" codec:"quota"`                             // 优惠限额（限制：[1,100000)）,京券无门槛传0
	ValidityType  uint   `json:"validityType" codec:"validityTYpe:"`              // 有效期类型 1相对时间 5绝对时间
	Days          uint   `json:"days" codec:"days"`                               // 有效期(validityType为1时必填)
	BeginTime     int64  `json:"beginTime,omitempty" codec:"beginTime,omitempty"` //有效期开始时间（validityType为5时必填）
	EndTime       int64  `json:"endTime,omitempty" codec:"endTime,omitempty"`     // 有效期结束时间（validityType为5时必填）,有效期开始结束时间是券使用时间，必须包含在活动时间内，且小于 90 天
	Member        uint   `json:"member" codec:"member"`                           // 会员等级 50注册会员 56铜牌 61银牌 62金牌 105钻石 110VIP 90企业会员
	TakeBeginTime int64  `json:"takeBeginTime" codec:"takeBeginTime"`             //领券开始时间（晚于当前）
	TakeEndTime   int64  `json:"takeEndTime" codec:"takeEndTime"`                 // 领券结束时间（活动时间最长90天）
	TakeRule      uint   `json:"takeRule" codec:"takeRule"`                       // 领券规则 5活动期间限领一张 4活动内每天限领一张
	TakeNum       uint   `json:"takeNum" codec:"takeNum"`                         // 限制条件内可以领取张数
	Display       uint   `json:"display" codec:"display"`                         // 是否公开 1不公开 3公开(grantType如设值5此参数必须为3)
	PlatformType  uint   `json:"platformType" codec:"platformType"`               // 使用平台 1全平台 3限平台
	Platform      string `json:"platform,omitempty" codec:"platform,omitempty"`   // 优惠券使用平台 0主站专享 1手机专享 3M版京东 4手Q专享 5微信专享 7京致衣橱（此参数需根据platformType设值，如限平台必填）
	ShareType     uint   `json:"shareType" codec:"shareType"`                     // 分享类型 1分享 2不分享（如设置京券type=0,此参数必填2不分享）
	SkuId         string `json:"skuId,omitempty" codec:"skuId,omitempty"`         // 商品sku编号(如设置bindType为2，此参数必填,需是有效sku)
}

type CouponWriteCreateResponse struct {
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CouponWriteCreateResponseData `json:"jingdong_seller_coupon_write_create_response,omitempty" codec:"jingdong_seller_coupon_write_create_response,omitempty"`
}

type CouponWriteCreateResponseData struct {
	CouponId uint64 `json:"couponId,omitempty" codec:"couponId,omitempty"`
}

func CouponWriteCreate(req *CouponWriteCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCouponWriteCreateRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	if req.Type == 0 || req.Type == 1 {
		r.SetType(req.Type)
	} else {
		return 0, errors.New("优惠券类型错误")
	}

	if req.BindType == 1 || req.BindType == 2 {
		r.SetBindType(req.BindType)
	} else {
		return 0, errors.New("绑定类型错误")
	}

	if req.GrantType == 3 || req.GrantType == 5 {
		r.SetGrantType(req.GrantType)
	} else {
		return 0, errors.New("发放类型错误")
	}
	if req.Num <= 1000000000 {
		r.SetNum(req.Num)
	} else {
		return 0, errors.New("优惠券数量（限制：(0,1000000000]）")
	}

	if req.Discount >= 1 && req.Discount < 5000 {
		r.SetDiscount(req.Discount)
	} else {
		return 0, errors.New("优惠券面额（限制：[1,5000)）")
	}

	if req.Type == 0 {
		req.SetQuota(0)
	} else if req.Quota > 1 && req.Quota < 100000 {
		r.SetQuota(req.Quota)
	} else {
		return 0, errors.New("优惠限额（限制：[1,100000)）")
	}

	if req.ValidityType == 1 || req.ValidityType == 5 {
		r.SetValidityType(req.ValidityType)
	} else {
		return 0, errors.New("有效期类型设置错误")
	}
	if req.Days > 0 {
		r.SetDays(req.Days)
	} else if req.ValidityType == 1 {
		return 0, errors.New("缺少相对时间有效期")
	} else if req.BeginTime > 0 && req.EndTime > 0 {
		if req.BeginTime >= req.EndTime {
			return 0, errors.New("绝对时间有效期开始时间应早于结束时间")
		} else if time.Duration(req.EndTime-req.BeginTime) >= time.Hour*90*24 {
			return 0, errors.New("绝对时间有效期应少于90天")
		}
		r.SetBeginTime(req.BeginTime)
		r.SetEndTime(req.EndTime)
	} else {
		return 0, errors.New("缺少绝对时间有效期")
	}

	if req.Member == 50 || req.Member == 56 || req.Member == 61 || req.Member == 62 || req.Member == 105 || req.Member == 110 || req.Member == 90 {
		r.SetMember(req.Member)
	} else {
		return 0, errors.New("会员等级设置错误")
	}
	if req.TakeBeginTime <= 0 || req.TakeEndTime <= 0 {
		return 0, errors.New("缺少领券时间")
	} else if req.TakeBeginTime < time.Now().Unix() {
		return 0, errors.New("领券开始时间应晚于当前时间")
	} else if req.TakeBeginTime <= req.TakeEndTime {
		return 0, errors.New("领券开始时间应早于领券结束时间")
	} else if time.Duration(req.TakeEndTime-req.TakeBeginTime) >= time.Hour*24*90 {
		return 0, errors.New("领券时间不应少于90天")
	} else {
		r.SetTakeBeginTime(req.TakeBeginTime)
		r.SetTakeEndTime(req.TakeEndTime)
	}

	if req.TakeRule == 4 || req.TakeRule == 5 {
		r.SetTakeRule(req.TakeRule)
	} else {
		return 0, errors.New("领券规则错误")
	}
	r.SetTakeNum(req.TakeNum)
	if req.GrantType == 5 {
		r.SetDisplay(3)
	} else if req.Display == 1 || req.Display == 5 {
		r.SetDisplay(req.Display)
	} else {
		return 0, errors.New("是否公开设置错误")
	}

	if req.PlatformType == 1 || req.PlatformType == 3 {
		r.SetPlatformType(req.PlatformType)
	} else {
		return 0, errors.New("使用平台错误")
	}
	if req.PlatformType == 3 && req.Platform != "" {
		r.SetPlatform(req.Platform)
	} else if req.PlatformType == 3 {
		return 0, errors.New("缺少使用平台设置")
	}
	if req.Type == 0 {
		r.SetShareType(2)
	} else if req.ShareType == 1 || req.ShareType == 2 {
		r.SetShareType(req.ShareType)
	} else {
		return 0, errors.New("分享设置错误")
	}

	if req.BindType == 2 || req.SkuId == "" {
		return 0, errors.New("缺少SKUID")
	} else if req.BindType == 2 {
		r.SetSkuId(req.SkuId)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response CouponWriteCreateResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}

	return response.Data.CouponId, nil
}
