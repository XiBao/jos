package secret

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/master"
	"github.com/XiBao/jos/api/secret"
	"github.com/XiBao/jos/api/voucher"
	"github.com/XiBao/jos/sdk/crypto"
)

type Client struct {
	AppKey      string
	AppSecret   string
	AccessToken string
	SdkVer      uint
	Host        string
	Env         string
	KeyStore    *crypto.KeyStore
	Enccnt      uint
	Deccnt      uint
	Encerrcnt   uint
	Decerrcnt   uint
	Debug       bool
	sync.RWMutex
}

func NewClient(appKey string, appSecret, accessToken string) *Client {
	client := &Client{
		AppKey:      appKey,
		AppSecret:   appSecret,
		AccessToken: accessToken,
		SdkVer:      2,
	}
	go client.StartReport()
	return client
}

func (this *Client) SetHost(host string) {
	this.Host = host
}

func (this *Client) SetEnv(env string) {
	this.Env = env
}

func (this *Client) StartReport() {
	hourlyTicker := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-hourlyTicker.C:
			this.ReportStatitics()
		}
	}
}

func (this *Client) GetVoucher() (voucherData voucher.VoucherData, err error) {
	req := &voucher.VoucherInfoGetRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
	}
	return voucher.VoucherInfoGet(req)
}

func (this *Client) GetMasterKey(tid string, voucherKey string) (keyStore *crypto.KeyStore, err error) {
	req := &master.MasterKeyGetRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
		Tid: tid,
		Key: voucherKey,
	}
	keyStore, err = master.MasterKeyGet(req)
	if err != nil {
		return nil, err
	}
	this.Lock()
	this.KeyStore = keyStore
	this.Unlock()
	return
}

func (this *Client) RefreshKeyStore() (*crypto.KeyStore, error) {
	voucher, err := this.GetVoucher()
	if err != nil {
		return nil, err
	}
	keyStore, err := this.GetMasterKey(voucher.Id, voucher.Key)
	if err != nil {
		return nil, err
	}
	this.Report(voucher.Service, INIT, INIT_TYPE, "0", "", "")
	return keyStore, nil
}

func (this *Client) GetKey(keyId string) (key crypto.Key, err error) {
	var found bool
	this.RLock()
	if this.KeyStore == nil {
		found = false
	} else {
		key, found = this.KeyStore.GetKey(keyId)
	}
	this.RUnlock()
	if found {
		return key, nil
	}
	keyStore, err := this.RefreshKeyStore()
	if err != nil {
		return key, err
	}
	key, found = keyStore.GetKey(keyId)
	if !found {
		return key, errors.New("not found key")
	}
	return key, nil
}

func (this *Client) GetCurrentKey() (key crypto.Key, err error) {
	var found bool
	this.RLock()
	if this.KeyStore == nil {
		found = false
	} else {
		key, found = this.KeyStore.GetCurrentKey()
	}
	this.RUnlock()
	if found {
		return key, nil
	}
	keyStore, err := this.RefreshKeyStore()
	if err != nil {
		return key, err
	}
	key, found = keyStore.GetCurrentKey()
	if !found {
		return key, errors.New("not found key")
	}
	return key, nil
}

func (this *Client) Decrypt(encryptedStr string, usePrivateEncrypt bool) (string, error) {
	if encryptedStr == "" {
		return "", nil
	}
	var keyData []byte
	ivStart := aes.BlockSize + CIPHER_LEN
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedStr)
	if err != nil {
		this.Lock()
		this.Decerrcnt += 1
		this.Unlock()
		return "", err
	}

	data := encryptedData[ivStart:len(encryptedData)]
	if usePrivateEncrypt {
		keyData = crypto.Sha256([]byte(this.AppKey))
	} else {
		keyId := base64.StdEncoding.EncodeToString(encryptedData[CIPHER_LEN:ivStart])
		key, err := this.GetKey(keyId)
		if err != nil {
			this.Lock()
			this.Decerrcnt += 1
			this.Unlock()
			return "", err
		}
		keyData, err = base64.StdEncoding.DecodeString(key.KeyString)
		if err != nil {
			this.Lock()
			this.Decerrcnt += 1
			this.Unlock()
			return "", err
		}
	}
	origData, err := crypto.AESCBCDecrypt(keyData, data)
	if err != nil {
		this.Lock()
		this.Decerrcnt += 1
		this.Unlock()
		return "", err
	}
	this.Lock()
	this.Deccnt += 1
	this.Unlock()
	return origData, nil
}

