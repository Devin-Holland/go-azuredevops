// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"azure-devops/5.1/client/agents"
	"azure-devops/5.1/client/counter_samples"
	"azure-devops/5.1/client/test_definitions"
	"azure-devops/5.1/client/test_drops"
	"azure-devops/5.1/client/test_runs"
)

// Default cloud load HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "vsclt.dev.azure.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new cloud load HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CloudLoad {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloud load HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CloudLoad {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloud load client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CloudLoad {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CloudLoad)
	cli.Transport = transport

	cli.Agents = agents.New(transport, formats)

	cli.CounterSamples = counter_samples.New(transport, formats)

	cli.TestDefinitions = test_definitions.New(transport, formats)

	cli.TestDrops = test_drops.New(transport, formats)

	cli.TestRuns = test_runs.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CloudLoad is a client for cloud load
type CloudLoad struct {
	Agents *agents.Client

	CounterSamples *counter_samples.Client

	TestDefinitions *test_definitions.Client

	TestDrops *test_drops.Client

	TestRuns *test_runs.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CloudLoad) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Agents.SetTransport(transport)

	c.CounterSamples.SetTransport(transport)

	c.TestDefinitions.SetTransport(transport)

	c.TestDrops.SetTransport(transport)

	c.TestRuns.SetTransport(transport)

}
