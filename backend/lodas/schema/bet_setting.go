package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type BetSetting struct {
	ent.Schema
}

func (BetSetting) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		//map bettype (enum) to scale
		field.JSON("values", &BetSettingMap{}),
		field.Time("created_time"),
	}
}

func (BetSetting) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("sheets", Sheet.Type),
	}
}

type BetSettingMap struct {
	LO2 int64
	LO3 int64
	DE2 int64
	DE3 int64
	DA  int64
}
