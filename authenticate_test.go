package azappconfig

import (
	"encoding/base64"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("tests Authentication function", func() {
	It("Test Parsing AccessKey", func() {
		parsed := parseAccessKey("Endpoint=https://haitchgo.azconfig.io;Id=readOnlyKey1;Secret=aGFpdGEK")
		Expect(parsed.Endpoint).To(BeEquivalentTo("https://haitchgo.azconfig.io"))
		Expect(parsed.ID).To(BeEquivalentTo("readOnlyKey1"))
		Expect(parsed.Secret).To(BeEquivalentTo("aGFpdGEK"))

		parsed = parseAccessKey("Endpoint=https://example.com;Id=dummyKeyID;Secret=aGFpdGFvCg==")
		Expect(parsed.Endpoint).To(BeEquivalentTo("https://example.com"))
		Expect(parsed.ID).To(BeEquivalentTo("dummyKeyID"))
		Expect(parsed.Secret).To(BeEquivalentTo("aGFpdGFvCg=="))
	})
	It("Test content hash", func() {
		parsed := parseAccessKey("Endpoint=https://haitchgo.azconfig.io;Id=fakeKeyId;Secret=UkVBTExZRkFLRVNFQ1JFVFMK")
		contentHash := getContentHashBase64(nil)
		fmt.Printf("x-ms-content-sha256 = %s \n", contentHash)
		stringToSign := getSigningContent("haitchgo.azconfig.io", "GET", "/keys", "Tue, 14 May 2019 07:36:46 GMT", contentHash)
		key, _ := base64.StdEncoding.DecodeString(parsed.Secret)
		signature := signRequest(stringToSign, key)
		fmt.Printf("signature = %s \n", signature)
	})
})
