package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
	"github.com/daviddengcn/ljson"
)

type ShopGiftDrawRequest struct {
	api.BaseRequest

	UserPin     string `json:"userPin" codec:"userPin"`             // 用户密文pin
	ActivityId  uint64 `json:"activityId" codec:"activityId"`       // 活动ID
	Ip          string `json:"ip" codec:"ip"`                       // 领取礼包的ip
	BussinessId string `json:"bussinessId" codec:"bussinessId"`     // 调用方的业务id
	Channel     uint8  `json:"channel" codec:"channel"`             // 领取礼包的渠道
	OpenIdBuyer string `json:"open_id_buyer" codec:"open_id_buyer"` // 用户密文pin
}

type ShopGiftDrawResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopGiftDrawData   `json:"jingdong_pop_crm_shopGift_drawShopGift_responce,omitempty" codec:"jingdong_pop_crm_shopGift_drawShopGift_responce,omitempty"`
}

type ShopGiftDrawData struct {
	Code      string              `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string              `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftDrawResult `json:"createshopgift_result,omitempty" codec:"createshopgift_result,omitempty"`
}

type ShopGiftDrawResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
}

func ShopGiftDraw(req *ShopGiftDrawRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewShopGiftDrawRequest()
	r.SetUserPin(req.UserPin)
	r.SetActivityId(req.ActivityId)
	r.SetIp(req.Ip)
	r.SetBussinessId(req.BussinessId)
	r.SetChannel(req.Channel)
	r.SetOpenIdBuyer(req.OpenIdBuyer)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("No result info.")
	}
	var response ShopGiftDrawResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return 0, err
	}
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return 0, errors.New("No draw result.")
	}
	if response.Data.Result.Code != "200" {
		if response.Data.Result.Desc == "" {
			return 0, errors.New("未知错误")
		} else {
			return 0, errors.New(response.Data.Result.Desc)
		}
	}
	return response.Data.Result.Data, nil
}
