// OHMAB
// Code generated by entc, DO NOT EDIT.

package address

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldDeletedAt, v))
}

// Addition applies equality check predicate on the "addition" field. It's identical to AdditionEQ.
func Addition(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddition, v))
}

// Street applies equality check predicate on the "street" field. It's identical to StreetEQ.
func Street(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreet, v))
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// Zip applies equality check predicate on the "zip" field. It's identical to ZipEQ.
func Zip(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldZip, v))
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldState, v))
}

// Country applies equality check predicate on the "country" field. It's identical to CountryEQ.
func Country(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// Locale applies equality check predicate on the "locale" field. It's identical to LocaleEQ.
func Locale(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLocale, v))
}

// Primary applies equality check predicate on the "primary" field. It's identical to PrimaryEQ.
func Primary(v bool) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPrimary, v))
}

// Telephone applies equality check predicate on the "telephone" field. It's identical to TelephoneEQ.
func Telephone(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldTelephone, v))
}

// Comment applies equality check predicate on the "comment" field. It's identical to CommentEQ.
func Comment(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldComment, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldDeletedAt))
}

// AdditionEQ applies the EQ predicate on the "addition" field.
func AdditionEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldAddition, v))
}

// AdditionNEQ applies the NEQ predicate on the "addition" field.
func AdditionNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldAddition, v))
}

// AdditionIn applies the In predicate on the "addition" field.
func AdditionIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldAddition, vs...))
}

// AdditionNotIn applies the NotIn predicate on the "addition" field.
func AdditionNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldAddition, vs...))
}

// AdditionGT applies the GT predicate on the "addition" field.
func AdditionGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldAddition, v))
}

// AdditionGTE applies the GTE predicate on the "addition" field.
func AdditionGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldAddition, v))
}

// AdditionLT applies the LT predicate on the "addition" field.
func AdditionLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldAddition, v))
}

// AdditionLTE applies the LTE predicate on the "addition" field.
func AdditionLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldAddition, v))
}

// AdditionContains applies the Contains predicate on the "addition" field.
func AdditionContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldAddition, v))
}

// AdditionHasPrefix applies the HasPrefix predicate on the "addition" field.
func AdditionHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldAddition, v))
}

// AdditionHasSuffix applies the HasSuffix predicate on the "addition" field.
func AdditionHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldAddition, v))
}

// AdditionIsNil applies the IsNil predicate on the "addition" field.
func AdditionIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldAddition))
}

// AdditionNotNil applies the NotNil predicate on the "addition" field.
func AdditionNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldAddition))
}

// AdditionEqualFold applies the EqualFold predicate on the "addition" field.
func AdditionEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldAddition, v))
}

// AdditionContainsFold applies the ContainsFold predicate on the "addition" field.
func AdditionContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldAddition, v))
}

// StreetEQ applies the EQ predicate on the "street" field.
func StreetEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldStreet, v))
}

// StreetNEQ applies the NEQ predicate on the "street" field.
func StreetNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldStreet, v))
}

// StreetIn applies the In predicate on the "street" field.
func StreetIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldStreet, vs...))
}

// StreetNotIn applies the NotIn predicate on the "street" field.
func StreetNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldStreet, vs...))
}

// StreetGT applies the GT predicate on the "street" field.
func StreetGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldStreet, v))
}

// StreetGTE applies the GTE predicate on the "street" field.
func StreetGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldStreet, v))
}

// StreetLT applies the LT predicate on the "street" field.
func StreetLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldStreet, v))
}

// StreetLTE applies the LTE predicate on the "street" field.
func StreetLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldStreet, v))
}

// StreetContains applies the Contains predicate on the "street" field.
func StreetContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldStreet, v))
}

// StreetHasPrefix applies the HasPrefix predicate on the "street" field.
func StreetHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldStreet, v))
}

// StreetHasSuffix applies the HasSuffix predicate on the "street" field.
func StreetHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldStreet, v))
}

// StreetIsNil applies the IsNil predicate on the "street" field.
func StreetIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldStreet))
}

// StreetNotNil applies the NotNil predicate on the "street" field.
func StreetNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldStreet))
}

// StreetEqualFold applies the EqualFold predicate on the "street" field.
func StreetEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldStreet, v))
}

// StreetContainsFold applies the ContainsFold predicate on the "street" field.
func StreetContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldStreet, v))
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCity, v))
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCity, v))
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCity, vs...))
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCity, vs...))
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCity, v))
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCity, v))
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCity, v))
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCity, v))
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCity, v))
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCity, v))
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCity, v))
}

// CityIsNil applies the IsNil predicate on the "city" field.
func CityIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldCity))
}

// CityNotNil applies the NotNil predicate on the "city" field.
func CityNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldCity))
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCity, v))
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCity, v))
}

// ZipEQ applies the EQ predicate on the "zip" field.
func ZipEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldZip, v))
}

// ZipNEQ applies the NEQ predicate on the "zip" field.
func ZipNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldZip, v))
}

// ZipIn applies the In predicate on the "zip" field.
func ZipIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldZip, vs...))
}

