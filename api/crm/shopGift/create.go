package crm

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
)

type ShopGiftCreateRequest struct {
	api.BaseRequest

	Name          string `json:"name" codec:"name"`                                   // 活动名称
	StartDate     string `json:"startDate" codec:"startDate"`                         // 活动开始时间
	EndDate       string `json:"endDate" codec:"endDate"`                             // 活动结束时间
	ModelIds      string `json:"modelIds" codec:"modelIds"`                           // 活动对应的人群包的id,用下划线分割;商家的人群包信息通过jingdong.pop.crm.shopGift.getGroupModelList查询，系统定义人群和系统推荐及自定义人群不能同时存在，且系统自定义人群单选
	Creator       string `json:"creator" codec:"creator"`                             // 活动创建者，商家用户pin
	CreateChannel string `json:"createChannel" codec:"createChannel"`                 // 活动创建渠道，isv侧填写固定值：SHOP_GIFT_ISV
	CloseLink     string `json:"closeLink" codec:"closeLink"`                         // 关闭活动链接，商家可在商家后台点击链接，关闭活动
	IsvName       string `json:"isvName" codec:"isvName"`                             // ISV服务商的名称/在京麦平台作为插件的名称
	ActivityLink  string `json:"activityLink" codec:"activityLink"`                   // ISV活动页链接，做奖项的展示，需要提前与京东app侧协议接入
	RequestPin    string `json:"requestPin" codec:"requestPin"`                       // 创建者的用户pin
	IsvValidate   string `json:"isvValidate,omitempty" codec:"isvValidate,omitempty"` // isv插件有效期
	SendNumber    string `json:"sendNumber,omitempty" codec:"sendNumber,omitempty"`   // 发放总量
	OpenIdSeller  string `json:"open_id_seller" codec:"open_id_seller"`               // 活动创建者，商家用户pin
	OpenIdIsv     string `json:"open_id_isv" codec:"open_id_isv"`                     // 创建者的用户pin
	Discount      string `json:"discount" codec:"discount"`                           // 优惠金额 单位奖项额度 京券、东券、京豆、积分； 例满100减1的东券，此处填写1，
	Quota         string `json:"quota,omitempty" codec:"quota,omitempty"`             // 满足优惠的消费金额 -东券；例满100减1的东券，此处填写100
	ValidateDay   string `json:"validateDay,omitempty" codec:"validateDay,omitempty"` // 有效期(天数) --东券 京券
	PrizeType     string `json:"prizeType" codec:"prizeType"`                         // 奖品类型code及代表类型如下： 0:京券 1:东券 4:京豆 6:积分
	SendCount     string `json:"sendCount" codec:"sendCount"`                         // 奖项发放人数
}

type ShopGiftCreateResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopGiftCreateData `json:"jingdong_pop_crm_shopGift_createShopGift_responce,omitempty" codec:"jingdong_pop_crm_shopGift_createShopGift_responce,omitempty"`
}

type ShopGiftCreateData struct {
	Code      string                `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftCreateResult `json:"createisvshopgift_result,omitempty" codec:"createisvshopgift_result,omitempty"`
}

type ShopGiftCreateResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"`
	Data uint64 `json:"data,omitempty" codec:"data,omitempty"`
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"`
}

func ShopGiftCreate(req *ShopGiftCreateRequest) (uint64, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewShopGiftCreateRequest()
	r.SetName(req.Name)
	r.SetStartDate(req.StartDate)
	r.SetEndDate(req.EndDate)
	r.SetModelIds(req.ModelIds)
	r.SetCreator(req.Creator)
	r.SetCreateChannel(req.CreateChannel)
	r.SetCloseLink(req.CloseLink)
	r.SetIsvName(req.IsvName)
	r.SetActivityLink(req.ActivityLink)
	r.SetRequestPin(req.RequestPin)
	r.SetOpenIdSeller(req.OpenIdSeller)
	r.SetOpenIdIsv(req.OpenIdIsv)
	r.SetDiscount(req.Discount)
	r.SetPrizeType(req.PrizeType)
	r.SetSendCount(req.SendCount)
	if req.IsvValidate != "" {
		r.SetIsvValidate(req.IsvValidate)
	}
	if req.SendNumber != "" {
		r.SetSendNumber(req.SendNumber)
	}
	if req.Quota != "" {
		r.SetQuota(req.Quota)
	}
	if req.ValidateDay != "" {
		r.SetValidateDay(req.ValidateDay)
	}
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return 0, err
	}
	if len(result) == 0 {
		return 0, errors.New("No result info.")
	}
	var response ShopGiftCreateResponse
	err = json.Unmarshal(result, &response)
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
		return 0, errors.New("No create result.")
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
