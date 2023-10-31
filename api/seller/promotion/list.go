package promotion

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type ListRequest struct {
	api.BaseRequest
	Ip          string `json:"ip" codec:"ip"`
	Port        string `json:"port" codec:"port"`
	PromoId     uint64 `json:"promo_id" codec:"promo_id"`
	Name        string `json:"name" codec:"name"`
	PType       int    `json:"type" codec:"type"`
	Page        int    `json:"page,omitempty" codec:"page,omitempty"`
	PageSize    int    `json:"page_size,omitempty" codec:"page_size,omitempty"`
	PromoStatus string `json:"page_status,omitempty" codec:"page_status,omitempty"`
	BeginTime   string `json:"begin_time,omitempty" codec:"begin_time,omitempty"`
	EndTime     string `json:"end_time,omitempty" codec:"end_time,omitempty"`
	StartId     uint64 `json:"start_id,omitempty" codec:"start_id,omitempty"`
}
type ListResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ListResponseData   `json:"jingdong_seller_promotion_v2_list_responce,omitempty" codec:"jingdong_seller_promotion_v2_list_responce,omitempty"`
}

func (r ListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ListResponseData struct {
	Code          string           `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string           `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromotionList []*PromotionList `json:"promotion_list,omitempty" codec:"promotion_list,omitempty"`
}

func (r ListResponseData) IsError() bool {
	return r.Code != "0"
}

func (r ListResponseData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

// 店铺促销查询
func List(req *ListRequest) ([]*PromotionList, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionListRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
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
	if req.PType > 0 {
		r.SetType(req.PType)
	} else {
		r.SetType(1)
	}
	if req.PromoId > 0 {
		r.SetPromoId(req.PromoId)
	}

	if len(req.BeginTime) > 0 {
		r.SetBeginTime(req.BeginTime)
	}

	if len(req.EndTime) > 0 {
		r.SetEndTime(req.EndTime)
	}

	if req.Name != "" {
		r.SetName(req.Name)
	}
	if req.PromoStatus != "" {
		r.SetPromoStatus(req.PromoStatus)
	}
	if req.StartId >= 0 {
		r.SetStartId(req.StartId)
	}

	var response ListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.PromotionList, nil
}
