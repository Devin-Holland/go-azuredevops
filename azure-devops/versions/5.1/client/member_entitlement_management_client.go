// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"azure-devops/5.1/client/group_entitlements"
	"azure-devops/5.1/client/members"
	"azure-devops/5.1/client/user_entitlement_summary"
	"azure-devops/5.1/client/user_entitlements"
)

// Default member entitlement management HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "vsaex.dev.azure.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new member entitlement management HTTP client.
func NewHTTPClient(formats strfmt.Registry) *MemberEntitlementManagement {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new member entitlement management HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *MemberEntitlementManagement {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new member entitlement management client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *MemberEntitlementManagement {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(MemberEntitlementManagement)
	cli.Transport = transport

	cli.GroupEntitlements = group_entitlements.New(transport, formats)

	cli.Members = members.New(transport, formats)

	cli.UserEntitlementSummary = user_entitlement_summary.New(transport, formats)

	cli.UserEntitlements = user_entitlements.New(transport, formats)

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

// MemberEntitlementManagement is a client for member entitlement management
type MemberEntitlementManagement struct {
	GroupEntitlements *group_entitlements.Client

	Members *members.Client

	UserEntitlementSummary *user_entitlement_summary.Client

	UserEntitlements *user_entitlements.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *MemberEntitlementManagement) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.GroupEntitlements.SetTransport(transport)

	c.Members.SetTransport(transport)

	c.UserEntitlementSummary.SetTransport(transport)

	c.UserEntitlements.SetTransport(transport)

}
