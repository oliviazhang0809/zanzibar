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

package baztchannelEndpoint

import (
	"context"

	"github.com/pkg/errors"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/baz/module"
	endpointsTchannelBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/tchannel/baz/baz"
	customBaz "github.com/uber/zanzibar/examples/example-gateway/endpoints/tchannel/baz"
)

// NewSimpleServiceCallHandler creates a handler to be registered with a thrift server.
func NewSimpleServiceCallHandler(deps *module.Dependencies) *SimpleServiceCallHandler {
	handler := &SimpleServiceCallHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewTChannelEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bazTChannel", "call", "SimpleService::Call",
		handler,
	)
	return handler
}

// SimpleServiceCallHandler is the handler for "SimpleService::Call".
type SimpleServiceCallHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.TChannelEndpoint
}

// Register adds the tchannel handler to the gateway's tchannel router
func (h *SimpleServiceCallHandler) Register(g *zanzibar.Gateway) error {
	g.TChannelRouter.Register(h.endpoint)
	// TODO: Register should return an error for route conflicts
	return nil
}

// Handle handles RPC call of "SimpleService::Call".
func (h *SimpleServiceCallHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	wfReqHeaders := zanzibar.ServerTChannelHeader(reqHeaders)
	if err := wfReqHeaders.Ensure([]string{"x-uuid", "x-token"}, h.endpoint.Logger); err != nil {
		return false, nil, nil, errors.Wrapf(
			err, "%s.%s (%s) missing request headers",
			h.endpoint.EndpointID, h.endpoint.HandlerID, h.endpoint.Method,
		)
	}

	var res endpointsTchannelBazBaz.SimpleService_Call_Result

	var req endpointsTchannelBazBaz.SimpleService_Call_Args
	if err := req.FromWire(*wireValue); err != nil {
		h.endpoint.Logger.Warn("Error converting request from wire", zap.Error(err))
		return false, nil, nil, errors.Wrapf(
			err, "Error converting %s.%s (%s) request from wire",
			h.endpoint.EndpointID, h.endpoint.HandlerID, h.endpoint.Method,
		)
	}
	workflow := customBaz.CallEndpoint{
		Clients: h.Clients,
		Logger:  h.endpoint.Logger,
	}

	wfResHeaders, err := workflow.Handle(ctx, wfReqHeaders, &req)

	resHeaders := map[string]string{}
	for _, key := range wfResHeaders.Keys() {
		resHeaders[key], _ = wfResHeaders.Get(key)
	}

	if err != nil {
		switch v := err.(type) {
		case *endpointsTchannelBazBaz.AuthErr:
			h.endpoint.Logger.Warn(
				"Handler returned non-nil error type *endpointsTchannelBazBaz.AuthErr but nil value",
				zap.Error(err),
			)
			if v == nil {
				return false, nil, resHeaders, errors.Errorf(
					"%s.%s (%s) handler returned non-nil error type *endpointsTchannelBazBaz.AuthErr but nil value",
					h.endpoint.EndpointID, h.endpoint.HandlerID, h.endpoint.Method,
				)
			}
			res.AuthErr = v
		default:
			h.endpoint.Logger.Warn("Handler returned error", zap.Error(err))
			return false, nil, resHeaders, errors.Wrapf(
				err, "%s.%s (%s) handler returned error",
				h.endpoint.EndpointID, h.endpoint.HandlerID, h.endpoint.Method,
			)
		}
	}
	if err := wfResHeaders.Ensure([]string{"some-res-header"}, h.endpoint.Logger); err != nil {
		return false, nil, nil, errors.Wrapf(
			err, "%s.%s (%s) missing response headers",
			h.endpoint.EndpointID, h.endpoint.HandlerID, h.endpoint.Method,
		)
	}

	return err == nil, &res, resHeaders, nil
}
