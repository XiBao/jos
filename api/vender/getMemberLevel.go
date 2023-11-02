package vender

import (
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

func (r GetMemberLevelResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetMemberLevelResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetMemberLevelData struct {
	ReturnType *MemberLevelReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r GetMemberLevelData) IsError() bool {
	return r.ReturnType == nil || r.ReturnType.IsError()
}

func (r GetMemberLevelData) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Error()
	}
	return "no result data"
}

type MemberLevelReturnType struct {
	Desc string           `json:"desc,omitempty" codec:"desc,omitempty"`
	Code string           `json:"code,omitempty" codec:"code,omitempty"`
	Info *MemberLevelInfo `json:"memberLevelInfo,omitempty" codec:"memberLevelInfo,omitempty"`
}

func (r MemberLevelReturnType) IsError() bool {
	return r.Code != "200" || r.Info == nil
}

func (r MemberLevelReturnType) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
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

	var response GetMemberLevelResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Data.ReturnType.Info, nil
}
