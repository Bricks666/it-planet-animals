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
	// EdgeVisitedLocationsLocation holds the string denoting the visited_locations_location edge name in mutations.
	EdgeVisitedLocationsLocation = "visited_locations_location"
	// Table holds the table name of the location in the database.
	Table = "locations"
	// VisitedLocationsLocationTable is the table that holds the visited_locations_location relation/edge. The primary key declared below.
	VisitedLocationsLocationTable = "animal_visited_locations_animals"
	// VisitedLocationsLocationInverseTable is the table name for the Animal entity.
	// It exists in this package in order to avoid circular dependency with the "animal" package.
	VisitedLocationsLocationInverseTable = "animals"
)

// Columns holds all SQL columns for location fields.
var Columns = []string{
	FieldID,
	FieldLatitude,
	FieldLongitude,
}

var (
	// VisitedLocationsLocationPrimaryKey and VisitedLocationsLocationColumn2 are the table columns denoting the
	// primary key for the visited_locations_location relation (M2M).
	VisitedLocationsLocationPrimaryKey = []string{"animal_id", "location_id"}
)

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
	IDValidator func(uint64) error
)
