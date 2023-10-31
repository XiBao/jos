package points

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
)

type QueryGiftSkuByExchangeGiftIdRequest struct {
	api.BaseRequest
	GiftId uint64 `json:"gift_id，omitempty" codec:"gift_id,omitempty"` //活动id
}

type QueryGiftSkuByExchangeGiftIdResponse struct {
	ErrorResp *api.ErrorResponnse               `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *QueryGiftSkuByExchangeGiftIdData `json:"jingdong_points_jos_queryGiftSkuByExchangeGiftId_responce,omitempty" codec:"jingdong_points_jos_queryGiftSkuByExchangeGiftId_responce,omitempty"`
}

func (r QueryGiftSkuByExchangeGiftIdResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r QueryGiftSkuByExchangeGiftIdResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type QueryGiftSkuByExchangeGiftIdData struct {
	JsfResult *QueryGiftSkuByExchangeGiftIdJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

func (r QueryGiftSkuByExchangeGiftIdData) IsError() bool {
	return r.JsfResult == nil || r.JsfResult.IsError()
}

func (r QueryGiftSkuByExchangeGiftIdData) Error() string {
	if r.JsfResult != nil {
		return r.JsfResult.Error()
	}
	return "no result data"
}

type QueryGiftSkuByExchangeGiftIdJsfResult struct {
	Code   string                      `json:"code,omitempty" codec:"code,omitempty"` //返回码
	Desc   string                      `json:"desc,omitempty" codec:"desc,omitempty"` //返回描述
	Result []*PointsExchangeGiftSkuDTO `json:"result,omitempty" codec:"result,omitempty"`
}

func (r QueryGiftSkuByExchangeGiftIdJsfResult) IsError() bool {
	return r.Code != "200"
}

func (r QueryGiftSkuByExchangeGiftIdJsfResult) Error() string {
	return fmt.Sprintf("code:%s, msg:%s", r.Code, r.Desc)
}

type PointsExchangeGiftSkuDTO struct {
	Id                   uint64  `json:"id,omitempty" codec:"id,omitempty"`                                     //主键
	ActivityStartTime    int64   `json:"activityStartTime,omitempty" codec:"activityStartTime,omitempty"`       //活动开始时间
	ActivityEndTime      int64   `json:"activityEndTime,omitempty" codec:"activityEndTime,omitempty"`           //活动结束时间
	ActivityPrice        float64 `json:"activityPrice,omitempty" codec:"activityPrice,omitempty"`               //活动价
	ActivityStatus       uint8   `json:"activityStatus,omitempty" codec:"activityStatus,omitempty"`             //活动状态：1-创建中、2-创建成功、3-失效中、 4-失效
	CreateTime           int64   `json:"createTime,omitempty" codec:"createTime,omitempty"`                     //创建时间
	HasExchange          bool    `json:"hasExchange,omitempty" codec:"hasExchange,omitempty"`                   //是否已兑换
	HasConsume           bool    `json:"hasConsume,omitempty" codec:"hasConsume,omitempty"`                     //是否已消费
	Points               uint64  `json:"points,omitempty" codec:"points,omitempty"`                             //兑换商品所需积分
	PointsExchangeGiftId uint64  `json:"pointsExchangeGiftId,omitempty" codec:"pointsExchangeGiftId,omitempty"` //积分兑换商品活动表主键
	PromoId              uint64  `json:"promoId,omitempty" codec:"promoId,omitempty"`                           // pop促销ID
	PromoStatus          uint8   `json:"promoStatus,omitempty" codec:"promoStatus,omitempty"`                   // 促销状态：1-创建中 2-待审核 3-通过 4-驳回 6-已删除 110-售完
	SkuId                uint64  `json:"skuId,omitempty" codec:"skuId,omitempty"`                               // SKU编码
	UpdateTime           int64   `json:"updateTime,omitempty" codec:"updateTime,omitempty"`                     // 更新时间
	VenderId             uint64  `json:"venderId,omitempty" codec:"venderId,omitempty"`                         // 商家id
	WareId               uint64  `json:"wareId,omitempty" codec:"wareId,omitempty"`                             // 	商品编号
}

// TODO  根据积分兑换商品活动id 获取商品信息
func QueryGiftSkuByExchangeGiftId(req *QueryGiftSkuByExchangeGiftIdRequest) ([]*PointsExchangeGiftSkuDTO, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewQueryGiftSkuByExchangeGiftIdRequest()

	if req.GiftId > 0 {
		r.SetGiftId(req.GiftId)
	}

	var response QueryGiftSkuByExchangeGiftIdResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.JsfResult.Result, nil

}
