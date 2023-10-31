package voucher

import (
	"encoding/base64"
	"fmt"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/sdk"
	"github.com/XiBao/jos/sdk/crypto"
	"github.com/XiBao/jos/sdk/request/voucher"
)

type VoucherInfoGetRequest struct {
	api.BaseRequest
	CustomerUserId string `json:"customer_user_id,omitempty" codec:"customer_user_id,omitempty"`
}

type VoucherInfoGetResponse struct {
	ErrorResp *api.ErrorResponnse  `json:"error_response,omitempty" codec:"error_response,omitempty"`
	Response  *VoucherInfoResponse `json:"jingdong_jos_voucher_info_get_responce,omitempty" codec:"jingdong_jos_voucher_info_get_responce,omitempty"`
}

func (r VoucherInfoGetResponse) IsError() bool {
	return r.ErrorResp != nil || r.Response == nil || r.Response.IsError()
}

func (r VoucherInfoGetResponse) Error() string {
	if r.ErrorResp != nil {
		return r.ErrorResp.Error()
	}
	if r.Response != nil {
		return r.Response.Error()
	}
	return "no result data"
}

type VoucherInfoResponse struct {
	Result *VoucherInfoResult `json:"response,omitempty" codec:"response,omitempty"`
}

func (r VoucherInfoResponse) IsError() bool {
	return r.Result == nil || r.Result.IsError()
}

func (r VoucherInfoResponse) Error() string {
	if r.Result != nil {
		return r.Result.Error()
	}
	return "no result data"
}

type VoucherInfoResult struct {
	Code      string          `json:"errorCode,omitempty" codec:"errorCode,omitempty"`
	ErrorDesc string          `json:"errorMsg,omitempty" codec:"errorMsg,omitempty"`
	Data      VoucherInfoData `json:"data,omitempty" codec:"data,omitempty"`
}

func (r VoucherInfoResult) IsError() bool {
	return r.Code != "0"
}

func (r VoucherInfoResult) Error() string {
	return fmt.Sprintf("code: %s, msg: %s", r.Code, r.ErrorDesc)
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
	Expired   int64  `json:"expired"`   //过期时间戳，客户端需要检查凭证是否已过期，已过期的凭证无法获取密钥
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
	r.SetCustomerUserId(req.CustomerUserId)

	var response VoucherInfoGetResponse
	if err = client.Execute(r.Request, req.Session, &response); err != nil {
		return
	}

	voucherBytes, err := base64.URLEncoding.DecodeString(response.Response.Result.Data.Voucher)
	var voucherInfo Voucher
	if err = client.Logger().DecodeJSON(voucherBytes, &voucherInfo); err != nil {
		return
	}
	if err = voucherInfo.Verify(); err != nil {
		return
	}
	return voucherInfo.Data, nil
}
