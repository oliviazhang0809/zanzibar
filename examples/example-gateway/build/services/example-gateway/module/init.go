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

package module

import (
	barClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/bar"
	barClientModule "github.com/uber/zanzibar/examples/example-gateway/build/clients/bar/module"
	bazClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz"
	bazClientModule "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz/module"
	contactsClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts"
	contactsClientModule "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts/module"
	googlenowClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now"
	googlenowClientModule "github.com/uber/zanzibar/examples/example-gateway/build/clients/google-now/module"
	multiClientGenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/multi"
	multiClientModule "github.com/uber/zanzibar/examples/example-gateway/build/clients/multi/module"
	barEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar"
	barEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
	bazEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz"
	bazEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/baz/module"
	contactsEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts"
	contactsEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/contacts/module"
	googlenowEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow"
	googlenowEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/googlenow/module"
	multiEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi"
	multiEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/multi/module"
	baztchannelEndpointGenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/baz"
	baztchannelEndpointModule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/baz/module"
	exampleMiddlewareGenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/example"
	exampleMiddlewareModule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/example/module"

	zanzibar "github.com/uber/zanzibar/runtime"
)

// DependenciesTree contains all deps for this service.
type DependenciesTree struct {
	Client     *ClientDependenciesNodes
	Middleware *MiddlewareDependenciesNodes
	Endpoint   *EndpointDependenciesNodes
}

// ClientDependenciesNodes contains client dependencies
type ClientDependenciesNodes struct {
	Bar       barClientGenerated.Client
	Baz       bazClientGenerated.Client
	Contacts  contactsClientGenerated.Client
	GoogleNow googlenowClientGenerated.Client
	Multi     multiClientGenerated.Client
}

// MiddlewareDependenciesNodes contains middleware dependencies
type MiddlewareDependenciesNodes struct {
	Example exampleMiddlewareGenerated.Middleware
}

// EndpointDependenciesNodes contains endpoint dependencies
type EndpointDependenciesNodes struct {
	Bar         barEndpointGenerated.Endpoint
	Baz         bazEndpointGenerated.Endpoint
	BazTChannel baztchannelEndpointGenerated.Endpoint
	Contacts    contactsEndpointGenerated.Endpoint
	Googlenow   googlenowEndpointGenerated.Endpoint
	Multi       multiEndpointGenerated.Endpoint
}

// InitializeDependencies fully initializes all dependencies in the dep tree
// for the example-gateway service
func InitializeDependencies(
	g *zanzibar.Gateway,
) (*DependenciesTree, *Dependencies) {
	tree := &DependenciesTree{}

	initializedDefaultDependencies := &zanzibar.DefaultDependencies{
		Logger:  g.Logger,
		Scope:   g.AllHostScope,
		Config:  g.Config,
		Channel: g.Channel,
	}

	initializedClientDependencies := &ClientDependenciesNodes{}
	tree.Client = initializedClientDependencies
	initializedClientDependencies.Bar = barClientGenerated.NewClient(&barClientModule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Baz = bazClientGenerated.NewClient(&bazClientModule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Contacts = contactsClientGenerated.NewClient(&contactsClientModule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.GoogleNow = googlenowClientGenerated.NewClient(&googlenowClientModule.Dependencies{
		Default: initializedDefaultDependencies,
	})
	initializedClientDependencies.Multi = multiClientGenerated.NewClient(&multiClientModule.Dependencies{
		Default: initializedDefaultDependencies,
	})

	initializedMiddlewareDependencies := &MiddlewareDependenciesNodes{}
	tree.Middleware = initializedMiddlewareDependencies
	initializedMiddlewareDependencies.Example = exampleMiddlewareGenerated.NewMiddleware(&exampleMiddlewareModule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &exampleMiddlewareModule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})

	initializedEndpointDependencies := &EndpointDependenciesNodes{}
	tree.Endpoint = initializedEndpointDependencies
	initializedEndpointDependencies.Bar = barEndpointGenerated.NewEndpoint(&barEndpointModule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &barEndpointModule.ClientDependencies{
			Bar: initializedClientDependencies.Bar,
		},
		Middleware: &barEndpointModule.MiddlewareDependencies{
			Example: initializedMiddlewareDependencies.Example,
		},
	})
	initializedEndpointDependencies.Baz = bazEndpointGenerated.NewEndpoint(&bazEndpointModule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &bazEndpointModule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedEndpointDependencies.BazTChannel = baztchannelEndpointGenerated.NewEndpoint(&baztchannelEndpointModule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &baztchannelEndpointModule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedEndpointDependencies.Contacts = contactsEndpointGenerated.NewEndpoint(&contactsEndpointModule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &contactsEndpointModule.ClientDependencies{
			Contacts: initializedClientDependencies.Contacts,
		},
	})
	initializedEndpointDependencies.Googlenow = googlenowEndpointGenerated.NewEndpoint(&googlenowEndpointModule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &googlenowEndpointModule.ClientDependencies{
			GoogleNow: initializedClientDependencies.GoogleNow,
		},
	})
	initializedEndpointDependencies.Multi = multiEndpointGenerated.NewEndpoint(&multiEndpointModule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &multiEndpointModule.ClientDependencies{
			Multi: initializedClientDependencies.Multi,
		},
	})

	return tree, &Dependencies{
		Default: initializedDefaultDependencies,
		Endpoint: &EndpointDependencies{
			Bar:         initializedEndpointDependencies.Bar,
			Baz:         initializedEndpointDependencies.Baz,
			BazTChannel: initializedEndpointDependencies.BazTChannel,
			Contacts:    initializedEndpointDependencies.Contacts,
			Googlenow:   initializedEndpointDependencies.Googlenow,
			Multi:       initializedEndpointDependencies.Multi,
		},
	}
}
