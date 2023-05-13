// Copyright (c) 2023 Z5Labs and Contributors
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package nquad

import (
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	rdfpb "github.com/z5labs/rdf/proto"
)

// Marshaler represents any type which can
// convert itself into a NQuad representation.
type Marshaler interface {
	MarshalNQuad() ([]byte, error)
}

// Unmarshaler represents any type which can
// construct itself from a NQuad representation.
type Unmarshaler interface {
	UnmarshalNQuad(data []byte) error
}

// Triple is a type alias for the underlying
// RDF Triple protobuf message which implements
// the Marshaler and Unmarshaler interfaces.
type Triple rdfpb.Triple

// MarshalNQuad implements the Marshaler interface.
func (t *Triple) MarshalNQuad() ([]byte, error) {
	var m marshaler
	return m.marshal((*rdfpb.Triple)(t))
}

type marshaler struct {
	buf []byte
}

func (m *marshaler) marshal(t *rdfpb.Triple) ([]byte, error) {
	switch x := t.Subject.Value.(type) {
	case *rdfpb.Subject_BlankNode:
		m.marshalBlankNode(x.BlankNode)
	case *rdfpb.Subject_Iri:
		m.marshalIri(x.Iri)
	}
	m.buf = append(m.buf, " "...)
	m.marshalIri(t.Predicate)
	m.buf = append(m.buf, " "...)
	switch x := t.Object.Value.(type) {
	case *rdfpb.Object_BlankNode:
		m.marshalBlankNode(x.BlankNode)
	case *rdfpb.Object_Iri:
		m.marshalIri(x.Iri)
	case *rdfpb.Object_Literal:
		m.marshalLiteral(x.Literal)
	}
	m.buf = append(m.buf, " ."...)
	return m.buf, nil
}

func (m *marshaler) marshalBlankNode(s string) {
	m.buf = append(m.buf, "_:"...)
	m.buf = append(m.buf, s...)
}

func (m *marshaler) marshalIri(s string) {
	m.buf = append(m.buf, "<"...)
	m.buf = append(m.buf, s...)
	m.buf = append(m.buf, ">"...)
}

func (m *marshaler) marshalLiteral(lit *rdfpb.Literal) {
	m.buf = append(m.buf, "\""...)
	switch x := lit.Value.(type) {
	case *rdfpb.Literal_String_:
		m.buf = append(m.buf, x.String_...)
	case *rdfpb.Literal_Int:
		m.buf = append(m.buf, strconv.FormatInt(x.Int, 10)...)
	case *rdfpb.Literal_Float64:
		m.buf = append(m.buf, strconv.FormatFloat(x.Float64, 'f', -1, 64)...)
	case *rdfpb.Literal_Bool:
		m.buf = append(m.buf, strconv.FormatBool(x.Bool)...)
	case *rdfpb.Literal_Bytes:
		base64.StdEncoding.Encode(m.buf, x.Bytes)
	}
	m.buf = append(m.buf, "\""...)
}

// UnmarshalNQuad implements the Unmarshaler interface.
func (t *Triple) UmarshalNQuad(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	s := string(data)
	terms := strings.Split(s, " ")
	if len(terms) > 4 {
		return errors.New("nquad statement can have a max of 4 terms in it")
	}
	u := unmarshaler{
		terms: terms,
	}
	return u.unmarshal(t)
}

type unmarshaler struct {
	terms []string
}

type stateAction func(*unmarshaler, *Triple) (stateAction, error)

func (u *unmarshaler) unmarshal(triple *Triple) (err error) {
	for next := unmarshalSubject; next != nil; {
		next, err = next(u, triple)
		if err != nil {
			return err
		}
	}
	return nil
}

func unmarshalSubject(u *unmarshaler, triple *Triple) (stateAction, error) {
	s := u.peek()
	r, _ := utf8.DecodeRuneInString(s)
	switch r {
	case '<':
		iri, err := u.unmarshalIri()
		if err != nil {
			return nil, err
		}
		triple.Subject = &rdfpb.Subject{
			Value: &rdfpb.Subject_Iri{
				Iri: iri,
			},
		}
	case '_':
		blankNode, err := u.unmarshalBlankNode()
		if err != nil {
			return nil, err
		}
		triple.Subject = &rdfpb.Subject{
			Value: &rdfpb.Subject_BlankNode{
				BlankNode: blankNode,
			},
		}
	default:
		return nil, fmt.Errorf("unexpected starting character for subject: %q", r)
	}
	return unmarshalPredicate, nil
}

func unmarshalPredicate(u *unmarshaler, triple *Triple) (stateAction, error) {
	s := u.peek()
	if len(s) == 0 {
		return nil, errors.New("missing predicate")
	}
	iri, err := u.unmarshalIri()
	if err != nil {
		return nil, err
	}
	triple.Predicate = iri
	return unmarshalObject, nil
}

func unmarshalObject(u *unmarshaler, triple *Triple) (stateAction, error) {
	s := u.peek()
	if len(s) == 0 {
		return nil, errors.New("missing object")
	}
	r, _ := utf8.DecodeRuneInString(s)
	switch r {
	case '<':
		iri, err := u.unmarshalIri()
		if err != nil {
			return nil, err
		}
		triple.Object = &rdfpb.Object{
			Value: &rdfpb.Object_Iri{
				Iri: iri,
			},
		}
	case '_':
		blankNode, err := u.unmarshalBlankNode()
		if err != nil {
			return nil, err
		}
		triple.Object = &rdfpb.Object{
			Value: &rdfpb.Object_BlankNode{
				BlankNode: blankNode,
			},
		}
	default:
		lit, err := u.unmarshalLiteral()
		if err != nil {
			return nil, err
		}
		triple.Object = &rdfpb.Object{
			Value: &rdfpb.Object_Literal{
				Literal: lit,
			},
		}
	}
	return unmarshalEndOfStatement, nil
}

func unmarshalEndOfStatement(u *unmarshaler, triple *Triple) (stateAction, error) {
	s, ok := u.next()
	if !ok || s != "." {
		return nil, errors.New("nquad statement should end with '.'")
	}
	return nil, nil
}

func (u *unmarshaler) unmarshalIri() (string, error) {
	s, _ := u.next()
	return strings.Trim(s, "<>"), nil
}

func (u *unmarshaler) unmarshalBlankNode() (string, error) {
	s, _ := u.next()
	return strings.TrimPrefix(s, "_:"), nil
}

func (u *unmarshaler) unmarshalLiteral() (*rdfpb.Literal, error) {
	return nil, nil
}

func (u *unmarshaler) peek() string {
	if len(u.terms) == 0 {
		return ""
	}
	return u.terms[0]
}

func (u *unmarshaler) next() (string, bool) {
	if len(u.terms) == 0 {
		return "", false
	}
	s := u.terms[0]
	u.terms = u.terms[1:]
	return s, true
}
