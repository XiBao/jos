package asset

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/asset"
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

func (r ActivityCreateResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response == nil || r.Response.IsError()
}

func (r ActivityCreateResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Response != nil {
		return r.Response.Error()
	}
	return "no result data"
}

type ActivityCreateRes struct {
	Code string              `json:"code,omitempty" codec:"code,omitempty"`
	Res  *ActivityCreateData `json:"response,omitempty" codec:"response,omitempty"`
}

func (r ActivityCreateRes) IsError() bool {
	return r.Res == nil || r.Res.IsError()
}

func (r ActivityCreateRes) Error() string {
	if r.Res != nil {
		return r.Res.Error()
	}
	return "no result data"
}

type ActivityCreateData struct {
	Code    string               `json:"code,omitempty" codec:"code,omitempty"`
	Message string               `json:"message,omitempty" codec:"message,omitempty"`
	Data    *ActivityCreateToken `json:"data,omitempty" codec:"data,omitempty"`
}

func (r ActivityCreateData) IsError() bool {
	return r.Code != "200"
}

func (r ActivityCreateData) Error() string {
	return fmt.Sprintf("code: %s, message: %s", r.Code, r.Message)
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

	var response ActivityCreateResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Response.Res.Data.Token, nil
}
