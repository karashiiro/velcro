package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.Time("timestamp"),
		field.Int("version"),
		field.Int("segment").Default(0),
		field.Int("opcode").Optional().Nillable(),
		field.String("source_address"),
		field.Int("source_port"),
		field.String("destination_address"),
		field.Int("destination_port"),
		field.Bytes("data"),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}
