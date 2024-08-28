package secret

import (
	"context"

	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/user"
)

func (this *Client) GetUserBaseInfo(ctx context.Context, pin string, loadType int, decrypt bool) (*user.UserInfo, error) {
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
	userInfo, err := user.GetUserBaseInfoByEncryPin(ctx, req)
	if err != nil {
		return nil, err
	}
	if !decrypt {
		err = this.DecryptUserInfo(ctx, userInfo, false)
		if err != nil {
			return nil, err
		}
	}

	return userInfo, nil
}

func (this *Client) DecryptUserInfo(ctx context.Context, userInfo *user.UserInfo, usePrivateKey bool) (err error) {
	userInfo.EncryptEmail = userInfo.Email
	if userInfo.Email, err = this.Decrypt(ctx, userInfo.EncryptEmail, usePrivateKey); err != nil {
		return err
	}
	userInfo.EncryptMobile = userInfo.Mobile
	if userInfo.Mobile, err = this.Decrypt(ctx, userInfo.EncryptMobile, usePrivateKey); err != nil {
		return err
	}

	userInfo.EncryptIntactMobile = userInfo.IntactMobile
	if userInfo.IntactMobile, err = this.Decrypt(ctx, userInfo.EncryptIntactMobile, usePrivateKey); err != nil {
		return err
	}

	return nil
}

func (this *Client) EncryptUserInfo(ctx context.Context, userInfo *user.UserInfo, usePrivateKey bool) (err error) {
	if userInfo.EncryptEmail, err = this.Decrypt(ctx, userInfo.Email, usePrivateKey); err != nil {
		return err
	}
	if userInfo.EncryptMobile, err = this.Decrypt(ctx, userInfo.Mobile, usePrivateKey); err != nil {
		return err
	}

	if userInfo.EncryptIntactMobile, err = this.Decrypt(ctx, userInfo.IntactMobile, usePrivateKey); err != nil {
		return err
	}

	return nil
}
