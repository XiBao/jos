package promotion

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type AppendPromoUsersRequest struct {
	api.BaseRequest
	Ip        string `json:"ip,omitempty" codec:"ip,omitempty"`
	Port      string `json:"port,omitempty" codec:"port,omitempty"`
	RequestId string `json:"request_id,omitempty" codec:"request_id,omitempty"`
	BeginTime string `json:"begin_time,omitempty" codec:"begin_time,omitempty"`
	EndTime   string `json:"end_time,omitempty" codec:"end_time,omitempty"`
	PromoId   uint64 `json:"promo_id" codec:"promo_id"`
	Pin       string `json:"pin" codec:"pin"`
}

type AppendPromoUsersResponse struct {
	ErrorResp *api.ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AppendPromoUsersResponseData `json:"jingdong_seller_promotion_appendPromoUsers_responce,omitempty" codec:"jingdong_seller_promotion_appendPromoUsers_responce,omitempty"`
}

type AppendPromoUsersResponseData struct {
	Code      string      `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string      `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    interface{} `json:"result,omitempty" codec:"result,omitempty"`
}

func AppendPromoUsers(req *AppendPromoUsersRequest) (interface{}, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionAppendPromoUsersRequest()

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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result.")
	}

	var response AppendPromoUsersResponse
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

	return response.Data.Result, nil
}
