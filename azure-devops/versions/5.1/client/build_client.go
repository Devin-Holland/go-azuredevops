// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"azure-devops/5.1/client/authorizedresources"
	"azure-devops/5.1/client/badge"
	"azure-devops/5.1/client/builds"
	"azure-devops/5.1/client/definitions"
	"azure-devops/5.1/client/folders"
	"azure-devops/5.1/client/metrics"
	"azure-devops/5.1/client/properties"
	"azure-devops/5.1/client/resources"
	"azure-devops/5.1/client/settings"
	"azure-devops/5.1/client/source_providers"
	"azure-devops/5.1/client/status"
	"azure-devops/5.1/client/tags"
	"azure-devops/5.1/client/templates"
)

// Default build HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "dev.azure.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new build HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Build {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new build HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Build {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new build client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Build {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Build)
	cli.Transport = transport

	cli.Authorizedresources = authorizedresources.New(transport, formats)

	cli.Badge = badge.New(transport, formats)

	cli.Builds = builds.New(transport, formats)

	cli.Definitions = definitions.New(transport, formats)

	cli.Folders = folders.New(transport, formats)

	cli.Metrics = metrics.New(transport, formats)

	cli.Properties = properties.New(transport, formats)

	cli.Resources = resources.New(transport, formats)

	cli.Settings = settings.New(transport, formats)

	cli.SourceProviders = source_providers.New(transport, formats)

	cli.Status = status.New(transport, formats)

	cli.Tags = tags.New(transport, formats)

	cli.Templates = templates.New(transport, formats)

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

// Build is a client for build
type Build struct {
	Authorizedresources *authorizedresources.Client

	Badge *badge.Client

	Builds *builds.Client

	Definitions *definitions.Client

	Folders *folders.Client

	Metrics *metrics.Client

	Properties *properties.Client

	Resources *resources.Client

	Settings *settings.Client

	SourceProviders *source_providers.Client

	Status *status.Client

	Tags *tags.Client

	Templates *templates.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Build) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Authorizedresources.SetTransport(transport)

	c.Badge.SetTransport(transport)

	c.Builds.SetTransport(transport)

	c.Definitions.SetTransport(transport)

	c.Folders.SetTransport(transport)

	c.Metrics.SetTransport(transport)

	c.Properties.SetTransport(transport)

	c.Resources.SetTransport(transport)

	c.Settings.SetTransport(transport)

	c.SourceProviders.SetTransport(transport)

	c.Status.SetTransport(transport)

	c.Tags.SetTransport(transport)

	c.Templates.SetTransport(transport)

}
