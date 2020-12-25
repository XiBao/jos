package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
)

type OrdermodeAddRequest struct {
	api.BaseRequest
	PromoId     uint64 `json:"promo_id" codec:"promo_id"`         // 促销编号
	FavorMode   uint8  `json:"favor_mode" codec:"favor_mode"`     // 订单规则类型，可选值：满赠（0），满减（1），每满减（2），满赠加价购（5），满M件减N件（6），M元任选N件（13），M件N折（15），满减送（元）（16）
	Quota       string `json:"quota" codec:"quota"`               // 订单额度；（满M件减N件或M件N折时为M的值，单位件，只支持正整数；M元任选N件时为M的值，单位元，支持小数点后一位，例：9.9元；满减送时为订单满金额，单位元，只支持正整数；）
	Rate        string `json:"rate" codec:"rate"`                 // 优惠力度；（满M件减N件、M元任选N件时为N的值，单位件,只支持正整数；M件N折时为N的值，单位折，支持小数点后一位，例：8.5折；满减送不支持此字段，除满减送之外其它促销为必填项）
	Plus        string `json:"plus" codec:"plus"`                 // 加价金额，只支持正整数；（只满减送有效，且为可选项，该字段设置了值，必须送赠品）
	Minus       string `json:"minus" codec:"minus"`               // 减金额，只支持正整数；（只满减送有效，且为可选项）
	Link        string `json:"link" codec:"link"`                 // 店铺活动链接地址
	FreePostage string `json:"free_postage" codec:"free_postage"` // 是否免邮（促销下订单规则列表的包邮必须保持一致）。0：不免邮1：免邮
}
type OrdermodeAddResponse struct {
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *OrdermodeAddResponseData `json:"jingdong_seller_promotion_ordermode_add_responce,omitempty" codec:"jingdong_seller_promotion_ordermode_add_responce,omitempty"`
}

type OrdermodeAddResponseData struct {
	Code      string   `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string   `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Ids       []uint64 `json:"ids,omitempty" codec:"ids,omitempty"`
}

// 新建总价促销订单规则
func OrdermodeAdd(req *OrdermodeAddRequest) ([]uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionOrdermodeAddRequest()
	r.SetPromoId(req.PromoId)
	r.SetFavorMode(req.FavorMode)
	r.SetQuota(req.Quota)
	if req.Rate != "" {
		r.SetRate(req.Rate)
	}
	if req.Plus != "" {
		r.SetPlus(req.Plus)
	}
	if req.Minus != "" {
		r.SetMinus(req.Minus)
	}
	if req.Link != "" {
		r.SetLink(req.Link)
	}
	if req.FreePostage != "" {
		r.SetFreePostage(req.FreePostage)
	}
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response OrdermodeAddResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Ids, nil
}
