package azappconfig

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAccessKey(t *testing.T) {
	parsed := parseAccessKey("Endpoint=https://haitchgo.azconfig.io;Id=readOnlyKey1;Secret=someKey")
	assert.Equal(t, "https://haitchgo.azconfig.io", parsed.Endpoint)
	assert.Equal(t, "readOnlyKey1", parsed.ID)
	assert.Equal(t, "someKey", parsed.Secret)
	fmt.Printf("Endpoing=%s;Id=%s;Secret=%s", parsed.Endpoint, parsed.ID, parsed.Secret)
}
