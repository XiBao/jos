package promotion

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type CountRequest struct {
	api.BaseRequest
	Ip          string `json:"ip,omitempty" codec:"ip,omitempty"`
	Port        string `json:"port,omitempty" codec:"port,omitempty"`
	PromoId     uint64 `json:"promo_id,omitempty" codec:"promo_id,omitempty"` // 促销ID
	Name        string `json:"name"`                                          // 促销名称
	Type        uint8  `json:"type"`                                          // 促销类型。1:单品促销,4:赠品促销,6:套装促销,10:总价促销
	FavorMode   uint8  `json:"favor_mode"`                                    // 0:满赠,1:满减,2:每满减, 3:百分比满减, 4: 阶梯满减,5:满赠加价购,6:满M件减N件,7:阶梯买M件减N件,13:M元任选N件,15:M件N折,16:满减送（元）
	BeginTime   string `json:"begin_time"`                                    // 开始时间。格式：yyyy-MM-dd HH:mm:ss
	EndTime     string `json:"end_time"`                                      // 结束时间。格式：yyyy-MM-dd HH:mm:ss
	PromoStatus uint8  `json:"promo_status"`                                  // 促销状态。1:驳回, 2:未审核, 3:人工审, 4:已审核, 5:已生效, 6:已暂停, 7:强制暂停
	WareId      string `json:"ware_id"`                                       // 商品ID
	SkuId       string `json:"sku_id"`                                        // sku ID
	SrcType     uint64 `json:"src_type"`                                      // 来源Id
}

type CountResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CountResponseData  `json:"jingdong_seller_promotion_v2_count_responce,omitempty" codec:"jingdong_seller_promotion_v2_count_responce,omitempty"`
}

func (r CountResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CountResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CountResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Count     uint64 `json:"promotion_count,omitempty" codec:"promotion_count,omitempty"`
}

func (r CountResponseData) IsError() bool {
	return r.Code != "0"
}

func (r CountResponseData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

func Count(req *CountRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCountRequest()

	if len(req.Ip) > 0 {
		r.SetIp(req.Ip)
	}

	if len(req.Port) > 0 {
		r.SetPort(req.Port)
	}

	if req.PromoId > 0 {
		r.SetPromoId(req.PromoId)
	}

	if len(req.Name) > 0 {
		r.SetName(req.Name)
	}

	if req.Type > 0 {
		r.SetType(req.Type)
	}

	if req.FavorMode > 0 {
		r.SetFavorMode(req.FavorMode)
	}

	if len(req.BeginTime) > 0 {
		r.SetBeginTime(req.BeginTime)
	}

	if len(req.EndTime) > 0 {
		r.SetEndTime(req.EndTime)
	}

	if req.PromoStatus > 0 {
		r.SetPromoStatus(req.PromoStatus)
	}

	if len(req.WareId) > 0 {
		r.SetWareId(req.WareId)
	}

	if len(req.SkuId) > 0 {
		r.SetSkuId(req.SkuId)
	}

	if req.SrcType > 0 {
		r.SetSrcType(req.SrcType)
	}

	var response CountResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Count, nil

}
