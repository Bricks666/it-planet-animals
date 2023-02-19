// Code generated by ent, DO NOT EDIT.

package location

const (
	// Label holds the string label denoting the location type in the database.
	Label = "location"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// Table holds the table name of the location in the database.
	Table = "locations"
)

// Columns holds all SQL columns for location fields.
var Columns = []string{
	FieldID,
	FieldLatitude,
	FieldLongitude,
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
	// LatitudeValidator is a validator for the "latitude" field. It is called by the builders before save.
	LatitudeValidator func(float64) error
	// LongitudeValidator is a validator for the "longitude" field. It is called by the builders before save.
	LongitudeValidator func(float64) error
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int64) error
)