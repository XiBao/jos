package sdk

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/XiBao/jos/sdk/internal/logger"
)

var (
	onceInit   sync.Once
	httpClient *http.Client
	tracerMap  = new(sync.Map)
)

func defaultHttpClient() *http.Client {
	onceInit.Do(func() {
		transport := http.DefaultTransport.(*http.Transport).Clone()
		transport.MaxIdleConns = 100
		transport.MaxConnsPerHost = 100
		transport.MaxIdleConnsPerHost = 100
		httpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Second * 60,
		}
	})
	return httpClient
}

type Request struct {
	Params     map[string]interface{}
	MethodName string
	IsLogGW    bool `json:"-"`
	IsUnionGW  bool `json:"-"`
}

type IResponse interface {
	error
	IsError() bool
}

type Response struct {
	Params     map[string]interface{}
	MethodName string
}

type Client struct {
	client    *http.Client
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
	clt := &Client{
		AppKey:    appKey,
		SecretKey: secretKey,
		client:    defaultHttpClient(),
	}
	clt.WithTracer("")
	return clt
}

func (c *Client) Logger() logger.Logger {
	if c.Debug {
		return logger.Debug
	}
	return logger.Default
}

// SetHttpClient 设置http.Client
func (c *Client) SetHttpClient(client *http.Client) {
	c.client = client
}

func (c *Client) WithTracer(namespace string) {
	tracerMap.LoadOrStore(c.AppKey, NewOtel(namespace, c.AppKey))
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
	response, err := c.client.Get(gatewayUrl)
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
	response, err := c.client.Get(gatewayUrl)
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

func (c *Client) Execute(ctx context.Context, req *Request, token string, rep IResponse) error {
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

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, gatewayUrl, nil)
	if err != nil {
		debug.DebugPrintError(err)
		return err
	}
	httpReq.Header.Set("content-type", "application/json")
	return c.WithSpan(ctx, req.MethodName, httpReq, rep, nil, c.fetch)
}

func (c *Client) PostExecute(ctx context.Context, req *Request, token string, rep IResponse) error {
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
	defer PutUrlValues(values)
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
	payload := values.Encode()
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, gwURL, strings.NewReader(payload))
	if err != nil {
		debug.DebugPrintError(err)
		return err
	}
	httpReq.Header.Set("content-type", "application/x-www-form-urlencoded")
	return c.WithSpan(ctx, req.MethodName, httpReq, rep, []byte(payload), c.fetch)
}

func (c *Client) fetch(httpReq *http.Request, rep IResponse) (*http.Response, error) {
	debug := c.Logger()
	response, err := c.client.Do(httpReq)
	if err != nil {
		debug.DebugPrintError(err)
		return response, err
	}
	defer response.Body.Close()
	if body, err := io.ReadAll(response.Body); err != nil {
		debug.DebugPrintError(err)
		return response, err
	} else if rep != nil {
		if err := debug.DecodeJSON(body, rep); err != nil {
			return response, errors.Join(Error{Code: 0, Msg: string(body)}, err)
		}
		if rep.IsError() {
			return response, rep
		}
	} else {
		debug.DebugPrintStringResponse(string(body))
	}
	return response, nil
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

func (c *Client) WithSpan(ctx context.Context, methodName string, req *http.Request, resp IResponse, payload []byte, fn func(*http.Request, IResponse) (*http.Response, error)) error {
	tracer, ok := tracerMap.Load(c.AppKey)
	if !ok {
		_, err := fn(req, resp)
		return err
	}
	return tracer.(*Otel).WithSpan(ctx, methodName, req, resp, payload, fn)
}
