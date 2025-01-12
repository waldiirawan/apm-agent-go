// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package apm // import "github.com/waldiirawan/apm-agent-go/v2"

import (
	"github.com/waldiirawan/apm-agent-go/v2/internal/wildcard"
	"github.com/waldiirawan/apm-agent-go/v2/model"
)

const redacted = "[REDACTED]"

var redactedValues = []string{redacted}

// sanitizeRequest sanitizes HTTP request data, redacting the
// values of cookies, headers and forms whose corresponding keys
// match any of the given wildcard patterns.
func sanitizeRequest(r *model.Request, matchers wildcard.Matchers) {
	for _, c := range r.Cookies {
		if !matchers.MatchAny(c.Name) {
			continue
		}
		c.Value = redacted
	}
	sanitizeHeaders(r.Headers, matchers)
	if r.Body != nil && r.Body.Form != nil {
		for key := range r.Body.Form {
			if !matchers.MatchAny(key) {
				continue
			}
			r.Body.Form[key] = redactedValues
		}
	}
}

// sanitizeResponse sanitizes HTTP response data, redacting
// the values of response headers whose corresponding keys
// match any of the given wildcard patterns.
func sanitizeResponse(r *model.Response, matchers wildcard.Matchers) {
	sanitizeHeaders(r.Headers, matchers)
}

func sanitizeHeaders(headers model.Headers, matchers wildcard.Matchers) {
	for i := range headers {
		h := &headers[i]
		if !matchers.MatchAny(h.Key) || len(h.Values) == 0 || h.Key == ":authority" {
			continue
		}
		// h.Values may hold the original value slice from a
		// net/http.Request, so it's important that we do not
		// modify it. Instead, just replace the values with a
		// shared, immutable slice.
		h.Values = redactedValues
	}
}
