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

package googlenowendpoint

import (
	"context"
	"fmt"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	clientsGooglenowGooglenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/googlenow/googlenow"
	endpointsGooglenowGooglenow "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/googlenow/googlenow"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/module"
)

// GoogleNowAddCredentialsHandler is the handler for "/googlenow/add-credentials"
type GoogleNowAddCredentialsHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewGoogleNowAddCredentialsHandler creates a handler
func NewGoogleNowAddCredentialsHandler(deps *module.Dependencies) *GoogleNowAddCredentialsHandler {
	handler := &GoogleNowAddCredentialsHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope, deps.Default.Tracer,
		"googlenow", "addCredentials",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *GoogleNowAddCredentialsHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"POST", "/googlenow/add-credentials",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/googlenow/add-credentials".
func (h *GoogleNowAddCredentialsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	if !req.CheckHeaders([]string{"x-uuid", "x-token"}) {
		return
	}
	var requestBody endpointsGooglenowGooglenow.GoogleNow_AddCredentials_Args
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	// log endpoint request to downstream services
	zfields := []zapcore.Field{
		zap.String("endpoint", h.endpoint.EndpointName),
	}

	// TODO: potential perf issue, use zap.Object lazy serialization
	zfields = append(zfields, zap.String("body", fmt.Sprintf("%#v", requestBody)))
	var headerOk bool
	var headerValue string
	headerValue, headerOk = req.Header.Get("X-Uuid")
	if headerOk {
		zfields = append(zfields, zap.String("X-Uuid", headerValue))
	}
	req.Logger.Debug("Endpoint request to downstream", zfields...)

	workflow := GoogleNowAddCredentialsEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	cliRespHeaders, err := workflow.Handle(ctx, req.Header, &requestBody)
	if err != nil {
		res.SendError(500, "Unexpected server error", err)
		return

	}
	// TODO(sindelar): implement check headers on response

	res.WriteJSONBytes(202, cliRespHeaders, nil)
}

// GoogleNowAddCredentialsEndpoint calls thrift client GoogleNow.AddCredentials
type GoogleNowAddCredentialsEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w GoogleNowAddCredentialsEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsGooglenowGooglenow.GoogleNow_AddCredentials_Args,
) (zanzibar.Header, error) {
	clientRequest := convertToAddCredentialsClientRequest(r)

	clientHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientHeaders["X-Uuid"] = h
	}

	cliRespHeaders, err := w.Clients.GoogleNow.AddCredentials(
		ctx, clientHeaders, clientRequest,
	)

	if err != nil {
		switch errValue := err.(type) {

		default:
			w.Logger.Warn("Could not make client request",
				zap.Error(errValue),
				zap.String("client", "GoogleNow"),
			)

			// TODO(sindelar): Consider returning partial headers

			return nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	resHeaders.Set("X-Uuid", cliRespHeaders["X-Uuid"])

	return resHeaders, nil
}

func convertToAddCredentialsClientRequest(in *endpointsGooglenowGooglenow.GoogleNow_AddCredentials_Args) *clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args {
	out := &clientsGooglenowGooglenow.GoogleNowService_AddCredentials_Args{}

	out.AuthCode = string(in.AuthCode)

	return out
}
