package link

import (
	"errors"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/wxsq/mjqj/link"
	"github.com/daviddengcn/ljson"
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

type GetOpenLinkData struct {
	Code      string `json:"code,omitempty" codec:"code,omitempty"`
	ErrorDesc string `json:"error_description,omitempty" codec:"error_description,omitempty"`
	Result    string `json:"result,omitempty" codec:"result,omitempty"`
}

func GetOpenLink(req *GetOpenLinkRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := link.NewGetOpenLinkRequest()
	r.SetJump(req.Jump)
	r.SetRUrl(req.RUrl)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	result = util.RemoveJsonSpace(result)

	var response GetOpenLinkResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	if response.ErrorResp != nil {
		return "", response.ErrorResp
	}
	if response.Data.Code != "0" {
		return "", errors.New(response.Data.ErrorDesc)
	}
	if response.Data.Result == "" {
		return "", errors.New("No open link info.")
	}

	return response.Data.Result, nil
}
