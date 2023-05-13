// Copyright (c) 2023 Z5Labs and Contributors
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package rdf

import (
	rdfpb "github.com/z5labs/rdf/proto"
)

func NewTriple(subject *rdfpb.Subject, predicate string, object *rdfpb.Object) *rdfpb.Triple {
	return &rdfpb.Triple{
		Subject:   subject,
		Predicate: predicate,
		Object:    object,
	}
}

func IriSubject(iri string) *rdfpb.Subject {
	return &rdfpb.Subject{
		Value: &rdfpb.Subject_Iri{
			Iri: iri,
		},
	}
}

func BlankNodeSubject(blankNode string) *rdfpb.Subject {
	return &rdfpb.Subject{
		Value: &rdfpb.Subject_BlankNode{
			BlankNode: blankNode,
		},
	}
}

func IriObject(iri string) *rdfpb.Object {
	return &rdfpb.Object{
		Value: &rdfpb.Object_Iri{
			Iri: iri,
		},
	}
}

func BlankNodeObject(blankNode string) *rdfpb.Object {
	return &rdfpb.Object{
		Value: &rdfpb.Object_BlankNode{
			BlankNode: blankNode,
		},
	}
}

func LiteralObject(literal *rdfpb.Literal) *rdfpb.Object {
	return &rdfpb.Object{
		Value: &rdfpb.Object_Literal{
			Literal: literal,
		},
	}
}

func StringLiteral(s string) *rdfpb.Literal {
	return &rdfpb.Literal{
		Value: &rdfpb.Literal_String_{
			String_: s,
		},
	}
}

func StringObject(s string) *rdfpb.Object {
	return LiteralObject(StringLiteral(s))
}

func IntLiteral(i int64) *rdfpb.Literal {
	return &rdfpb.Literal{
		Value: &rdfpb.Literal_Int{
			Int: i,
		},
	}
}

func IntObject(i int64) *rdfpb.Object {
	return LiteralObject(IntLiteral(i))
}

func Float64Literal(f float64) *rdfpb.Literal {
	return &rdfpb.Literal{
		Value: &rdfpb.Literal_Float64{
			Float64: f,
		},
	}
}

func Float64Object(f float64) *rdfpb.Object {
	return LiteralObject(Float64Literal(f))
}

func BoolLiteral(b bool) *rdfpb.Literal {
	return &rdfpb.Literal{
		Value: &rdfpb.Literal_Bool{
			Bool: b,
		},
	}
}

func BoolObject(b bool) *rdfpb.Object {
	return LiteralObject(BoolLiteral(b))
}

func BytesLiteral(b []byte) *rdfpb.Literal {
	return &rdfpb.Literal{
		Value: &rdfpb.Literal_Bytes{
			Bytes: b,
		},
	}
}

func BytesObject(b []byte) *rdfpb.Object {
	return LiteralObject(BytesLiteral(b))
}
