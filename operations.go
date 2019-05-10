package azappconfig

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/tracing"
)

// List all app configs
func (client OperationsClient) ListKeys(ctx context.Context) (err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, "ListKeys")
		defer func() {
			sc := -1
			/* 			if result.olr.Response.Response != nil {
				sc = result.olr.Response.Response.StatusCode
			} */
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "azAppConfig.OperationsClient", "ListKeys", nil, "Failure preparing request")
		return
	}

	fmt.Printf(req.URL.String())
	return nil
}

// preparer prepares the List request.
func (client OperationsClient) preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.config.Endpoint),
		autorest.WithPath("/keys"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// sender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client OperationsClient) sender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}
