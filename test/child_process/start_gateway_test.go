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

package startGatewayTest

import (
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"syscall"
	"testing"

	"github.com/uber-go/zap"

	"github.com/uber/zanzibar/examples/example-gateway/clients"
	"github.com/uber/zanzibar/examples/example-gateway/endpoints"
	"github.com/uber/zanzibar/runtime"
)

var cachedServer *zanzibar.Gateway

func TestMain(m *testing.M) {
	if os.Getenv("GATEWAY_RUN_CHILD_PROCESS_TEST") != "" {
		listenOnSignals()

		code := m.Run()
		os.Exit(code)
	} else {
		os.Exit(0)
	}
}

func listenOnSignals() {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGUSR2)

	go func() {
		_ = <-sigs

		if cachedServer != nil {
			cachedServer.Close()
		}
	}()
}

func getProjectDir() string {
	goPath := os.Getenv("GOPATH")
	return path.Join(goPath, "src", "github.com", "uber", "zanzibar")
}

func TestStartGateway(t *testing.T) {
	testLogger := zap.New(
		zap.NewJSONEncoder(),
		zap.Output(os.Stderr),
	)

	config := zanzibar.NewStaticConfigOrDie([]string{
		filepath.Join(getProjectDir(), "config", "production.json"),
		filepath.Join(
			getProjectDir(),
			"examples",
			"example-gateway",
			"config",
			"production.json",
		),
		filepath.Join(os.Getenv("CONFIG_DIR"), "production.json"),
	}, nil)

	clients := clients.CreateClients(config)

	server, err := zanzibar.CreateGateway(config, &zanzibar.Options{
		Clients: clients,
	})
	if err != nil {
		testLogger.Error(
			"Failed to CreateGateway in TestStartGateway()",
			zap.String("error", err.Error()),
		)
		// ?
		return
	}

	cachedServer = server
	err = server.Bootstrap(endpoints.Register)
	if err != nil {
		testLogger.Error(
			"Failed to Bootstrap in TestStartGateway()",
			zap.String("error", err.Error()),
		)
		// ?
		return
	}

	server.Logger.Info("Started Gateway",
		zap.String("realAddr", server.RealAddr),
		zap.Object("config", config.InspectOrDie()),
	)

	server.Wait()
}
