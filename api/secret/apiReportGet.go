package secret

import (
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/secret"
)

type SecretApiReportGetRequest struct {
	api.BaseRequest
	AccessToken    string `json:"access_token,omitempty" codec:"access_token,omitempty"`
	CustomerUserId string `json:"customer_user_id,omitempty" codec:"customer_user_id,omitempty"`
	BusinessId     string `json:"businessId,omitempty" codec:"businessId,omitempty"`
	Text           string `json:"text,omitempty" codec:"text,omitempty"`
	Attribute      string `json:"attribute,omitempty" codec:"attribute,omitempty"`
	ServerUrl      string `json:"server_url,omitempty" codec:"server_url,omitempty"`
}

type SecretApiReportGetResponse struct {
	ErrorResp *api.ErrorResponnse      `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *SecretApiReportResponse `json:"jingdong_jos_secret_api_report_get_responce,omitempty" codec:"jingdong_jos_secret_api_report_get_responce,omitempty"`
}

func (r SecretApiReportGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response == nil || r.Response.IsError()
}

func (r SecretApiReportGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Response != nil {
		return r.Response.Error()
	}
	return "no result data"
}

type SecretApiReportResponse struct {
	Result SecretApiReportResult `json:"response,omitempty" codec:"response,omitempty"`
}

func (r SecretApiReportResponse) IsError() bool {
	return r.Result.IsError()
}

func (r SecretApiReportResponse) Error() string {
	return r.Result.Error()
}

type SecretApiReportResult struct {
	Code      int    `json:"errorCode,omitempty" codec:"errorCode,omitempty"`
	ErrorDesc string `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
}

func (r SecretApiReportResult) IsError() bool {
	return r.Code != 0
}

func (r SecretApiReportResult) Error() string {
	return fmt.Sprintf("code:%d, msg:%s", r.Code, r.ErrorDesc)
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

	var response SecretApiReportGetResponse
	if err := client.Execute(r.Request, req.Session, &response); err != nil {
		return err
	}
	return nil
}
