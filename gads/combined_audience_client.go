// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package gads

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	resourcespb "github.com/opteo/google-ads-go/resources"
	servicespb "github.com/opteo/google-ads-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newCombinedAudienceClientHook clientHook

// CombinedAudienceCallOptions contains the retry settings for each method of CombinedAudienceClient.
type CombinedAudienceCallOptions struct {
	GetCombinedAudience []gax.CallOption
}

func defaultCombinedAudienceGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("googleads.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("googleads.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://googleads.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		option.WithGRPCDialOption(grpc.WithDisableServiceConfig()),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
		grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCombinedAudienceCallOptions() *CombinedAudienceCallOptions {
	return &CombinedAudienceCallOptions{
		GetCombinedAudience: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
					codes.DeadlineExceeded,
				}, gax.Backoff{
					Initial:    5000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalCombinedAudienceClient is an interface that defines the methods availaible from Google Ads API.
type internalCombinedAudienceClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetCombinedAudience(context.Context, *servicespb.GetCombinedAudienceRequest, ...gax.CallOption) (*resourcespb.CombinedAudience, error)
}

// CombinedAudienceClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage combined audiences. This service can be used to list all
// your combined audiences with metadata, but won’t show the structure and
// components of the combined audience.
type CombinedAudienceClient struct {
	// The internal transport-dependent client.
	internalClient internalCombinedAudienceClient

	// The call options for this service.
	CallOptions *CombinedAudienceCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *CombinedAudienceClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *CombinedAudienceClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *CombinedAudienceClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetCombinedAudience returns the requested combined audience in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *CombinedAudienceClient) GetCombinedAudience(ctx context.Context, req *servicespb.GetCombinedAudienceRequest, opts ...gax.CallOption) (*resourcespb.CombinedAudience, error) {
	return c.internalClient.GetCombinedAudience(ctx, req, opts...)
}

// combinedAudienceGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type combinedAudienceGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing CombinedAudienceClient
	CallOptions **CombinedAudienceCallOptions

	// The gRPC API client.
	combinedAudienceClient servicespb.CombinedAudienceServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewCombinedAudienceClient creates a new combined audience service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage combined audiences. This service can be used to list all
// your combined audiences with metadata, but won’t show the structure and
// components of the combined audience.
func NewCombinedAudienceClient(ctx context.Context, opts ...option.ClientOption) (*CombinedAudienceClient, error) {
	clientOpts := defaultCombinedAudienceGRPCClientOptions()
	if newCombinedAudienceClientHook != nil {
		hookOpts, err := newCombinedAudienceClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := CombinedAudienceClient{CallOptions: defaultCombinedAudienceCallOptions()}

	c := &combinedAudienceGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		combinedAudienceClient: servicespb.NewCombinedAudienceServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *combinedAudienceGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *combinedAudienceGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *combinedAudienceGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *combinedAudienceGRPCClient) GetCombinedAudience(ctx context.Context, req *servicespb.GetCombinedAudienceRequest, opts ...gax.CallOption) (*resourcespb.CombinedAudience, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetCombinedAudience[0:len((*c.CallOptions).GetCombinedAudience):len((*c.CallOptions).GetCombinedAudience)], opts...)
	var resp *resourcespb.CombinedAudience
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.combinedAudienceClient.GetCombinedAudience(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
