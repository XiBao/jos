package points

import (
	"encoding/json"
	"errors"

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

type GetCouponInfoData struct {
	JsfResult *GetCouponInfoJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

type GetCouponInfoJsfResult struct {
	Code   string              `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	Desc   string              `json:"desc,omitempty" codec:"desc,omitempty"`     //返回描述
	Result []*PointsCouponInfo `json:"result,omitempty" codec:"result,omitempty"` //优惠券信息
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
func GetCouponInfo(req *GetCouponInfoRequest) ([]*PointsCouponInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewGetCouponInfoRequest()

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}
	var response GetCouponInfoResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.JsfResult.Code != "200" {
		return nil, errors.New(response.Data.JsfResult.Desc)
	}

	return response.Data.JsfResult.Result, nil

}
