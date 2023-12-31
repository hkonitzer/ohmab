// OHMAB
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/goodsign/monday"
)

func (t *Timetable) I18nDate(field string) string {
	l := t.Edges.Address.Locale
	if l == "" {
		l = "en_US"
	}
	switch field {

	case "created_at":
		return monday.Format(t.CreatedAt, "Monday 02.01.2006 15:04", monday.Locale(l))

	case "updated_at":
		return monday.Format(t.UpdatedAt, "Monday 02.01.2006 15:04", monday.Locale(l))

	case "deleted_at":
		return monday.Format(t.DeletedAt, "Monday 02.01.2006 15:04", monday.Locale(l))

	case "datetime_from":
		return monday.Format(t.DatetimeFrom, "Monday 02.01.2006 15:04", monday.Locale(l))

	case "datetime_to":
		return monday.Format(t.DatetimeTo, "Monday 02.01.2006 15:04", monday.Locale(l))

	default:
		return ""
	}
}
