package promotion

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
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

func (r UserListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r UserListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type UserListResponseData struct {
	Code              string               `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc         string               `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromotionUserList []*PromotionUserList `json:"getpromouserlist_result,omitempty" codec:"getpromouserlist_result,omitempty"`
}

func (r UserListResponseData) IsError() bool {
	return r.Code != "0"
}

func (r UserListResponseData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
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

	var response UserListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.PromotionUserList, nil
}
