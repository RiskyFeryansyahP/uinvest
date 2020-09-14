package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Indexes hold index field of user entity
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email", "password").
			Unique(),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("firstname").
			NotEmpty(),
		field.String("lastname").
			NotEmpty(),
		field.String("email").
			NotEmpty().
			Unique(),
		field.String("phonenumber").
			Unique().
			NotEmpty(),
		field.String("password").
			NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
