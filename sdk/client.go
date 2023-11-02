package sdk

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/XiBao/jos/sdk/internal/logger"
)

type Request struct {
	MethodName string
	Params     map[string]interface{}
	IsLogGW    bool `json:"-"`
	IsUnionGW  bool `json:"-"`
}

type IResponse interface {
	error
	IsError() bool
}

type Response struct {
	MethodName string
	Params     map[string]interface{}
}

type Client struct {
	AppKey    string
	SecretKey string

	Dev   bool
	Debug bool
}

func GetOauthURL(appKey, rURI, state, scope string) string {
	values := GetUrlValues()
	values.Set("app_key", appKey)
	values.Set("response_type", "code")
	values.Set("redirect_uri", rURI)
	values.Set("state", state)
	values.Set("scope", scope)
	enc := values.Encode()
	PutUrlValues(values)
	return StringsJoin("https://open-oauth.jd.com/oauth2/to_login?", enc)
}

// create new client
func NewClient(appKey string, secretKey string) *Client {
	return &Client{
		AppKey:    appKey,
		SecretKey: secretKey,
	}
}

func (c *Client) Logger() logger.Logger {
	if c.Debug {
		return logger.Debug
	}
	return logger.Default
}

func (c *Client) GetAccessTokenNew(code string) (string, error) {
	values := GetUrlValues()
	values.Set("app_key", c.AppKey)
	values.Set("app_secret", c.SecretKey)
	values.Set("grant_type", "authorization_code")
	values.Set("code", code)
	payload := values.Encode()
	PutUrlValues(values)
	gatewayUrl := StringsJoin("https://open-oauth.jd.com/oauth2/access_token?", payload)
	debug := c.Logger()
	debug.DebugPrintGetRequest(gatewayUrl)
	response, err := http.DefaultClient.Get(gatewayUrl)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	res := string(body)
	debug.DebugPrintStringResponse(res)
	return res, nil
}

func (c *Client) GetAccessToken(code, state, redirectUri string) (string, error) {
	values := GetUrlValues()
	values.Set("grant_type", "authorization_code")
	values.Set("client_id", c.AppKey)
	values.Set("redirect_uri", redirectUri)
	values.Set("code", code)
	values.Set("state", state)
	values.Set("client_secret", c.SecretKey)
	payload := values.Encode()
	PutUrlValues(values)
	gatewayUrl := StringsJoin(`https://oauth.jd.com/oauth/token?`, payload)
	debug := c.Logger()
	debug.DebugPrintGetRequest(gatewayUrl)
	response, err := http.DefaultClient.Get(gatewayUrl)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		debug.DebugPrintError(err)
		return "", err
	}
	res := string(body)
	debug.DebugPrintStringResponse(res)
	return res, nil
}

func (c *Client) SetDev(dev bool) {
	c.Dev = dev
}

func (c *Client) Execute(req *Request, token string, rep IResponse) error {
	sysParams := make(map[string]string, 7)
	if paramJson, err := json.Marshal(req.Params); err != nil {
		return err
	} else if req.IsUnionGW {
		sysParams["360buy_param_json"] = string(paramJson)
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
	values := GetUrlValues()
	for k, v := range sysParams {
		values.Set(k, v)
	}
	payload := values.Encode()
	PutUrlValues(values)
	gwURL := GATEWAY_URL
	if c.Dev {
		gwURL = GATEWAY_DEV_URL
	} else if req.IsLogGW {
		gwURL = LOG_GATEWAY_URL
	} else if req.IsUnionGW {
		gwURL = UNION_GATEWAY_URL
	}
	debug := c.Logger()
	debug.DebugPrintPostJSONRequest(gwURL, Json(sysParams))
	gatewayUrl := StringsJoin(gwURL, "?", payload)
	debug.DebugPrintGetRequest(gatewayUrl)

	response, err := http.DefaultClient.Get(gatewayUrl)
	if err != nil {
		debug.DebugPrintError(err)
		return err
	}
	defer response.Body.Close()
	if body, err := io.ReadAll(response.Body); err != nil {
		debug.DebugPrintError(err)
		return err
	} else if rep != nil {
		if err := debug.DecodeJSON(body, rep); err != nil {
			return errors.Join(Error{Code: 0, Msg: string(body)}, err)
		}
		if rep.IsError() {
			return rep
		}
	} else {
		debug.DebugPrintStringResponse(string(body))
	}
	return nil
}

func (c *Client) PostExecute(req *Request, token string, rep IResponse) error {
	sysParams := make(map[string]string, 7)
	if paramJson, err := json.Marshal(req.Params); err != nil {
		return err
	} else if req.IsUnionGW {
		sysParams["360buy_param_json"] = string(paramJson)
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
	values := GetUrlValues()
	for k, v := range sysParams {
		values.Set(k, v)
	}
	gwURL := GATEWAY_URL
	if c.Dev {
		gwURL = GATEWAY_DEV_URL
	} else if req.IsLogGW {
		gwURL = LOG_GATEWAY_URL
	} else if req.IsUnionGW {
		gwURL = UNION_GATEWAY_URL
	}
	debug := c.Logger()
	debug.DebugPrintPostJSONRequest(gwURL, Json(sysParams))

	response, err := http.PostForm(gwURL, values)
	PutUrlValues(values)
	if err != nil {
		debug.DebugPrintError(err)
		return err
	}
	defer response.Body.Close()
	if body, err := io.ReadAll(response.Body); err != nil {
		debug.DebugPrintError(err)
		return err
	} else if rep != nil {
		if err := debug.DecodeJSON(body, rep); err != nil {
			return errors.Join(Error{Code: 0, Msg: string(body)}, err)
		}
		if rep.IsError() {
			return rep
		}
	} else {
		debug.DebugPrintStringResponse(string(body))
	}
	return nil
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
