// OHMAB
// Code generated by entc, DO NOT EDIT.

package auditlog

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLTE(FieldID, id))
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldUser, v))
}

// Action applies equality check predicate on the "action" field. It's identical to ActionEQ.
func Action(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldAction, v))
}

// EntitySchema applies equality check predicate on the "entity_schema" field. It's identical to EntitySchemaEQ.
func EntitySchema(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldEntitySchema, v))
}

// EntityUUID applies equality check predicate on the "entity_uuid" field. It's identical to EntityUUIDEQ.
func EntityUUID(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldEntityUUID, v))
}

// Timestamp applies equality check predicate on the "timestamp" field. It's identical to TimestampEQ.
func Timestamp(v time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldTimestamp, v))
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldUser, v))
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNEQ(FieldUser, v))
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIn(FieldUser, vs...))
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotIn(FieldUser, vs...))
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGT(FieldUser, v))
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGTE(FieldUser, v))
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLT(FieldUser, v))
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLTE(FieldUser, v))
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContains(FieldUser, v))
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasPrefix(FieldUser, v))
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasSuffix(FieldUser, v))
}

// UserEqualFold applies the EqualFold predicate on the "user" field.
func UserEqualFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEqualFold(FieldUser, v))
}

// UserContainsFold applies the ContainsFold predicate on the "user" field.
func UserContainsFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContainsFold(FieldUser, v))
}

// ActionEQ applies the EQ predicate on the "action" field.
func ActionEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldAction, v))
}

// ActionNEQ applies the NEQ predicate on the "action" field.
func ActionNEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNEQ(FieldAction, v))
}

// ActionIn applies the In predicate on the "action" field.
func ActionIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIn(FieldAction, vs...))
}

// ActionNotIn applies the NotIn predicate on the "action" field.
func ActionNotIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotIn(FieldAction, vs...))
}

// ActionGT applies the GT predicate on the "action" field.
func ActionGT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGT(FieldAction, v))
}

// ActionGTE applies the GTE predicate on the "action" field.
func ActionGTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGTE(FieldAction, v))
}

// ActionLT applies the LT predicate on the "action" field.
func ActionLT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLT(FieldAction, v))
}

// ActionLTE applies the LTE predicate on the "action" field.
func ActionLTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLTE(FieldAction, v))
}

// ActionContains applies the Contains predicate on the "action" field.
func ActionContains(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContains(FieldAction, v))
}

// ActionHasPrefix applies the HasPrefix predicate on the "action" field.
func ActionHasPrefix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasPrefix(FieldAction, v))
}

// ActionHasSuffix applies the HasSuffix predicate on the "action" field.
func ActionHasSuffix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasSuffix(FieldAction, v))
}

// ActionEqualFold applies the EqualFold predicate on the "action" field.
func ActionEqualFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEqualFold(FieldAction, v))
}

// ActionContainsFold applies the ContainsFold predicate on the "action" field.
func ActionContainsFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContainsFold(FieldAction, v))
}

// EntitySchemaEQ applies the EQ predicate on the "entity_schema" field.
func EntitySchemaEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldEntitySchema, v))
}

// EntitySchemaNEQ applies the NEQ predicate on the "entity_schema" field.
func EntitySchemaNEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNEQ(FieldEntitySchema, v))
}

// EntitySchemaIn applies the In predicate on the "entity_schema" field.
func EntitySchemaIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIn(FieldEntitySchema, vs...))
}

// EntitySchemaNotIn applies the NotIn predicate on the "entity_schema" field.
func EntitySchemaNotIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotIn(FieldEntitySchema, vs...))
}

// EntitySchemaGT applies the GT predicate on the "entity_schema" field.
func EntitySchemaGT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGT(FieldEntitySchema, v))
}

// EntitySchemaGTE applies the GTE predicate on the "entity_schema" field.
func EntitySchemaGTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGTE(FieldEntitySchema, v))
}

// EntitySchemaLT applies the LT predicate on the "entity_schema" field.
func EntitySchemaLT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLT(FieldEntitySchema, v))
}

// EntitySchemaLTE applies the LTE predicate on the "entity_schema" field.
func EntitySchemaLTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLTE(FieldEntitySchema, v))
}

// EntitySchemaContains applies the Contains predicate on the "entity_schema" field.
func EntitySchemaContains(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContains(FieldEntitySchema, v))
}

// EntitySchemaHasPrefix applies the HasPrefix predicate on the "entity_schema" field.
func EntitySchemaHasPrefix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasPrefix(FieldEntitySchema, v))
}

// EntitySchemaHasSuffix applies the HasSuffix predicate on the "entity_schema" field.
func EntitySchemaHasSuffix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasSuffix(FieldEntitySchema, v))
}

