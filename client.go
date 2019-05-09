package azappconfig

import (
	"github.com/azure/go-autorest/autorest"
)

const (
	// DefaultUserAgent is default value for User-Agent header
	DefaultUserAgent = "github.com/haitch/azAppConfig"
)

// BaseClient is base client for Azure App Configuration
type BaseClient struct {
	autorest.Client
	config accessKey
}

// New creates an instance of BaseClient
func New(accessKey string) BaseClient {
	return BaseClient{
		Client: autorest.NewClientWithUserAgent(DefaultUserAgent),
		config: parseAccessKey(accessKey),
	}
}
