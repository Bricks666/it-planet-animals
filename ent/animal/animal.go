// Code generated by ent, DO NOT EDIT.

package animal

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the animal type in the database.
	Label = "animal"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldHeight holds the string denoting the height field in the database.
	FieldHeight = "height"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldLifeStatus holds the string denoting the life_status field in the database.
	FieldLifeStatus = "life_status"
	// FieldChippingDateTime holds the string denoting the chipping_date_time field in the database.
	FieldChippingDateTime = "chipping_date_time"
	// FieldChipperID holds the string denoting the chipper_id field in the database.
	FieldChipperID = "chipper_id"
	// FieldChippingLocationID holds the string denoting the chipping_location_id field in the database.
	FieldChippingLocationID = "chipping_location_id"
	// FieldDeathDateTime holds the string denoting the death_date_time field in the database.
	FieldDeathDateTime = "death_date_time"
	// EdgeChipper holds the string denoting the chipper edge name in mutations.
	EdgeChipper = "chipper"
	// EdgeTypes holds the string denoting the types edge name in mutations.
	EdgeTypes = "types"
	// EdgeChippingLocation holds the string denoting the chipping_location edge name in mutations.
	EdgeChippingLocation = "chipping_location"
	// EdgeLocations holds the string denoting the locations edge name in mutations.
	EdgeLocations = "locations"
	// EdgeAnimalTypes holds the string denoting the animal_types edge name in mutations.
	EdgeAnimalTypes = "animal_types"
	// EdgeVisitedLocations holds the string denoting the visited_locations edge name in mutations.
	EdgeVisitedLocations = "visited_locations"
	// Table holds the table name of the animal in the database.
	Table = "animals"
	// ChipperTable is the table that holds the chipper relation/edge.
	ChipperTable = "animals"
	// ChipperInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ChipperInverseTable = "users"
	// ChipperColumn is the table column denoting the chipper relation/edge.
	ChipperColumn = "chipper_id"
	// TypesTable is the table that holds the types relation/edge. The primary key declared below.
	TypesTable = "animal_types"
	// TypesInverseTable is the table name for the AnimalTag entity.
	// It exists in this package in order to avoid circular dependency with the "animaltag" package.
	TypesInverseTable = "animal_tags"
	// ChippingLocationTable is the table that holds the chipping_location relation/edge.
	ChippingLocationTable = "animals"
	// ChippingLocationInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	ChippingLocationInverseTable = "locations"
	// ChippingLocationColumn is the table column denoting the chipping_location relation/edge.
	ChippingLocationColumn = "chipping_location_id"
	// LocationsTable is the table that holds the locations relation/edge. The primary key declared below.
	LocationsTable = "visited_locations"
	// LocationsInverseTable is the table name for the Location entity.
	// It exists in this package in order to avoid circular dependency with the "location" package.
	LocationsInverseTable = "locations"
	// AnimalTypesTable is the table that holds the animal_types relation/edge.
	AnimalTypesTable = "animal_types"
	// AnimalTypesInverseTable is the table name for the AnimalType entity.
	// It exists in this package in order to avoid circular dependency with the "animaltype" package.
	AnimalTypesInverseTable = "animal_types"
	// AnimalTypesColumn is the table column denoting the animal_types relation/edge.
	AnimalTypesColumn = "animal_id"
	// VisitedLocationsTable is the table that holds the visited_locations relation/edge.
	VisitedLocationsTable = "visited_locations"
	// VisitedLocationsInverseTable is the table name for the VisitedLocation entity.
	// It exists in this package in order to avoid circular dependency with the "visitedlocation" package.
	VisitedLocationsInverseTable = "visited_locations"
	// VisitedLocationsColumn is the table column denoting the visited_locations relation/edge.
	VisitedLocationsColumn = "animal_id"
)

// Columns holds all SQL columns for animal fields.
var Columns = []string{
	FieldID,
	FieldWeight,
	FieldLength,
	FieldHeight,
	FieldGender,
	FieldLifeStatus,
	FieldChippingDateTime,
	FieldChipperID,
	FieldChippingLocationID,
	FieldDeathDateTime,
}

var (
	// TypesPrimaryKey and TypesColumn2 are the table columns denoting the
	// primary key for the types relation (M2M).
	TypesPrimaryKey = []string{"animal_id", "type_id"}
	// LocationsPrimaryKey and LocationsColumn2 are the table columns denoting the
	// primary key for the locations relation (M2M).
	LocationsPrimaryKey = []string{"animal_id", "location_id"}
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
	// WeightValidator is a validator for the "weight" field. It is called by the builders before save.
	WeightValidator func(float32) error
	// LengthValidator is a validator for the "length" field. It is called by the builders before save.
	LengthValidator func(float32) error
	// HeightValidator is a validator for the "height" field. It is called by the builders before save.
	HeightValidator func(float32) error
	// DefaultChippingDateTime holds the default value on creation for the "chipping_date_time" field.
	DefaultChippingDateTime func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(uint64) error
)

// Gender defines the type for the "gender" enum field.
type Gender string

// Gender values.
const (
	GenderMALE   Gender = "MALE"
	GenderFEMALE Gender = "FEMALE"
	GenderOTHER  Gender = "OTHER"
)

func (ge Gender) String() string {
	return string(ge)
}

// GenderValidator is a validator for the "gender" field enum values. It is called by the builders before save.
func GenderValidator(ge Gender) error {
	switch ge {
	case GenderMALE, GenderFEMALE, GenderOTHER:
		return nil
	default:
		return fmt.Errorf("animal: invalid enum value for gender field: %q", ge)
	}
}

// LifeStatus defines the type for the "life_status" enum field.
type LifeStatus string

// LifeStatusALIVE is the default value of the LifeStatus enum.
const DefaultLifeStatus = LifeStatusALIVE

// LifeStatus values.
const (
	LifeStatusALIVE LifeStatus = "ALIVE"
	LifeStatusDEAD  LifeStatus = "DEAD"
)

func (ls LifeStatus) String() string {
	return string(ls)
}

// LifeStatusValidator is a validator for the "life_status" field enum values. It is called by the builders before save.
func LifeStatusValidator(ls LifeStatus) error {
	switch ls {
	case LifeStatusALIVE, LifeStatusDEAD:
		return nil
	default:
		return fmt.Errorf("animal: invalid enum value for life_status field: %q", ls)
	}
}