// EntitySchemaEqualFold applies the EqualFold predicate on the "entity_schema" field.
func EntitySchemaEqualFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEqualFold(FieldEntitySchema, v))
}

// EntitySchemaContainsFold applies the ContainsFold predicate on the "entity_schema" field.
func EntitySchemaContainsFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContainsFold(FieldEntitySchema, v))
}

// EntityValuesIsNil applies the IsNil predicate on the "entity_values" field.
func EntityValuesIsNil() predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIsNull(FieldEntityValues))
}

// EntityValuesNotNil applies the NotNil predicate on the "entity_values" field.
func EntityValuesNotNil() predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotNull(FieldEntityValues))
}

// EntityUUIDEQ applies the EQ predicate on the "entity_uuid" field.
func EntityUUIDEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldEntityUUID, v))
}

// EntityUUIDNEQ applies the NEQ predicate on the "entity_uuid" field.
func EntityUUIDNEQ(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNEQ(FieldEntityUUID, v))
}

// EntityUUIDIn applies the In predicate on the "entity_uuid" field.
func EntityUUIDIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIn(FieldEntityUUID, vs...))
}

// EntityUUIDNotIn applies the NotIn predicate on the "entity_uuid" field.
func EntityUUIDNotIn(vs ...string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotIn(FieldEntityUUID, vs...))
}

// EntityUUIDGT applies the GT predicate on the "entity_uuid" field.
func EntityUUIDGT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGT(FieldEntityUUID, v))
}

// EntityUUIDGTE applies the GTE predicate on the "entity_uuid" field.
func EntityUUIDGTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGTE(FieldEntityUUID, v))
}

// EntityUUIDLT applies the LT predicate on the "entity_uuid" field.
func EntityUUIDLT(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLT(FieldEntityUUID, v))
}

// EntityUUIDLTE applies the LTE predicate on the "entity_uuid" field.
func EntityUUIDLTE(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLTE(FieldEntityUUID, v))
}

// EntityUUIDContains applies the Contains predicate on the "entity_uuid" field.
func EntityUUIDContains(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContains(FieldEntityUUID, v))
}

// EntityUUIDHasPrefix applies the HasPrefix predicate on the "entity_uuid" field.
func EntityUUIDHasPrefix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasPrefix(FieldEntityUUID, v))
}

// EntityUUIDHasSuffix applies the HasSuffix predicate on the "entity_uuid" field.
func EntityUUIDHasSuffix(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldHasSuffix(FieldEntityUUID, v))
}

// EntityUUIDIsNil applies the IsNil predicate on the "entity_uuid" field.
func EntityUUIDIsNil() predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIsNull(FieldEntityUUID))
}

// EntityUUIDNotNil applies the NotNil predicate on the "entity_uuid" field.
func EntityUUIDNotNil() predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotNull(FieldEntityUUID))
}

// EntityUUIDEqualFold applies the EqualFold predicate on the "entity_uuid" field.
func EntityUUIDEqualFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEqualFold(FieldEntityUUID, v))
}

// EntityUUIDContainsFold applies the ContainsFold predicate on the "entity_uuid" field.
func EntityUUIDContainsFold(v string) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldContainsFold(FieldEntityUUID, v))
}

// TimestampEQ applies the EQ predicate on the "timestamp" field.
func TimestampEQ(v time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldEQ(FieldTimestamp, v))
}

// TimestampNEQ applies the NEQ predicate on the "timestamp" field.
func TimestampNEQ(v time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNEQ(FieldTimestamp, v))
}

// TimestampIn applies the In predicate on the "timestamp" field.
func TimestampIn(vs ...time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldIn(FieldTimestamp, vs...))
}

// TimestampNotIn applies the NotIn predicate on the "timestamp" field.
func TimestampNotIn(vs ...time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldNotIn(FieldTimestamp, vs...))
}

// TimestampGT applies the GT predicate on the "timestamp" field.
func TimestampGT(v time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGT(FieldTimestamp, v))
}

// TimestampGTE applies the GTE predicate on the "timestamp" field.
func TimestampGTE(v time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldGTE(FieldTimestamp, v))
}

// TimestampLT applies the LT predicate on the "timestamp" field.
func TimestampLT(v time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLT(FieldTimestamp, v))
}

// TimestampLTE applies the LTE predicate on the "timestamp" field.
func TimestampLTE(v time.Time) predicate.AuditLog {
	return predicate.AuditLog(sql.FieldLTE(FieldTimestamp, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AuditLog) predicate.AuditLog {
	return predicate.AuditLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AuditLog) predicate.AuditLog {
	return predicate.AuditLog(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AuditLog) predicate.AuditLog {
	return predicate.AuditLog(func(s *sql.Selector) {
		p(s.Not())
	})
}
