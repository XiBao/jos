package promotion

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type CreateRequest struct {
	api.BaseRequest
	Ip                    string `json:"ip,omitempty" codec:"ip,omitempty"`
	Port                  string `json:"port,omitempty" codec:"port,omitempty"`
	RequestId             string `json:"request_id,omitempty" codec:"request_id,omitempty"`
	Name                  string `json:"name" codec:"name"`
	BeginTime             string `json:"beginTime" codec:"beginTime"`                                 // 促销开始时间（yyyy-MM-dd HH:mm:ss），开始时间最长可设置为180天后
	EndTime               string `json:"endTime" codec:"endTime"`                                     // 促销结束时间（yyyy-MM-dd HH:mm:ss），结束时间要晚于开始时间至少一分钟
	Bound                 uint8  `json:"bound" codec:"bound"`                                         // 促销范围（1、部分商品参加）
	Member                uint   `json:"member" codec:"member"`                                       // 京东会员级别(注册会员：50 VIP：110 企业会员：90)
	Slogan                string `json:"slogan,omitempty" codec:"slogan,omitempty"`                   // 促销广告词，不能超过20个字符
	Comment               string `json:"comment,omitempty" codec:"comment,omitempty"`                 // 活动备注
	Platform              uint8  `json:"platform" codec:"platform"`                                   // 平台类型（1、全平台 2、APP专享 3、微信专享 5、手Q专享）
	FavorMode             uint   `json:"favorMode" codec:"favorMode"`                                 // 可取值：101表示令牌
	ShopMember            uint   `json:"shopMember,omitempty" codec:"shopMember,omitempty"`           // 店铺会员级别：0或null:非店铺会员（5001：一星会员 5002：二星会员 5003：三星会员 5004：四星会员 5005：五星会员）
	QqMember              uint8  `json:"qqMember,omitempty" codec:"qqMember,omitempty"`               // QQ会员级别：0或null非会员、1 qq会员
	PlusMember            uint8  `json:"plusMember,omitempty" codec:"plusMember,omitempty"`           // plus会员级别 1 付费用户专享 2 付费与试用用户均可享受
	SamMember             uint8  `json:"samMember,omitempty" codec:"samMember,omitempty"`             // Sam店会员级别：0或null非会员、1 Sam会员
	TokenId               uint64 `json:"tokenId,omitempty" codec:"tokenId,omitempty"`                 // 令牌编号
	PromoChannel          string `json:"promoChannel,omitempty" codec:"promoChannel,omitempty"`       // 业务渠道
	MemberLevelOnly       bool   `json:"memberLevelOnly,omitempty" codec:"memberLevelOnly,omitempty"` // 是否是会员级别专享：true 是、false 不是
	TokenUseNum           uint   `json:"tokenUseNum" codec:"tokenUseNum"`                             // 用户使用令牌的次数(0表示不限)，默认为0
	AllowOthersOperate    bool   `json:"allowOthersOperate" codec:"allowOthersOperate"`               // 是否允许其他来源操作该促销
	AllowOthersCheck      bool   `json:"allowOthersCheck" codec:"allowOthersCheck"`                   // 是否允许其他来源审核该促销
	AllowOtherUserOperate bool   `json:"allowOtherUserOperate" codec:"allowOtherUserOperate"`         // 是否允许其他人操作该促销
	AllowOtherUserCheck   bool   `json:"allowOtherUserCheck" codec:"allowOtherUserCheck"`             // 是否允许其他人审核该促销
	NeedManualCheck       bool   `json:"needManualCheck" codec:"needManualCheck"`                     // 促销是否需要人工审核
	ShowTokenPrice        bool   `json:"showTokenPrice" codec:"showTokenPrice"`                       // 在单品页是否展示令牌促销价格
	SkuId                 string `json:"skuId" codec:"skuId"`                                         // sku编号(数组)
	BindType              string `json:"bindType" codec:"bindType"`                                   // 绑定商品类型：1主商品(数组)
	PromoPrice            string `json:"promoPrice" codec:"promoPrice"`                               // sku的促销价,促销价格小于京东价(数组)
	Num                   string `json:"num" codec:"num"`                                             // 该sku的限购数量(数组)
	WareId                string `json:"wareId,omitempty" codec:"wareId,omitempty"`                   // 商品编号(数组)
	SkuName               string `json:"skuName,omitempty" codec:"skuName,omitempty"`                 // sku名称(数组)
	JdPrice               string `json:"jdPrice,omitempty" codec:"jdPrice,omitempty"`                 // sku京东价(数组)
	ItemNum               string `json:"itemNum,omitempty" codec:"itemNum,omitempty"`                 // 商品货号
	FreqBound             uint8  `json:"freqBound,omitempty" codec:"freqBound,omitempty"`             // 限购一次：0，不限，1、限ip、账号 2、限ip 3、限账号
	PerMaxNum             uint   `json:"perMaxNum,omitempty" codec:"perMaxNum,omitempty"`             // 单次最大购买数量：0、不限
	PerMinNum             uint   `json:"perMinNum,omitempty" codec:"perMinNum,omitempty"`             // 单次最小购买数量：0、不限
	Pin                   string `json:"pin,omitempty" codec:"pin,omitempty"`                         // 用户PIN(数组)
	UseBeginTime          string `json:"useBeginTime,omitempty" codec:"useBeginTime,omitempty"`       // 拥有令牌的开始时间(数组)
	UseEndTime            string `json:"useEndTime,omitempty" codec:"useEndTime,omitempty"`           // 拥有令牌的结束时间(数组)
}

type CreateResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CreateResponseData `json:"jingdong_seller_promotion_create_responce,omitempty" codec:"jingdong_seller_promotion_create_responce,omitempty"`
}

func (r CreateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CreateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CreateResponseData struct {
	Code         string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc    string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	CreateResult uint64 `json:"create_result,omitempty" codec:"create_result,omitempty"`
}

func (r CreateResponseData) IsError() bool {
	return r.Code != "0"
}

func (r CreateResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func Create(ctx context.Context, req *CreateRequest) (interface{}, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCreateRequest()

	r.SetName(req.Name)
	r.SetBeginTime(req.BeginTime)
	r.SetEndTime(req.EndTime)
	r.SetBound(req.Bound)
	r.SetMember(req.Member)
	r.SetPlatform(req.Platform)
	r.SetFavorMode(req.FavorMode)
	r.SetTokenUseNum(req.TokenUseNum)
	r.SetAllowOthersOperate(req.AllowOthersOperate)
	r.SetAllowOthersCheck(req.AllowOthersCheck)
	r.SetAllowOtherUserOperate(req.AllowOtherUserOperate)
	r.SetAllowOtherUserCheck(req.AllowOtherUserCheck)
	r.SetNeedManualCheck(req.NeedManualCheck)
	r.SetSkuId(req.SkuId)
	r.SetBindType(req.BindType)
	r.SetPromoPrice(req.PromoPrice)
	r.SetNum(req.Num)
	r.SetFreqBound(req.FreqBound)
	r.SetPerMaxNum(req.PerMaxNum)
	r.SetPerMinNum(req.PerMinNum)
	r.SetMemberLevelOnly(req.MemberLevelOnly)
	r.SetShowTokenPrice(req.ShowTokenPrice)

	if len(req.Ip) > 0 {
		r.SetIp(req.Ip)
	}
	if len(req.Port) > 0 {
		r.SetPort(req.Port)
	}
	if len(req.RequestId) > 0 {
		r.SetRequestId(req.RequestId)
	}
	if len(req.Slogan) > 0 {
		r.SetSlogan(req.Slogan)
	}
	if len(req.Comment) > 0 {
		r.SetComment(req.Comment)
	}
	if req.ShopMember > 0 {
		r.SetShopMember(req.ShopMember)
	}
	if req.QqMember > 0 {
		r.SetQqMember(req.QqMember)
	}
	if req.PlusMember > 0 {
		r.SetPlusMember(req.PlusMember)
	}
	if req.SamMember > 0 {
		r.SetSamMember(req.SamMember)
	}
	if req.TokenId > 0 {
		r.SetTokenId(req.TokenId)
	}
	if len(req.PromoChannel) > 0 {
		r.SetPromoChannel(req.PromoChannel)
	}
	if req.TokenUseNum > 0 {
		r.SetTokenUseNum(req.TokenUseNum)
	}
	if len(req.WareId) > 0 {
		r.SetWareId(req.WareId)
	}
	if len(req.SkuName) > 0 {
		r.SetSkuName(req.SkuName)
	}
	if len(req.JdPrice) > 0 {
		r.SetJdPrice(req.JdPrice)
	}
	if len(req.ItemNum) > 0 {
		r.SetItemNum(req.ItemNum)
	}
	if len(req.UseBeginTime) > 0 {
		r.SetUseBeginTime(req.UseBeginTime)
	}
	if len(req.UseEndTime) > 0 {
		r.SetUseEndTime(req.UseEndTime)
	}

	var response CreateResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}

	return response.Data.CreateResult, nil
}
