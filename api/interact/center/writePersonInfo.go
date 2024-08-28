package center

import (
	"context"

	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/interact/center"
)

type WritePersonInfoRequest struct {
	BaseRequest
	AppName     string `json:"appName,omitempty" codec:"appName,omitempty"`             // ISV服务商名称
	Channel     uint8  `json:"channel,omitempty" codec:"channel,omitempty"`             // PC:1, APP:2, 任务中心:3,发现频道:4, 上海运营模板::5 , 微信: 6, QQ: 7, ARVR: 9
	Pin         string `json:"pin,omitempty" codec:"pin,omitempty"`                     // 用户pin
	OpenIdBuyer string `json:"open_id_buyer,omitempty" codec:"open_id_buyer,omitempty"` // 用户pin
	XIdBuyer    string `json:"xid_buyer,omitempty" codec:"xid_buyer,omitempty"`         // 用户pin
	ProfileUrl  string `json:"profileUrl,omitempty" codec:"profileUrl,omitempty"`       // 用户头像连接 非必填
	ActivityId  uint64 `json:"activityId,omitempty" codec:"activityId,omitempty"`       // 活动id
	Created     string `json:"created,omitempty" codec:"created,omitempty"`             // 创建时间
	StartTime   string `json:"startTime,omitempty" codec:"startTime,omitempty"`         // 活动开始时间
	EndTime     string `json:"endTime,omitempty" codec:"endTime,omitempty"`             // 结束时间
	Id          string `json:"id,omitempty" codec:"id,omitempty"`                       // id（isv自己的活动id） 非必填
	Type        uint8  `json:"type,omitempty" codec:"type,omitempty"`                   // 活动类型
	/*
	   投票有礼 6：购物车红包 8: 盖楼有礼 11：拼购定向投放 19: 分享有礼
	    20: 集卡有礼 23 锦鲤圈抽奖 25 抽奖 26 加购 27 签到 28 积分兑换 29 商品收藏 30 游戏 31 砍价拼团 32 专享价
	    33 组队 34 知识超人 35：粉丝互动 36：限时抢券 37：N元试用 38：前N名优惠 39：定制活动 40评价有礼
	    41 买家秀征集 42 关注有礼 43邀请有礼 44 浏览有礼(其他类型请联系产品)
	*/
	ActionType string `json:"actionType,omitempty" codec:"actionType,omitempty"`
	//  行为动作类型 1: 参与 2：分享 3： 助力 4：浏览店铺 5：浏览商品 6：关注店铺 7：关注商品 8:开卡会员 9：加购
}

type WritePersonInfoResponse struct {
	ErrorResp *ErrorResponnse          `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  WritePersonInfoResponse1 `json:"jingdong_interact_center_api_service_write_writePersonInfo_responce,omitempty" codec:"jingdong_interact_center_api_service_write_writePersonInfo_responce,omitempty"`
}

func (r WritePersonInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response.IsError()
}

func (r WritePersonInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return r.Response.Error()
}

type WritePersonInfoResponse1 struct {
	Result WritePersonInfoResult `json:"giftActivityResults" codec:"giftActivityResults"`
}

func (r WritePersonInfoResponse1) IsError() bool {
	return r.Result.IsError()
}

func (r WritePersonInfoResponse1) Error() string {
	return r.Result.Error()
}

type WritePersonInfoResult struct {
	Data bool   `json:"data" codec:"data"` // 请求是否成功
	Code int    `json:"code" codec:"code"` // 返回状态码
	Msg  string `json:"msg" codec:"msg"`
}

func (r WritePersonInfoResult) IsError() bool {
	return r.Code != 200 && r.Code != 0
}

func (r WritePersonInfoResult) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

func WritePersonInfo(ctx context.Context, req *WritePersonInfoRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewWritePersonInfoRequest()
	r.SetAppName(req.AppName)
	r.SetChannel(req.Channel)
	r.SetPin(req.Pin)
	r.SetOpenIdBuyer(req.OpenIdBuyer)
	r.SetXIdBuyer(req.XIdBuyer)
	r.SetActivityId(req.ActivityId)
	r.SetCreated(req.Created)
	r.SetStartTime(req.StartTime)
	r.SetEndTime(req.EndTime)
	r.SetType(req.Type)
	r.SetActionType(req.ActionType)

	if len(req.ProfileUrl) > 0 {
		r.SetProfileUrl(req.ProfileUrl)
	}

	var response WritePersonInfoResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}

	return response.Response.Result.Data, nil
}
