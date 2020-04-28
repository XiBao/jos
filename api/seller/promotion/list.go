package promotion

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
	"github.com/daviddengcn/ljson"
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
}
type ListResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ListResponseData   `json:"jingdong_seller_promotion_v2_list_responce,omitempty" codec:"jingdong_seller_promotion_v2_list_responce,omitempty"`
}

type ListResponseData struct {
	Code          string           `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc     string           `json:"error_description,omitempty" codec:"error_description,omitempty"`
	PromotionList []*PromotionList `json:"promotion_list,omitempty" codec:"promotion_list,omitempty"`
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
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response ListResponse
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

	return response.Data.PromotionList, nil
}
