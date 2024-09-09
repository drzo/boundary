// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package roles

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"slices"
	"strconv"
	"time"

	"github.com/hashicorp/boundary/api"
	"github.com/hashicorp/boundary/api/scopes"
)

type Role struct {
	Id                string            `json:"id,omitempty"`
	ScopeId           string            `json:"scope_id,omitempty"`
	Scope             *scopes.ScopeInfo `json:"scope,omitempty"`
	Name              string            `json:"name,omitempty"`
	Description       string            `json:"description,omitempty"`
	CreatedTime       time.Time         `json:"created_time,omitempty"`
	UpdatedTime       time.Time         `json:"updated_time,omitempty"`
	Version           uint32            `json:"version,omitempty"`
	GrantScopeIds     []string          `json:"grant_scope_ids,omitempty"`
	PrincipalIds      []string          `json:"principal_ids,omitempty"`
	Principals        []*Principal      `json:"principals,omitempty"`
	GrantStrings      []string          `json:"grant_strings,omitempty"`
	Grants            []*Grant          `json:"grants,omitempty"`
	AuthorizedActions []string          `json:"authorized_actions,omitempty"`
}

type RoleReadResult struct {
	Item     *Role
	Response *api.Response
}

func (n RoleReadResult) GetItem() *Role {
	return n.Item
}

func (n RoleReadResult) GetResponse() *api.Response {
	return n.Response
}

type RoleCreateResult = RoleReadResult
type RoleUpdateResult = RoleReadResult

type RoleDeleteResult struct {
	Response *api.Response
}

// GetItem will always be nil for RoleDeleteResult
func (n RoleDeleteResult) GetItem() any {
	return nil
}

func (n RoleDeleteResult) GetResponse() *api.Response {
	return n.Response
}

type RoleListResult struct {
	Items        []*Role  `json:"items,omitempty"`
	EstItemCount uint     `json:"est_item_count,omitempty"`
	RemovedIds   []string `json:"removed_ids,omitempty"`
	ListToken    string   `json:"list_token,omitempty"`
	ResponseType string   `json:"response_type,omitempty"`
	Response     *api.Response

	// The following fields are used for cached information when client-directed
	// pagination is used.
	Recursive     bool
	PageSize      uint32
	AllRemovedIds []string
	scopeId       string
}

func (n RoleListResult) GetItems() []*Role {
	return n.Items
}

func (n RoleListResult) GetEstItemCount() uint {
	return n.EstItemCount
}

func (n RoleListResult) GetRemovedIds() []string {
	return n.RemovedIds
}

func (n RoleListResult) GetListToken() string {
	return n.ListToken
}

func (n RoleListResult) GetResponseType() string {
	return n.ResponseType
}

func (n RoleListResult) GetResponse() *api.Response {
	return n.Response
}

// Client is a client for this collection
type Client struct {
	client *api.Client
}

// Creates a new client for this collection. The submitted API client is cloned;
// modifications to it after generating this client will not have effect. If you
// need to make changes to the underlying API client, use ApiClient() to access
// it.
func NewClient(c *api.Client) *Client {
	return &Client{client: c.Clone()}
}

// ApiClient returns the underlying API client
func (c *Client) ApiClient() *api.Client {
	return c.client
}

func (c *Client) Create(ctx context.Context, scopeId string, opt ...Option) (*RoleCreateResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into Create request")
	}

	opts, apiOpts := getOpts(opt...)

	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts.postMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "POST", "roles", opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Create request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Create call: %w", err)
	}

	target := new(RoleCreateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Create response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) Read(ctx context.Context, id string, opt ...Option) (*RoleReadResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Read request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "GET", fmt.Sprintf("roles/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Read request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Read call: %w", err)
	}

	target := new(RoleReadResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Read response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) Update(ctx context.Context, id string, version uint32, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Update request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into Update request and automatic versioning not specified")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	req, err := c.client.NewRequest(ctx, "PATCH", fmt.Sprintf("roles/%s", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Update request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Update call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding Update response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) Delete(ctx context.Context, id string, opt ...Option) (*RoleDeleteResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into Delete request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	req, err := c.client.NewRequest(ctx, "DELETE", fmt.Sprintf("roles/%s", url.PathEscape(id)), nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating Delete request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during Delete call: %w", err)
	}

	apiErr, err := resp.Decode(nil)
	if err != nil {
		return nil, fmt.Errorf("error decoding Delete response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	target := &RoleDeleteResult{
		Response: resp,
	}
	return target, nil
}

