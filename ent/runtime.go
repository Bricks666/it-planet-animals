// Code generated by ent, DO NOT EDIT.

package ent

import (
	"animals/ent/animaltype"
	"animals/ent/location"
	"animals/ent/schema"
	"animals/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	animaltypeFields := schema.AnimalType{}.Fields()
	_ = animaltypeFields
	// animaltypeDescType is the schema descriptor for type field.
	animaltypeDescType := animaltypeFields[1].Descriptor()
	// animaltype.TypeValidator is a validator for the "type" field. It is called by the builders before save.
	animaltype.TypeValidator = animaltypeDescType.Validators[0].(func(string) error)
	// animaltypeDescID is the schema descriptor for id field.
	animaltypeDescID := animaltypeFields[0].Descriptor()
	// animaltype.IDValidator is a validator for the "id" field. It is called by the builders before save.
	animaltype.IDValidator = animaltypeDescID.Validators[0].(func(int64) error)
	locationFields := schema.Location{}.Fields()
	_ = locationFields
	// locationDescLatitude is the schema descriptor for latitude field.
	locationDescLatitude := locationFields[1].Descriptor()
	// location.LatitudeValidator is a validator for the "latitude" field. It is called by the builders before save.
	location.LatitudeValidator = locationDescLatitude.Validators[0].(func(float64) error)
	// locationDescLongitude is the schema descriptor for longitude field.
	locationDescLongitude := locationFields[2].Descriptor()
	// location.LongitudeValidator is a validator for the "longitude" field. It is called by the builders before save.
	location.LongitudeValidator = locationDescLongitude.Validators[0].(func(float64) error)
	// locationDescID is the schema descriptor for id field.
	locationDescID := locationFields[0].Descriptor()
	// location.IDValidator is a validator for the "id" field. It is called by the builders before save.
	location.IDValidator = locationDescID.Validators[0].(func(int64) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescFirstName is the schema descriptor for firstName field.
	userDescFirstName := userFields[3].Descriptor()
	// user.FirstNameValidator is a validator for the "firstName" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for lastName field.
	userDescLastName := userFields[4].Descriptor()
	// user.LastNameValidator is a validator for the "lastName" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(uint32) error)
}
