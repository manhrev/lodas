package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

type Record struct {
	ent.Schema
}

func (Record) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Strings("numbers"),
		field.Int64("cash_amount"),
		field.Int64("bet_type").Default(int64(lodas_pb.BetType_BET_TYPE_UNSPECIFIED)),
		field.Ints("prize"),
		field.Int64("cash_in").Optional(),
		field.Int64("cash_out").Optional(),
		field.Time("created_time"),
		field.JSON("win_info", &PrizeMap{}).Optional(),
	}
}

func (Record) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("sheet", Sheet.Type).
			Ref("records").
			Unique(),
	}
}
