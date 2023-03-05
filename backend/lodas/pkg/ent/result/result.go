// Code generated by ent, DO NOT EDIT.

package result

import (
	"time"
)

const (
	// Label holds the string label denoting the result type in the database.
	Label = "result"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldPrizeMap holds the string denoting the prize_map field in the database.
	FieldPrizeMap = "prize_map"
	// Table holds the table name of the result in the database.
	Table = "results"
)

// Columns holds all SQL columns for result fields.
var Columns = []string{
	FieldID,
	FieldProvince,
	FieldCreatedTime,
	FieldPrizeMap,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultProvince holds the default value on creation for the "province" field.
	DefaultProvince int64
	// DefaultCreatedTime holds the default value on creation for the "created_time" field.
	DefaultCreatedTime time.Time
)
