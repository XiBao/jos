package fw

import (
	"encoding/json"
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/fw"
)

type MarketPaymentoutRequest struct {
	api.BaseRequest
	RequestNo      string `json:"requestNo,omitempty" codec:"requestNo,omitempty"`           // 请求唯一ID
	ActivityId     uint   `json:"activityId,omitempty" codec:"activityId,omitempty"`         // 活动ID
	AppId          string `json:"appId,omitempty" codec:"appId,omitempty"`                   // 开普勒小程序ID（仅限开普勒业务）
	Price          uint   `json:"price,omitempty" codec:"price,omitempty"`                   // 服务价格
	IsMainService  bool   `json:"isMainService,omitempty" codec:"isMainService,omitempty"`   // 是否主服务
	ServiceCycle   uint   `json:"serviceCycle,omitempty" codec:"serviceCycle,omitempty"`     // 服务周期
	SkuId          uint64 `json:"skuId,omitempty" codec:"skuId,omitempty"`                   // 订购skuId
	ServiceCode    string `json:"serviceCode,omitempty" codec:"serviceCode,omitempty"`       // 订购服务编码
	OrderNum       uint   `json:"orderNum,omitempty" codec:"orderNum,omitempty"`             // 订购数量
	ItemCode       string `json:"itemCode,omitempty" codec:"itemCode,omitempty"`             // 订购项目id
	OutOrderId     uint   `json:"outOrderId,omitempty" codec:"outOrderId,omitempty"`         // 外部系统订单号
	Value1         string `json:"value1,omitempty" codec:"value1,omitempty"`                 // 扩展属性
	ResultPageType uint   `json:"resultPageType,omitempty" codec:"resultPageType,omitempty"` // 结算页类型【1：PC；2：移动版H5；3：移动版小程序（预留）】
	SuccessUrl     string `json:"successUrl,omitempty" codec:"successUrl,omitempty"`         // 支付成功跳转地址
	Ip             string `json:"ip,omitempty" codec:"ip,omitempty"`                         // ip
}

type MarketPaymentoutResponse struct {
	ErrorResp *api.ErrorResponnse   `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *MarketPaymentoutData `json:"jingdong_fw_market_paymentout_responce,omitempty" codec:"jingdong_fw_market_paymentout_responce,omitempty"`
}

type MarketPaymentoutData struct {
	ReturnType *MarketPaymentoutReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
	Code       string                      `json:"code"`
}

type MarketPaymentoutReturnType struct {
	ErrorCode uint64                          `json:"errorCode,omitempty" codec:"errorCode,omitempty"` //错误码
	Success   bool                            `json:"success,omitempty" codec:"success,omitempty"`     //是否成功
	ErrorMsg  string                          `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`   // 错误信息
	Body      *MarketPaymentoutReturnTypeBody `json:"body,omitempty" codec:"body,omitempty"`           //订单列表
}

type MarketPaymentoutReturnTypeBody struct {
	RequestId      string `json:"requestId,omitempty" codec:"requestId,omitempty"`
	SettlementUrl  string `json:"settlementUrl,omitempty" codec:"settlementUrl,omitempty"`
	OrderId        uint64 `json:"orderId,omitempty" codec:"orderId,omitempty"`
	ResultPageType uint   `json:"resultPageType,omitempty" codec:"resultPageType,omitempty"`
}

func MarketPaymentout(req *MarketPaymentoutRequest) (*MarketPaymentoutReturnType, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := fw.NewMarketPaymentoutRequest()
	r.SetRequestNo(req.RequestNo)
	r.SetSkuId(req.SkuId)
	r.SetServiceCode(req.ServiceCode)
	r.SetOrderNum(req.OrderNum)
	r.SetResultPageType(req.ResultPageType)
	r.SetIsMainService(req.IsMainService)
	r.SetIp(req.Ip)

	if req.ActivityId > 0 {
		r.SetActivityId(req.ActivityId)
	}

	if req.AppId != "" {
		r.SetAppId(req.AppId)
	}

	if req.Price > 0 {
		r.SetPrice(req.Price)
	}

	if req.ServiceCycle > 0 {
		r.SetServiceCycle(req.ServiceCycle)
	}

	if req.ItemCode != "" {
		r.SetItemCode(req.ItemCode)
	}

	if req.OutOrderId > 0 {
		r.SetOutOrderId(req.OutOrderId)
	}

	if req.Value1 != "" {
		r.SetValue1(req.Value1)
	}

	if req.SuccessUrl != "" {
		r.SetSuccessUrl(req.SuccessUrl)
	}

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("No result info.")
	}
	var response MarketPaymentoutResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return nil, err
	}

	if response.ErrorResp != nil {
		return nil, response.ErrorResp
	}

	if response.Data.ReturnType.ErrorCode != 0 {
		return nil, errors.New(response.Data.ReturnType.ErrorMsg)
	}
	return response.Data.ReturnType, nil
}
