package points

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
)

type GetCouponInfoRequest struct {
	api.BaseRequest
}

type GetCouponInfoResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetCouponInfoData  `json:"jingdong_points_jos_getCouponInfo_responce,omitempty" codec:"jingdong_points_jos_getCouponInfo_responce,omitempty"`
}

func (r GetCouponInfoResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetCouponInfoResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data.IsError() {
		return r.Data.Error()
	}
	return "no result data"
}

type GetCouponInfoData struct {
	JsfResult *GetCouponInfoJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

func (r GetCouponInfoData) IsError() bool {
	return r.JsfResult == nil || r.JsfResult.IsError()
}

func (r GetCouponInfoData) Error() string {
	if r.IsError() {
		return r.JsfResult.Error()
	}
	return "no result data"
}

type GetCouponInfoJsfResult struct {
	Code   string             `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	Desc   string             `json:"desc,omitempty" codec:"desc,omitempty"`     //返回描述
	Result []PointsCouponInfo `json:"result,omitempty" codec:"result,omitempty"` //优惠券信息
}

func (r GetCouponInfoJsfResult) IsError() bool {
	return r.Code != "200"
}

func (r GetCouponInfoJsfResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

type PointsCouponInfo struct {
	BatchId           uint64   `json:"batch_id,omitempty" codec:"batch_id,omitempty"`                   // 批次id
	BatchKey          string   `json:"batchKey,omitempty" codec:"batchKey,omitempty"`                   // 批次Key
	VenderId          uint64   `json:"venderId,omitempty" codec:"venderId,omitempty"`                   // 	商家ID
	Create            string   `json:"create,omitempty" codec:"create,omitempty"`                       // 优惠券创建时间
	Discount          uint64   `json:"discount,omitempty" codec:"discount,omitempty"`                   // 积分券面额
	Condition         uint64   `json:"condition,omitempty" codec:"condition,omitempty"`                 // 	满减金额
	CouponType        uint8    `json:"couponType,omitempty" codec:"couponType,omitempty"`               // 优惠券类型 0:京券 1:东券
	Points            uint64   `json:"points,omitempty" codec:"points,omitempty"`                       // 所需积分值
	UsePlatList       []uint8  `json:"usePlatList,omitempty" codec:"usePlatList,omitempty"`             // 使用平台
	PlatFormDesc      []string `json:"platFormDesc,omitempty" codec:"platFormDesc,omitempty"`           //使用平台描述
	Period            uint64   `json:"period,omitempty" codec:"period,omitempty"`                       //优惠券有效期
	SendCount         uint64   `json:"sendCount,omitempty" codec:"sendCount,omitempty"`                 //发行量
	TradeCount        uint64   `json:"tradeCount,omitempty" codec:"tradeCount,omitempty"`               //已经领取量
	RemainingCount    uint64   `json:"remainingCount,omitempty" codec:"remainingCount,omitempty"`       //	剩余量
	FullPlat          uint8    `json:"fullPlat,omitempty" codec:"fullPlat,omitempty"`                   //是否全平台使用 1：全平台 3：限平台
	ActivityStartTime string   `json:"activityStartTime,omitempty" codec:"activityStartTime,omitempty"` //活动开始时间
	ActivityEndTime   string   `json:"activityEndTime,omitempty" codec:"activityEndTime,omitempty"`     //活动结束时间
	RealCouponId      uint64   `json:"realCouponId,omitempty" codec:"realCouponId,omitempty"`           //卡券组ID
}

// TODO 通过venderId查询商家设置的积分可兑换优惠券信息
func GetCouponInfo(req *GetCouponInfoRequest) ([]PointsCouponInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewGetCouponInfoRequest()

	var response GetCouponInfoResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.JsfResult.Result, nil

}
