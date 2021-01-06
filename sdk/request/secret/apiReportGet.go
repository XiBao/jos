package secret

import (
	"github.com/XiBao/jos/sdk"
)

type SecretApiReportGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewSecretApiReportGetRequest() (req *SecretApiReportGetRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.secret.api.report.get", Params: make(map[string]interface{}, 6)}
	req = &SecretApiReportGetRequest{
		Request: &request,
	}
	return
}

func (req *SecretApiReportGetRequest) SetAccessToken(accessToken string) {
	req.Request.Params["access_token"] = accessToken
}

func (req *SecretApiReportGetRequest) GetAccessToken() string {
	accessToken, found := req.Request.Params["access_token"]
	if found {
		return accessToken.(string)
	}
	return ""
}

func (req *SecretApiReportGetRequest) SetBusinessId(businessId string) {
	req.Request.Params["businessId"] = businessId
}

func (req *SecretApiReportGetRequest) GetBusinessId() string {
	businessId, found := req.Request.Params["businessId"]
	if found {
		return businessId.(string)
	}
	return ""
}

func (req *SecretApiReportGetRequest) SetText(text string) {
	req.Request.Params["text"] = text
}

func (req *SecretApiReportGetRequest) GetText() string {
	text, found := req.Request.Params["text"]
	if found {
		return text.(string)
	}
	return ""
}

func (req *SecretApiReportGetRequest) SetAttribute(attribute string) {
	req.Request.Params["attribute"] = attribute
}

func (req *SecretApiReportGetRequest) GetAttribute() string {
	attribute, found := req.Request.Params["attribute"]
	if found {
		return attribute.(string)
	}
	return ""
}

func (req *SecretApiReportGetRequest) SetServerUrl(serverUrl string) {
	req.Request.Params["server_url"] = serverUrl
}

func (req *SecretApiReportGetRequest) GetServerUrl() string {
	serverUrl, found := req.Request.Params["server_url"]
	if found {
		return serverUrl.(string)
	}
	return ""
}

func (req *SecretApiReportGetRequest) SetCustomerUserId(customerUserId string) {
	req.Request.Params["customer_user_id"] = customerUserId
}

func (req *SecretApiReportGetRequest) GetCustomerUserId() string {
	customerUserId, found := req.Request.Params["customer_user_id"]
	if found {
		return customerUserId.(string)
	}
	return ""
}