func (c *Client) List(ctx context.Context, scopeId string, opt ...Option) (*RoleListResult, error) {
	if scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value passed into List request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["scope_id"] = scopeId

	req, err := c.client.NewRequest(ctx, "GET", "roles", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call: %w", err)
	}

	target := new(RoleListResult)
	apiErr, err := resp.Decode(target)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	if target.ResponseType == "complete" || target.ResponseType == "" {
		return target, nil
	}

	// In case we shortcut out due to client directed pagination, ensure these
	// are set
	target.Recursive = opts.withRecursive
	target.PageSize = opts.withPageSize
	target.scopeId = scopeId
	target.AllRemovedIds = target.RemovedIds
	if opts.withClientDirectedPagination {
		return target, nil
	}

	allItems := make([]*Role, 0, target.EstItemCount)
	allItems = append(allItems, target.Items...)

	// If there are more results, automatically fetch the rest of the results.
	// idToIndex keeps a map from the ID of an item to its index in target.Items.
	// This is used to update updated items in-place and remove deleted items
	// from the result after pagination is done.
	idToIndex := map[string]int{}
	for i, item := range allItems {
		idToIndex[item.Id] = i
	}

	// If we're here there are more pages and the client does not want to
	// paginate on their own; fetch them as this call returns all values.
	currentPage := target
	for {
		nextPage, err := c.ListNextPage(ctx, currentPage, opt...)
		if err != nil {
			return nil, fmt.Errorf("error getting next page in List call: %w", err)
		}

		for _, item := range nextPage.Items {
			if i, ok := idToIndex[item.Id]; ok {
				// Item has already been seen at index i, update in-place
				allItems[i] = item
			} else {
				allItems = append(allItems, item)
				idToIndex[item.Id] = len(allItems) - 1
			}
		}

		currentPage = nextPage

		if currentPage.ResponseType == "complete" {
			break
		}
	}

	// The current page here is the final page of the results, that is, the
	// response type is "complete"

	// Remove items that were deleted since the end of the last iteration.
	// If a Role has been updated and subsequently removed, we don't want
	// it to appear both in the Items and RemovedIds, so we remove it from the Items.
	for _, removedId := range currentPage.RemovedIds {
		if i, ok := idToIndex[removedId]; ok {
			// Remove the item at index i without preserving order
			// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
			allItems[i] = allItems[len(allItems)-1]
			allItems = allItems[:len(allItems)-1]
			// Update the index of the previously last element
			idToIndex[allItems[i].Id] = i
		}
	}
	// Sort the results again since in-place updates and deletes
	// may have shuffled items. We sort by created time descending
	// (most recently created first), same as the API.
	slices.SortFunc(allItems, func(i, j *Role) int {
		return j.CreatedTime.Compare(i.CreatedTime)
	})
	// Since we paginated to the end, we can avoid confusion
	// for the user by setting the estimated item count to the
	// length of the items slice. If we don't set this here, it
	// will equal the value returned in the last response, which is
	// often much smaller than the total number returned.
	currentPage.EstItemCount = uint(len(allItems))
	// Set items to the full list we have collected here
	currentPage.Items = allItems
	// Set the returned value to the last page with calculated values
	target = currentPage
	// Finally, since we made at least 2 requests to the server to fulfill this
	// function call, resp.Body and resp.Map will only contain the most recent response.
	// Overwrite them with the true response.
	target.Response.Body.Reset()
	if err := json.NewEncoder(target.Response.Body).Encode(target); err != nil {
		return nil, fmt.Errorf("error encoding final JSON list response: %w", err)
	}
	if err := json.Unmarshal(target.Response.Body.Bytes(), &target.Response.Map); err != nil {
		return nil, fmt.Errorf("error encoding final map list response: %w", err)
	}
	// Note: the HTTP response body is consumed by resp.Decode in the loop,
	// so it doesn't need to be updated (it will always be, and has always been, empty).
	return target, nil
}

