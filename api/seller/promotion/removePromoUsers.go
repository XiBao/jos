package promotion

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type RemovePromoUsersRequest struct {
	api.BaseRequest
	Ip        string `json:"ip,omitempty" codec:"ip,omitempty"`
	Port      string `json:"port,omitempty" codec:"port,omitempty"`
	RequestId string `json:"request_id,omitempty" codec:"request_id,omitempty"`
	BeginTime string `json:"begin_time,omitempty" codec:"begin_time,omitempty"`
	EndTime   string `json:"end_time,omitempty" codec:"end_time,omitempty"`
	PromoId   uint64 `json:"promo_id" codec:"promo_id"`
	Pin       string `json:"pin" codec:"pin"`
}

type RemovePromoUsersResponse struct {
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *RemovePromoUsersResponseData `json:"jingdong_seller_promotion_removePromoUsers_responce,omitempty" codec:"jingdong_seller_promotion_removePromoUsers_responce,omitempty"`
}

func (r RemovePromoUsersResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r RemovePromoUsersResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type RemovePromoUsersResponseData struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    interface{} `json:"result,omitempty" codec:"result,omitempty"`
}

func (r RemovePromoUsersResponseData) IsError() bool {
	return r.Code != "0"
}

func (r RemovePromoUsersResponseData) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.ErrorDesc)
}

func RemovePromoUsers(req *RemovePromoUsersRequest) (interface{}, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionRemovePromoUsersRequest()

	r.SetPromoId(req.PromoId)
	r.SetPin(req.Pin)

	if len(req.Ip) > 0 {
		r.SetIp(req.Ip)
	}
	if len(req.Port) > 0 {
		r.SetPort(req.Port)
	}
	if len(req.RequestId) > 0 {
		r.SetRequestId(req.RequestId)
	}
	if len(req.BeginTime) > 0 {
		r.SetBeginTime(req.BeginTime)
	}
	if len(req.EndTime) > 0 {
		r.SetEndTime(req.EndTime)
	}

	var response RemovePromoUsersResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result, nil
}
