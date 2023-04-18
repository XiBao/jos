package promotion

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type GetRequest struct {
	api.BaseRequest
	Ip        string `json:"ip" codec:"ip"`
	Port      string `json:"port" codec:"port"`
	PromoId   uint64 `json:"promo_id" codec:"promo_id"`
	PromoType uint8  `json:"promo_type" codec:"promo_type"`
}
type GetResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetResponseData    `json:"jingdong_seller_promotion_v2_get_responce,omitempty" codec:"jingdong_seller_promotion_v2_get_responce,omitempty"`
}

type GetResponseData struct {
	Code         string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc    string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	JosPromotion *PromotionList `json:"jos_promotion,omitempty" codec:"jos_promotion,omitempty"`
}

// 促销详情查询
func Get(req *GetRequest) (*PromotionList, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionGetRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetPromoId(req.PromoId)
	r.SetPromoType(req.PromoType)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response GetResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.JosPromotion, nil
}
