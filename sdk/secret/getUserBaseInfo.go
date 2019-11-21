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
		return userInfo, nil
	}
	userInfo.EncryptEmail = userInfo.Email
	if userInfo.Email, err = this.Decrypt(userInfo.EncryptEmail, false); err != nil {
		return nil, err
	}
	userInfo.EncryptMobile = userInfo.Mobile
	if userInfo.Mobile, err = this.Decrypt(userInfo.EncryptMobile, false); err != nil {
		return nil, err
	}

	userInfo.EncryptIntactMobile = userInfo.IntactMobile
	if userInfo.IntactMobile, err = this.Decrypt(userInfo.EncryptIntactMobile, false); err != nil {
		return nil, err
	}
	return userInfo, nil
}
