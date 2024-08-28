package secret

import (
	"bytes"
	"context"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
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

func NewClient(ctx context.Context, appKey string, appSecret, accessToken string) *Client {
	client := &Client{
		AppKey:      appKey,
		AppSecret:   appSecret,
		AccessToken: accessToken,
		SdkVer:      2,
	}
	go client.StartReport(ctx)
	return client
}

func (this *Client) SetHost(host string) {
	this.Host = host
}

func (this *Client) SetEnv(env string) {
	this.Env = env
}

func (this *Client) StartReport(ctx context.Context) {
	hourlyTicker := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-hourlyTicker.C:
			this.ReportStatitics(ctx)
		}
	}
}

func (this *Client) GetVoucher(ctx context.Context) (voucherData voucher.VoucherData, err error) {
	req := &voucher.VoucherInfoGetRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
	}
	return voucher.VoucherInfoGet(ctx, req)
}

func (this *Client) GetMasterKey(ctx context.Context, tid string, voucherKey string) (keyStore *crypto.KeyStore, err error) {
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
	if keyStores, err := master.MasterKeyGet(ctx, req); err != nil {
		return nil, err
	} else if len(keyStores) == 0 {
		return nil, errors.New("empty keystore")
	} else {
		keyStore = &keyStores[0]
	}
	this.Lock()
	this.KeyStore = keyStore
	this.Unlock()
	return
}

func (this *Client) RefreshKeyStore(ctx context.Context) (*crypto.KeyStore, error) {
	voucher, err := this.GetVoucher(ctx)
	if err != nil {
		return nil, err
	}
	keyStore, err := this.GetMasterKey(ctx, voucher.Id, voucher.Key)
	if err != nil {
		return nil, err
	}
	this.Report(ctx, voucher.Service, INIT, INIT_TYPE, "0", "", "")
	return keyStore, nil
}

func (this *Client) GetKey(ctx context.Context, keyId string) (key crypto.Key, err error) {
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
	keyStore, err := this.RefreshKeyStore(ctx)
	if err != nil {
		return key, err
	}
	key, found = keyStore.GetKey(keyId)
	if !found {
		return key, errors.New("not found key")
	}
	return key, nil
}

func (this *Client) GetCurrentKey(ctx context.Context) (key crypto.Key, err error) {
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
	keyStore, err := this.RefreshKeyStore(ctx)
	if err != nil {
		return key, err
	}
	key, found = keyStore.GetCurrentKey()
	if !found {
		return key, errors.New("not found key")
	}
	return key, nil
}

func (this *Client) GetFirstKey(ctx context.Context) (key crypto.Key, err error) {
	var found bool
	this.RLock()
	if this.KeyStore == nil {
		found = false
	} else {
		key, found = this.KeyStore.GetFirstKey()
	}
	this.RUnlock()
	if found {
		return key, nil
	}
	keyStore, err := this.RefreshKeyStore(ctx)
	if err != nil {
		return key, err
	}
	key, found = keyStore.GetCurrentKey()
	if !found {
		return key, errors.New("not found key")
	}
	return key, nil
}

func (this *Client) Salt(ctx context.Context, usePrivateEncrypt bool) ([]byte, error) {
	var keyData []byte
	if usePrivateEncrypt {
		keyData = crypto.Sha256([]byte(this.AppKey))
	} else {
		key, err := this.GetFirstKey(ctx)
		if err != nil {
			return nil, err
		}
		keyData, err = base64.StdEncoding.DecodeString(key.KeyString)
	}
	return crypto.AESCBCEncryptZeroIV(keyData)
}

func (this *Client) Hash(ctx context.Context, oriData string, usePrivateEncrypt bool) (string, error) {
	salt, err := this.Salt(ctx, usePrivateEncrypt)
	if err != nil {
		return "", err
	}
	oriDataBytes := []byte(oriData)
	data := make([]byte, 0, len(oriDataBytes)+len(salt))
	buf := bytes.NewBuffer(data)
	buf.Write(oriDataBytes)
	buf.Write(salt)
	return hex.EncodeToString(crypto.Sha256(buf.Bytes())), nil
}

func (this *Client) Decrypt(ctx context.Context, encryptedStr string, usePrivateEncrypt bool) (string, error) {
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

	data := encryptedData[ivStart:]
	if usePrivateEncrypt {
		keyData = crypto.Sha256([]byte(this.AppKey))
	} else {
		keyId := base64.StdEncoding.EncodeToString(encryptedData[CIPHER_LEN:ivStart])
		key, err := this.GetKey(ctx, keyId)
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

func (this *Client) Encrypt(ctx context.Context, str string, usePrivateEncrypt bool) (string, error) {
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
		key, err := this.GetCurrentKey(ctx)
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

func (this *Client) Report(ctx context.Context, service string, reportText ReportText, reportType ReportType, code string, msg string, heap string) error {
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
	return secret.SecretApiReportGet(ctx, req)
}

func (this *Client) ReportStatitics(ctx context.Context) error {
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
	return secret.SecretApiReportGet(ctx, req)
}
