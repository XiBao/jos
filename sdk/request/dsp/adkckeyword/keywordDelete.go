package adkckeyword

import "github.com/XiBao/jos/sdk"

type KeywordDeleteRequest struct {
	Request *sdk.Request
}

func NewKeywordDeleteRequest() (req *KeywordDeleteRequest) {
	request := sdk.Request{MethodName: "jingdong.dsp.adkckeyword.keyword.delete", Params: make(map[string]interface{}, 2)}
	req = &KeywordDeleteRequest{
		Request: &request,
	}
	return
}

func (req *KeywordDeleteRequest) SetKeyWordsName(keyWordsName string) {
	req.Request.Params["keyWordsName"] = keyWordsName
}

func (req *KeywordDeleteRequest) GetKeyWordsName() string {
	keyWordsName, found := req.Request.Params["keyWordsName"]
	if found {
		return keyWordsName.(string)
	}
	return ""
}

func (req *KeywordDeleteRequest) SetAdGroupId(adGroupId uint64) {
	req.Request.Params["adGroupId"] = adGroupId
}

func (req *KeywordDeleteRequest) GetAdGroupId() uint64 {
	adGroupId, found := req.Request.Params["adGroupId"]
	if found {
		return adGroupId.(uint64)
	}
	return 0
}
