package azappconfig

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"strings"
)

func getSigningContent(host string, verb string, pathAndQuery string, utcTimestamp string, contentHashBase64 string) string {
	message := fmt.Sprintf("%s\n%s\n%s;%s;%s", strings.ToUpper(verb), pathAndQuery, utcTimestamp, host, contentHashBase64)
	return message
}

func signRequest(content string, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	signed := mac.Sum(nil)
	return signed
}

func parseAccessKey(accessKeyStr string) accessKey {
	kvps := strings.Split(accessKeyStr, ";")
	result := accessKey{}
	for _, kvp := range kvps {
		kv := strings.Split(kvp, "=")
		k := kv[0]
		v := kv[1]
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
