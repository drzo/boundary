// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package targets

import (
	"encoding/json"
)

type SessionSecret struct {
	Raw     json.RawMessage        `json:"raw,omitempty"`
	Decoded map[string]interface{} `json:"decoded,omitempty"`
}
