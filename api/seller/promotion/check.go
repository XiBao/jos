package promotion

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type CheckRequest struct {
	api.BaseRequest
	PromoId uint64 `json:"promo_id" codec:"promo_id"` // 促销Id
	Status  uint8  `json:"status" codec:"status"`     // 审核状态。1:驳回,4:通过
}
type CheckResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CheckResponseData  `json:"jingdong_seller_promotion_check_responce,omitempty" codec:"jingdong_seller_promotion_check_responce,omitempty"`
}

type CheckResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Count     uint   `json:"count,omitempty" codec:"count,omitempty"`
}

// 促销审核,只能对人工审状态的促销进行审核
func Check(req *CheckRequest) (uint, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCheckRequest()
	r.SetPromoId(req.PromoId)
	r.SetStatus(req.Status)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("no result.")
	}

	var response CheckResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Count, nil
}
