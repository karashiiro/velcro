package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LogEvent holds the schema definition for the LogEvent entity.
type LogEvent struct {
	ent.Schema
}

// Fields of the LogEvent.
func (LogEvent) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp"),
		field.Int("level"),
		field.String("message"),
	}
}

// Edges of the LogEvent.
func (LogEvent) Edges() []ent.Edge {
	return nil
}
