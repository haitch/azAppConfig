package azappconfig

import (
	"fmt"
	"context"
	"os"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var realConnectionString = os.Getenv("AzAppConfig_ConnectionString")

var _ = Describe("tests End to End scenarios", func() {
	BeforeEach(func() {
		if realConnectionString == "" {
			Skip("real connection string is not set. The end to end test is skipped.")
		}
	})
	It("Test List Keys", func() {
		client := New(realConnectionString)
		ctx := context.Background()
		response, err := client.ListKeys(ctx)
		Expect(err).To(BeNil())
		Expect(response.Items).NotTo(BeEmpty())
		fmt.Printf("%+v", response.Items)
	})
})
