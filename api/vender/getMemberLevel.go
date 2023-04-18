package vender

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/vender"
)

type GetMemberLevelRequest struct {
	api.BaseRequest
	CustomerPin string
}

type GetMemberLevelResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetMemberLevelData `json:"jingdong_pop_vender_getMemberLevel_responce,omitempty" codec:"jingdong_pop_vender_getMemberLevel_responce,omitempty"`
}

type GetMemberLevelData struct {
	ReturnType *MemberLevelReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

type MemberLevelReturnType struct {
	Desc string           `json:"desc,omitempty" codec:"desc,omitempty"`
	Code string           `json:"code,omitempty" codec:"code,omitempty"`
	Info *MemberLevelInfo `json:"memberLevelInfo,omitempty" codec:"memberLevelInfo,omitempty"`
}

type MemberLevelInfo struct {
	LevelAtShop        uint8   `json:"levelAtShop,omitempty" codec:"levelAtShop,omitempty"`               //等级
	AvgOrderPrice      float64 `json:"avgOrderPrice,omitempty" codec:"avgOrderPrice,omitempty"`           //平均客单价
	RefundOrderCount   uint64  `json:"refundOrderCount,omitempty" codec:"refundOrderCount,omitempty"`     //退单次数
	TotalGoodsCount    uint64  `json:"totalGoodsCount,omitempty" codec:"totalGoodsCount,omitempty"`       //商品数量
	ChangedOrderCount  uint64  `json:"changedOrderCount,omitempty" codec:"changedOrderCount,omitempty"`   //换货次数
	CustomerPin        string  `json:"customerPin,omitempty" codec:"customerPin,omitempty"`               //客户pin
	CanceledOrderCount uint64  `json:"canceledOrderCount,omitempty" codec:"canceledOrderCount,omitempty"` //取消订单数
	RefundOrderPrice   float64 `json:"refundOrderPrice,omitempty" codec:"refundOrderPrice,omitempty"`     //退换货金额
	VenderId           uint64  `json:"venderId,omitempty" codec:"venderId,omitempty"`                     //商家Id
	LevelAtShopName    string  `json:"levelAtShopName,omitempty" codec:"levelAtShopName,omitempty"`       //等级名称
	TotalOrderPrice    float64 `json:"totalOrderPrice,omitempty" codec:"totalOrderPrice,omitempty"`       //订单总价格
	TotalOrderCount    uint64  `json:"totalOrderCount,omitempty" codec:"totalOrderCount,omitempty"`       //订单总数量
	NickName           string  `json:"nickName,omitempty" codec:"nickName,omitempty"`                     //用户昵称
	BindingTime        string  `json:"bindingTime,omitempty" codec:"bindingTime,omitempty"`               //绑定时间
	BindingType        uint8   `json:"bindingType,omitempty" codec:"bindingType,omitempty"`               //绑定类型
}

// TODO 查询会员等级及会员信息  交易数据 T+1 更新
func GetMemberLevel(req *GetMemberLevelRequest) (*MemberLevelInfo, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := vender.NewVenderGetMemberLevelRequest()

	if len(req.CustomerPin) > 0 {
		r.SetCustomerPin(req.CustomerPin)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no result info")
	}
	var response GetMemberLevelResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.ReturnType.Code != "200" {
		return nil, errors.New(response.Data.ReturnType.Desc)
	}

	return response.Data.ReturnType.Info, nil
}