// ZipNotIn applies the NotIn predicate on the "zip" field.
func ZipNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldZip, vs...))
}

// ZipGT applies the GT predicate on the "zip" field.
func ZipGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldZip, v))
}

// ZipGTE applies the GTE predicate on the "zip" field.
func ZipGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldZip, v))
}

// ZipLT applies the LT predicate on the "zip" field.
func ZipLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldZip, v))
}

// ZipLTE applies the LTE predicate on the "zip" field.
func ZipLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldZip, v))
}

// ZipContains applies the Contains predicate on the "zip" field.
func ZipContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldZip, v))
}

// ZipHasPrefix applies the HasPrefix predicate on the "zip" field.
func ZipHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldZip, v))
}

// ZipHasSuffix applies the HasSuffix predicate on the "zip" field.
func ZipHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldZip, v))
}

// ZipIsNil applies the IsNil predicate on the "zip" field.
func ZipIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldZip))
}

// ZipNotNil applies the NotNil predicate on the "zip" field.
func ZipNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldZip))
}

// ZipEqualFold applies the EqualFold predicate on the "zip" field.
func ZipEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldZip, v))
}

// ZipContainsFold applies the ContainsFold predicate on the "zip" field.
func ZipContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldZip, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldState, vs...))
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldState, v))
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldState, v))
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldState, v))
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldState, v))
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldState, v))
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldState, v))
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldState, v))
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldState))
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldState))
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldState, v))
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldState, v))
}

// CountryEQ applies the EQ predicate on the "country" field.
func CountryEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldCountry, v))
}

// CountryNEQ applies the NEQ predicate on the "country" field.
func CountryNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldCountry, v))
}

// CountryIn applies the In predicate on the "country" field.
func CountryIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldCountry, vs...))
}

// CountryNotIn applies the NotIn predicate on the "country" field.
func CountryNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldCountry, vs...))
}

// CountryGT applies the GT predicate on the "country" field.
func CountryGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldCountry, v))
}

// CountryGTE applies the GTE predicate on the "country" field.
func CountryGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldCountry, v))
}

// CountryLT applies the LT predicate on the "country" field.
func CountryLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldCountry, v))
}

// CountryLTE applies the LTE predicate on the "country" field.
func CountryLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldCountry, v))
}

// CountryContains applies the Contains predicate on the "country" field.
func CountryContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldCountry, v))
}

// CountryHasPrefix applies the HasPrefix predicate on the "country" field.
func CountryHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldCountry, v))
}

// CountryHasSuffix applies the HasSuffix predicate on the "country" field.
func CountryHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldCountry, v))
}

// CountryIsNil applies the IsNil predicate on the "country" field.
func CountryIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldCountry))
}

// CountryNotNil applies the NotNil predicate on the "country" field.
func CountryNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldCountry))
}

// CountryEqualFold applies the EqualFold predicate on the "country" field.
func CountryEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldCountry, v))
}

// CountryContainsFold applies the ContainsFold predicate on the "country" field.
func CountryContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldCountry, v))
}

// LocaleEQ applies the EQ predicate on the "locale" field.
func LocaleEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldLocale, v))
}

// LocaleNEQ applies the NEQ predicate on the "locale" field.
func LocaleNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldLocale, v))
}

// LocaleIn applies the In predicate on the "locale" field.
func LocaleIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldLocale, vs...))
}

// LocaleNotIn applies the NotIn predicate on the "locale" field.
func LocaleNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldLocale, vs...))
}

// LocaleGT applies the GT predicate on the "locale" field.
func LocaleGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldLocale, v))
}

// LocaleGTE applies the GTE predicate on the "locale" field.
func LocaleGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldLocale, v))
}

// LocaleLT applies the LT predicate on the "locale" field.
func LocaleLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldLocale, v))
}

// LocaleLTE applies the LTE predicate on the "locale" field.
func LocaleLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldLocale, v))
}

// LocaleContains applies the Contains predicate on the "locale" field.
func LocaleContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldLocale, v))
}

// LocaleHasPrefix applies the HasPrefix predicate on the "locale" field.
func LocaleHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldLocale, v))
}

// LocaleHasSuffix applies the HasSuffix predicate on the "locale" field.
func LocaleHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldLocale, v))
}

// LocaleEqualFold applies the EqualFold predicate on the "locale" field.
func LocaleEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldLocale, v))
}

// LocaleContainsFold applies the ContainsFold predicate on the "locale" field.
func LocaleContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldLocale, v))
}

// PrimaryEQ applies the EQ predicate on the "primary" field.
func PrimaryEQ(v bool) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldPrimary, v))
}

// PrimaryNEQ applies the NEQ predicate on the "primary" field.
func PrimaryNEQ(v bool) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldPrimary, v))
}

// TelephoneEQ applies the EQ predicate on the "telephone" field.
func TelephoneEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldTelephone, v))
}

// TelephoneNEQ applies the NEQ predicate on the "telephone" field.
func TelephoneNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldTelephone, v))
}

