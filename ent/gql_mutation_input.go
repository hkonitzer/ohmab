// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/hkonitzer/ohmab/ent/timetable"
)

// CreateBusinessInput represents a mutation input for creating businesses.
type CreateBusinessInput struct {
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
	Name1      string
	Name2      *string
	Alias      string
	Telephone  *string
	Email      *string
	Website    *string
	Comment    *string
	Active     *bool
	AddressIDs []uuid.UUID
	TagIDs     []uuid.UUID
	UserIDs    []uuid.UUID
}

// Mutate applies the CreateBusinessInput on the BusinessMutation builder.
func (i *CreateBusinessInput) Mutate(m *BusinessMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	m.SetName1(i.Name1)
	if v := i.Name2; v != nil {
		m.SetName2(*v)
	}
	m.SetAlias(i.Alias)
	if v := i.Telephone; v != nil {
		m.SetTelephone(*v)
	}
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Website; v != nil {
		m.SetWebsite(*v)
	}
	if v := i.Comment; v != nil {
		m.SetComment(*v)
	}
	if v := i.Active; v != nil {
		m.SetActive(*v)
	}
	if v := i.AddressIDs; len(v) > 0 {
		m.AddAddressIDs(v...)
	}
	if v := i.TagIDs; len(v) > 0 {
		m.AddTagIDs(v...)
	}
	if v := i.UserIDs; len(v) > 0 {
		m.AddUserIDs(v...)
	}
}

// SetInput applies the change-set in the CreateBusinessInput on the BusinessCreate builder.
func (c *BusinessCreate) SetInput(i CreateBusinessInput) *BusinessCreate {
	i.Mutate(c.Mutation())
	return c
}

// CreateTimetableInput represents a mutation input for creating timetables.
type CreateTimetableInput struct {
	CreatedAt              *time.Time
	UpdatedAt              *time.Time
	DeletedAt              *time.Time
	TimetableType          *timetable.TimetableType
	DatetimeFrom           *time.Time
	Duration               *uint8
	DatetimeTo             time.Time
	TimeWholeDay           *bool
	Comment                *string
	AvailabilityByPhone    *string
	AvailabilityByEmail    *string
	AvailabilityBySms      *string
	AvailabilityByWhatsapp *string
	AddressID              uuid.UUID
	UsersOnDutyIDs         []uuid.UUID
}

// Mutate applies the CreateTimetableInput on the TimetableMutation builder.
func (i *CreateTimetableInput) Mutate(m *TimetableMutation) {
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.UpdatedAt; v != nil {
		m.SetUpdatedAt(*v)
	}
	if v := i.DeletedAt; v != nil {
		m.SetDeletedAt(*v)
	}
	if v := i.TimetableType; v != nil {
		m.SetTimetableType(*v)
	}
	if v := i.DatetimeFrom; v != nil {
		m.SetDatetimeFrom(*v)
	}
	if v := i.Duration; v != nil {
		m.SetDuration(*v)
	}
	m.SetDatetimeTo(i.DatetimeTo)
	if v := i.TimeWholeDay; v != nil {
		m.SetTimeWholeDay(*v)
	}
	if v := i.Comment; v != nil {
		m.SetComment(*v)
	}
	if v := i.AvailabilityByPhone; v != nil {
		m.SetAvailabilityByPhone(*v)
	}
	if v := i.AvailabilityByEmail; v != nil {
		m.SetAvailabilityByEmail(*v)
	}
	if v := i.AvailabilityBySms; v != nil {
		m.SetAvailabilityBySms(*v)
	}
	if v := i.AvailabilityByWhatsapp; v != nil {
		m.SetAvailabilityByWhatsapp(*v)
	}
	m.SetAddressID(i.AddressID)
	if v := i.UsersOnDutyIDs; len(v) > 0 {
		m.AddUsersOnDutyIDs(v...)
	}
}

// SetInput applies the change-set in the CreateTimetableInput on the TimetableCreate builder.
func (c *TimetableCreate) SetInput(i CreateTimetableInput) *TimetableCreate {
	i.Mutate(c.Mutation())
	return c
}
