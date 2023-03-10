package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

type Sheet struct {
	ent.Schema
}

func (Sheet) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("status").Default(int64(lodas_pb.SheetStatus_SHEET_STATUS_NOT_SUBMITTED)),
		field.String("name"),
		field.Int64("area").Default(int64(lodas_pb.Area_AREA_UNSPECIFIED)),
		field.Int64("province").Default(int64(lodas_pb.Province_UNSPECIFIED)),
		field.Float("ratio").Default(0),
		field.Float("win_ratio").Default(1),
		field.Int64("total_cashin").Default(0),
		field.Int64("total_cashout").Default(0),
		field.Time("result_time"),
		field.Time("created_time"),
		field.Time("updated_time"),
		field.Int64("user_id"),
	}
}

func (Sheet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("records", Record.Type),
		edge.From("bet_setting", BetSetting.Type).
			Ref("sheets").
			Unique(),
	}
}
