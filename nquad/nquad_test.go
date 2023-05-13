// Copyright (c) 2023 Z5Labs and Contributors
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package nquad

import (
	"testing"

	"github.com/stretchr/testify/assert"
	rdfpb "github.com/z5labs/rdf/proto"
)

func TestMarshalThenUnmarshalIdentity(t *testing.T) {
	testCases := []struct {
		Name   string
		Triple *rdfpb.Triple
	}{
		{
			Name: "BlankNodeToBlankNode",
			Triple: &rdfpb.Triple{
				Subject: &rdfpb.Subject{
					Value: &rdfpb.Subject_BlankNode{
						BlankNode: "bob",
					},
				},
				Predicate: "knows",
				Object: &rdfpb.Object{
					Value: &rdfpb.Object_BlankNode{
						BlankNode: "alice",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			srcTriple := (*Triple)(testCase.Triple)
			b, err := srcTriple.MarshalNQuad()
			if !assert.Nil(t, err) {
				return
			}

			var uTriple Triple
			err = uTriple.UmarshalNQuad(b)
			if !assert.Nil(t, err) {
				return
			}

			if !assert.Equal(t, srcTriple, &uTriple) {
				return
			}
		})
	}
}

func TestUnmarshalThenMarshalIdentity(t *testing.T) {
	testCases := []struct {
		Name  string
		NQuad string
	}{
		{
			Name:  "BlankNodeToBlankNode",
			NQuad: "_:bob <knows> _:alice .",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			var triple Triple
			err := triple.UmarshalNQuad([]byte(testCase.NQuad))
			if !assert.Nil(t, err) {
				return
			}

			b, err := triple.MarshalNQuad()
			if !assert.Nil(t, err) {
				return
			}

			if !assert.Equal(t, testCase.NQuad, string(b)) {
				return
			}
		})
	}
}

func BenchmarkTriple_MarshalNQuad(b *testing.B) {
	triple := &Triple{
		Subject: &rdfpb.Subject{
			Value: &rdfpb.Subject_BlankNode{
				BlankNode: "bob",
			},
		},
		Predicate: "knows",
		Object: &rdfpb.Object{
			Value: &rdfpb.Object_BlankNode{
				BlankNode: "alice",
			},
		},
	}

	for i := 0; i < b.N; i++ {
		buf, err := triple.MarshalNQuad()
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

func BenchmarkTriple_UnmarshalNQuad(b *testing.B) {
	nquad := []byte("_:bob <knows> _:alice .")

	for i := 0; i < b.N; i++ {
		var triple Triple
		err := triple.UmarshalNQuad(nquad)
		if err != nil {
			b.Error(err)
			return
		}
		if triple.Subject == nil || triple.Predicate == "" || triple.Object == nil {
			b.Fail()
			return
		}
	}
}
