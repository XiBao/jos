package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
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

type SkuListResponseData struct {
	Code             string              `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc        string              `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromotionSkuList []*PromotionSkuList `json:"promotion_sku_list,omitempty" codec:"promotion_sku_list,omitempty"`
}

// 店铺促销商品查询
func SkuList(req *SkuListRequest) ([]*PromotionSkuList, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionSkuListRequest()
	if req.PageSize > 0 {
		r.SetPageSize(req.PageSize)
	} else {
		r.SetPageSize(100)
	}

	if req.Page > 0 {
		r.SetPage(req.Page)
	} else {
		r.SetPage(1)
	}

	if req.PromoId > 0 {
		r.SetPromoId(req.PromoId)
	} else {
		return nil, errors.New("no promo id")
	}

	if req.PromoType > 0 {
		r.SetPromoType(req.PromoType)
	} else {
		r.SetPromoType(1)
	}

	if req.BindType > 0 {
		r.SetBindType(req.BindType)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response SkuListResponse
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

	return response.Data.PromotionSkuList, nil
}
