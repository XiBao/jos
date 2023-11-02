package item

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/kpl/open/item"
)

type FindJoinActivitiesRequest struct {
	api.BaseRequest
	Uid string `json:"uid,omitempty" codec:"uid,omitempty"`
	Sku uint64 `json:"sku,omitempty" codec:"sku,omitempty"`
}

type FindJoinActivitiesResponse struct {
	ErrorResp *api.ErrorResponnse     `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *FindJoinActivitiesData `json:"jd_kpl_open_item_findjoinactives_response,omitempty" codec:"jd_kpl_open_item_findjoinactives_responseomitempty"`
}

func (r FindJoinActivitiesResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r FindJoinActivitiesResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type FindJoinActivitiesData struct {
	Result FindJoinActivitiesResult `json:"findjoinactivesResult" codec:"findjoinactivesResult"`
}

func (r FindJoinActivitiesData) IsError() bool {
	return r.Result.IsError()
}

func (r FindJoinActivitiesData) Error() string {
	return r.Result.Error()
}

type FindJoinActivitiesResult struct {
	Err     string     `json:"error,omitempty" codec:"error,omitempty"`
	Coupons []CouponVo `json:"coupons,omitempty" codec:"coupons,omitempty"`
}

func (r FindJoinActivitiesResult) IsError() bool {
	return r.Err != ""
}

func (r FindJoinActivitiesResult) Error() string {
	return r.Err
}

type CouponVo struct {
	Coupon Coupon `json:"batchActiveVo,omitempty" codec:"batchActiveVo,omitempty"`
}

type CouponType uint

const (
	JING_QUAN          CouponType = 0
	DONG_QUAN          CouponType = 1
	FREE_SHIPPING_QUAN CouponType = 2
)

type Coupon struct {
	RoleId        uint64     `json:"roleId,omitempty" codec:"roleId,omitempty"`
	EncryptedKey  string     `json:"encryptedKey,omitempty" codec:"encryptedKey,omitempty"`
	ToUrl         string     `json:"toUrl,omitempty" codec:"toUrl,omitempty"`
	UserClass     string     `json:"userClass,omitempty" codec:"userClass,omitempty"`
	UserRiskLevel int        `json:"userRiskLevel,omitempty" codec:"userRiskLevel,omitempty"`
	BeginTime     int64      `json:"beginTime,omitempty" codec:"beginTime,omitempty"`
	EndTime       int64      `json:"endTime,omitempty" codec:"endTime,omitempty"`
	BatchId       uint64     `json:"batchId,omitempty" codec:"batchId,omitempty"`
	Discount      float64    `json:"discount,omitempty" codec:"discount,omitempty"`
	Quota         float64    `json:"quota,omitempty" codec:"quota,omitempty"`
	CouponType    CouponType `json:"couponType,omitempty" codec:"couponType,omitempty"`
	Name          string     `json:"name,omitempty" codec:"name,omitempty"`
	LimitType     uint       `json:"limitType,omitempty" codec:"limitType,omitempty"`
	AddDays       uint       `json:"addDays,omitempty" codec:"addDays,omitempty"`
	BatchCount    uint       `json:"batchCount,omitempty" codec:"batchCount,omitempty"`
	NowCount      uint       `json:"nowCount,omitempty" codec:"nowCount,omitempty"`
	Url           string     `json:"url,omitempty" codec:"url,omitempty"`
	MUrl          string     `json:"mUrl,omitempty" codec:"mUrl,omitempty"`
}

// 输入单个订单id，得到所有相关订单信息
func FindJoinActivities(req *FindJoinActivitiesRequest) ([]Coupon, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := item.NewFindJoinActivitiesRequest()
	r.SetUid(req.Uid)
	r.SetSku(req.Sku)

	var response FindJoinActivitiesResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	var coupons []Coupon
	for _, i := range response.Data.Result.Coupons {
		coupons = append(coupons, i.Coupon)
	}
	return coupons, nil
}
