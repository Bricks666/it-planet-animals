// Code generated by ent, DO NOT EDIT.

package animaltype

const (
	// Label holds the string label denoting the animaltype type in the database.
	Label = "animal_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// Table holds the table name of the animaltype in the database.
	Table = "animal_types"
)

// Columns holds all SQL columns for animaltype fields.
var Columns = []string{
	FieldID,
	FieldType,
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
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)
