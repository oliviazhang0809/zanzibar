// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package multiEndpoint

import (
	"context"
	"encoding/json"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi/module"
)

// ServiceBFrontHelloHandler is the handler for "/multi/serviceB_f/hello"
type ServiceBFrontHelloHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewServiceBFrontHelloHandler creates a handler
func NewServiceBFrontHelloHandler(deps *module.Dependencies) *ServiceBFrontHelloHandler {
	handler := &ServiceBFrontHelloHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"multi", "helloB",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *ServiceBFrontHelloHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/multi/serviceB_f/hello",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/multi/serviceB_f/hello".
func (h *ServiceBFrontHelloHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	workflow := ServiceBFrontHelloEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		switch errValue := err.(type) {

		default:
			req.Logger.Warn("Workflow for endpoint returned error", zap.Error(errValue))
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		req.Logger.Warn("Unable to marshal response into json", zap.Error(err))
		res.SendErrorString(500, "Unexpected server error")
		return
	}
	res.WriteJSONBytes(200, cliRespHeaders, bytes)
}

// ServiceBFrontHelloEndpoint calls thrift client Multi.HelloB
type ServiceBFrontHelloEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w ServiceBFrontHelloEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (string, zanzibar.Header, error) {

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Multi.HelloB(
		ctx, clientHeaders,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request", zap.Error(errValue))
			// TODO(sindelar): Consider returning partial headers

			return "", nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertServiceBBackHelloClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertServiceBBackHelloClientResponse(in string) string {
	out := in

	return out
}
