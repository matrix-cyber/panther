// Code generated by go-swagger; DO NOT EDIT.

// Panther is a scalable, powerful, cloud-native SIEM written in Golang/React.
// Copyright (C) 2020 Panther Labs Inc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
//

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListRemediations(params *ListRemediationsParams) (*ListRemediationsOK, error)

	RemediateResource(params *RemediateResourceParams) (*RemediateResourceOK, error)

	RemediateResourceAsync(params *RemediateResourceAsyncParams) (*RemediateResourceAsyncOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListRemediations retrieves available remediations
*/
func (a *Client) ListRemediations(params *ListRemediationsParams) (*ListRemediationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRemediationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListRemediations",
		Method:             "GET",
		PathPattern:        "/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListRemediationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRemediationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ListRemediations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemediateResource synchronouslies remediate resource for an account
*/
func (a *Client) RemediateResource(params *RemediateResourceParams) (*RemediateResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemediateResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemediateResource",
		Method:             "POST",
		PathPattern:        "/remediate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemediateResourceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemediateResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RemediateResource: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RemediateResourceAsync asynchronouslies remediate resource for an account
*/
func (a *Client) RemediateResourceAsync(params *RemediateResourceAsyncParams) (*RemediateResourceAsyncOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRemediateResourceAsyncParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "RemediateResourceAsync",
		Method:             "POST",
		PathPattern:        "/remediateasync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RemediateResourceAsyncReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RemediateResourceAsyncOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for RemediateResourceAsync: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
