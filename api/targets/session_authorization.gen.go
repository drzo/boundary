// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package targets

import (
	"time"

	"github.com/hashicorp/boundary/api/scopes"
)

type SessionAuthorization struct {
	SessionId          string               `json:"session_id,omitempty"`
	TargetId           string               `json:"target_id,omitempty"`
	Scope              *scopes.ScopeInfo    `json:"scope,omitempty"`
	CreatedTime        time.Time            `json:"created_time,omitempty"`
	UserId             string               `json:"user_id,omitempty"`
	HostSetId          string               `json:"host_set_id,omitempty"`
	HostId             string               `json:"host_id,omitempty"`
	Type               string               `json:"type,omitempty"`
	AuthorizationToken string               `json:"authorization_token,omitempty"`
	Endpoint           string               `json:"endpoint,omitempty"`
	Expiration         time.Time            `json:"expiration,omitempty"`
	Credentials        []*SessionCredential `json:"credentials,omitempty"`
}
