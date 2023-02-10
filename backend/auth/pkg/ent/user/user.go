// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldDisplayName holds the string denoting the display_name field in the database.
	FieldDisplayName = "display_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldPassword,
	FieldDisplayName,
	FieldEmail,
	FieldPhone,
	FieldRole,
	FieldAge,
	FieldCreatedAt,
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
	// DefaultDisplayName holds the default value on creation for the "display_name" field.
	DefaultDisplayName string
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// DefaultRole holds the default value on creation for the "role" field.
	DefaultRole int32
	// DefaultAge holds the default value on creation for the "age" field.
	DefaultAge int32
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)
