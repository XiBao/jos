package promotion

import (
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

func (r GetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetResponseData struct {
	Code         string         `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc    string         `json:"error_description,omitempty" codec:"error_description,omitempty"`
	JosPromotion *PromotionList `json:"jos_promotion,omitempty" codec:"jos_promotion,omitempty"`
}

func (r GetResponseData) IsError() bool {
	return r.Code != "0" || r.JosPromotion == nil
}

func (r GetResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
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

	var response GetResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.JosPromotion, nil
}
