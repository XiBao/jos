package crm

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
	"github.com/daviddengcn/ljson"
)

type ShopGiftValidActivityRequest struct {
	api.BaseRequest

	Channel uint8 `json:"channel" codec:"channel"` // 渠道来源，pc:1 app:2 等
}

type ShopGiftValidActivityResponse struct {
	ErrorResp *api.ErrorResponnse        `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopGiftValidActivityData `json:"jingdong_pop_crm_shopGift_getValidActivity_responce,omitempty" codec:"jingdong_pop_crm_shopGift_getValidActivity_responce,omitempty"`
}

type ShopGiftValidActivityData struct {
	Code      string                       `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                       `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftValidActivityResult `json:"commonResult,omitempty" codec:"commonResult,omitempty"`
}

type ShopGiftValidActivityResult struct {
	Code string                  `json:"code,omitempty" codec:"code,omitempty"`
	Desc string                  `json:"msg,omitempty" codec:"desc,omitempty"`
	Data *ShopGiftActivityDomain `json:"data,omitempty" codec:"data,omitempty"`
}

type ShopGiftActivityDomain struct {
	Id                uint64 `json:"id"`                // 活动id
	VenderId          uint64 `json:"venderId"`          // 商家id
	ModelId           uint64 `json:"modelId"`           // 人群模型id
	ActivityName      string `json:"activityName"`      // 活动名称
	ActivityStartTime uint64 `json:"activityStartTime"` // 活动开始时间
	ActivityEndTime   uint64 `json:"activityEndTime"`   // 活动结束时间
}

func ShopGiftValidActivity(req *ShopGiftValidActivityRequest) (*ShopGiftActivityDomain, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewShopGiftValidActivityRequest()
	r.SetChannel(req.Channel)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result info.")
	}
	var response ShopGiftValidActivityResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}
	if response.Data.Code != "0" {
		return nil, errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == nil {
		return nil, errors.New("No result.")
	}
	if response.Data.Result.Code != "200" {
		if response.Data.Result.Desc == "" {
			return nil, errors.New("未知错误")
		} else {
			return nil, errors.New(response.Data.Result.Desc)
		}
	}
	return response.Data.Result.Data, nil
}
