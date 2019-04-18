// Copyright 2019 Google LLC
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

// Code generated by gapic-generator. DO NOT EDIT.

package talent_test

import (
	"context"

	talent "cloud.google.com/go/talent/apiv4beta1"
	"google.golang.org/api/iterator"
	talentpb "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"
)

func ExampleNewJobClient() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleJobClient_CreateJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.CreateJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_GetJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.GetJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_UpdateJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.UpdateJobRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleJobClient_DeleteJob() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.DeleteJobRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteJob(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleJobClient_ListJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.ListJobsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListJobs(ctx, req)
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

func ExampleJobClient_BatchDeleteJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.BatchDeleteJobsRequest{
		// TODO: Fill request struct fields.
	}
	err = c.BatchDeleteJobs(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleJobClient_SearchJobs() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.SearchJobsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchJobs(ctx, req)
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

func ExampleJobClient_SearchJobsForAlert() {
	ctx := context.Background()
	c, err := talent.NewJobClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &talentpb.SearchJobsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.SearchJobsForAlert(ctx, req)
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
