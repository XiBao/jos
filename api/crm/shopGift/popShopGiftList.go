package crm

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
)

type PopShopGiftListRequest struct {
	api.BaseRequest

	UserPin     string `json:"userPin" codec:"userPin"`
	Channel     uint8  `json:"channel" codec:"channel"` // 渠道来源，pc:1 app:2 等
	OpenIdBuyer string `json:"open_id_buyer" codec:"open_id_buyer"`
}

type PopShopGiftListResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *PopShopGiftListData `json:"jingdong_pop_crm_shopGift_getPopShopGiftList_responce,omitempty" codec:"jingdong_pop_crm_shopGift_getPopShopGiftList_responce,omitempty"`
}

func (r PopShopGiftListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r PopShopGiftListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type PopShopGiftListData struct {
	Code      string                 `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                 `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *PopShopGiftListResult `json:"commonResult,omitempty" codec:"commonResult,omitempty"`
}

func (r PopShopGiftListData) IsError() bool {
	return r.Code != "0" || r.Result == nil || r.Result.IsError()
}

func (r PopShopGiftListData) Error() string {
	if r.Result != nil && r.Result.IsError() {
		return r.Result.Error()
	}
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

type PopShopGiftListResult struct {
	Code string                 `json:"code,omitempty" codec:"code,omitempty"`
	Desc string                 `json:"desc,omitempty" codec:"desc,omitempty"`
	Data []PopShopGiftListModel `json:"data,omitempty" codec:"data,omitempty"`
}

func (r PopShopGiftListResult) IsError() bool {
	return r.Code != "200"
}

func (r PopShopGiftListResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

type PopShopGiftListModel struct {
	Id            uint64   `json:"id"`            // 奖励Id
	VenderId      uint64   `json:"venderId"`      // 商家id
	ActivityId    uint64   `json:"activityId"`    // 活动id
	PrizeType     uint8    `json:"prizeType"`     // 奖品类型
	Discount      uint     `json:"discount"`      // 积分/京豆：单位发放数量；优惠金额: 优惠券面值
	Quota         uint     `json:"quota"`         // 满足优惠的最低消费额
	ModelNameList []string `json:"modelNameList"` // 人群列表
}

func PopShopGiftList(ctx context.Context, req *PopShopGiftListRequest) ([]PopShopGiftListModel, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewPopShopGiftListRequest()
	r.SetChannel(req.Channel)
	r.SetUserPin(req.UserPin)
	r.SetOpenIdBuyer(req.OpenIdBuyer)

	var response PopShopGiftListResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.Result.Data, nil
}