func (c *Client) ListNextPage(ctx context.Context, currentPage *RoleListResult, opt ...Option) (*RoleListResult, error) {
	if currentPage == nil {
		return nil, fmt.Errorf("empty currentPage value passed into ListNextPage request")
	}
	if currentPage.scopeId == "" {
		return nil, fmt.Errorf("empty scopeId value in currentPage passed into ListNextPage request")
	}
	if c.client == nil {
		return nil, fmt.Errorf("nil client")
	}
	if currentPage.ResponseType == "complete" || currentPage.ResponseType == "" {
		return nil, fmt.Errorf("no more pages available in ListNextPage request")
	}

	opts, apiOpts := getOpts(opt...)
	opts.queryMap["scope_id"] = currentPage.scopeId

	// Don't require them to re-specify recursive
	if currentPage.Recursive {
		opts.queryMap["recursive"] = "true"
	}

	if currentPage.PageSize != 0 {
		opts.queryMap["page_size"] = strconv.FormatUint(uint64(currentPage.PageSize), 10)
	}

	req, err := c.client.NewRequest(ctx, "GET", "targets", nil, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating List request: %w", err)
	}

	opts.queryMap["list_token"] = currentPage.ListToken
	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during List call during ListNextPage: %w", err)
	}

	nextPage := new(RoleListResult)
	apiErr, err := resp.Decode(nextPage)
	if err != nil {
		return nil, fmt.Errorf("error decoding List response during ListNextPage: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}

	// Ensure values are carried forward to the next call
	nextPage.scopeId = currentPage.scopeId

	nextPage.Recursive = currentPage.Recursive

	nextPage.PageSize = currentPage.PageSize
	// Cache the removed IDs from this page
	nextPage.AllRemovedIds = append(currentPage.AllRemovedIds, nextPage.RemovedIds...)
	// Set the response body to the current response
	nextPage.Response = resp
	// If we're done iterating, pull the full set of removed IDs into the last
	// response
	if nextPage.ResponseType == "complete" {
		// Collect up the last values
		nextPage.RemovedIds = nextPage.AllRemovedIds
		// For now, removedIds will only be populated if this pagination cycle
		// was the result of a "refresh" operation (i.e., the caller provided a
		// list token option to this call).
		//
		// Sort to make response deterministic
		slices.Sort(nextPage.RemovedIds)
		// Remove any duplicates
		nextPage.RemovedIds = slices.Compact(nextPage.RemovedIds)
	}

	return nextPage, nil
}

func (c *Client) AddGrantScopes(ctx context.Context, id string, version uint32, grantScopeIds []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into AddGrantScopes request")
	}

	if len(grantScopeIds) == 0 {
		return nil, errors.New("empty grantScopeIds passed into AddGrantScopes request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into AddGrantScopes request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["grant_scope_ids"] = grantScopeIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:add-grant-scopes", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating AddGrantScopes request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during AddGrantScopes call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding AddGrantScopes response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) AddGrants(ctx context.Context, id string, version uint32, grantStrings []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into AddGrants request")
	}

	if len(grantStrings) == 0 {
		return nil, errors.New("empty grantStrings passed into AddGrants request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into AddGrants request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["grant_strings"] = grantStrings

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:add-grants", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating AddGrants request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during AddGrants call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding AddGrants response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) AddPrincipals(ctx context.Context, id string, version uint32, principalIds []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into AddPrincipals request")
	}

	if len(principalIds) == 0 {
		return nil, errors.New("empty principalIds passed into AddPrincipals request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into AddPrincipals request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["principal_ids"] = principalIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:add-principals", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating AddPrincipals request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during AddPrincipals call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding AddPrincipals response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) SetGrantScopes(ctx context.Context, id string, version uint32, grantScopeIds []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into SetGrantScopes request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into SetGrantScopes request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["grant_scope_ids"] = grantScopeIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:set-grant-scopes", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating SetGrantScopes request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during SetGrantScopes call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding SetGrantScopes response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) SetGrants(ctx context.Context, id string, version uint32, grantStrings []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into SetGrants request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into SetGrants request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["grant_strings"] = grantStrings

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:set-grants", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating SetGrants request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during SetGrants call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding SetGrants response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) SetPrincipals(ctx context.Context, id string, version uint32, principalIds []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into SetPrincipals request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into SetPrincipals request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["principal_ids"] = principalIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:set-principals", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating SetPrincipals request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during SetPrincipals call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding SetPrincipals response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) RemoveGrantScopes(ctx context.Context, id string, version uint32, grantScopeIds []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into RemoveGrantScopes request")
	}

	if len(grantScopeIds) == 0 {
		return nil, errors.New("empty grantScopeIds passed into RemoveGrantScopes request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into RemoveGrantScopes request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["grant_scope_ids"] = grantScopeIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:remove-grant-scopes", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating RemoveGrantScopes request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during RemoveGrantScopes call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding RemoveGrantScopes response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) RemoveGrants(ctx context.Context, id string, version uint32, grantStrings []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into RemoveGrants request")
	}

	if len(grantStrings) == 0 {
		return nil, errors.New("empty grantStrings passed into RemoveGrants request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into RemoveGrants request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["grant_strings"] = grantStrings

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:remove-grants", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating RemoveGrants request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during RemoveGrants call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding RemoveGrants response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}

func (c *Client) RemovePrincipals(ctx context.Context, id string, version uint32, principalIds []string, opt ...Option) (*RoleUpdateResult, error) {
	if id == "" {
		return nil, fmt.Errorf("empty id value passed into RemovePrincipals request")
	}

	if len(principalIds) == 0 {
		return nil, errors.New("empty principalIds passed into RemovePrincipals request")
	}

	if c.client == nil {
		return nil, errors.New("nil client")
	}

	opts, apiOpts := getOpts(opt...)

	if version == 0 {
		if !opts.withAutomaticVersioning {
			return nil, errors.New("zero version number passed into RemovePrincipals request")
		}
		existingTarget, existingErr := c.Read(ctx, id, append([]Option{WithSkipCurlOutput(true)}, opt...)...)
		if existingErr != nil {
			if api.AsServerError(existingErr) != nil {
				return nil, fmt.Errorf("error from controller when performing initial check-and-set read: %w", existingErr)
			}
			return nil, fmt.Errorf("error performing initial check-and-set read: %w", existingErr)
		}
		if existingTarget == nil {
			return nil, errors.New("nil resource response found when performing initial check-and-set read")
		}
		if existingTarget.Item == nil {
			return nil, errors.New("nil resource found when performing initial check-and-set read")
		}
		version = existingTarget.Item.Version
	}

	opts.postMap["version"] = version

	opts.postMap["principal_ids"] = principalIds

	req, err := c.client.NewRequest(ctx, "POST", fmt.Sprintf("roles/%s:remove-principals", url.PathEscape(id)), opts.postMap, apiOpts...)
	if err != nil {
		return nil, fmt.Errorf("error creating RemovePrincipals request: %w", err)
	}

	if len(opts.queryMap) > 0 {
		q := url.Values{}
		for k, v := range opts.queryMap {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing client request during RemovePrincipals call: %w", err)
	}

	target := new(RoleUpdateResult)
	target.Item = new(Role)
	apiErr, err := resp.Decode(target.Item)
	if err != nil {
		return nil, fmt.Errorf("error decoding RemovePrincipals response: %w", err)
	}
	if apiErr != nil {
		return nil, apiErr
	}
	target.Response = resp
	return target, nil
}
