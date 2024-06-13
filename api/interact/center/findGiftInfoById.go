package center

import (
	. "github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/interact/center"
)

type FindGiftInfoByIdRequest struct {
	BaseRequest
	AppName    string `json:"appName" codec:"appName"`       // ISV服务商名称
	AppId      uint64 `json:"appId" codec:"appId"`           // 应用id，jos默认为1784
	Channel    uint8  `json:"channel" codec:"channel"`       // PC:1, APP:2, 任务中心:3,发现频道:4, 上海运营模板::5 , 微信: 6, QQ: 7, ARVR: 9
	Type       uint8  `json:"type" codec:"type"`             // 活动类型
	ActivityId uint64 `json:"activityId" codec:"activityId"` // 活动id
}

type FindGiftInfoByIdResponse struct {
	ErrorResp *ErrorResponnse           `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  FindGiftInfoByIdResponse1 `json:"jingdong_interact_center_api_service_read_GiftActivityResultReadService_findGiftInfoById_responce,omitempty" codec:"jingdong_interact_center_api_service_read_GiftActivityResultReadService_findGiftInfoById_responce,omitempty"`
}

func (r FindGiftInfoByIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response.IsError()
}

func (r FindGiftInfoByIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	return r.Response.Error()
}

type FindGiftInfoByIdResponse1 struct {
	Result FindGiftInfoByIdResult `json:"returnType" codec:"returnType"`
}

func (r FindGiftInfoByIdResponse1) IsError() bool {
	return r.Result.IsError()
}

func (r FindGiftInfoByIdResponse1) Error() string {
	return r.Result.Error()
}

type FindGiftInfoByIdResult struct {
	Data *GiftInfo `json:"data" codec:"data"` //互动活动信息
	Code int       `json:"code" codec:"code"` //返回状态码
	Msg  string    `json:"msg" codec:"msg"`
}

func (r FindGiftInfoByIdResult) IsError() bool {
	return r.Code != 200 && r.Code != 0
}

func (r FindGiftInfoByIdResult) Error() string {
	return sdk.ErrorString(r.Code, r.Msg)
}

type GiftInfo struct {
	IsPrize         uint8  `json:"isPrize"`                   // 是否为有奖活动： 0：无奖活动，1：有奖活动
	Modifier        string `json:"modifier,omitempty"`        // 商家修改活动的用户pin
	VenderId        uint64 `json:"venderId"`                  // 商家id
	SourceLink      string `json:"sourceLink,omitempty"`      // 其他活动连接
	IsSinglePrize   uint8  `json:"isSinglePrize,omitempty"`   // 是否是单独奖项(单独奖项在同一个活动中，同一个pin可以领取多种类型奖项) 0或null(默认): 同一个活动，一个pin只能领取一次(和collectTimes无关) 1是单独奖项：在同一个活动中，同一个pin可以领取多种类型奖项
	Source          uint8  `json:"source"`                    // 系统来源(PC:1, APP:2, 任务中心:3,发现:4, 上海运营模板::5 , 微信: 6, QQ: 7, ISV:8,ARVR: 9)
	Type            uint8  `json:"type"`                      // 活动类型（1：投票有礼 6：购物车红包 8: 盖楼有礼 11：拼购定向投放 19: 分享有礼 20: 集卡有礼 23 锦鲤圈抽奖 25 抽奖 26 加购 27 签到 28 积分兑换 29 商品收藏 30 游戏 31 砍价拼团 32 专享价 33 组队 34 知识超人 35：粉丝互动 36：限时抢券 37：N元试用 38：前N名优惠 39：定制活动 40评价有礼 41 买家秀征集 42 关注有礼 43邀请有礼 44 浏览有礼(其他类型请联系产品)）
	ModelIds        string `json:"modelIds,omitempty"`        // 用户人群Id,以下划线分割
	Modified        uint64 `json:"modified"`                  // 活动修改时间
	RfId            int64  `json:"rfId,omitempty"`            // 业务id
	StartTime       uint64 `json:"startTime"`                 // 活动开始时间
	Id              uint64 `json:"id"`                        // 活动id
	Validate        string `json:"validate,omitempty"`        // 插件有效期
	IsEverydayAward uint8  `json:"isEverydayAward,omitempty"` // 是否每天发放 0或null: 默认为活动维度发放，一个活动中，一个Pin只能领取一次 1: 按照每天发放
	Creator         string `json:"creator"`                   // 商家创建活动用户pin
	SubtitleName    string `json:"subtitleName,omitempty"`    // 活动副标题
	Created         uint64 `json:"created"`                   // 活动创建时间
	TaskIds         string `json:"taskIds,omitempty"`         // 用户人群taskIds
	Name            string `json:"name"`                      // 活动名称
	SourceCloseLink string `json:"sourceCloseLink,omitempty"` // 其他活动关闭连接
	PictureLink     string `json:"pictureLink,omitempty"`     // 活动图片连接
	EndTime         uint64 `json:"endTime"`                   // 活动结束时间
	SourceName      string `json:"sourceName"`                // 官方活动: interact-center 其他活动:isv
	Status          uint8  `json:"status"`                    // 活动状态(仅说明，写接口不需要传该值)：1：正在创建中 2: 已经创建 4：活动被商家取消 5: 活动已经结束(包含奖项发完及活动自然结束) 6：活动创建异常 9：进行中
}

func FindGiftInfoById(req *FindGiftInfoByIdRequest) (*GiftInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := center.NewFindGiftInfoByIdRequest()
	r.SetAppName(req.AppName)
	r.SetAppId(req.AppId)
	r.SetChannel(req.Channel)
	r.SetType(req.Type)
	r.SetActivityId(req.ActivityId)

	var response FindGiftInfoByIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}

	return response.Response.Result.Data, nil
}
