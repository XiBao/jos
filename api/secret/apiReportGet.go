package secret

import (
	"errors"
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/secret"
	"github.com/daviddengcn/ljson"
)

type SecretApiReportGetRequest struct {
	api.BaseRequest
	AccessToken    string `json:"access_token,omitempty" codec:"access_token,omitempty"`
	CustomerUserId uint64 `json:"customer_user_id,omitempty" codec:"customer_user_id,omitempty"`
	BusinessId     string `json:"businessId,omitempty" codec:"businessId,omitempty"`
	Text           string `json:"text,omitempty" codec:"text,omitempty"`
	Attribute      string `json:"attribute,omitempty" codec:"attribute,omitempty"`
	ServerUrl      string `json:"server_url,omitempty" codec:"server_url,omitempty"`
}

type SecretApiReportGetResponse struct {
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *SecretApiReportResponse `json:"jingdong_jos_secret_api_report_get_responce,omitempty" codec:"jingdong_jos_secret_api_report_get_responce,omitempty"`
}

type SecretApiReportResponse struct {
	Result SecretApiReportResult `json:"response,omitempty" codec:"response,omitempty"`
}

type SecretApiReportResult struct {
	Code      int    `json:"errorCode,omitempty" codec:"errorCode,omitempty"`
	ErrorDesc string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

// 对加解密等调用信息上报，不包含敏感信息，只负责统计系统性能
func SecretApiReportGet(req *SecretApiReportGetRequest) error {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := secret.NewSecretApiReportGetRequest()
	r.SetAccessToken(req.AccessToken)
	r.SetCustomerUserId(req.CustomerUserId)
	r.SetText(req.Text)
	r.SetAttribute(req.Attribute)
	r.SetServerUrl(req.ServerUrl)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("No result.")
	}
	var response SecretApiReportGetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return err
	}
	fmt.Println(string(result))
	if response.ErrorResp != nil {
		return response.ErrorResp
	}
	if response.Response.Result.Code != 0 {
		return errors.New(response.Response.Result.ErrorDesc)
	}
	return nil
}
