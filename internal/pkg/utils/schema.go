package utils

import (
	"github.com/hkonitzer/ohmab/ent/address"
	"github.com/hkonitzer/ohmab/ent/business"
	"github.com/hkonitzer/ohmab/ent/schema/constants"
	"strings"
)

func GetBusinessFields(headersFromSchemas []string, withSchemaNameAsPrefix bool, toUpper bool) []string {
	for _, field := range business.Columns {
		if strings.HasSuffix(field, constants.IDFieldName) {
			continue
		}
		if field == constants.TimeMixinCreatedAtFieldName ||
			field == constants.TimeMixinUpdatedAtFieldName ||
			field == constants.TimeMixinDeletedAtFieldName ||
			field == constants.ActiveFieldName {
			continue
		}
		if toUpper {
			field = strings.ToUpper(field)
		}
		if withSchemaNameAsPrefix {
			headersFromSchemas = append(headersFromSchemas, "BUSINESS$"+field)
		} else {
			headersFromSchemas = append(headersFromSchemas, field)
		}

	}
	return headersFromSchemas
}

func GetAddressFields(headersFromSchemas []string, withSchemaNameAsPrefix bool, toUpper bool) []string {
	for _, field := range address.Columns {
		if strings.HasSuffix(field, constants.IDFieldName) {
			continue
		}
		if field == constants.TimeMixinCreatedAtFieldName ||
			field == constants.TimeMixinUpdatedAtFieldName ||
			field == constants.TimeMixinDeletedAtFieldName ||
			field == constants.ActiveFieldName {
			continue
		}
		if toUpper {
			field = strings.ToUpper(field)
		}
		if withSchemaNameAsPrefix {
			headersFromSchemas = append(headersFromSchemas, "ADDRESS$"+field)
		} else {
			headersFromSchemas = append(headersFromSchemas, field)
		}

	}
	return headersFromSchemas
}
