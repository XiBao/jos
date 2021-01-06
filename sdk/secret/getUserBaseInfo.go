package secret

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/user"
)

func (this *Client) GetUserBaseInfo(pin string, loadType int, decrypt bool) (*user.UserInfo, error) {
	req := &user.GetUserBaseInfoByEncryPinRequest{
		BaseRequest: api.BaseRequest{
			AnApiKey: &api.ApiKey{
				Key:    this.AppKey,
				Secret: this.AppSecret,
			},
			Session: this.AccessToken,
		},
		Pin:      pin,
		LoadType: loadType,
	}
	userInfo, err := user.GetUserBaseInfoByEncryPin(req)
	if err != nil {
		return nil, err
	}
	if !decrypt {
		err = this.DecryptUserInfo(userInfo, false)
		if err != nil {
			return nil, err
		}
	}

	return userInfo, nil
}

func (this *Client) DecryptUserInfo(userInfo *user.UserInfo, usePrivateKey bool) (err error) {
	userInfo.EncryptEmail = userInfo.Email
	if userInfo.Email, err = this.Decrypt(userInfo.EncryptEmail, usePrivateKey); err != nil {
		return err
	}
	userInfo.EncryptMobile = userInfo.Mobile
	if userInfo.Mobile, err = this.Decrypt(userInfo.EncryptMobile, usePrivateKey); err != nil {
		return err
	}

	userInfo.EncryptIntactMobile = userInfo.IntactMobile
	if userInfo.IntactMobile, err = this.Decrypt(userInfo.EncryptIntactMobile, usePrivateKey); err != nil {
		return err
	}

	return nil
}

func (this *Client) EncryptUserInfo(userInfo *user.UserInfo, usePrivateKey bool) (err error) {
	if userInfo.EncryptEmail, err = this.Decrypt(userInfo.Email, usePrivateKey); err != nil {
		return err
	}
	if userInfo.EncryptMobile, err = this.Decrypt(userInfo.Mobile, usePrivateKey); err != nil {
		return err
	}

	if userInfo.EncryptIntactMobile, err = this.Decrypt(userInfo.IntactMobile, usePrivateKey); err != nil {
		return err
	}

	return nil
}
