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
	resourcespb "github.com/kritzware/google-ads-go/resources"
	servicespb "github.com/kritzware/google-ads-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newKeywordThemeConstantClientHook clientHook

// KeywordThemeConstantCallOptions contains the retry settings for each method of KeywordThemeConstantClient.
type KeywordThemeConstantCallOptions struct {
	GetKeywordThemeConstant []gax.CallOption
	SuggestKeywordThemeConstants []gax.CallOption
}

func defaultKeywordThemeConstantGRPCClientOptions() []option.ClientOption {
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

func defaultKeywordThemeConstantCallOptions() *KeywordThemeConstantCallOptions {
	return &KeywordThemeConstantCallOptions{
		GetKeywordThemeConstant: []gax.CallOption{
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
		SuggestKeywordThemeConstants: []gax.CallOption{
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

// internalKeywordThemeConstantClient is an interface that defines the methods availaible from Google Ads API.
type internalKeywordThemeConstantClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetKeywordThemeConstant(context.Context, *servicespb.GetKeywordThemeConstantRequest, ...gax.CallOption) (*resourcespb.KeywordThemeConstant, error)
	SuggestKeywordThemeConstants(context.Context, *servicespb.SuggestKeywordThemeConstantsRequest, ...gax.CallOption) (*servicespb.SuggestKeywordThemeConstantsResponse, error)
}

// KeywordThemeConstantClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to fetch Smart Campaign keyword themes.
type KeywordThemeConstantClient struct {
	// The internal transport-dependent client.
	internalClient internalKeywordThemeConstantClient

	// The call options for this service.
	CallOptions *KeywordThemeConstantCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *KeywordThemeConstantClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *KeywordThemeConstantClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *KeywordThemeConstantClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetKeywordThemeConstant returns the requested keyword theme constant in full detail.
func (c *KeywordThemeConstantClient) GetKeywordThemeConstant(ctx context.Context, req *servicespb.GetKeywordThemeConstantRequest, opts ...gax.CallOption) (*resourcespb.KeywordThemeConstant, error) {
	return c.internalClient.GetKeywordThemeConstant(ctx, req, opts...)
}

// SuggestKeywordThemeConstants returns KeywordThemeConstant suggestions by keyword themes.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *KeywordThemeConstantClient) SuggestKeywordThemeConstants(ctx context.Context, req *servicespb.SuggestKeywordThemeConstantsRequest, opts ...gax.CallOption) (*servicespb.SuggestKeywordThemeConstantsResponse, error) {
	return c.internalClient.SuggestKeywordThemeConstants(ctx, req, opts...)
}

// keywordThemeConstantGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type keywordThemeConstantGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing KeywordThemeConstantClient
	CallOptions **KeywordThemeConstantCallOptions

	// The gRPC API client.
	keywordThemeConstantClient servicespb.KeywordThemeConstantServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewKeywordThemeConstantClient creates a new keyword theme constant service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to fetch Smart Campaign keyword themes.
func NewKeywordThemeConstantClient(ctx context.Context, opts ...option.ClientOption) (*KeywordThemeConstantClient, error) {
	clientOpts := defaultKeywordThemeConstantGRPCClientOptions()
	if newKeywordThemeConstantClientHook != nil {
		hookOpts, err := newKeywordThemeConstantClientHook(ctx, clientHookParams{})
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
	client := KeywordThemeConstantClient{CallOptions: defaultKeywordThemeConstantCallOptions()}

	c := &keywordThemeConstantGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		keywordThemeConstantClient: servicespb.NewKeywordThemeConstantServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *keywordThemeConstantGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *keywordThemeConstantGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *keywordThemeConstantGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *keywordThemeConstantGRPCClient) GetKeywordThemeConstant(ctx context.Context, req *servicespb.GetKeywordThemeConstantRequest, opts ...gax.CallOption) (*resourcespb.KeywordThemeConstant, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetKeywordThemeConstant[0:len((*c.CallOptions).GetKeywordThemeConstant):len((*c.CallOptions).GetKeywordThemeConstant)], opts...)
	var resp *resourcespb.KeywordThemeConstant
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordThemeConstantClient.GetKeywordThemeConstant(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *keywordThemeConstantGRPCClient) SuggestKeywordThemeConstants(ctx context.Context, req *servicespb.SuggestKeywordThemeConstantsRequest, opts ...gax.CallOption) (*servicespb.SuggestKeywordThemeConstantsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).SuggestKeywordThemeConstants[0:len((*c.CallOptions).SuggestKeywordThemeConstants):len((*c.CallOptions).SuggestKeywordThemeConstants)], opts...)
	var resp *servicespb.SuggestKeywordThemeConstantsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.keywordThemeConstantClient.SuggestKeywordThemeConstants(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
