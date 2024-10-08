// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new provisioning API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new provisioning API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new provisioning API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for provisioning API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithAccept allows the client to force the Accept header
// to negotiate a specific Producer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithAccept(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ProducesMediaTypes = []string{mime}
	}
}

// WithAcceptApplicationJSON sets the Accept header to "application/json".
func WithAcceptApplicationJSON(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"application/json"}
}

// WithAcceptTextPlain sets the Accept header to "text/plain".
func WithAcceptTextPlain(r *runtime.ClientOperation) {
	r.ProducesMediaTypes = []string{"text/plain"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	Get5BridgeLocation(params *Get5BridgeLocationParams, opts ...ClientOption) (*Get5BridgeLocationOK, error)

	Get5Bridges(params *Get5BridgesParams, opts ...ClientOption) (*Get5BridgesOK, error)

	Get5GatewayLocation(params *Get5GatewayLocationParams, opts ...ClientOption) (*Get5GatewayLocationOK, error)

	Get5Gateways(params *Get5GatewaysParams, opts ...ClientOption) (*Get5GatewaysOK, error)

	Get5OpenvpnCert(params *Get5OpenvpnCertParams, opts ...ClientOption) (*Get5OpenvpnCertOK, error)

	Get5OpenvpnConfig(params *Get5OpenvpnConfigParams, opts ...ClientOption) (*Get5OpenvpnConfigOK, error)

	Get5Service(params *Get5ServiceParams, opts ...ClientOption) (*Get5ServiceOK, error)

	GetAutoconf(params *GetAutoconfParams, opts ...ClientOption) (*GetAutoconfOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Get5BridgeLocation gets bridges

fetch bridges by location
*/
func (a *Client) Get5BridgeLocation(params *Get5BridgeLocationParams, opts ...ClientOption) (*Get5BridgeLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet5BridgeLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get5BridgeLocation",
		Method:             "GET",
		PathPattern:        "/5/bridge/{location}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Get5BridgeLocationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get5BridgeLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get5BridgeLocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get5Bridges gets all bridges

Fetch all bridges. This is an optional API endpoint for compatibility with vpnweb, but do not count on all the providers to have it enabled since it makes it easier to enumerate resources. On the other hand, if the service has "open" VPN endpoints, they can enumerate them here freely. Bridges, however, should be more restricted as a general rule.
*/
func (a *Client) Get5Bridges(params *Get5BridgesParams, opts ...ClientOption) (*Get5BridgesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet5BridgesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get5Bridges",
		Method:             "GET",
		PathPattern:        "/5/bridges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Get5BridgesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get5BridgesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get5Bridges: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get5GatewayLocation gets gateways by location

fetch random gateways for a given location
*/
func (a *Client) Get5GatewayLocation(params *Get5GatewayLocationParams, opts ...ClientOption) (*Get5GatewayLocationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet5GatewayLocationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get5GatewayLocation",
		Method:             "GET",
		PathPattern:        "/5/gateway/{location}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Get5GatewayLocationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get5GatewayLocationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get5GatewayLocation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get5Gateways gets all gateways

Fetch all gateways. This is an optional API endpoint for compatibility with vpnweb, but do not count on all the providers to have it enabled since it makes it easier to enumerate resources. On the other hand, if the service has "open" VPN endpoints, they can enumerate them here freely. Bridges, however, should be more restricted as a general rule.
*/
func (a *Client) Get5Gateways(params *Get5GatewaysParams, opts ...ClientOption) (*Get5GatewaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet5GatewaysParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get5Gateways",
		Method:             "GET",
		PathPattern:        "/5/gateways",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Get5GatewaysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get5GatewaysOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get5Gateways: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get5OpenvpnCert gets openvpn cert

Fetch a new key and cert.
*/
func (a *Client) Get5OpenvpnCert(params *Get5OpenvpnCertParams, opts ...ClientOption) (*Get5OpenvpnCertOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet5OpenvpnCertParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get5OpenvpnCert",
		Method:             "GET",
		PathPattern:        "/5/openvpn/cert",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Get5OpenvpnCertReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get5OpenvpnCertOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get5OpenvpnCert: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get5OpenvpnConfig fetches open v p n config file

fetch a working config file for OpenVPN service.
*/
func (a *Client) Get5OpenvpnConfig(params *Get5OpenvpnConfigParams, opts ...ClientOption) (*Get5OpenvpnConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet5OpenvpnConfigParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get5OpenvpnConfig",
		Method:             "GET",
		PathPattern:        "/5/openvpn/config",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Get5OpenvpnConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get5OpenvpnConfigOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get5OpenvpnConfig: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
Get5Service gets service info

fetch service information, including location and common tunnel config
*/
func (a *Client) Get5Service(params *Get5ServiceParams, opts ...ClientOption) (*Get5ServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGet5ServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get5Service",
		Method:             "GET",
		PathPattern:        "/5/service",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &Get5ServiceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Get5ServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for Get5Service: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetAutoconf fetches open v p n config file

fetch a working config file for OpenVPN service.
*/
func (a *Client) GetAutoconf(params *GetAutoconfParams, opts ...ClientOption) (*GetAutoconfOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAutoconfParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAutoconf",
		Method:             "GET",
		PathPattern:        "/autoconf",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAutoconfReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAutoconfOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAutoconf: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
