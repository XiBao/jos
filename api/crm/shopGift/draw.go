package crm

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
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

func (r ShopGiftDrawResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ShopGiftDrawResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ShopGiftDrawData struct {
	Code      string              `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string              `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftDrawResult `json:"createshopgift_result,omitempty" codec:"createshopgift_result,omitempty"`
}

func (r ShopGiftDrawData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r ShopGiftDrawData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return fmt.Sprintf("code: %s, msg: %s", r.Code, r.ErrorDesc)
}

type ShopGiftDrawResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
}

func (r ShopGiftDrawResult) IsError() bool {
	return r.Code != "200"
}

func (r ShopGiftDrawResult) Error() string {
	return fmt.Sprintf("code: %s, msg: %s", r.Code, r.Desc)
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

	var response ShopGiftDrawResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return 0, err
	}
	return response.Data.Result.Data, nil
}
