// Code generated by ent, DO NOT EDIT.

package record

const (
	// Label holds the string label denoting the record type in the database.
	Label = "record"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldNumbers holds the string denoting the numbers field in the database.
	FieldNumbers = "numbers"
	// FieldCashAmount holds the string denoting the cash_amount field in the database.
	FieldCashAmount = "cash_amount"
	// FieldBetType holds the string denoting the bet_type field in the database.
	FieldBetType = "bet_type"
	// FieldPrize holds the string denoting the prize field in the database.
	FieldPrize = "prize"
	// FieldCashIn holds the string denoting the cash_in field in the database.
	FieldCashIn = "cash_in"
	// FieldCashOut holds the string denoting the cash_out field in the database.
	FieldCashOut = "cash_out"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldWinInfo holds the string denoting the win_info field in the database.
	FieldWinInfo = "win_info"
	// EdgeSheet holds the string denoting the sheet edge name in mutations.
	EdgeSheet = "sheet"
	// Table holds the table name of the record in the database.
	Table = "records"
	// SheetTable is the table that holds the sheet relation/edge.
	SheetTable = "records"
	// SheetInverseTable is the table name for the Sheet entity.
	// It exists in this package in order to avoid circular dependency with the "sheet" package.
	SheetInverseTable = "sheets"
	// SheetColumn is the table column denoting the sheet relation/edge.
	SheetColumn = "sheet_records"
)

// Columns holds all SQL columns for record fields.
var Columns = []string{
	FieldID,
	FieldNumbers,
	FieldCashAmount,
	FieldBetType,
	FieldPrize,
	FieldCashIn,
	FieldCashOut,
	FieldCreatedTime,
	FieldWinInfo,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "records"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sheet_records",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultBetType holds the default value on creation for the "bet_type" field.
	DefaultBetType int64
)
