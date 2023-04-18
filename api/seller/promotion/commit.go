package promotion

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type CommitRequest struct {
	api.BaseRequest
	PromoId uint64 `json:"promo_id" codec:"promo_id"` // 促销Id
}
type CommitResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *CommitResponseData `json:"jingdong_seller_promotion_commit_responce,omitempty" codec:"jingdong_seller_promotion_commit_responce,omitempty"`
}

type CommitResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Success   bool   `json:"success,omitempty" codec:"success,omitempty"`
}

// 促销创建完毕,提交保存促销命令
func Commit(req *CommitRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCommitRequest()
	r.SetPromoId(req.PromoId)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result.")
	}

	var response CommitResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}
	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return false, errors.New(response.Data.ErrorDesc)
	}

	return response.Data.Success, nil
}
