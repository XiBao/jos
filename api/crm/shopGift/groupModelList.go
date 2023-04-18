package crm

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	crm "github.com/XiBao/jos/sdk/request/crm/shopGift"
)

type ShopGiftGroupModelListRequest struct {
	api.BaseRequest

	Channel uint8 `json:"channel" codec:"channel"` // 渠道来源，pc:1 app:2 等
}

type ShopGiftGroupModelListResponse struct {
	ErrorResp *api.ErrorResponnse         `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *ShopGiftGroupModelListData `json:"jingdong_pop_crm_shopGift_getGroupModelList_responce,omitempty" codec:"jingdong_pop_crm_shopGift_getGroupModelList_responce,omitempty"`
}

type ShopGiftGroupModelListData struct {
	Code      string                        `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string                        `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    *ShopGiftGroupModelListResult `json:"commonResult,omitempty" codec:"commonResult,omitempty"`
}

type ShopGiftGroupModelListResult struct {
	Code string                `json:"code,omitempty" codec:"code,omitempty"`
	Desc string                `json:"desc,omitempty" codec:"desc,omitempty"`
	Data []*ShopGiftGroupModel `json:"data,omitempty" codec:"data,omitempty"`
}

type ShopGiftGroupModel struct {
	Id            uint64   `json:"id"`               // 人群id
	VenderId      int      `json:"venderId"`         // 商家id
	ModelId       int      `json:"modelId"`          // 人群模型id
	ModelName     string   `json:"modelName"`        // 人群名称
	ModelDesc     string   `json:"modelDesc"`        // 人群名称描述
	SRange        string   `json:"range,omitempty"`  // 范围
	RuleExp       string   `json:"ruleExp"`          // 规则
	IsDelete      uint8    `json:"isDelete"`         // 是否删除
	CreateTime    uint64   `json:"createTime"`       // 创建时间
	ModifyTime    uint64   `json:"modifyTime"`       // 修改时间
	ModelNum      uint     `json:"modelNum"`         // 模型数
	ModelType     uint     `json:"modelType"`        // 模型类型
	TaskId        uint64   `json:"taskId,omitempty"` // 任务id
	Status        int      `json:"status"`           // 状态
	ModelDescList []string `json:"modelDescList"`    // 人群信息描述
}

func ShopGiftGroupModelList(req *ShopGiftGroupModelListRequest) ([]*ShopGiftGroupModel, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := crm.NewShopGiftGroupModelListRequest()
	r.SetChannel(req.Channel)
	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result info.")
	}
	var response ShopGiftGroupModelListResponse
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
