// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/xmlking/grpc-starter-kit/ent/profile"
	"github.com/xmlking/grpc-starter-kit/ent/schema"
	"github.com/xmlking/grpc-starter-kit/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	profileMixin := schema.Profile{}.Mixin()
	profileMixinFields0 := profileMixin[0].Fields()
	_ = profileMixinFields0
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescCreateTime is the schema descriptor for create_time field.
	profileDescCreateTime := profileMixinFields0[0].Descriptor()
	// profile.DefaultCreateTime holds the default value on creation for the create_time field.
	profile.DefaultCreateTime = profileDescCreateTime.Default.(func() time.Time)
	// profileDescUpdateTime is the schema descriptor for update_time field.
	profileDescUpdateTime := profileMixinFields0[1].Descriptor()
	// profile.DefaultUpdateTime holds the default value on creation for the update_time field.
	profile.DefaultUpdateTime = profileDescUpdateTime.Default.(func() time.Time)
	// profile.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	profile.UpdateDefaultUpdateTime = profileDescUpdateTime.UpdateDefault.(func() time.Time)
	// profileDescAge is the schema descriptor for age field.
	profileDescAge := profileFields[1].Descriptor()
	// profile.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	profile.AgeValidator = profileDescAge.Validators[0].(func(int) error)
	// profileDescID is the schema descriptor for id field.
	profileDescID := profileFields[0].Descriptor()
	// profile.DefaultID holds the default value on creation for the id field.
	profile.DefaultID = profileDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[2].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[3].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[4].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = func() func(string) error {
		validators := userDescEmail.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
			validators[3].(func(string) error),
		}
		return func(email string) error {
			for _, fn := range fns {
				if err := fn(email); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescTenant is the schema descriptor for tenant field.
	userDescTenant := userFields[5].Descriptor()
	// user.DefaultTenant holds the default value on creation for the tenant field.
	user.DefaultTenant = userDescTenant.Default.(string)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
