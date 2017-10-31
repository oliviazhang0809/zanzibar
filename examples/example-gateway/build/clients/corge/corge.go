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

package corgeClient

import (
	"context"
	"errors"
	"strconv"
	"time"

	"go.uber.org/zap"

	tchannel "github.com/uber/tchannel-go"
	zanzibar "github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/corge/module"
	clientsCorgeCorge "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/corge/corge"
)

// Client defines corge client interface.
type Client interface {
	EchoString(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsCorgeCorge.Corge_EchoString_Args,
	) (string, map[string]string, error)
}

// NewClient returns a new TChannel client for service corge.
func NewClient(
	g *zanzibar.Gateway,
	deps *module.Dependencies,
) Client {
	serviceName := deps.Default.Config.MustGetString("clients.corge.serviceName")
	var routingKey string
	if deps.Default.Config.ContainsKey("clients.corge.routingKey") {
		routingKey = deps.Default.Config.MustGetString("clients.corge.routingKey")
	}
	sc := deps.Default.Channel.GetSubChannel(serviceName, tchannel.Isolated)

	ip := deps.Default.Config.MustGetString("sidecarRouter.default.tchannel.ip")
	port := deps.Default.Config.MustGetInt("sidecarRouter.default.tchannel.port")
	sc.Peers().Add(ip + ":" + strconv.Itoa(int(port)))

	timeout := time.Millisecond * time.Duration(
		deps.Default.Config.MustGetInt("clients.corge.timeout"),
	)
	timeoutPerAttempt := time.Millisecond * time.Duration(
		deps.Default.Config.MustGetInt("clients.corge.timeoutPerAttempt"),
	)

	methodNames := map[string]string{
		"Corge::echoString": "EchoString",
	}

	client := zanzibar.NewTChannelClient(
		deps.Default.Channel,
		deps.Default.Logger,
		deps.Default.Scope,
		&zanzibar.TChannelClientOption{
			ServiceName:       serviceName,
			ClientID:          "corge",
			MethodNames:       methodNames,
			Timeout:           timeout,
			TimeoutPerAttempt: timeoutPerAttempt,
			RoutingKey:        &routingKey,
		},
	)

	return &corgeClient{
		client: client,
	}
}

// corgeClient is the TChannel client for downstream service.
type corgeClient struct {
	client *zanzibar.TChannelClient
}

// EchoString is a client RPC call for method "Corge::echoString"
func (c *corgeClient) EchoString(
	ctx context.Context,
	reqHeaders map[string]string,
	args *clientsCorgeCorge.Corge_EchoString_Args,
) (string, map[string]string, error) {
	var result clientsCorgeCorge.Corge_EchoString_Result
	var resp string

	logger := c.client.Loggers["Corge::echoString"]

	success, respHeaders, err := c.client.Call(
		ctx, "Corge", "echoString", reqHeaders, args, &result,
	)

	if err == nil && !success {
		switch {
		default:
			err = errors.New("corgeClient received no result or unknown exception for EchoString")
		}
	}
	if err != nil {
		logger.Warn("TChannel client call returned error", zap.Error(err))
		return resp, nil, err
	}

	resp, err = clientsCorgeCorge.Corge_EchoString_Helper.UnwrapResponse(&result)
	if err != nil {
		logger.Warn("Unable to unwrap client response", zap.Error(err))
	}
	return resp, respHeaders, err
}
