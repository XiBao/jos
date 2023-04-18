package adkcunit

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/dsp"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/dsp/adkcunit"
)

type AdkcunitAdgroupListRequest struct {
	api.BaseRequest
	PageNum    int    `json:"page_num,omitempty" codec:"page_num,omitempty"`
	PageSize   int    `json:"page_size,omitempty" codec:"page_size,omitempty"`
	CampaignId uint64 `json:"campaign_id,omitempty" codec:"campaign_id,omitempty"` // 计划id
}

type AdkcunitAdgroupListResponse struct {
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *AdkcunitAdgroupListData `json:"jingdong_dsp_adkcunit_adgroup_list_responce,omitempty" codec:"jingdong_dsp_adkcunit_adgroup_list_responce,omitempty"`
}

type AdkcunitAdgroupListData struct {
	Result *AdkcunitAdgroupListResult `json:"querylistbyparam_result,omitempty" codec:"querylistbyparam_result,omitempty"`
}

type AdkcunitAdgroupListResult struct {
	ErrorMsg   string                    `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	ResultCode string                    `json:"resultCode,omitempty" codec:"resultCode,omitempty"`
	Success    bool                      `json:"success,omitempty" codec:"success,omitempty"`
	Value      *AdkcunitAdgroupListValue `json:"data,omitempty" codec:"data,omitempty"`
}

type AdkcunitAdgroupListValue struct {
	Paginator *dsp.Paginator  `json:"paginator,omitempty" codec:"paginator,omitempty"` // 分页组件
	Datas     []*ADGroupQuery `json:"datas,omitempty" codec:"datas,omitempty"`
}

type ADGroupQuery struct {
	NewAreaId      string `json:"newAreaId,omitempty" codec:"newAreaId,omitempty"`           // 推广区域id
	Id             uint64 `json:"id,omitemmpty" codec:"id,omitempty"`                        // 单元ID
	Position       string `json:"position,omitempty" codec:"position,omitempty"`             // 广告位展示
	OuterFeeStr    string `json:"outerFeeStr,omitempty" codec:"outerFeeStr,omitempty"`       // 站外出价
	CampaignId     uint64 `json:"campaignId,omitempty" codec:"campaignId,omitepty"`          // 计划ID
	Status         uint8  `json:"status,omitempty" codec:"status,omitempty"`                 // status
	Name           string `json:"name,omitempty" codec:"name,omitempty"`                     // 推广单元名
	BillingType    uint8  `json:"billingType,omitempty" codec:"billingType,omitempty"`       //	计费类型
	GroupDirection string `json:"groupDirection,omitempty" codec:"groupDirection,omitempty"` // 人群定向
	InSearchFee    uint64 `json:"inSearchFee,omitempty" codec:"inSearchFee,omitempty"`       //	搜索出价
	FeeStr         string `json:"feeStr,omitempty" codec:"feeStr,omitempty"`                 // 站内出价
	Area           string `json:"area,omitempty" codec:"area,omitempty"`                     // 推广区域

	CreatedTime string `json:"createdTime,omitempty" codec:"createdTime,omitempty"` // 推广区域
	PutType     int8   `json:"putType,omitempty" codec:"putType,omitempty"`         // 推广类型
}

// 获取计划下的推广单元列表
func AdkcunitAdgroupList(req *AdkcunitAdgroupListRequest) ([]*ADGroupQuery, int, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := adkcunit.NewAdkcunitAdgroupListRequest()
	r.SetCampaignId(req.CampaignId)
	r.SetPageSize(req.PageSize)
	r.SetPageNum(req.PageNum)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, 0, err
	}
	if len(result) == 0 {
		return nil, 0, errors.New("no result info")
	}
	var response AdkcunitAdgroupListResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, 0, err
	}
	if response.ErrorResp != nil {
		return nil, 0, response.ErrorResp
	}

	if !response.Data.Result.Success {
		return nil, 0, errors.New(response.Data.Result.ErrorMsg)
	}

	return response.Data.Result.Value.Datas, response.Data.Result.Value.Paginator.Items, nil

}
