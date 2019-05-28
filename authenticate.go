package azappconfig

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Azure/go-autorest/autorest"
)

// HMACAuthorizer implements the bearer authorization
type HMACAuthorizer struct {
	KeyID string
	Key   []byte
}

// NewHMACAuthorizer crates a BearerAuthorizer using the given token provider
func NewHMACAuthorizer(id string, key string) *HMACAuthorizer {
	keyBytes, _ := base64.StdEncoding.DecodeString(key)
	return &HMACAuthorizer{KeyID: id, Key: keyBytes}
}

// WithAuthorization returns a PrepareDecorator that adds an HTTP Authorization header whose
// value is "Bearer " followed by the token.
//
// By default, the token will be automatically refreshed through the Refresher interface.
func (hmacA *HMACAuthorizer) WithAuthorization() autorest.PrepareDecorator {
	return func(p autorest.Preparer) autorest.Preparer {
		return autorest.PreparerFunc(func(r *http.Request) (*http.Request, error) {
			r, _ = p.Prepare(r)
			timestamp := time.Now().UTC().Format(http.TimeFormat)
			contentHashBase64 := getContentHashBase64(r)
			stringToSign := getSigningContent(r.URL.Host, r.Method, r.URL.Path, timestamp, contentHashBase64)
			signature := signRequest(stringToSign, hmacA.Key)
			authStr := fmt.Sprintf("HMAC-SHA256 Credential=%s, SignedHeaders=x-ms-date;host;x-ms-content-sha256, Signature=%s", hmacA.KeyID, signature)
			return autorest.Prepare(
				r,
				autorest.WithHeader("x-ms-date", timestamp),
				autorest.WithHeader("x-ms-content-sha256", contentHashBase64),
				autorest.WithHeader("Authorization", authStr),
			)
		})
	}
}

func getContentHashBase64(r *http.Request) string {
	buf := new(bytes.Buffer)
	if r != nil && r.Body != nil {
		buf.ReadFrom(r.Body)
	}
	hasher := sha256.New()
	hasher.Write(buf.Bytes())
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

func getSigningContent(host string, verb string, pathAndQuery string, timestamp string, contentHashBase64 string) string {
	return fmt.Sprintf("%s\n%s\n%s;%s;%s", strings.ToUpper(verb), pathAndQuery, timestamp, host, contentHashBase64)
}

func signRequest(content string, key []byte) string {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func parseAccessKey(accessKeyStr string) *accessKey {
	kvps := strings.Split(accessKeyStr, ";")
	result := &accessKey{}
	for _, kvp := range kvps {
		index := strings.Index(kvp, "=")
		k := kvp[0:index]
		v := kvp[index+1:]
		if strings.EqualFold(k, "Endpoint") {
			result.Endpoint = v
		}
		if strings.EqualFold(k, "Id") {
			result.ID = v
		}
		if strings.EqualFold(k, "Secret") {
			result.Secret = v
		}
	}
	return result
}
