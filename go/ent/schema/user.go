package schema

import (
	valueobject "ESTIHUB/services/user/domain/valueObject"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(func() uuid.UUID {
				id, err := uuid.NewV7()
				if err != nil {
					panic("failed to generate UUID: " + err.Error())
				}
				return id
			}),
		field.String("first_name").
			MinLen(valueobject.FirstNameMinLength).
			MaxLen(valueobject.FirstNameMaxLength),
		field.String("last_name").
			MinLen(valueobject.LastNameMinLength).
			MaxLen(valueobject.LastNameMaxLength),
		field.Enum("role").
			Values("general", "admin", "super_user").
			Default("general"),
		field.Time("updated_at").Default(time.Now),
	}
}
