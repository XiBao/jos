package voucher

import (
	"github.com/XiBao/jos/sdk"
)

type VoucherInfoGetRequest struct {
	Request *sdk.Request
}

// create new request
func NewVoucherInfoGet() (req *VoucherInfoGetRequest) {
	request := sdk.Request{MethodName: "jingdong.jos.voucher.info.get", Params: make(map[string]interface{}, 2)}
	req = &VoucherInfoGetRequest{
		Request: &request,
	}
	return
}

func (req *VoucherInfoGetRequest) SetAccessToken(accessToken string) {
	req.Request.Params["access_token"] = accessToken
}

func (req *VoucherInfoGetRequest) GetAccessToken() string {
	accessToken, found := req.Request.Params["access_token"]
	if found {
		return accessToken.(string)
	}
	return ""
}

func (req *VoucherInfoGetRequest) SetCustomerUserId(customerUserId string) {
	req.Request.Params["customer_user_id"] = customerUserId
}

func (req *VoucherInfoGetRequest) GetCustomerUserId() string {
	customerUserId, found := req.Request.Params["customer_user_id"]
	if found {
		return customerUserId.(string)
	}
	return ""
}
