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

package gads_test

import (
	"context"

	gads "github.com/opteo/google-ads-go"
	"google.golang.org/api/iterator"
	servicespb "github.com/opteo/google-ads-go/services"
)

func ExampleNewCampaignDraftClient() {
	ctx := context.Background()
	c, err := gads.NewCampaignDraftClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	// TODO: Use client.
	_ = c
}

func ExampleCampaignDraftClient_GetCampaignDraft() {
	ctx := context.Background()
	c, err := gads.NewCampaignDraftClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.GetCampaignDraftRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCampaignDraft(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCampaignDraftClient_MutateCampaignDrafts() {
	ctx := context.Background()
	c, err := gads.NewCampaignDraftClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.MutateCampaignDraftsRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.MutateCampaignDrafts(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleCampaignDraftClient_PromoteCampaignDraft() {
	ctx := context.Background()
	c, err := gads.NewCampaignDraftClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.PromoteCampaignDraftRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.PromoteCampaignDraft(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleCampaignDraftClient_ListCampaignDraftAsyncErrors() {
	ctx := context.Background()
	c, err := gads.NewCampaignDraftClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &servicespb.ListCampaignDraftAsyncErrorsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCampaignDraftAsyncErrors(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
