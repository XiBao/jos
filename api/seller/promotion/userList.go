package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
)

type UserListRequest struct {
	api.BaseRequest
	PromoId  uint64 `json:"promo_id" codec:"promo_id"`
	Page     int    `json:"page,omitempty" codec:"page,omitempty"`
	PageSize int    `json:"page_size,omitempty" codec:"page_size,omitempty"`
}
type UserListResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *UserListResponseData `json:"jingdong_pop_market_retrieve_promotion_getPromoUserList_responce,omitempty" codec:"jingdong_pop_market_retrieve_promotion_getPromoUserList_responce,omitempty"`
}

type UserListResponseData struct {
	Code              string               `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc         string               `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromotionUserList []*PromotionUserList `json:"getpromouserlist_result,omitempty" codec:"getpromouserlist_result,omitempty"`
}

// 店铺促销用户查询
func UserList(req *UserListRequest) ([]*PromotionUserList, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionUserListRequest()
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
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response SellerPromotionUserListResponse
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

	return response.Data.PromotionUserList, nil
}
