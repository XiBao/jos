package promotion

import (
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

func (r CheckResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CheckResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CheckResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Count     uint   `json:"count,omitempty" codec:"count,omitempty"`
}

func (r CheckResponseData) IsError() bool {
	return r.Code != "0"
}

func (r CheckResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 促销审核,只能对人工审状态的促销进行审核
func Check(req *CheckRequest) (uint, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCheckRequest()
	r.SetPromoId(req.PromoId)
	r.SetStatus(req.Status)

	var response CheckResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Count, nil
}
