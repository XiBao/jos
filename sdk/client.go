package sdk

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/XiBao/jos/sdk/internal/debug"
)

type Request struct {
	MethodName string
	Params     map[string]interface{}
	IsLogGW    bool `json:"-"`
	IsUnionGW  bool `json:"-"`
}

type Response struct {
	MethodName string
	Params     map[string]interface{}
}

type Client struct {
	AppKey    string
	SecretKey string

	Debug bool
}

func GetOauthURL(appKey, rURI, state, scope string) string {
	return fmt.Sprintf("https://open-oauth.jd.com/oauth2/to_login?app_key=%s&response_type=code&redirect_uri=%s&state=%s&scope=%s", appKey, url.QueryEscape(rURI), state, scope)
}

//create new client
func NewClient(appKey string, secretKey string) *Client {
	return &Client{
		AppKey:    appKey,
		SecretKey: secretKey,
	}
}

func (c *Client) GetAccessTokenNew(code string) (string, error) {
	gatewayUrl := fmt.Sprintf("https://open-oauth.jd.com/oauth2/access_token?app_key=%s&app_secret=%s&grant_type=authorization_code&code=%s", c.AppKey, c.SecretKey, code)
	debug.DebugPrintGetRequest(gatewayUrl)
	response, err := http.DefaultClient.Get(gatewayUrl)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	res := string(body)
	debug.DebugPrintStringResponse(res)
	return res, nil
}

func (c *Client) GetAccessToken(code, state, redirectUri string) (string, error) {
	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("client_id", c.AppKey)
	values.Add("redirect_uri", redirectUri)
	values.Add("code", code)
	values.Add("state", state)
	values.Add("client_secret", c.SecretKey)

	gatewayUrl := fmt.Sprintf(`%s?%s`, `https://oauth.jd.com/oauth/token`, values.Encode())
	debug.DebugPrintGetRequest(gatewayUrl)
	response, err := http.DefaultClient.Get(gatewayUrl)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	res := string(body)
	debug.DebugPrintStringResponse(res)
	return res, nil
}

func (c *Client) Execute(req *Request, token string) (result []byte, err error) {
	sysParams := make(map[string]string, 7)
	if paramJson, e := json.Marshal(req.Params); e != nil {
		return nil, e
	} else if req.IsUnionGW {
		sysParams["param_json"] = string(paramJson)
	} else {
		sysParams["360buy_param_json"] = string(paramJson)
	}
	sysParams["method"] = req.MethodName
	if token != "" {
		sysParams["access_token"] = token
	}
	sysParams["app_key"] = c.AppKey
	sysParams["timestamp"] = time.Now().Local().Format("2006-01-02 15:04:05")
	sysParams["format"] = "json"
	if req.IsUnionGW {
		sysParams["v"] = "1.0"
	} else {
		sysParams["v"] = API_VERSION
	}
	rawSign := c.GenerateRawSign(sysParams)
	sysParams["sign"] = c.GenerateSign(rawSign)
	values := url.Values{}
	for k, v := range sysParams {
		values.Add(k, v)
	}
	gwURL := GATEWAY_URL
	if c.Debug {
		gwURL = GATEWAY_DEV_URL
	} else if req.IsLogGW {
		gwURL = LOG_GATEWAY_URL
	} else if req.IsUnionGW {
		gwURL = UNION_GATEWAY_URL
	}
	debug.DebugPrintPostJSONRequest(gwURL, Json(sysParams))
	gatewayUrl := fmt.Sprintf(`%s?%s`, gwURL, values.Encode())
	if c.Debug {
		unescapeUrl, _ := url.QueryUnescape(gatewayUrl)
		fmt.Println(unescapeUrl)
	}
	debug.DebugPrintGetRequest(gatewayUrl)
	var (
		response *http.Response
		e        error
	)
	tryCnt := 3
	for {
		response, e = http.DefaultClient.Get(gatewayUrl)
		if e != nil {
			debug.DebugPrintError(err)
			if tryCnt <= 0 {
				return nil, Error{Code: 0, Msg: "HTTP Response Error"}
			} else {
				tryCnt--
				continue
			}
		}
		break
	}
	defer response.Body.Close()
	res, e := ioutil.ReadAll(response.Body)
	if e != nil {
		debug.DebugPrintError(err)
		return nil, Error{Code: 0, Msg: fmt.Sprintf("ReadAll on response.Body: %v", e)}
	}
	debug.DebugPrintStringResponse(string(res))
	return res, nil
}

func (c *Client) GenerateRawSign(params map[string]string) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	stringToBeSigned := c.SecretKey
	for _, k := range keys {
		stringToBeSigned += k + params[k]
	}
	stringToBeSigned += c.SecretKey
	return stringToBeSigned
}

func (c *Client) GenerateSign(stringToBeSigned string) string {
	h := md5.New()
	io.WriteString(h, stringToBeSigned)
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
