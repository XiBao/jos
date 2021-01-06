package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
	"github.com/daviddengcn/ljson"
)

type ShopGiftCloseRequest struct {
	api.BaseRequest

	ActivityId uint64 `json:"activityId" codec:"activityId"` // 活动ID
}

type ShopGiftCloseResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopGiftCloseData  `json:"jingdong_pop_crm_shopGift_closeShopGiftCallBack_responce,omitempty" codec:"jingdong_pop_crm_shopGift_closeShopGiftCallBack_responce,omitempty"`
}

type ShopGiftCloseData struct {
	Code      string               `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string               `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftCloseResult `json:"closeshopgiftcallback_result,omitempty" codec:"closeshopgiftcallback_result,omitempty"`
}

type ShopGiftCloseResult struct {
	Code string `json:"code,omitempty" codec:"code,omiempty"`
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,oomitempty"`
}

func ShopGiftClose(req *ShopGiftCloseRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewShopGiftCloseRequest()
	r.SetActivityId(req.ActivityId)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("No result info.")
	}
	var response ShopGiftCloseResponse
	err = ljson.Unmarshal(result, &response)
	if response.ErrorResp != nil {
		return 0, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return 0, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return 0, errors.New("No close result.")
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
