package azappconfig

import (
	"encoding/base64"
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

func TestContentHash(t *testing.T) {
	parsed := parseAccessKey("Endpoint=https://haitchgo.azconfig.io;Id=2-l1-s0:XijEoV2A/dcM7LGKwcWP;Secret=Dd7JU24xI8b2SpWpAVtz70AEbc9sWF6nGK9Uqat+fsU=")
	contentHash := getContentHashBase64(nil)
	fmt.Printf("x-ms-content-sha256 = %s \n", contentHash)
	stringToSign := getSigningContent("haitchgo.azconfig.io", "GET", "/keys", "Tue, 14 May 2019 07:36:46 GMT", contentHash)
	key, _ := base64.StdEncoding.DecodeString(parsed.Secret)
	signature := signRequest(stringToSign, key)
	fmt.Printf("signature = %s \n", signature)
}
