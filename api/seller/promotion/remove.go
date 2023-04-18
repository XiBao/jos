package promotion

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/seller/promotion"
)

type RemoveRequest struct {
	api.BaseRequest
	Ip        string `json:"ip,omitempty" codec:"ip,omitempty"`                 // 调用方IP
	Port      string `json:"port,omitempty" codec:"port,omitempty"`             // 调用方端口
	RequestId string `json:"request_id,omitempty" codec:"request_id,omitempty"` // 防重码。业务唯一值
	PromoId   uint64 `json:"promo_id,omitempty" codec:"promo_id,omitempty"`     // 促销编号
	PromoType uint8  `json:"promo_type,omitempty" codec:"promo_type,omitempty"` // 促销类型。1:单品促销,4:赠品促销,6:套装促销,10:总价促销
}

type RemoveResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *RemoveData         `json:"jingdong_seller_promotion_v2_remove_responce,omitempty" codec:"jingdong_seller_promotion_v2_remove_responce,omitempty"`
}

type RemoveData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`

	RemoveResult bool `json:"remove_result,omitempty" codec:"remove_result,omitempty"`
}

func Remove(req *RemoveRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := promotion.NewSellerPromotionRemoveRequest()
	r.SetIp(req.Ip)
	r.SetPort(req.Port)
	r.SetPromoId(req.PromoId)
	r.SetPromoType(req.PromoType)
	if req.RequestId != "" {
		r.SetRequestId(req.RequestId)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	result = util.RemoveJsonSpace(result)

	var response RemoveResponse
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

	return response.Data.RemoveResult, nil
}