func (this *Client) Encrypt(str string, usePrivateEncrypt bool) (string, error) {
	if str == "" {
		return "", nil
	}
	var (
		keyData   []byte
		keyIdData []byte
	)
	if usePrivateEncrypt {
		keyData = crypto.Sha256([]byte(this.AppKey))
		keyIdData = crypto.Sha256([]byte(this.AppSecret))[:aes.BlockSize]
	} else {
		key, err := this.GetCurrentKey()
		if err != nil {
			this.Lock()
			this.Encerrcnt += 1
			this.Unlock()
			return "", err
		}
		keyData, err = base64.StdEncoding.DecodeString(key.KeyString)
		if err != nil {
			this.Lock()
			this.Encerrcnt += 1
			this.Unlock()
			return "", err
		}
		keyIdData, err = base64.StdEncoding.DecodeString(key.Id)
		if err != nil {
			this.Lock()
			this.Encerrcnt += 1
			this.Unlock()
			return "", err
		}
	}
	encrypted, err := crypto.AESCBCEncrypt(keyData, str)
	if err != nil {
		this.Lock()
		this.Encerrcnt += 1
		this.Unlock()
		return "", err
	}
	data := make([]byte, 0, CIPHER_LEN+len(keyIdData)+len(encrypted))
	buf := bytes.NewBuffer(data)
	buf.Write([]byte{0, 4})
	buf.Write(keyIdData)
	buf.Write(encrypted)
	this.Lock()
	this.Enccnt += 1
	this.Unlock()
	return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
}

func (this *Client) Report(service string, reportText ReportText, reportType ReportType, code string, msg string, heap string) error {
	level := INFO_LEVEL
	if reportType == EXCEPTION_TYPE {
		level = ERROR_LEVEL
	}
	attr := ReportAttribute{
		Type:    reportType,
		Host:    this.Host,
		Level:   level,
		Service: service,
		SdkVer:  this.SdkVer,
		Env:     this.Env,
		Ts:      time.Now().UnixNano() / 1000000,
		Code:    code,
		Msg:     msg,
		Heap:    heap,
	}
	buf, _ := json.Marshal(attr)
	req := &secret.SecretApiReportGetRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
		BusinessId: NewBusinessId(),
		Text:       reportText,
		Attribute:  string(buf),
		ServerUrl:  DefaultServerUrl,
	}
	return secret.SecretApiReportGet(req)
}

func (this *Client) ReportStatitics() error {
	this.Lock()
	if this.KeyStore == nil {
		this.Unlock()
		return nil
	}
	attr := ReportStatistcAttribute{
		Type:      STATISTIC_TYPE,
		Host:      this.Host,
		Level:     INFO_LEVEL,
		Service:   this.KeyStore.Service,
		SdkVer:    this.SdkVer,
		Env:       this.Env,
		Ts:        time.Now().UnixNano() / 1000000,
		Enccnt:    this.Enccnt,
		Deccnt:    this.Deccnt,
		Encerrcnt: this.Encerrcnt,
		Decerrcnt: this.Decerrcnt,
	}
	this.Enccnt = 0
	this.Deccnt = 0
	this.Encerrcnt = 0
	this.Decerrcnt = 0
	this.Unlock()
	buf, _ := json.Marshal(attr)
	req := &secret.SecretApiReportGetRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
		BusinessId: NewBusinessId(),
		Text:       STATISTIC,
		Attribute:  string(buf),
		ServerUrl:  DefaultServerUrl,
	}
	return secret.SecretApiReportGet(req)
}
