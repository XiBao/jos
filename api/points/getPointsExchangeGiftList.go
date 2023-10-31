package points

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
)

type GetPointsExchangeGiftListRequest struct {
	api.BaseRequest
}

type GetPointsExchangeGiftListResponse struct {
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetPointsExchangeGiftListData `json:"jingdong_points_jos_getPointsExchangeGiftList_responce,omitempty" codec:"jingdong_points_jos_getPointsExchangeGiftList_responce,omitempty"`
}

func (r GetPointsExchangeGiftListResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetPointsExchangeGiftListResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetPointsExchangeGiftListData struct {
	JsfResult *GetPointsExchangeGiftListJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

func (r GetPointsExchangeGiftListData) IsError() bool {
	return r.JsfResult == nil || r.JsfResult.IsError()
}

func (r GetPointsExchangeGiftListData) Error() string {
	if r.JsfResult == nil {
		return r.JsfResult.Error()
	}
	return "no result data"
}

type GetPointsExchangeGiftListJsfResult struct {
	Code   string                   `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	Desc   string                   `json:"desc,omitempty" codec:"desc,omitempty"`     //返回描述
	Result []*PointsExchangeGiftDTO `json:"result,omitempty" codec:"result,omitempty"` //活动列表
}

func (r GetPointsExchangeGiftListJsfResult) IsError() bool {
	return r.Code != "200"
}

func (r GetPointsExchangeGiftListJsfResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Desc)
}

type PointsExchangeGiftDTO struct {
	Id                   uint64 `json:"id,omitempty" codec:"id,omitempty"`                                     //活动id
	VenderId             uint64 `json:"venderId,omitempty" codec:"venderId,omitempty"`                         //商家ID
	ActivityName         string `json:"activityName,omitempty" codec:"activityName,omitempty"`                 //活动名称
	ActivityStartTime    int64  `json:"activityStartTime,omitempty" codec:"activityStartTime,omitempty"`       //活动开始时间
	ActivityEndTime      int64  `json:"activityEndTime,omitempty" codec:"activityEndTime,omitempty"`           //活动结束时间
	ActivityStatus       uint16 `json:"activityStatus,omitempty" codec:"activityStatus,omitempty"`             //活动状态
	CreateTime           int64  `json:"createTime,omitempty" codec:"createTime,omitempty"`                     //创建时间
	UpdateTime           int64  `json:"updateTime,omitempty" codec:"updateTime,omitempty"`                     //更新时间
	ActivityStatusString string `json:"activityStatusString,omitempty" codec:"activityStatusString,omitempty"` //活动页面显示状态
}

// TODO 查询正在进行中的积分兑换商品活动列表
func GetPointsExchangeGiftList(req *GetPointsExchangeGiftListRequest) ([]*PointsExchangeGiftDTO, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewGetPointsExchangeGiftListRequest()

	var response GetPointsExchangeGiftListResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.JsfResult.Result, nil

}
