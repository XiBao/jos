package voucher

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/crypto"
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

type Voucher struct {
	Sig  string      `json:"sig" codec:"sig"`
	Data VoucherData `json:"data" codec:"data"`
}

type VoucherData struct {
	Id        string `json:"id"`        // 凭证id
	Key       string `json:"key"`       // 该凭证对应的密钥，请求密钥时，需要使用该key对业务入参签名，用于获取加密密钥
	Service   string `json:"service"`   //服务识别码
	Act       string `json:"act"`       //ignore
	Effective int64  `json:"effective"` //生效时间戳，客户端需要检查凭证是否已生效，未生效的凭证无法获取密钥
	Expired   int64  `"expired"`        //过期时间戳，客户端需要检查凭证是否已过期，已过期的凭证无法获取密钥
	SType     int    `json:"stype"`     //ignore
}

func (this *Voucher) Verify() error {
	js := []byte(fmt.Sprintf(`{"act":"%s","effective":%d,"expired":%d,"id":"%s","key":"%s","service":"%s","stype":%d}`, this.Data.Act, this.Data.Effective, this.Data.Expired, this.Data.Id, this.Data.Key, this.Data.Service, this.Data.SType))
	return crypto.RSAVerifySignWithSha256([]byte(crypto.PublicKey), js, this.Sig)
}

// 凭证获取
func VoucherInfoGet(req *VoucherInfoGetRequest) (voucherData VoucherData, err error) {
	client := sdk.NewClient(req.AnApiKey.Key, req.AnApiKey.Secret)
	client.Debug = req.Debug
	r := voucher.NewVoucherInfoGet()
	r.SetAccessToken(req.AccessToken)
	r.SetCustomerUserId(req.CustomerUserId)

	result, err := client.Execute(r.Request, req.Session)
	if err != nil {
		return
	}
	if len(result) == 0 {
		err = errors.New("No result.")
		return
	}
	var response VoucherInfoGetResponse
	err = ljson.Unmarshal(result, &response)
	if err != nil {
		return
	}
	if response.ErrorResp != nil {
		err = response.ErrorResp
		return
	}
	if response.Response.Result == nil {
		err = errors.New("No result.")
		return
	}
	if response.Response.Result.Code != "0" {
		err = errors.New(response.Response.Result.ErrorDesc)
		return
	}

	voucherBytes, err := base64.URLEncoding.DecodeString(response.Response.Result.Data.Voucher)
	var voucherInfo Voucher
	err = json.Unmarshal(voucherBytes, &voucherInfo)
	if err != nil {
		return
	}
	err = voucherInfo.Verify()
	if err != nil {
		return
	}
	return voucherInfo.Data, nil
}
