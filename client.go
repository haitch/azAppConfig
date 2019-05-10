package azappconfig

import (
	"github.com/Azure/go-autorest/autorest"
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

// OperationsClient is the Azure App Configuration Client.
type OperationsClient struct {
	BaseClient
}

// New creates an instance of BaseClient
func newBase(accessKey string) BaseClient {
	return BaseClient{
		Client: autorest.NewClientWithUserAgent(DefaultUserAgent),
		config: parseAccessKey(accessKey),
	}
}

func New(accessKey string) OperationsClient {
	return OperationsClient{newBase(accessKey)}
}
