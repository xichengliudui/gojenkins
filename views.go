// Copyright 2015 Vadim Kravcenko
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package gojenkins

import (
	"context"
	"errors"
	"strconv"
)

type View struct {
	Raw     *ViewResponse
	Jenkins *Jenkins
	Base    string
}

type ViewResponse struct {
	Description string        `json:"description"`
	Jobs        []InnerJob    `json:"jobs"`
	Name        string        `json:"name"`
	Property    []interface{} `json:"property"`
	URL         string        `json:"url"`
}

type JenkinsViewType string

const (
	LIST_VIEW      JenkinsViewType = "hudson.model.ListView"
	NESTED_VIEW    JenkinsViewType = "hudson.plugins.nested_view.NestedView"
	MY_VIEW        JenkinsViewType = "hudson.model.MyView"
	DASHBOARD_VIEW JenkinsViewType = "hudson.plugins.view.dashboard.Dashboard"
	PIPELINE_VIEW  JenkinsViewType = "au.com.centrumsystems.hudson.plugin.buildpipeline.BuildPipelineView"
)

// Returns True if successfully added Job, otherwise false
func (v *View) AddJob(ctx context.Context, name string) (bool, error) {
	url := "/addJobToView"
	qr := map[string]string{"name": name}
	resp, err := v.Jenkins.Requester.Post(ctx, v.Base+url, nil, nil, qr)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	}
	return false, errors.New(strconv.Itoa(resp.StatusCode))
}

func (v *View) DeleteView(ctx context.Context) (bool, error) {
	url := "/doDelete"
	resp, err := v.Jenkins.Requester.Post(ctx, v.Base+url, nil, nil, nil)
	if err != nil {
		return false, err
	}
	return resp.StatusCode == 200, nil
}

// Returns True if successfully deleted Job, otherwise false
func (v *View) DeleteJob(ctx context.Context, name string) (bool, error) {
	url := "/removeJobFromView"
	qr := map[string]string{"name": name}
	resp, err := v.Jenkins.Requester.Post(ctx, v.Base+url, nil, nil, qr)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	}
	return false, errors.New(strconv.Itoa(resp.StatusCode))
}

func (v *View) GetDescription() string {
	return v.Raw.Description
}

func (v *View) GetJobs() []InnerJob {
	return v.Raw.Jobs
}

func (v *View) GetName() string {
	return v.Raw.Name
}

func (v *View) GetUrl() string {
	return v.Raw.URL
}

func (v *View) Poll(ctx context.Context) (int, error) {
	response, err := v.Jenkins.Requester.GetJSON(ctx, v.Base, v.Raw, nil)
	if err != nil {
		return 0, err
	}
	return response.StatusCode, nil
}
