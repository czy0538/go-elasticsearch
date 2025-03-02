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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// IngestPipeline type.
//
// https://github.com/elastic/elasticsearch-specification/blob/6e0fb6b929f337b62bf0676bdf503e061121fad2/specification/ingest/_types/Pipeline.ts#L23-L45
type IngestPipeline struct {
	// Description Description of the ingest pipeline.
	Description *string `json:"description,omitempty"`
	// Meta_ Arbitrary metadata about the ingest pipeline. This map is not automatically
	// generated by Elasticsearch.
	Meta_ Metadata `json:"_meta"`
	// OnFailure Processors to run immediately after a processor failure.
	OnFailure []ProcessorContainer `json:"on_failure,omitempty"`
	// Processors Processors used to perform transformations on documents before indexing.
	// Processors run sequentially in the order specified.
	Processors []ProcessorContainer `json:"processors,omitempty"`
	// Version Version number used by external systems to track ingest pipelines.
	Version *int64 `json:"version,omitempty"`
}

func (s *IngestPipeline) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "_meta":
			if err := dec.Decode(&s.Meta_); err != nil {
				return err
			}

		case "on_failure":
			if err := dec.Decode(&s.OnFailure); err != nil {
				return err
			}

		case "processors":
			if err := dec.Decode(&s.Processors); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewIngestPipeline returns a IngestPipeline.
func NewIngestPipeline() *IngestPipeline {
	r := &IngestPipeline{}

	return r
}
