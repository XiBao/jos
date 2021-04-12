package asset

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/util"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/asset"
	"github.com/daviddengcn/ljson"
)

type ActivityCreateRequest struct {
	api.BaseRequest

	ActivityId   string `json:"activity_id" codec:"activity_id"`     // 活动ID
	ActivityName string `json:"activity_name" codec:"activity_name"` // 活动名称
	BeginDate    string `json:"begin_date" codec:"begin_date"`       // 活动开始时间
	EndDate      string `json:"end_date" codec:"end_date"`           // 活动结束时间
	Tool         string `json:"tool" codec:"tool"`                   // 工具名称：自定义的一个固定值，用于区分创建活动的渠道
	Details      string `json:"details" codec:"details"`             // JSON格式的活动配置, itemId:资产项ID[1:流量(1M流量), 2:E卡(1元), 3:E卡(5元), 4:E卡(20元), 5:E卡(50元), 6:E卡(100元), 7:自营卡(PLUS会员), 8:自营卡(爱奇艺会员(月)), 12:自营卡(爱奇艺会员(季)), 13:自营卡(爱奇艺会员(半年)), 14:自营卡(爱奇艺会员(年)), 15:红包(分), 16:短信(条)], 可通过查询账户余额信息接口获得; quantity:活动需要冻结的资产项的数量。如果活动中包含红包，请参考:https://docs.qq.com/doc/DVEl0T01xTHZEdFBM
}

type ActivityCreateResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *ActivityCreateRes  `json:"jingdong_asset_activity_create_responce,omitempty" codec:"jingdong_asset_activity_create_responce,omitempty"`
}

type ActivityCreateRes struct {
	Code string              `json:"code,omitempty" codec:"code,omitempty"`
	Res  *ActivityCreateData `json:"response,omitempty" codec:"response,omitempty"`
}

type ActivityCreateData struct {
	Code    string               `json:"code,omitempty" codec:"code,omitempty"`
	Message string               `json:"message,omitempty" codec:"message,omitempty"`
	Data    *ActivityCreateToken `json:"data,omitempty" codec:"data,omitempty"`
}

type ActivityCreateToken struct {
	Token string `json:"token" codec:"token"`
}

func ActivityCreate(req *ActivityCreateRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := asset.NewActivityCreateRequest()
	r.SetActivityId(req.ActivityId)
	r.SetActivityName(req.ActivityName)
	r.SetBeginDate(req.BeginDate)
	r.SetEndDate(req.EndDate)
	r.SetTool(req.Tool)
	r.SetDetails(req.Details)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	result = util.RemoveJsonSpace(result)

	var response ActivityCreateResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	if response.ErrorResp != nil {
		return "", response.ErrorResp
	}
	if response.Response == nil || response.Response.Res == nil {
		return "", errors.New("No response.")
	}
	if response.Response.Res.Code != "200" && response.Response.Res.Message != "" {
		return "", errors.New(response.Response.Res.Message)
	}
	if response.Response.Res.Data == nil {
		return "", errors.New("No result.")
	}

	return response.Response.Res.Data.Token, nil
}
