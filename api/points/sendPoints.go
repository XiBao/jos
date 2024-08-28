package points

import (
	"context"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
)

type JosSendPointsRequest struct {
	api.BaseRequest
	Pin        string `json:"pin,omitempty" codec:"pin,omitempty"`               // 用户Pin
	BusinessId string `json:"businessId,omitempty" codec:"businessId,omitempty"` // 防重ID
	SourceType uint8  `json:"sourceType,omitempty" codec:"sourceType,omitempty"` // 渠道类型：26-消费积分 27-发放积分
	Points     int64  `json:"points,omitempty" codec:"points,omitempty"`         // 积分变更值
}

type JosSendPointsResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *JosSendPointsData  `json:"jingdong_points_jos_sendPoints_responce,omitempty" codec:"jingdong_points_jos_sendPoints_responce,omitempty"`
}

func (r JosSendPointsResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r JosSendPointsResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type JosSendPointsData struct {
	JsfResult *JosSendPointsJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

func (r JosSendPointsData) IsError() bool {
	return r.JsfResult == nil || r.JsfResult.IsError()
}

func (r JosSendPointsData) Error() string {
	if r.JsfResult != nil {
		return r.JsfResult.Error()
	}
	return "no result data"
}

type JosSendPointsJsfResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"` // 返回码
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"` // 返回描述
}

func (r JosSendPointsJsfResult) IsError() bool {
	return r.Code != "200"
}

func (r JosSendPointsJsfResult) Error() string {
	return sdk.ErrorString(r.Code, r.Desc)
}

// TODO 积分变更开放接口 开放请求的渠道为：26-消费积分 27-发放积分
func JosSendPoints(ctx context.Context, req *JosSendPointsRequest) (bool, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := points.NewSendPointsRequest()

	if len(req.Pin) > 0 {
		r.SetPin(req.Pin)
	}
	if req.Points != 0 {
		r.SetPoints(req.Points)
	}
	if req.SourceType > 0 {
		r.SetSourceType(req.SourceType)
	}
	if req.BusinessId == "" {
		req.BusinessId = sdk.StringsJoin("jdvip", time.Now().Format("20060102150405"))
	}

	r.SetBusinessId(req.BusinessId)

	var response JosSendPointsResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return false, err
	}
	return true, nil
}
