package points

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/points"
)

type JosSendPointsRequest struct {
	api.BaseRequest
	Pin        string `json:"pin,omitempty" codec:"pin,omitempty"`               //用户Pin
	BusinessId string `json:"businessId,omitempty" codec:"businessId,omitempty"` //防重ID
	SourceType uint8  `json:"sourceType,omitempty" codec:"sourceType,omitempty"` //渠道类型：26-消费积分 27-发放积分
	Points     int64  `json:"points,omitempty" codec:"points,omitempty"`         //积分变更值
}

type JosSendPointsResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *JosSendPointsData  `json:"jingdong_points_jos_sendPoints_responce,omitempty" codec:"jingdong_points_jos_sendPoints_responce,omitempty"`
}

type JosSendPointsData struct {
	JsfResult *JosSendPointsJsfResult `json:"jsfResult,omitempty" codec:"jsfResult,omitempty"`
}

type JosSendPointsJsfResult struct {
	Code string `json:"code,omitempty" codec:"code,omitempty"` //返回码
	Desc string `json:"desc,omitempty" codec:"desc,omitempty"` //返回描述
}

// TODO 积分变更开放接口 开放请求的渠道为：26-消费积分 27-发放积分
func JosSendPoints(req *JosSendPointsRequest) (bool, error) {
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
		req.BusinessId = fmt.Sprintf("jdvip%s", time.Now().Format("20060102150405"))
	}

	r.SetBusinessId(req.BusinessId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return false, err
	}
	if len(result) == 0 {
		return false, errors.New("no result info")
	}
	var response JosSendPointsResponse
	err = json.Unmarshal(result, &response)
	if err != nil {
		return false, err
	}

	if response.ErrorResp != nil {
		return false, response.ErrorResp
	}

	if response.Data.JsfResult.Code != "200" {
		return false, errors.New(response.Data.JsfResult.Desc)
	}

	return true, nil

}
