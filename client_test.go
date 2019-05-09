package azappconfig

import (
	"fmt"
	"testing"
)

func TestBaseClient(t *testing.T) {
	client := New("Endpoint=https://haitchgo.azconfig.io;Id=readOnlyKey1;Secret=someKey")
	fmt.Printf("%s", client.config.Endpoint)
}
