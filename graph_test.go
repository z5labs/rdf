// Copyright (c) 2023 Z5Labs and Contributors
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package rdf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryEncodingIdentity(t *testing.T) {
	t.Run("id == MarshalBinary . UnmarshalBinary", func(t *testing.T) {
		buf := []byte{27, 0, 0, 0, 0, 0, 0, 0, 10, 7, 18, 5, 95, 58, 98, 111, 98, 18, 5,
			107, 110, 111, 119, 115, 26, 9, 42, 7, 95, 58, 97, 108, 105, 99, 101, 27, 0, 0, 0, 0, 0, 0, 0, 10, 9, 18, 7, 95,
			58, 97, 108, 105, 99, 101, 18, 5, 107, 110, 111, 119, 115, 26, 7, 42, 5, 95, 58, 98, 111, 98}

		var g Graph
		err := g.UnmarshalBinary(buf)
		if !assert.Nil(t, err) {
			return
		}
		gb, err := g.MarshalBinary()
		if !assert.Nil(t, err) {
			return
		}

		if !assert.Equal(t, buf, gb) {
			return
		}
	})

	t.Run("id == UnmarshalBinary . MarshalBinary", func(t *testing.T) {
		var g Graph
		g.
			Add(NewTriple(BlankNodeSubject("_:bob"), Predicate("knows"), BlankNodeObject("_:alice"))).
			Add(NewTriple(BlankNodeSubject("_:alice"), Predicate("knows"), BlankNodeObject("_:bob")))

		gb, err := g.MarshalBinary()
		if !assert.Nil(t, err) {
			return
		}

		var h Graph
		err = h.UnmarshalBinary(gb)
		if !assert.Nil(t, err) {
			return
		}

		if !assert.Len(t, h.Triples(), len(g.Triples())) {
			return
		}
	})
}

func TestJSONEncodingIdentity(t *testing.T) {
	t.Run("id == MarshalJSON . UnmarshalJSON", func(t *testing.T) {
		buf := `{"triples":[{"subject":{"blankNode":"_:bob"},"predicate":"knows","object":{"blankNode":"_:alice"}},{"subject":{"blankNode":"_:alice"},"predicate":"knows","object":{"blankNode":"_:bob"}}]}`

		var g Graph
		err := g.UnmarshalJSON([]byte(buf))
		if !assert.Nil(t, err) {
			return
		}
		gb, err := g.MarshalJSON()
		if !assert.Nil(t, err) {
			return
		}

		if !assert.Equal(t, buf, string(gb)) {
			return
		}
	})

	t.Run("id == UnmarshalJSON . MarshalJSON", func(t *testing.T) {
		var g Graph
		g.
			Add(NewTriple(BlankNodeSubject("_:bob"), Predicate("knows"), BlankNodeObject("_:alice"))).
			Add(NewTriple(BlankNodeSubject("_:alice"), Predicate("knows"), BlankNodeObject("_:bob")))

		gb, err := g.MarshalJSON()
		if !assert.Nil(t, err) {
			return
		}

		var h Graph
		err = h.UnmarshalJSON(gb)
		if !assert.Nil(t, err) {
			return
		}

		if !assert.Len(t, h.Triples(), len(g.Triples())) {
			return
		}
	})
}

func BenchmarkGraph_MarshalBinary(b *testing.B) {
	var g Graph
	g.
		Add(NewTriple(BlankNodeSubject("_:bob"), Predicate("knows"), BlankNodeObject("_:alice"))).
		Add(NewTriple(BlankNodeSubject("_:alice"), Predicate("knows"), BlankNodeObject("_:bob")))

	for i := 0; i < b.N; i++ {
		buf, err := g.MarshalBinary()
		if err != nil {
			b.Error(err)
			return
		}
		if len(buf) == 0 {
			b.Fail()
			return
		}
	}
}

func BenchmarkGraph_UnmarshalBinary(b *testing.B) {
	buf := []byte{27, 0, 0, 0, 0, 0, 0, 0, 10, 7, 18, 5, 95, 58, 98, 111, 98, 18, 5,
		107, 110, 111, 119, 115, 26, 9, 42, 7, 95, 58, 97, 108, 105, 99, 101, 27, 0, 0, 0, 0, 0, 0, 0, 10, 9, 18, 7, 95,
		58, 97, 108, 105, 99, 101, 18, 5, 107, 110, 111, 119, 115, 26, 7, 42, 5, 95, 58, 98, 111, 98}

	for i := 0; i < b.N; i++ {
		var g Graph
		err := g.UnmarshalBinary(buf)
		if err != nil {
			b.Error(err)
			return
		}
		if len(g) != 2 {
			b.Fail()
			return
		}
	}
}

func BenchmarkGraph_MarshalJSON(b *testing.B) {
	var g Graph
	g.
		Add(NewTriple(BlankNodeSubject("_:bob"), Predicate("knows"), BlankNodeObject("_:alice"))).
		Add(NewTriple(BlankNodeSubject("_:alice"), Predicate("knows"), BlankNodeObject("_:bob")))

	for i := 0; i < b.N; i++ {
		buf, err := g.MarshalJSON()
		if err != nil {
			b.Error(err)
			return
		}
		if len(buf) == 0 {
			b.Fail()
			return
		}
	}
}

func BenchmarkGraph_UnmarshalJSON(b *testing.B) {
	buf := []byte(`{"triples":[{"subject":{"blankNode":"_:bob"},"predicate":"knows","object":{"blankNode":"_:alice"}},{"subject":{"blankNode":"_:alice"},"predicate":"knows","object":{"blankNode":"_:bob"}}]}`)

	for i := 0; i < b.N; i++ {
		var g Graph
		err := g.UnmarshalJSON(buf)
		if err != nil {
			b.Error(err)
			return
		}
		if len(g) != 2 {
			b.Fail()
			return
		}
	}
}
