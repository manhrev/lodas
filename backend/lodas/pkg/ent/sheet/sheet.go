// Code generated by ent, DO NOT EDIT.

package sheet

const (
	// Label holds the string label denoting the sheet type in the database.
	Label = "sheet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldArea holds the string denoting the area field in the database.
	FieldArea = "area"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldRatio holds the string denoting the ratio field in the database.
	FieldRatio = "ratio"
	// FieldWinRatio holds the string denoting the win_ratio field in the database.
	FieldWinRatio = "win_ratio"
	// FieldTotalCashin holds the string denoting the total_cashin field in the database.
	FieldTotalCashin = "total_cashin"
	// FieldTotalCashout holds the string denoting the total_cashout field in the database.
	FieldTotalCashout = "total_cashout"
	// FieldResultTime holds the string denoting the result_time field in the database.
	FieldResultTime = "result_time"
	// FieldCreatedTime holds the string denoting the created_time field in the database.
	FieldCreatedTime = "created_time"
	// FieldUpdatedTime holds the string denoting the updated_time field in the database.
	FieldUpdatedTime = "updated_time"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// EdgeRecords holds the string denoting the records edge name in mutations.
	EdgeRecords = "records"
	// EdgeBetSetting holds the string denoting the bet_setting edge name in mutations.
	EdgeBetSetting = "bet_setting"
	// Table holds the table name of the sheet in the database.
	Table = "sheets"
	// RecordsTable is the table that holds the records relation/edge.
	RecordsTable = "records"
	// RecordsInverseTable is the table name for the Record entity.
	// It exists in this package in order to avoid circular dependency with the "record" package.
	RecordsInverseTable = "records"
	// RecordsColumn is the table column denoting the records relation/edge.
	RecordsColumn = "sheet_records"
	// BetSettingTable is the table that holds the bet_setting relation/edge.
	BetSettingTable = "sheets"
	// BetSettingInverseTable is the table name for the BetSetting entity.
	// It exists in this package in order to avoid circular dependency with the "betsetting" package.
	BetSettingInverseTable = "bet_settings"
	// BetSettingColumn is the table column denoting the bet_setting relation/edge.
	BetSettingColumn = "bet_setting_sheets"
)

// Columns holds all SQL columns for sheet fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldName,
	FieldArea,
	FieldProvince,
	FieldRatio,
	FieldWinRatio,
	FieldTotalCashin,
	FieldTotalCashout,
	FieldResultTime,
	FieldCreatedTime,
	FieldUpdatedTime,
	FieldUserID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sheets"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"bet_setting_sheets",
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultArea holds the default value on creation for the "area" field.
	DefaultArea int64
	// DefaultProvince holds the default value on creation for the "province" field.
	DefaultProvince int64
	// DefaultRatio holds the default value on creation for the "ratio" field.
	DefaultRatio float64
	// DefaultWinRatio holds the default value on creation for the "win_ratio" field.
	DefaultWinRatio float64
	// DefaultTotalCashin holds the default value on creation for the "total_cashin" field.
	DefaultTotalCashin int64
	// DefaultTotalCashout holds the default value on creation for the "total_cashout" field.
	DefaultTotalCashout int64
)
