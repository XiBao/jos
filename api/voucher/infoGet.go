package voucher

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/request/voucher"
	"github.com/daviddengcn/ljson"
)

type VoucherInfoGetRequest struct {
	api.BaseRequest
	AccessToken    string `json:"access_token,omitempty" codec:"access_token,omitempty"`
	CustomerUserId string `json:"customer_user_id,omitempty" codec:"customer_user_id,omitempty"`
}

type VoucherInfoGetResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *VoucherInfoResponse `json:"jingdong_jos_voucher_info_get_responce,omitempty" codec:"jingdong_jos_voucher_info_get_responce,omitempty"`
}

type VoucherInfoResponse struct {
	Result *VoucherInfoResult `json:"response,omitempty" codec:"response,omitempty"`
}

type VoucherInfoResult struct {
	Code      string          `json:"errorCode,omitempty" codec:"errorCode,omitempty"`
	ErrorDesc string          `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Data      VoucherInfoData `json:"data,omitempty" codec:"data,omitempty"`
}

type VoucherInfoData struct {
	Voucher string `json:"voucher,omitempty" codec:"voucher,omitempty"`
}

// 凭证获取
func VoucherInfoGet(req *VoucherInfoGetRequest) (string, error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := voucher.NewVoucherInfoGet()
	r.SetAccessToken(req.AccessToken)
	r.SetCustomerUserId(req.CustomerUserId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", errors.New("No result.")
	}
	var response VoucherInfoGetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	fmt.Println(string(result))
	if response.ErrorResp != nil {
		return "", response.ErrorResp
	}
	if response.Response.Result == nil {
		return "", errors.New("No result.")
	}
	if response.Response.Result.Code != "0" {
		return "", errors.New(response.Response.Result.ErrorDesc)
	}

	voucher, err := base64.URLEncoding.DecodeString(response.Response.Result.Data.Voucher)
	return string(voucher), err
}
