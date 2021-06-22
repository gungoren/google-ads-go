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

var newFeedItemClientHook clientHook

// FeedItemCallOptions contains the retry settings for each method of FeedItemClient.
type FeedItemCallOptions struct {
	GetFeedItem []gax.CallOption
	MutateFeedItems []gax.CallOption
}

func defaultFeedItemGRPCClientOptions() []option.ClientOption {
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

func defaultFeedItemCallOptions() *FeedItemCallOptions {
	return &FeedItemCallOptions{
		GetFeedItem: []gax.CallOption{
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
		MutateFeedItems: []gax.CallOption{
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

// internalFeedItemClient is an interface that defines the methods availaible from Google Ads API.
type internalFeedItemClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	GetFeedItem(context.Context, *servicespb.GetFeedItemRequest, ...gax.CallOption) (*resourcespb.FeedItem, error)
	MutateFeedItems(context.Context, *servicespb.MutateFeedItemsRequest, ...gax.CallOption) (*servicespb.MutateFeedItemsResponse, error)
}

// FeedItemClient is a client for interacting with Google Ads API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service to manage feed items.
type FeedItemClient struct {
	// The internal transport-dependent client.
	internalClient internalFeedItemClient

	// The call options for this service.
	CallOptions *FeedItemCallOptions

}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *FeedItemClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *FeedItemClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *FeedItemClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// GetFeedItem returns the requested feed item in full detail.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// HeaderError (at )
// InternalError (at )
// QuotaError (at )
// RequestError (at )
func (c *FeedItemClient) GetFeedItem(ctx context.Context, req *servicespb.GetFeedItemRequest, opts ...gax.CallOption) (*resourcespb.FeedItem, error) {
	return c.internalClient.GetFeedItem(ctx, req, opts...)
}

// MutateFeedItems creates, updates, or removes feed items. Operation statuses are
// returned.
//
// List of thrown errors:
// AuthenticationError (at )
// AuthorizationError (at )
// CollectionSizeError (at )
// CriterionError (at )
// DatabaseError (at )
// DateError (at )
// DistinctError (at )
// FeedItemError (at )
// FieldError (at )
// FieldMaskError (at )
// HeaderError (at )
// IdError (at )
// InternalError (at )
// ListOperationError (at )
// MutateError (at )
// NotEmptyError (at )
// NullError (at )
// OperatorError (at )
// QuotaError (at )
// RangeError (at )
// RequestError (at )
// SizeLimitError (at )
// StringFormatError (at )
// StringLengthError (at )
// UrlFieldError (at )
func (c *FeedItemClient) MutateFeedItems(ctx context.Context, req *servicespb.MutateFeedItemsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemsResponse, error) {
	return c.internalClient.MutateFeedItems(ctx, req, opts...)
}

// feedItemGRPCClient is a client for interacting with Google Ads API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type feedItemGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing FeedItemClient
	CallOptions **FeedItemCallOptions

	// The gRPC API client.
	feedItemClient servicespb.FeedItemServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewFeedItemClient creates a new feed item service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service to manage feed items.
func NewFeedItemClient(ctx context.Context, opts ...option.ClientOption) (*FeedItemClient, error) {
	clientOpts := defaultFeedItemGRPCClientOptions()
	if newFeedItemClientHook != nil {
		hookOpts, err := newFeedItemClientHook(ctx, clientHookParams{})
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
	client := FeedItemClient{CallOptions: defaultFeedItemCallOptions()}

	c := &feedItemGRPCClient{
		connPool:    connPool,
		disableDeadlines: disableDeadlines,
		feedItemClient: servicespb.NewFeedItemServiceClient(connPool),
		CallOptions: &client.CallOptions,

	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *feedItemGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *feedItemGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *feedItemGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *feedItemGRPCClient) GetFeedItem(ctx context.Context, req *servicespb.GetFeedItemRequest, opts ...gax.CallOption) (*resourcespb.FeedItem, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "resource_name", url.QueryEscape(req.GetResourceName())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetFeedItem[0:len((*c.CallOptions).GetFeedItem):len((*c.CallOptions).GetFeedItem)], opts...)
	var resp *resourcespb.FeedItem
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemClient.GetFeedItem(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *feedItemGRPCClient) MutateFeedItems(ctx context.Context, req *servicespb.MutateFeedItemsRequest, opts ...gax.CallOption) (*servicespb.MutateFeedItemsResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 3600000 * time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "customer_id", url.QueryEscape(req.GetCustomerId())))
	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).MutateFeedItems[0:len((*c.CallOptions).MutateFeedItems):len((*c.CallOptions).MutateFeedItems)], opts...)
	var resp *servicespb.MutateFeedItemsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.feedItemClient.MutateFeedItems(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
