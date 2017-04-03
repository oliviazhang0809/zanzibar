// Code generated by zanzibar
// @generated

package bar

import (
	"context"
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/uber-go/zap"
	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"

	endpoints_bar_bar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/github.com/uber/zanzibar/endpoints/bar/bar"

	clients_bar_bar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/github.com/uber/zanzibar/clients/bar/bar"
)

// HandleNoRequestRequest handles "/bar/no-request-path".
func HandleNoRequestRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
	clients *clients.Clients,
) {

	// Handle request body.
	clientResp, err := clients.Bar.NoRequest(ctx)
	if err != nil {
		req.Logger.Error("Could not make client request",
			zap.String("error", err.Error()),
		)
		res.SendError(500, errors.Wrap(err, "could not make client request:"))
		return
	}

	defer func() {
		if cerr := clientResp.Body.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()

	// Handle client respnse.
	expectedStatusCode := []int{200}
	if !res.IsOKResponse(clientResp.StatusCode, expectedStatusCode) {
		req.Logger.Warn("Unknown response status code",
			zap.Int("status code", clientResp.StatusCode),
		)
	}
	b, err := ioutil.ReadAll(clientResp.Body)
	if err != nil {
		res.SendError(500, errors.Wrap(err, "could not read client response body:"))
		return
	}
	var clientRespBody clients_bar_bar.BarResponse
	if err := clientRespBody.UnmarshalJSON(b); err != nil {
		res.SendError(500, errors.Wrap(err, "could not unmarshal client response body:"))
		return
	}
	response := convertNoRequestClientResponse(&clientRespBody)
	res.WriteJSON(200, response)
}

func convertNoRequestClientResponse(body *clients_bar_bar.BarResponse) *endpoints_bar_bar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := &endpoints_bar_bar.BarResponse{}
	return downstreamResponse
}
