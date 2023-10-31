package user

type UserInfo struct {
	UserId         int64  `json:"userId,omitempty" codec:"userId,omitempty"`
	Nickname       string `json:"nickname,omitempty" codec:"nickname,omitempty"`
	NicknameShow   string `json:"nicknameShow,omitempty" codec:"nicknameShow,omitempty"`
	UserLevel      int    `json:"userLevel,omitempty" codec:"userLevel,omitempty"`
	Gendar         uint8  `json:"gendar,omitempty" codec:"gendar,omitempty"`
	YunBigImageUrl string `json:"yunBigImageUrl,omitempty" codec:"yunBigImageUrl,omitempty"`
	Pin            string `json:"pin,omitempty" codec:"pin,omitempty"`           // 明文pin
	Birthday       string `json:"birthday,omitempty" codec:"birthday,omitempty"` // 生日
	OpenId         string `json:"open_id,omitempty"`                             // openId
	CountryCode    string `json:"countryCode,omitempty" codec:"countryCode,omitempty"`
	County         int64  `json:"county,omitempty" codec:"county,omitempty"`
	Province       int64  `json:"province,omitempty" codec:"province,omitempty"`
	City           int64  `json:"city,omitempty" codec:"city,omitempty"`

	Email        string `json:"email,omitempty" codec:"email,omitempty"`
	Mobile       string `json:"mobile,omitempty" codec:"mobile,omitempty"`
	IntactMobile string `json:"intactMobile,omitempty" codec:"intactMobile,omitempty"`

	EncryptEmail        string `json:"encrypt_email,omitempty" codec:"encrypt_email,omitempty"`
	EncryptMobile       string `json:"encrypt_mobile,omitempty" codec:"encrypt_mobile,omitempty"`
	EncryptIntactMobile string `json:"encrypt_intactMobile,omitempty" codec:"encrypt_intactMobile,omitempty"`
	EncryptAccountId    uint64 `json:"encrypt_account_id,omitempty" codec:"encrypt_account_id" `

	Created     int64 `json:"created,omitempty" codec:"created,omitempty"`
	Modified    int64 `json:"modified,omitempty" codec:"modified,omitempty"`
	RegTime     int64 `json:"regTime,omitempty" codec:"regTime,omitempty"`
	UpgradeTime int64 `json:"upgradeTime,omitempty" codec:"upgradeTime,omitempty"`

	EncryPin string `json:"encry_pin,omitempty" codec:"encry_pin,omitempty"`
}

type SocialInfo struct {
	Birthday       string `json:"birthday,omitempty"`
	Gender         int    `json:"gender,omitempty"`
	YunBigImageUrl string `json:"yunBigImageUrl,omitempty" codec:"yunBigImageUrl,omitempty"`
	Nickname       string `json:"nickname,omitempty" codec:"nickname,omitempty"`
	EncryPin       string `json:"encry_pin,omitempty" codec:"encry_pin,omitempty"`
}
