package promotion

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type AddRequest struct {
	api.BaseRequest
	Name      string `json:"name" codec:"name"`             // 促销名称，字符串长度小于等于10
	Type      uint8  `json:"type" codec:"type"`             // 促销类型，可选值：单品促销（1），赠品促销（4），套装促销（6），总价促销（10）
	BeginTime string `json:"begin_time" codec:"begin_time"` // 促销开始时间，格式为yyyy-MM-dd HH:mm:ss，精确到分钟，最长可设置为距当前时间180天之内的时间点
	EndTime   string `json:"end_time" codec:"end_time"`     // 促销结束时间，格式为yyyy-MM-dd HH:mm:ss，精确到分钟，必须大于开始时间至少一分钟，且晚于当前时间，建议至少晚10分钟，且和开始时间最大间隔不能超过180天
	Bound     uint8  `json:"bound" codec:"bound"`           // 促销范围，总价促销为必填项，其它促销类型无效，可选值：部分商品参加（1）、全场参加（2）、部分商品不参加（3），注：M元任选N件只支持部分商品参加
	Member    uint8  `json:"member" codec:"member"`         // 会员限制，默认值：注册会员（50），可选值：注册会员（50）、铜牌（56）、银牌（61）、金牌（62）、钻石（105）、VIP（110）
	Slogan    string `json:"slogan" codec:"slogan"`         // 广告语，字符串长度小于等于50
	Comment   string `json:"comment" codec:"comment"`       // 活动备注，不超过200字节
	FavorMode uint8  `json:"favor_mode" codec:"favor_mode"` // 总价促销订单规则类型（总价促销时为必填项，orderMdoe需要和此值保持一致），满赠（0），满减（1），每满减（2），满赠加价购（5），满M件减N件（6），M元任选N件（13），M件N折（15），满减送（元）（16）
}
type AddResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AddResponseData    `json:"jingdong_seller_promotion_add_responce,omitempty" codec:"jingdong_seller_promotion_add_responce,omitempty"`
}

func (r AddResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r AddResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type AddResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromoId   uint64 `json:"promo_id,omitempty" codec:"promo_id,omitempty"`
}

func (r AddResponseData) IsError() bool {
	return r.Code != "0"
}

func (r AddResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 创建促销
func Add(ctx context.Context, req *AddRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionAddRequest()
	r.SetName(req.Name)
	r.SetType(req.Type)
	r.SetBeginTime(req.BeginTime)
	r.SetEndTime(req.EndTime)
	r.SetFavorMode(req.FavorMode)
	if req.Bound > 0 {
		r.SetBound(req.Bound)
	}
	if req.Member > 0 {
		r.SetMember(req.Member)
	}
	if req.Slogan != "" {
		r.SetSlogan(req.Slogan)
	}
	if req.Comment != "" {
		r.SetComment(req.Comment)
	}

	var response AddResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return 0, err
	}

	return response.Data.PromoId, nil
}
