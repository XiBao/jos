package jm

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/jm"
)

// 风险信息查询服务，需要授权才能使用。
type RiskQueryServiceRequest struct {
	api.BaseRequest
	Param1 *jm.RiskQueryServiceRequestParam1 `json:"param1,omitempty" codec:"param1,omitempty"` // 业务参数
}

type RiskQueryServiceResponse struct {
	Responce  *RiskQueryServiceResponce `json:"jingdong_risk_RiskQueryService_responce,omitempty" codec:"jingdong_risk_RiskQueryService_responce,omitempty"`
	ErrorResp *api.ErrorResponnse       `json:"error_response,omitempty" codec:"error_response,omitempty"`
}

func (r RiskQueryServiceResponse) IsError() bool {
	return r.ErrorResp != nil || r.Responce == nil || r.Responce.IsError()
}

func (r RiskQueryServiceResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Responce != nil {
		return r.Responce.Error()
	}
	return "no result data"
}

type RiskQueryServiceResponce struct {
	ReturnType *RiskQueryServiceResponseReturnType `json:"returnType,omitempty" codec:"returnType,omitempty"`
}

func (r RiskQueryServiceResponce) IsError() bool {
	return r.ReturnType != nil && !r.ReturnType.Success
}

func (r RiskQueryServiceResponce) Error() string {
	if r.ReturnType != nil {
		return r.ReturnType.Msg
	}
	return "no result data"
}

type RiskQueryServiceResponseReturnType struct {
	Msg     string                                    `json:"msg,omitempty" codec:"msg,omitempty"`         // 状态码描述
	Code    string                                    `json:"code,omitempty" codec:"code,omitempty"`       // 状态码
	Success bool                                      `json:"success,omitempty" codec:"success,omitempty"` // 请求是否成功
	Result  *RiskQueryServiceResponseReturnTypeResult `json:"result,omitempty" codec:"result,omitempty"`   // 风控识别结果
	Output  *RiskQueryServiceResponseReturnTypeOutput `json:"output,omitempty" codec:"output,omitempty"`   // 扩展字段
}

type RiskQueryServiceResponseReturnTypeResult struct {
	Strategy string `json:"strategy"` // 风控识别结果命中标识 risk代表有风险，white代表白名单， none代表无风险
}

type RiskQueryServiceResponseReturnTypeOutput struct {
	Addition string `json:"addition"` // 扩展信息
}

func RiskQueryService(ctx context.Context, req *RiskQueryServiceRequest) (*RiskQueryServiceResponseReturnTypeResult, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := jm.NewRiskQueryServiceRequest()
	r.SetParam1(req.Param1)

	var response RiskQueryServiceResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return nil, err
	}
	return response.Responce.ReturnType.Result, nil
}
