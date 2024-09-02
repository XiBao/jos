package promotion

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type ResumeRequest struct {
	api.BaseRequest
	Ip        string `json:"ip,omitempty" codec:"ip,omitempty"`                 // 调用方IP
	Port      string `json:"port,omitempty" codec:"port,omitempty"`             // 调用方端口
	RequestId string `json:"request_id,omitempty" codec:"request_id,omitempty"` // 防重码。业务唯一值
	PromoId   uint64 `json:"promo_id,omitempty" codec:"promo_id,omitempty"`     // 促销编号
	PromoType uint8  `json:"promo_type,omitempty" codec:"promo_type,omitempty"` // 促销类型。1:单品促销,4:赠品促销,6:套装促销,10:总价促销
}

type ResumeResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ResumeData         `json:"jingdong_seller_promotion_v2_resume_responce,omitempty" codec:"jingdong_seller_promotion_v2_resume_responce,omitempty"`
}

func (r ResumeResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ResumeResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ResumeData struct {
	Code         string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc    string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	ResumeResult bool   `json:"resume_result,omitempty" codec:"resume_result,omitempty"`
}

func (r ResumeData) IsError() bool {
	return r.Code != "0"
}

func (r ResumeData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func Resume(ctx context.Context, req *ResumeRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionResumeRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetPromoId(req.PromoId)
	r.SetPromoType(req.PromoType)
	if req.RequestId != "" {
		r.SetRequestId(req.RequestId)
	}

	var response ResumeResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.ResumeResult, nil
}
