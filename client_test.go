package azappconfig

import (
	"context"
	"fmt"
	"testing"
)

func TestBaseClient(t *testing.T) {
	client := New("Endpoint=https://haitchgo.azconfig.io;Id=2-l1-s0:XijEoV2A/dcM7LGKwcWP;Secret=Dd7JU24xI8b2SpWpAVtz70AEbc9sWF6nGK9Uqat+fsU=")
	fmt.Printf("%s\n", client.config.Endpoint)
	ctx := context.Background()
	response, err := client.ListKeys(ctx)
	fmt.Println(err.Error())
	fmt.Println(response.Status)
	fmt.Println(response.Header["Www-Authenticate"][0])
}
