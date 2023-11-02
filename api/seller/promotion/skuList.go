package promotion

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type SkuListRequest struct {
	api.BaseRequest
	Ip        string `json:"ip" codec:"ip"`
	Port      string `json:"port" codec:"port"`
	PromoId   uint64 `json:"promo_id" codec:"promo_id"`
	PromoType int    `json:"promo_type" codec:"promo_type"`
	Page      int    `json:"page,omitempty" codec:"page,omitempty"`
	PageSize  int    `json:"page_size,omitempty" codec:"page_size,omitempty"`
	BindType  uint8  `json:"bind_type,omitempey" codec:"bind_type,omitempty"`
}
type SkuListResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *SkuListResponseData `json:"jingdong_seller_promotion_v2_sku_list_responce,omitempty" codec:"jingdong_seller_promotion_v2_sku_list_responce,omitempty"`
}

func (r SkuListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r SkuListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data.IsError() {
		return r.Data.Error()
	}
	return "no result data"
}

type SkuListResponseData struct {
	Code             string             `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc        string             `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromotionSkuList []PromotionSkuList `json:"promotion_sku_list,omitempty" codec:"promotion_sku_list,omitempty"`
}

func (r SkuListResponseData) IsError() bool {
	return r.Code != "0"
}

func (r SkuListResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 店铺促销商品查询
func SkuList(req *SkuListRequest) ([]PromotionSkuList, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionSkuListRequest()
	r.SetPromoId(req.PromoId)
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	} else {
		r.SetPageSize(20)
	}

	if req.Page > 0 {
		r.SetPage(req.Page)
	} else {
		r.SetPage(1)
	}

	if req.PromoType > 0 {
		r.SetPromoType(req.PromoType)
	} else {
		r.SetPromoType(1)
	}

	if req.BindType > 0 {
		r.SetBindType(req.BindType)
	}

	var response SkuListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.PromotionSkuList, nil
}
