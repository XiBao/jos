package secret

import (
	"github.com/XiBao/jos/api"
	"github.com/XiBao/jos/api/user"
)

func (this *Client) GetUserBaseInfo(pin string, loadType int) (*user.UserInfo, error) {
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
	if userInfo.EncryptEmail != "" {
		if userInfo.Email, err = this.Decrypt(userInfo.EncryptEmail, false); err != nil {
			return nil, err
		}
	}
	if userInfo.EncryptMobile != "" {
		if userInfo.Mobile, err = this.Decrypt(userInfo.EncryptMobile, false); err != nil {
			return nil, err
		}
	}
	if userInfo.EncryptIntactMobile != "" {
		if userInfo.IntactMobile, err = this.Decrypt(userInfo.EncryptIntactMobile, false); err != nil {
			return nil, err
		}
	}
	return userInfo, nil
}
