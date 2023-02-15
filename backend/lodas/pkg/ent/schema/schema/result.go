package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	lodas_pb "github.com/manhrev/lodas/backend/lodas/pkg/api"
)

type Result struct {
	ent.Schema
}

func (Result) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Unique(),
		field.Int64("province").Default(int64(lodas_pb.Province_UNSPECIFIED)),
		field.Time("created_time"),
		//map prize (enum) to string
		field.JSON("prize_map", &PrizeMap{}),
	}
}

type PrizeMap struct {
	PRIZE_1DB string
	PRIZE_1G1 string

	PRIZE_1G2 string
	PRIZE_2G2 string

	PRIZE_1G3 string
	PRIZE_2G3 string
	PRIZE_3G3 string
	PRIZE_4G3 string
	PRIZE_5G3 string
	PRIZE_6G3 string

	PRIZE_1G4 string
	PRIZE_2G4 string
	PRIZE_3G4 string
	PRIZE_4G4 string
	PRIZE_5G4 string
	PRIZE_6G4 string
	PRIZE_7G4 string

	PRIZE_1G5 string
	PRIZE_2G5 string
	PRIZE_3G5 string
	PRIZE_4G5 string
	PRIZE_5G5 string
	PRIZE_6G5 string

	PRIZE_1G6 string
	PRIZE_2G6 string
	PRIZE_3G6 string

	PRIZE_1G7 string
	PRIZE_2G7 string
	PRIZE_3G7 string
	PRIZE_4G7 string

	PRIZE_1G8 string
}
