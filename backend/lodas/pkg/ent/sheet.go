// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/betsetting"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
)

// Sheet is the model entity for the Sheet schema.
type Sheet struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Status holds the value of the "status" field.
	Status int64 `json:"status,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Area holds the value of the "area" field.
	Area int64 `json:"area,omitempty"`
	// Province holds the value of the "province" field.
	Province int64 `json:"province,omitempty"`
	// Ratio holds the value of the "ratio" field.
	Ratio float64 `json:"ratio,omitempty"`
	// WinRatio holds the value of the "win_ratio" field.
	WinRatio float64 `json:"win_ratio,omitempty"`
	// TotalCashin holds the value of the "total_cashin" field.
	TotalCashin int64 `json:"total_cashin,omitempty"`
	// TotalCashout holds the value of the "total_cashout" field.
	TotalCashout int64 `json:"total_cashout,omitempty"`
	// ResultTime holds the value of the "result_time" field.
	ResultTime time.Time `json:"result_time,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// UpdatedTime holds the value of the "updated_time" field.
	UpdatedTime time.Time `json:"updated_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SheetQuery when eager-loading is set.
	Edges              SheetEdges `json:"edges"`
	bet_setting_sheets *int64
}

// SheetEdges holds the relations/edges for other nodes in the graph.
type SheetEdges struct {
	// Records holds the value of the records edge.
	Records []*Record `json:"records,omitempty"`
	// BetSetting holds the value of the bet_setting edge.
	BetSetting *BetSetting `json:"bet_setting,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RecordsOrErr returns the Records value or an error if the edge
// was not loaded in eager-loading.
func (e SheetEdges) RecordsOrErr() ([]*Record, error) {
	if e.loadedTypes[0] {
		return e.Records, nil
	}
	return nil, &NotLoadedError{edge: "records"}
}

// BetSettingOrErr returns the BetSetting value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SheetEdges) BetSettingOrErr() (*BetSetting, error) {
	if e.loadedTypes[1] {
		if e.BetSetting == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: betsetting.Label}
		}
		return e.BetSetting, nil
	}
	return nil, &NotLoadedError{edge: "bet_setting"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Sheet) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case sheet.FieldRatio, sheet.FieldWinRatio:
			values[i] = new(sql.NullFloat64)
		case sheet.FieldID, sheet.FieldStatus, sheet.FieldArea, sheet.FieldProvince, sheet.FieldTotalCashin, sheet.FieldTotalCashout, sheet.FieldUserID:
			values[i] = new(sql.NullInt64)
		case sheet.FieldName:
			values[i] = new(sql.NullString)
		case sheet.FieldResultTime, sheet.FieldCreatedTime, sheet.FieldUpdatedTime:
			values[i] = new(sql.NullTime)
		case sheet.ForeignKeys[0]: // bet_setting_sheets
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Sheet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Sheet fields.
func (s *Sheet) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case sheet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case sheet.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = value.Int64
			}
		case sheet.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case sheet.FieldArea:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field area", values[i])
			} else if value.Valid {
				s.Area = value.Int64
			}
		case sheet.FieldProvince:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				s.Province = value.Int64
			}
		case sheet.FieldRatio:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field ratio", values[i])
			} else if value.Valid {
				s.Ratio = value.Float64
			}
		case sheet.FieldWinRatio:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field win_ratio", values[i])
			} else if value.Valid {
				s.WinRatio = value.Float64
			}
		case sheet.FieldTotalCashin:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_cashin", values[i])
			} else if value.Valid {
				s.TotalCashin = value.Int64
			}
		case sheet.FieldTotalCashout:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_cashout", values[i])
			} else if value.Valid {
				s.TotalCashout = value.Int64
			}
		case sheet.FieldResultTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field result_time", values[i])
			} else if value.Valid {
				s.ResultTime = value.Time
			}
		case sheet.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				s.CreatedTime = value.Time
			}
		case sheet.FieldUpdatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_time", values[i])
			} else if value.Valid {
				s.UpdatedTime = value.Time
			}
		case sheet.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				s.UserID = value.Int64
			}
		case sheet.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field bet_setting_sheets", value)
			} else if value.Valid {
				s.bet_setting_sheets = new(int64)
				*s.bet_setting_sheets = int64(value.Int64)
			}
		}
	}
	return nil
}

// QueryRecords queries the "records" edge of the Sheet entity.
func (s *Sheet) QueryRecords() *RecordQuery {
	return NewSheetClient(s.config).QueryRecords(s)
}

// QueryBetSetting queries the "bet_setting" edge of the Sheet entity.
func (s *Sheet) QueryBetSetting() *BetSettingQuery {
	return NewSheetClient(s.config).QueryBetSetting(s)
}

// Update returns a builder for updating this Sheet.
// Note that you need to call Sheet.Unwrap() before calling this method if this Sheet
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Sheet) Update() *SheetUpdateOne {
	return NewSheetClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Sheet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Sheet) Unwrap() *Sheet {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Sheet is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Sheet) String() string {
	var builder strings.Builder
	builder.WriteString("Sheet(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("area=")
	builder.WriteString(fmt.Sprintf("%v", s.Area))
	builder.WriteString(", ")
	builder.WriteString("province=")
	builder.WriteString(fmt.Sprintf("%v", s.Province))
	builder.WriteString(", ")
	builder.WriteString("ratio=")
	builder.WriteString(fmt.Sprintf("%v", s.Ratio))
	builder.WriteString(", ")
	builder.WriteString("win_ratio=")
	builder.WriteString(fmt.Sprintf("%v", s.WinRatio))
	builder.WriteString(", ")
	builder.WriteString("total_cashin=")
	builder.WriteString(fmt.Sprintf("%v", s.TotalCashin))
	builder.WriteString(", ")
	builder.WriteString("total_cashout=")
	builder.WriteString(fmt.Sprintf("%v", s.TotalCashout))
	builder.WriteString(", ")
	builder.WriteString("result_time=")
	builder.WriteString(s.ResultTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_time=")
	builder.WriteString(s.CreatedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_time=")
	builder.WriteString(s.UpdatedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", s.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// Sheets is a parsable slice of Sheet.
type Sheets []*Sheet

func (s Sheets) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
