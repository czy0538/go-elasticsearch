// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/6e0fb6b929f337b62bf0676bdf503e061121fad2

// Allow to shrink an existing index into a new index with fewer primary shards.
package shrink

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	indexMask = iota + 1

	targetMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Shrink struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index  string
	target string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewShrink type alias for index.
type NewShrink func(index, target string) *Shrink

// NewShrinkFunc returns a new instance of Shrink with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewShrinkFunc(tp elastictransport.Interface) NewShrink {
	return func(index, target string) *Shrink {
		n := New(tp)

		n._index(index)

		n._target(target)

		return n
	}
}

// Allow to shrink an existing index into a new index with fewer primary shards.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shrink-index.html
func New(tp elastictransport.Interface) *Shrink {
	r := &Shrink{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Shrink) Raw(raw io.Reader) *Shrink {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Shrink) Request(req *Request) *Shrink {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Shrink) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Shrink: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|targetMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_shrink")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "target", r.target)
		}
		path.WriteString(r.target)

		method = http.MethodPut
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Shrink) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.shrink")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "indices.shrink")
		if reader := instrument.RecordRequestBody(ctx, "indices.shrink", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.shrink")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Shrink query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a shrink.Response
func (r Shrink) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.shrink")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the Shrink headers map.
func (r *Shrink) Header(key, value string) *Shrink {
	r.headers.Set(key, value)

	return r
}

// Index Name of the source index to shrink.
// API Name: index
func (r *Shrink) _index(index string) *Shrink {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Target Name of the target index to create.
// API Name: target
func (r *Shrink) _target(target string) *Shrink {
	r.paramSet |= targetMask
	r.target = target

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Shrink) MasterTimeout(duration string) *Shrink {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *Shrink) Timeout(duration string) *Shrink {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Shrink) WaitForActiveShards(waitforactiveshards string) *Shrink {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// Aliases The key is the alias name.
// Index alias names support date math.
// API name: aliases
func (r *Shrink) Aliases(aliases map[string]types.Alias) *Shrink {

	r.req.Aliases = aliases

	return r
}

// Settings Configuration options for the target index.
// API name: settings
func (r *Shrink) Settings(settings map[string]json.RawMessage) *Shrink {

	r.req.Settings = settings

	return r
}
