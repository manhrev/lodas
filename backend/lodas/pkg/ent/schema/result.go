package schema

import (
	"encoding/json"
	"time"

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
		field.Time("created_time").Default(time.Now().In(time.FixedZone("UTC+7", 7*60*60))),
		//map prize (enum) to string
		field.JSON("prize_map", &PrizeMap{}),
	}
}

type PrizeMap map[lodas_pb.Prize]string

func (p *PrizeMap) UnmarshalByte(data []byte) error {
	return json.Unmarshal(data, p)
}
