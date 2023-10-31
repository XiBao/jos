package crm

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
)

type ShopGiftValidActivityRequest struct {
	api.BaseRequest

	Channel uint8 `json:"channel" codec:"channel"` // 渠道来源，pc:1 app:2 等
}

type ShopGiftValidActivityResponse struct {
	ErrorResp *api.ErrorResponnse        `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopGiftValidActivityData `json:"jingdong_pop_crm_shopGift_getValidActivity_responce,omitempty" codec:"jingdong_pop_crm_shopGift_getValidActivity_responce,omitempty"`
}

func (r ShopGiftValidActivityResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r ShopGiftValidActivityResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type ShopGiftValidActivityData struct {
	Code      string                       `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                       `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftValidActivityResult `json:"commonResult,omitempty" codec:"commonResult,omitempty"`
}

func (r ShopGiftValidActivityData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r ShopGiftValidActivityData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return fmt.Sprintf("code: %s, msg: %s", r.Code, r.ErrorDesc)
}

type ShopGiftValidActivityResult struct {
	Code string                  `json:"code,omitempty" codec:"code,omitempty"`
	Desc string                  `json:"msg,omitempty" codec:"desc,omitempty"`
	Data *ShopGiftActivityDomain `json:"data,omitempty" codec:"data,omitempty"`
}

func (r ShopGiftValidActivityResult) IsError() bool {
	return r.Code != "200"
}

func (r ShopGiftValidActivityResult) Error() string {
	return fmt.Sprintf("code: %s, msg: %s", r.Code, r.Desc)
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

	var response ShopGiftValidActivityResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data, nil
}
