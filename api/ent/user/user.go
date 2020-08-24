// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstname holds the string denoting the firstname field in the database.
	FieldFirstname = "firstname"
	// FieldLastname holds the string denoting the lastname field in the database.
	FieldLastname = "lastname"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhonenumber holds the string denoting the phonenumber field in the database.
	FieldPhonenumber = "phonenumber"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFirstname,
	FieldLastname,
	FieldEmail,
	FieldPhonenumber,
	FieldPassword,
}

var (
	// FirstnameValidator is a validator for the "firstname" field. It is called by the builders before save.
	FirstnameValidator func(string) error
	// LastnameValidator is a validator for the "lastname" field. It is called by the builders before save.
	LastnameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PhonenumberValidator is a validator for the "phonenumber" field. It is called by the builders before save.
	PhonenumberValidator func(string) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
)