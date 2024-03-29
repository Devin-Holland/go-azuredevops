// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"azure-devops/5.1/client/artifact_details"
	"azure-devops/5.1/client/change_tracking"
	"azure-devops/5.1/client/feed_management"
	"azure-devops/5.1/client/recycle_bin"
	"azure-devops/5.1/client/retention_policies"
	"azure-devops/5.1/client/service_settings"
)

// Default feed HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "feeds.dev.azure.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new feed HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Feed {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new feed HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Feed {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new feed client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Feed {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Feed)
	cli.Transport = transport

	cli.ArtifactDetails = artifact_details.New(transport, formats)

	cli.ChangeTracking = change_tracking.New(transport, formats)

	cli.FeedManagement = feed_management.New(transport, formats)

	cli.RecycleBin = recycle_bin.New(transport, formats)

	cli.RetentionPolicies = retention_policies.New(transport, formats)

	cli.ServiceSettings = service_settings.New(transport, formats)

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

// Feed is a client for feed
type Feed struct {
	ArtifactDetails *artifact_details.Client

	ChangeTracking *change_tracking.Client

	FeedManagement *feed_management.Client

	RecycleBin *recycle_bin.Client

	RetentionPolicies *retention_policies.Client

	ServiceSettings *service_settings.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Feed) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.ArtifactDetails.SetTransport(transport)

	c.ChangeTracking.SetTransport(transport)

	c.FeedManagement.SetTransport(transport)

	c.RecycleBin.SetTransport(transport)

	c.RetentionPolicies.SetTransport(transport)

	c.ServiceSettings.SetTransport(transport)

}
