package crm

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
)

type ShopGiftCloseRequest struct {
	api.BaseRequest

	ActivityId uint64 `json:"activityId" codec:"activityId"` // 活动ID
}

type ShopGiftCloseResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopGiftCloseData  `json:"jingdong_pop_crm_shopGift_closeShopGiftCallBack_responce,omitempty" codec:"jingdong_pop_crm_shopGift_closeShopGiftCallBack_responce,omitempty"`
}

func (r ShopGiftCloseResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ShopGiftCloseResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ShopGiftCloseData struct {
	Code      string               `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string               `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftCloseResult `json:"closeshopgiftcallback_result,omitempty" codec:"closeshopgiftcallback_result,omitempty"`
}

func (r ShopGiftCloseData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r ShopGiftCloseData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type ShopGiftCloseResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Data bool   `json:"data,omitempty" codec:"data,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
}

func (r ShopGiftCloseResult) IsError() bool {
	return r.Code != "200"
}

func (r ShopGiftCloseResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

func ShopGiftClose(req *ShopGiftCloseRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewShopGiftCloseRequest()
	r.SetActivityId(req.ActivityId)

	var response ShopGiftCloseResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return response.Data.Result.Data, nil
}
