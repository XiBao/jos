package points

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
	"github.com/daviddengcn/ljson"
)

type GetPointsExchangeGiftListRequest struct {
	api.BaseRequest
}

type GetPointsExchangeGiftListResponse struct {
	ErrorResp *api.ErrorResponnse            `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetPointsExchangeGiftListData `json:"jingdong_points_jos_getPointsExchangeGiftList_responce,omitempty" codec:"jingdong_points_jos_getPointsExchangeGiftList_responce,omitempty"`
}

type GetPointsExchangeGiftListData struct {
	JsfResult *GetPointsExchangeGiftListJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

type GetPointsExchangeGiftListJsfResult struct {
	Code   string                   `json:"code,omitempty" codec:"code,omitempty"`     //返回码
	Desc   string                   `json:"desc,omitempty" codec:"desc,omitempty"`     //返回描述
	Result []*PointsExchangeGiftDTO `json:"result,omitempty" codec:"result,omitempty"` //活动列表
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

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}
	var response GetPointsExchangeGiftListResponse
	err = ljson.Unmarshal(result, &response)
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
