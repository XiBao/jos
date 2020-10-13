package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
)

type FullCreateRequest struct {
	api.BaseRequest
	Ip                    string `json:"ip"`                       // 调用方IP
	Port                  string `json:"port"`                     // 调用方端口
	RequestId             string `json:"request_id"`               // 防重码。业务唯一值
	PromoName             string `json:"promo_name"`               // 促销名称。最大长度20个字符
	BeginTime             string `json:"begin_time"`               // 促销开始时间。格式：yyyy-MM-dd HH:mm:ss
	EndTime               string `json:"end_time"`                 // 促销结束时间。格式：yyyy-MM-dd HH:mm:ss
	Slogan                string `json:"slogan"`                   // 促销宣传语。最大长度70个字符
	Comment               string `json:"comment"`                  // 促销备注信息。最大长度100个字符
	Link                  string `json:"link"`                     // 活动链接
	PlusMember            int64  `json:"plusMember"`               // plus会员级别（0 非plus会员,1 付费及试用用户）
	AllowOthersOperate    bool   `json:"allow_others_operate"`     // 是否允许其他来源操作该促销
	AllowOthersCheck      bool   `json:"allow_others_check"`       // 是否允许其他来源审核该促销
	AllowOtherUserOperate bool   `json:"allow_other_user_operate"` // 是否允许其他人操作该促销
	AllowOtherUserCheck   bool   `json:"allow_other_user_check"`   // 是否允许其他人审核该促销
	NeedManualCheck       bool   `json:"need_manual_check"`        // 促销是否需要人工审核
	FreqBound             int64  `json:"freq_bound"`               // 促销限购一次形式：（0，不限，1、限ip、账号 2、限ip 3、限账号）
	PerMaxNum             int64  `json:"per_max_num"`              // 单次最大购买数量：0、不限
	PerMinNum             int64  `json:"per_min_num"`              // 单次最小购买数量：0、不限
	PropType              uint8  `json:"prop_type"`                // 道具类型：2、京豆，10 、店铺京券
	PropNum               uint   `json:"prop_num"`                 // 道具数值
	PropUsedWay           uint8  `json:"prop_used_way"`            // 道具使用方式：0、消耗，2、奖励
	CouponValidDays       int64  `json:"coupon_valid_days"`        // 优惠券的有效天数
	TokenUseNum           int64  `json:"token_use_num"`            // 用户使用令牌的次数
	UserPins              string `json:"user_pins"`                // 令牌用户列表
	PromoAreaType         int64  `json:"promo_area_type"`          // 促销区域类型： 2 白名单 3 黑名单
	PromoAreas            string `json:"promo_areas"`              // 促销区域列表（英文分号分隔，若地址是多级，则每一级的编号都必须传，各级之间用英文逗号分隔，eg：1,2,3;2,5,8;）
	SkuId                 string `json:"sku_id"`                   // sku ID
	PromoPrice            string `json:"promo_price"`              // 促销价
	LimitNum              string `json:"limit_num"`                // 限购数量。0:不限
}

type FullCreateResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FullCreateResponseData `json:"jingdong_seller_promotion_v2_unit_full_create_responce,omitempty" codec:"jingdong_seller_promotion_v2_unit_full_create_responce,omitempty"`
}

type FullCreateResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromoId   uint64 `json:"promo_id,omitempty" codec:"promo_id,omitempty"`
}

func FullCreate(req *FullCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionFullCreateRequest()

	if len(req.Ip) > 0 {
		r.SetIp(req.Ip)
	}

	if len(req.Port) > 0 {
		r.SetPort(req.Port)
	}

	if len(req.RequestId) > 0 {
		r.SetRequestId(req.RequestId)
	}

	if len(req.PromoName) > 0 {
		r.SetPromoName(req.PromoName)
	}

	if len(req.BeginTime) > 0 {
		r.SetBeginTime(req.BeginTime)
	}

	if len(req.EndTime) > 0 {
		r.SetEndTime(req.EndTime)
	}

	if len(req.Slogan) > 0 {
		r.SetSlogan(req.Slogan)
	}

	if len(req.Comment) > 0 {
		r.SetComment(req.Comment)
	}

	if len(req.Link) > 0 {
		r.SetLink(req.Link)
	}

	if req.PlusMember >= 0 {
		r.SetPlusMember(req.PlusMember)
	}

	r.SetAllowOthersOperate(req.AllowOthersOperate)

	r.SetAllowOthersCheck(req.AllowOthersCheck)

	r.SetAllowOtherUserOperate(req.AllowOtherUserOperate)

	r.SetAllowOtherUserCheck(req.AllowOtherUserCheck)

	r.SetNeedManualCheck(req.NeedManualCheck)

	if req.FreqBound > 0 {
		r.SetFreqBound(req.FreqBound)
	}

	if req.PerMaxNum > 0 {
		r.SetPerMaxNum(req.PerMaxNum)
	}

	if req.PerMinNum > 0 {
		r.SetPerMinNum(req.PerMinNum)
	}

	if req.PropType > 0 {
		r.SetPropType(req.PropType)
	}

	if req.PropNum > 0 {
		r.SetPropNum(req.PropNum)
	}

	if req.PropType > 0 && req.PropUsedWay >= 0 {
		r.SetPropUsedWay(req.PropUsedWay)
	}

	if req.CouponValidDays > 0 {
		r.SetCouponValidDays(req.CouponValidDays)
	}

	if req.TokenUseNum > 0 {
		r.SetTokenUseNum(req.TokenUseNum)
	}

	if len(req.UserPins) > 0 {
		r.SetUserPins(req.UserPins)
	}

	if req.PromoAreaType > 0 {
		r.SetPromoAreaType(req.PromoAreaType)
	}

	if len(req.PromoAreas) > 0 {
		r.SetPromoAreas(req.PromoAreas)
	}

	if len(req.SkuId) > 0 {
		r.SetSkuId(req.SkuId)
	}

	if len(req.PromoPrice) > 0 {
		r.SetPromoPrice(req.PromoPrice)
	}

	if len(req.LimitNum) > 0 {
		r.SetLimitNum(req.LimitNum)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response FullCreateResponse
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
	if response.Data.PromoId == 0 {
		return 0, errors.New("No promotion id.")
	}

	return response.Data.PromoId, nil
}
