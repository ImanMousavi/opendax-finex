package schema

import (
	"errors"
	"regexp"
	"strings"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Asset struct {
	ent.Schema
}

func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			Unique().
			Immutable().
			NotEmpty().
			MinLen(3).
			MaxLen(5).
			Match(regexp.MustCompile("^[A-Z]+$")).
			Annotations(entproto.Field(1)),
		field.String("name").
			NotEmpty().
			MinLen(3).
			MaxLen(100).
			Match(regexp.MustCompile("^[a-zA-Z ]+$")).
			Validate(func(s string) error {
				if strings.TrimSpace(s) != s {
					return errors.New("asset name must not begin or end with white spaces")
				}
				if strings.ToLower(s) == s {
					return errors.New("asset name must begin with uppercase")
				}
				return nil
			}).
			Annotations(entproto.Field(2)),
		field.Uint32("index").
			Unique().
			Annotations(entproto.Field(3)),
	}
}

func (Asset) Edges() []ent.Edge {
	return nil
}

func (Asset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