// TelephoneIn applies the In predicate on the "telephone" field.
func TelephoneIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldTelephone, vs...))
}

// TelephoneNotIn applies the NotIn predicate on the "telephone" field.
func TelephoneNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldTelephone, vs...))
}

// TelephoneGT applies the GT predicate on the "telephone" field.
func TelephoneGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldTelephone, v))
}

// TelephoneGTE applies the GTE predicate on the "telephone" field.
func TelephoneGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldTelephone, v))
}

// TelephoneLT applies the LT predicate on the "telephone" field.
func TelephoneLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldTelephone, v))
}

// TelephoneLTE applies the LTE predicate on the "telephone" field.
func TelephoneLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldTelephone, v))
}

// TelephoneContains applies the Contains predicate on the "telephone" field.
func TelephoneContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldTelephone, v))
}

// TelephoneHasPrefix applies the HasPrefix predicate on the "telephone" field.
func TelephoneHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldTelephone, v))
}

// TelephoneHasSuffix applies the HasSuffix predicate on the "telephone" field.
func TelephoneHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldTelephone, v))
}

// TelephoneIsNil applies the IsNil predicate on the "telephone" field.
func TelephoneIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldTelephone))
}

// TelephoneNotNil applies the NotNil predicate on the "telephone" field.
func TelephoneNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldTelephone))
}

// TelephoneEqualFold applies the EqualFold predicate on the "telephone" field.
func TelephoneEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldTelephone, v))
}

// TelephoneContainsFold applies the ContainsFold predicate on the "telephone" field.
func TelephoneContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldTelephone, v))
}

// CommentEQ applies the EQ predicate on the "comment" field.
func CommentEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldEQ(FieldComment, v))
}

// CommentNEQ applies the NEQ predicate on the "comment" field.
func CommentNEQ(v string) predicate.Address {
	return predicate.Address(sql.FieldNEQ(FieldComment, v))
}

// CommentIn applies the In predicate on the "comment" field.
func CommentIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldIn(FieldComment, vs...))
}

// CommentNotIn applies the NotIn predicate on the "comment" field.
func CommentNotIn(vs ...string) predicate.Address {
	return predicate.Address(sql.FieldNotIn(FieldComment, vs...))
}

// CommentGT applies the GT predicate on the "comment" field.
func CommentGT(v string) predicate.Address {
	return predicate.Address(sql.FieldGT(FieldComment, v))
}

// CommentGTE applies the GTE predicate on the "comment" field.
func CommentGTE(v string) predicate.Address {
	return predicate.Address(sql.FieldGTE(FieldComment, v))
}

// CommentLT applies the LT predicate on the "comment" field.
func CommentLT(v string) predicate.Address {
	return predicate.Address(sql.FieldLT(FieldComment, v))
}

// CommentLTE applies the LTE predicate on the "comment" field.
func CommentLTE(v string) predicate.Address {
	return predicate.Address(sql.FieldLTE(FieldComment, v))
}

// CommentContains applies the Contains predicate on the "comment" field.
func CommentContains(v string) predicate.Address {
	return predicate.Address(sql.FieldContains(FieldComment, v))
}

// CommentHasPrefix applies the HasPrefix predicate on the "comment" field.
func CommentHasPrefix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasPrefix(FieldComment, v))
}

// CommentHasSuffix applies the HasSuffix predicate on the "comment" field.
func CommentHasSuffix(v string) predicate.Address {
	return predicate.Address(sql.FieldHasSuffix(FieldComment, v))
}

// CommentIsNil applies the IsNil predicate on the "comment" field.
func CommentIsNil() predicate.Address {
	return predicate.Address(sql.FieldIsNull(FieldComment))
}

// CommentNotNil applies the NotNil predicate on the "comment" field.
func CommentNotNil() predicate.Address {
	return predicate.Address(sql.FieldNotNull(FieldComment))
}

// CommentEqualFold applies the EqualFold predicate on the "comment" field.
func CommentEqualFold(v string) predicate.Address {
	return predicate.Address(sql.FieldEqualFold(FieldComment, v))
}

// CommentContainsFold applies the ContainsFold predicate on the "comment" field.
func CommentContainsFold(v string) predicate.Address {
	return predicate.Address(sql.FieldContainsFold(FieldComment, v))
}

// HasBusiness applies the HasEdge predicate on the "business" edge.
func HasBusiness() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BusinessTable, BusinessColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBusinessWith applies the HasEdge predicate on the "business" edge with a given conditions (other predicates).
func HasBusinessWith(preds ...predicate.Business) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newBusinessStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTimetables applies the HasEdge predicate on the "timetables" edge.
func HasTimetables() predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TimetablesTable, TimetablesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTimetablesWith applies the HasEdge predicate on the "timetables" edge with a given conditions (other predicates).
func HasTimetablesWith(preds ...predicate.Timetable) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		step := newTimetablesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
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
func Not(p predicate.Address) predicate.Address {
	return predicate.Address(func(s *sql.Selector) {
		p(s.Not())
	})
}
