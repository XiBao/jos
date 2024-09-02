package secret

import (
	"context"
	"sync"

	"github.com/XiBao/jos/sdk"
)

type Service struct {
	clients map[string]*Client
	host    string
	env     string
	sync.RWMutex
}

func NewService(host string, env string) *Service {
	return &Service{
		clients: make(map[string]*Client),
		host:    host,
		env:     env,
	}
}

func instanceKey(appKey string, accessToken string) string {
	return sdk.StringsJoin(appKey, "-", accessToken)
}

func (this *Service) Instance(ctx context.Context, appKey string, secret string, accessToken string) *Client {
	key := instanceKey(appKey, accessToken)
	this.RLock()
	if client, found := this.clients[key]; found {
		this.RUnlock()
		return client
	}
	this.RUnlock()
	this.Lock()
	client := NewClient(ctx, appKey, secret, accessToken)
	client.SetHost(this.host)
	client.SetEnv(this.env)
	this.clients[key] = client
	this.Unlock()
	return client
}
