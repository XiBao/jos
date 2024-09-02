package crypto

type KeyStore struct {
	Service           string `json:"service"`             // 服务识别码
	GrantUsage        string `json:"grant_usage"`         // ignore，密钥权限，E为加密，D为解密
	Keys              []Key  `json:"keys"`                // 加解密密钥，可能是多个
	CurrentKeyVersion uint   `json:"current_key_version"` // 最新密钥版本号
}

type Key struct {
	Id           string `json:"id"`            // 密钥id
	KeyString    string `json:"key_string"`    // 密钥base64编码
	KeyType      string `json:"key_type"`      // ignore，密钥种类
	KeyDigest    string `json:"key_digest"`    // 密钥digest(SHA256)，客户端可以校验密钥完整性
	KeyExp       int64  `json:"key_exp"`       // ignore，unix timestamp，密钥过期时间
	KeyEffective int64  `json:"key_effective"` // ignore，unix timestamp，密钥生效日期
	Version      uint   `json:"version"`       // 密钥版本号(从0开始)
	KeyStatus    int    `json:"key_status"`    // ignore，密钥状态, 0为可用
}

func (this *KeyStore) GetKey(keyId string) (Key, bool) {
	for _, key := range this.Keys {
		if key.Id == keyId {
			return key, true
		}
	}
	return Key{}, false
}

func (this *KeyStore) GetFirstKey() (Key, bool) {
	for _, key := range this.Keys {
		if key.Version == 0 {
			return key, true
		}
	}
	return Key{}, false
}

func (this *KeyStore) GetCurrentKey() (Key, bool) {
	for _, key := range this.Keys {
		if key.Version == this.CurrentKeyVersion {
			return key, true
		}
	}
	return Key{}, false
}
