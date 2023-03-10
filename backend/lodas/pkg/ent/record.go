// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/record"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/schema"
	"github.com/manhrev/lodas/backend/lodas/pkg/ent/sheet"
)

// Record is the model entity for the Record schema.
type Record struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Numbers holds the value of the "numbers" field.
	Numbers []string `json:"numbers,omitempty"`
	// CashAmount holds the value of the "cash_amount" field.
	CashAmount int64 `json:"cash_amount,omitempty"`
	// BetType holds the value of the "bet_type" field.
	BetType int64 `json:"bet_type,omitempty"`
	// Prize holds the value of the "prize" field.
	Prize []int `json:"prize,omitempty"`
	// CashIn holds the value of the "cash_in" field.
	CashIn int64 `json:"cash_in,omitempty"`
	// CashOut holds the value of the "cash_out" field.
	CashOut int64 `json:"cash_out,omitempty"`
	// CreatedTime holds the value of the "created_time" field.
	CreatedTime time.Time `json:"created_time,omitempty"`
	// WinInfo holds the value of the "win_info" field.
	WinInfo *schema.PrizeMap `json:"win_info,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RecordQuery when eager-loading is set.
	Edges         RecordEdges `json:"edges"`
	sheet_records *int64
}

// RecordEdges holds the relations/edges for other nodes in the graph.
type RecordEdges struct {
	// Sheet holds the value of the sheet edge.
	Sheet *Sheet `json:"sheet,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// SheetOrErr returns the Sheet value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RecordEdges) SheetOrErr() (*Sheet, error) {
	if e.loadedTypes[0] {
		if e.Sheet == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: sheet.Label}
		}
		return e.Sheet, nil
	}
	return nil, &NotLoadedError{edge: "sheet"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Record) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case record.FieldNumbers, record.FieldPrize, record.FieldWinInfo:
			values[i] = new([]byte)
		case record.FieldID, record.FieldCashAmount, record.FieldBetType, record.FieldCashIn, record.FieldCashOut:
			values[i] = new(sql.NullInt64)
		case record.FieldCreatedTime:
			values[i] = new(sql.NullTime)
		case record.ForeignKeys[0]: // sheet_records
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Record", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Record fields.
func (r *Record) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int64(value.Int64)
		case record.FieldNumbers:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field numbers", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Numbers); err != nil {
					return fmt.Errorf("unmarshal field numbers: %w", err)
				}
			}
		case record.FieldCashAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cash_amount", values[i])
			} else if value.Valid {
				r.CashAmount = value.Int64
			}
		case record.FieldBetType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bet_type", values[i])
			} else if value.Valid {
				r.BetType = value.Int64
			}
		case record.FieldPrize:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field prize", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Prize); err != nil {
					return fmt.Errorf("unmarshal field prize: %w", err)
				}
			}
		case record.FieldCashIn:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cash_in", values[i])
			} else if value.Valid {
				r.CashIn = value.Int64
			}
		case record.FieldCashOut:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cash_out", values[i])
			} else if value.Valid {
				r.CashOut = value.Int64
			}
		case record.FieldCreatedTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_time", values[i])
			} else if value.Valid {
				r.CreatedTime = value.Time
			}
		case record.FieldWinInfo:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field win_info", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.WinInfo); err != nil {
					return fmt.Errorf("unmarshal field win_info: %w", err)
				}
			}
		case record.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field sheet_records", value)
			} else if value.Valid {
				r.sheet_records = new(int64)
				*r.sheet_records = int64(value.Int64)
			}
		}
	}
	return nil
}

// QuerySheet queries the "sheet" edge of the Record entity.
func (r *Record) QuerySheet() *SheetQuery {
	return NewRecordClient(r.config).QuerySheet(r)
}

// Update returns a builder for updating this Record.
// Note that you need to call Record.Unwrap() before calling this method if this Record
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Record) Update() *RecordUpdateOne {
	return NewRecordClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Record entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Record) Unwrap() *Record {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Record is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Record) String() string {
	var builder strings.Builder
	builder.WriteString("Record(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("numbers=")
	builder.WriteString(fmt.Sprintf("%v", r.Numbers))
	builder.WriteString(", ")
	builder.WriteString("cash_amount=")
	builder.WriteString(fmt.Sprintf("%v", r.CashAmount))
	builder.WriteString(", ")
	builder.WriteString("bet_type=")
	builder.WriteString(fmt.Sprintf("%v", r.BetType))
	builder.WriteString(", ")
	builder.WriteString("prize=")
	builder.WriteString(fmt.Sprintf("%v", r.Prize))
	builder.WriteString(", ")
	builder.WriteString("cash_in=")
	builder.WriteString(fmt.Sprintf("%v", r.CashIn))
	builder.WriteString(", ")
	builder.WriteString("cash_out=")
	builder.WriteString(fmt.Sprintf("%v", r.CashOut))
	builder.WriteString(", ")
	builder.WriteString("created_time=")
	builder.WriteString(r.CreatedTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("win_info=")
	builder.WriteString(fmt.Sprintf("%v", r.WinInfo))
	builder.WriteByte(')')
	return builder.String()
}

// Records is a parsable slice of Record.
type Records []*Record

func (r Records) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
