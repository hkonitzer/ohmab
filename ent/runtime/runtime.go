// OHMAB
// Code generated by entc, DO NOT EDIT.

package runtime

import (
	"time"

	"github.com/google/uuid"
	"hynie.de/ohmab/ent/address"
	"hynie.de/ohmab/ent/auditlog"
	"hynie.de/ohmab/ent/business"
	"hynie.de/ohmab/ent/schema"
	"hynie.de/ohmab/ent/tag"
	"hynie.de/ohmab/ent/timetable"
	"hynie.de/ohmab/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	addressMixin := schema.Address{}.Mixin()
	addressHooks := schema.Address{}.Hooks()
	address.Hooks[0] = addressHooks[0]
	addressMixinFields0 := addressMixin[0].Fields()
	_ = addressMixinFields0
	addressFields := schema.Address{}.Fields()
	_ = addressFields
	// addressDescCreatedAt is the schema descriptor for created_at field.
	addressDescCreatedAt := addressMixinFields0[0].Descriptor()
	// address.DefaultCreatedAt holds the default value on creation for the created_at field.
	address.DefaultCreatedAt = addressDescCreatedAt.Default.(func() time.Time)
	// addressDescUpdatedAt is the schema descriptor for updated_at field.
	addressDescUpdatedAt := addressMixinFields0[1].Descriptor()
	// address.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	address.DefaultUpdatedAt = addressDescUpdatedAt.Default.(func() time.Time)
	// address.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	address.UpdateDefaultUpdatedAt = addressDescUpdatedAt.UpdateDefault.(func() time.Time)
	// addressDescPrimary is the schema descriptor for primary field.
	addressDescPrimary := addressFields[7].Descriptor()
	// address.DefaultPrimary holds the default value on creation for the primary field.
	address.DefaultPrimary = addressDescPrimary.Default.(bool)
	// addressDescID is the schema descriptor for id field.
	addressDescID := addressFields[0].Descriptor()
	// address.DefaultID holds the default value on creation for the id field.
	address.DefaultID = addressDescID.Default.(func() uuid.UUID)
	auditlogFields := schema.AuditLog{}.Fields()
	_ = auditlogFields
	// auditlogDescUser is the schema descriptor for user field.
	auditlogDescUser := auditlogFields[1].Descriptor()
	// auditlog.UserValidator is a validator for the "user" field. It is called by the builders before save.
	auditlog.UserValidator = auditlogDescUser.Validators[0].(func(string) error)
	// auditlogDescAction is the schema descriptor for action field.
	auditlogDescAction := auditlogFields[2].Descriptor()
	// auditlog.ActionValidator is a validator for the "action" field. It is called by the builders before save.
	auditlog.ActionValidator = auditlogDescAction.Validators[0].(func(string) error)
	// auditlogDescEntitySchema is the schema descriptor for entity_schema field.
	auditlogDescEntitySchema := auditlogFields[3].Descriptor()
	// auditlog.EntitySchemaValidator is a validator for the "entity_schema" field. It is called by the builders before save.
	auditlog.EntitySchemaValidator = auditlogDescEntitySchema.Validators[0].(func(string) error)
	// auditlogDescTimestamp is the schema descriptor for timestamp field.
	auditlogDescTimestamp := auditlogFields[6].Descriptor()
	// auditlog.DefaultTimestamp holds the default value on creation for the timestamp field.
	auditlog.DefaultTimestamp = auditlogDescTimestamp.Default.(func() time.Time)
	// auditlogDescID is the schema descriptor for id field.
	auditlogDescID := auditlogFields[0].Descriptor()
	// auditlog.DefaultID holds the default value on creation for the id field.
	auditlog.DefaultID = auditlogDescID.Default.(func() uuid.UUID)
	businessMixin := schema.Business{}.Mixin()
	businessHooks := schema.Business{}.Hooks()
	business.Hooks[0] = businessHooks[0]
	business.Hooks[1] = businessHooks[1]
	businessMixinFields0 := businessMixin[0].Fields()
	_ = businessMixinFields0
	businessFields := schema.Business{}.Fields()
	_ = businessFields
	// businessDescCreatedAt is the schema descriptor for created_at field.
	businessDescCreatedAt := businessMixinFields0[0].Descriptor()
	// business.DefaultCreatedAt holds the default value on creation for the created_at field.
	business.DefaultCreatedAt = businessDescCreatedAt.Default.(func() time.Time)
	// businessDescUpdatedAt is the schema descriptor for updated_at field.
	businessDescUpdatedAt := businessMixinFields0[1].Descriptor()
	// business.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	business.DefaultUpdatedAt = businessDescUpdatedAt.Default.(func() time.Time)
	// business.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	business.UpdateDefaultUpdatedAt = businessDescUpdatedAt.UpdateDefault.(func() time.Time)
	// businessDescName1 is the schema descriptor for name1 field.
	businessDescName1 := businessFields[1].Descriptor()
	// business.Name1Validator is a validator for the "name1" field. It is called by the builders before save.
	business.Name1Validator = businessDescName1.Validators[0].(func(string) error)
	// businessDescAlias is the schema descriptor for alias field.
	businessDescAlias := businessFields[3].Descriptor()
	// business.AliasValidator is a validator for the "alias" field. It is called by the builders before save.
	business.AliasValidator = businessDescAlias.Validators[0].(func(string) error)
	// businessDescActive is the schema descriptor for active field.
	businessDescActive := businessFields[8].Descriptor()
	// business.DefaultActive holds the default value on creation for the active field.
	business.DefaultActive = businessDescActive.Default.(bool)
	// businessDescID is the schema descriptor for id field.
	businessDescID := businessFields[0].Descriptor()
	// business.DefaultID holds the default value on creation for the id field.
	business.DefaultID = businessDescID.Default.(func() uuid.UUID)
	tagMixin := schema.Tag{}.Mixin()
	tagMixinFields0 := tagMixin[0].Fields()
	_ = tagMixinFields0
	tagFields := schema.Tag{}.Fields()
	_ = tagFields
	// tagDescCreatedAt is the schema descriptor for created_at field.
	tagDescCreatedAt := tagMixinFields0[0].Descriptor()
	// tag.DefaultCreatedAt holds the default value on creation for the created_at field.
	tag.DefaultCreatedAt = tagDescCreatedAt.Default.(func() time.Time)
	// tagDescUpdatedAt is the schema descriptor for updated_at field.
	tagDescUpdatedAt := tagMixinFields0[1].Descriptor()
	// tag.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	tag.DefaultUpdatedAt = tagDescUpdatedAt.Default.(func() time.Time)
	// tag.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	tag.UpdateDefaultUpdatedAt = tagDescUpdatedAt.UpdateDefault.(func() time.Time)
	// tagDescName is the schema descriptor for name field.
	tagDescName := tagFields[1].Descriptor()
	// tag.NameValidator is a validator for the "name" field. It is called by the builders before save.
	tag.NameValidator = tagDescName.Validators[0].(func(string) error)
	// tagDescID is the schema descriptor for id field.
	tagDescID := tagFields[0].Descriptor()
	// tag.DefaultID holds the default value on creation for the id field.
	tag.DefaultID = tagDescID.Default.(func() uuid.UUID)
	timetableMixin := schema.Timetable{}.Mixin()
	timetableHooks := schema.Timetable{}.Hooks()
	timetable.Hooks[0] = timetableHooks[0]
	timetableMixinFields0 := timetableMixin[0].Fields()
	_ = timetableMixinFields0
	timetableFields := schema.Timetable{}.Fields()
	_ = timetableFields
	// timetableDescCreatedAt is the schema descriptor for created_at field.
	timetableDescCreatedAt := timetableMixinFields0[0].Descriptor()
	// timetable.DefaultCreatedAt holds the default value on creation for the created_at field.
	timetable.DefaultCreatedAt = timetableDescCreatedAt.Default.(func() time.Time)
	// timetableDescUpdatedAt is the schema descriptor for updated_at field.
	timetableDescUpdatedAt := timetableMixinFields0[1].Descriptor()
	// timetable.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	timetable.DefaultUpdatedAt = timetableDescUpdatedAt.Default.(func() time.Time)
	// timetable.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	timetable.UpdateDefaultUpdatedAt = timetableDescUpdatedAt.UpdateDefault.(func() time.Time)
	// timetableDescTimeWholeDay is the schema descriptor for time_whole_day field.
	timetableDescTimeWholeDay := timetableFields[4].Descriptor()
	// timetable.DefaultTimeWholeDay holds the default value on creation for the time_whole_day field.
	timetable.DefaultTimeWholeDay = timetableDescTimeWholeDay.Default.(bool)
	// timetableDescID is the schema descriptor for id field.
	timetableDescID := timetableFields[0].Descriptor()
	// timetable.DefaultID holds the default value on creation for the id field.
	timetable.DefaultID = timetableDescID.Default.(func() uuid.UUID)
	userMixin := schema.User{}.Mixin()
	userHooks := schema.User{}.Hooks()
	user.Hooks[0] = userHooks[0]
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userMixinFields0[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userMixinFields0[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescSurname is the schema descriptor for surname field.
	userDescSurname := userFields[1].Descriptor()
	// user.SurnameValidator is a validator for the "surname" field. It is called by the builders before save.
	user.SurnameValidator = userDescSurname.Validators[0].(func(string) error)
	// userDescActive is the schema descriptor for active field.
	userDescActive := userFields[7].Descriptor()
	// user.DefaultActive holds the default value on creation for the active field.
	user.DefaultActive = userDescActive.Default.(bool)
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[8].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(int)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}

const (
	Version = "v0.12.3"                                         // Version of ent codegen.
	Sum     = "h1:N5lO2EOrHpCH5HYfiMOCHYbo+oh5M8GjT0/cx5x6xkk=" // Sum of ent codegen.
)
