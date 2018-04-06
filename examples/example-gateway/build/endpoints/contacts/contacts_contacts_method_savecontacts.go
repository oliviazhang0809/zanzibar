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

package contactsendpoint

import (
	"context"
	"encoding/json"
	"fmt"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	endpointsContactsContacts "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/contacts/contacts"
	customContacts "github.com/uber/zanzibar/examples/example-gateway/endpoints/contacts"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts/module"
)

// ContactsSaveContactsHandler is the handler for "/contacts/:userUUID/contacts"
type ContactsSaveContactsHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewContactsSaveContactsHandler creates a handler
func NewContactsSaveContactsHandler(deps *module.Dependencies) *ContactsSaveContactsHandler {
	handler := &ContactsSaveContactsHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"contacts", "saveContacts",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *ContactsSaveContactsHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"POST", "/contacts/:userUUID/contacts",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/contacts/:userUUID/contacts".
func (h *ContactsSaveContactsHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {
	var requestBody endpointsContactsContacts.SaveContactsRequest
	if ok := req.ReadAndUnmarshalBody(&requestBody); !ok {
		return
	}

	requestBody.UserUUID = req.Params.ByName("userUUID")

	// log endpoint request to downstream services
	if ce := req.Logger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		zfields = append(zfields, zap.String("body", fmt.Sprintf("%s", req.GetRawBody())))
		for _, k := range req.Header.Keys() {
			if val, ok := req.Header.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		req.Logger.Debug("endpoint request to downstream", zfields...)
	}

	w := customContacts.NewContactsSaveContactsWorkflow(h.Clients, req.Logger)

	response, cliRespHeaders, err := w.Handle(ctx, req.Header, &requestBody)

	// log downstream response to endpoint
	if ce := req.Logger.Check(zapcore.DebugLevel, "stub"); ce != nil {
		zfields := []zapcore.Field{
			zap.String("endpoint", h.endpoint.EndpointName),
		}
		if response != nil {
			if body, err := json.Marshal(response); err == nil {
				zfields = append(zfields, zap.String("body", fmt.Sprintf("%s", body)))
			}
		}
		for _, k := range cliRespHeaders.Keys() {
			if val, ok := cliRespHeaders.Get(k); ok {
				zfields = append(zfields, zap.String(k, val))
			}
		}
		req.Logger.Debug("downstream service response", zfields...)
	}

	if err != nil {
		res.SendError(500, "Unexpected server error", err)
		return

	}

	res.WriteJSON(202, cliRespHeaders, response)
}
