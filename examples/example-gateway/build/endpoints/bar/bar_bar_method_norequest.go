// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package bar

import (
	"context"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"
)

// NoRequestHandler is the handler for "/bar/no-request-path"
type NoRequestHandler struct {
	Clients *clients.Clients
}

// NewNoRequestEndpoint creates a handler
func NewNoRequestEndpoint(
	gateway *zanzibar.Gateway,
) *NoRequestHandler {
	return &NoRequestHandler{
		Clients: gateway.Clients.(*clients.Clients),
	}
}

// HandleRequest handles "/bar/no-request-path".
func (handler *NoRequestHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	workflow := NoRequestEndpoint{
		Clients: handler.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBarBar.BarException:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error",
				zap.String("error", errValue.Error()),
			)
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}

	res.WriteJSON(200, cliRespHeaders, response)
}

// NoRequestEndpoint calls thrift client Bar.NoRequest
type NoRequestEndpoint struct {
	Clients *clients.Clients
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w NoRequestEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.NoRequest(
		ctx, clientHeaders,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertNoRequestBarException(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request",
				zap.String("error", errValue.Error()),
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertNoRequestClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertNoRequestBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}

func convertNoRequestClientResponse(body *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	// TODO: Add response fields mapping here.
	downstreamResponse := (*endpointsBarBar.BarResponse)(body)
	return downstreamResponse
}
