// Copyright (c) 2023 Z5Labs and Contributors
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package rdf

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"

	rdfpb "github.com/z5labs/rdf/proto"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// Graph represents a set of RDF Triples.
type Graph []*rdfpb.Triple

// Add the given RDF Triple to the graph.
func (g *Graph) Add(t *rdfpb.Triple) *Graph {
	*g = append(*g, t)
	return g
}

// Triples returns all the returns contained within the graph.
func (g Graph) Triples() []*rdfpb.Triple {
	return g
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
//
// Binary format = repeat 8 bytes + n bytes
// 8 bytes = length of RDF Triple protobuf message
// n bytes = RDF Triple protobuf message
func (g Graph) MarshalBinary() ([]byte, error) {
	if len(g) == 0 {
		return nil, nil
	}

	var buf bytes.Buffer
	err := marshalSequentially(&buf, g)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func marshalSequentially(w io.Writer, g Graph) error {
	for _, t := range g {
		b, err := proto.Marshal(t)
		if err != nil {
			return err
		}

		length := len(b)
		err = binary.Write(w, binary.LittleEndian, uint64(length))
		if err != nil {
			return err
		}
		n, err := w.Write(b)
		if err != nil {
			return err
		}
		if n != length {
			return errors.New("failed to write all of rdf triple proto message")
		}
	}
	return nil
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
//
// Binary format = repeat 8 bytes + n bytes
// 8 bytes = length of RDF Triple protobuf message
// n bytes = RDF Triple protobuf message
func (g *Graph) UnmarshalBinary(data []byte) error {
	for {
		if len(data) == 0 {
			return nil
		}
		if len(data) < 8 {
			return errors.New("malformed rdf triple length")
		}
		length := binary.LittleEndian.Uint64(data[0:8])
		data = data[8:]
		if uint64(len(data)) < length {
			return errors.New("triple length does not match remaining bytes")
		}

		var triple rdfpb.Triple
		b := data[0:length]
		err := proto.Unmarshal(b, &triple)
		if err != nil {
			return err
		}
		g.Add(&triple)
		data = data[length:]
	}
}

type jsonGraph struct {
	Triples []json.RawMessage `json:"triples"`
}

// MarshalJSON implements json.Marshaler interface.
//
// CAUTION!! Always prefer MarshalBinary over this.
// This relies in protojson which does not
// define a stable version.
func (g Graph) MarshalJSON() ([]byte, error) {
	if len(g) == 0 {
		return nil, nil
	}

	triples := make([]json.RawMessage, len(g))
	for i := range g {
		b, err := protojson.Marshal(g[i])
		if err != nil {
			return nil, err
		}
		triples[i] = b
	}

	b, err := json.Marshal(jsonGraph{
		Triples: triples,
	})
	if err != nil {
		return nil, err
	}
	return b, nil
}

// UnmarshalJSON implements json.Unmarshaler interface.
//
// CAUTION!! Always prefer UnmarshalBinary over this.
// This relies in protojson which does not
// define a stable version.
func (g *Graph) UnmarshalJSON(data []byte) error {
	var jg jsonGraph
	err := json.Unmarshal(data, &jg)
	if err != nil {
		return err
	}
	for _, b := range jg.Triples {
		var triple rdfpb.Triple
		err = protojson.Unmarshal(b, &triple)
		if err != nil {
			return err
		}
		g.Add(&triple)
	}
	return nil
}
