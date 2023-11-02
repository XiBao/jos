package promotion

import (
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

func (r CommitResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r CommitResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type CommitResponseData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Success   bool   `json:"success,omitempty" codec:"success,omitempty"`
}

func (r CommitResponseData) IsError() bool {
	return r.Code != "0"
}

func (r CommitResponseData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

// 促销创建完毕,提交保存促销命令
func Commit(req *CommitRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionCommitRequest()
	r.SetPromoId(req.PromoId)

	var response CommitResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}

	return response.Data.Success, nil
}
