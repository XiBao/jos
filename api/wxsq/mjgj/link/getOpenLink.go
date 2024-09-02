package link

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/wxsq/mjgj/link"
)

type GetOpenLinkRequest struct {
	api.BaseRequest
	Jump uint8  `json:"jump" codec:"jump"`
	RUrl string `json:"rurl" codec:"rurl"`
}

type GetOpenLinkResponse struct {
	ErrorResp *api.ErrorResponnse `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Data      *GetOpenLinkData    `json:"jingdong_new_ware_mobilebigfield_get_responce,omitempty" codec:"jingdong_new_ware_mobilebigfield_get_responce,omitempty"`
}

func (r GetOpenLinkResponse) IsError() bool {
	return r.ErrorResp != nil || r.Data == nil || r.Data.IsError()
}

func (r GetOpenLinkResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Data != nil {
		return r.Data.Error()
	}
	return "no result data"
}

type GetOpenLinkData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    string `json:"result,omitempty" codec:"result,omitempty"`
}

func (r GetOpenLinkData) IsError() bool {
	return r.Code != "0" || r.Result == ""
}

func (r GetOpenLinkData) Error() string {
	return sdk.ErrorString(r.Code, r.ErrorDesc)
}

func GetOpenLink(ctx context.Context, req *GetOpenLinkRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := link.NewGetOpenLinkRequest()
	r.SetJump(req.Jump)
	r.SetRUrl(req.RUrl)

	var response GetOpenLinkResponse
	if err := client.Execute(ctx, r.Request, req.Session, &response); err != nil {
		return "", err
	}
	return response.Data.Result, nil
}
