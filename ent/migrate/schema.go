// OHMAB
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "addition", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "street", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "city", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "zip", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "state", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "country", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "locale", Type: field.TypeString, Size: 5, Default: "en_US"},
		{Name: "primary", Type: field.TypeBool, Default: false},
		{Name: "telephone", Type: field.TypeString, Unique: true, Nullable: true, Size: 2147483647},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "business_addresses", Type: field.TypeUUID, Nullable: true},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "addresses_businesses_addresses",
				Columns:    []*schema.Column{AddressesColumns[14]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// AuditLogsColumns holds the columns for the "audit_logs" table.
	AuditLogsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "user", Type: field.TypeString, Size: 2147483647},
		{Name: "action", Type: field.TypeString, Size: 2147483647},
		{Name: "entity_schema", Type: field.TypeString, Size: 2147483647},
		{Name: "entity_values", Type: field.TypeJSON, Nullable: true},
		{Name: "entity_uuid", Type: field.TypeString, Nullable: true},
		{Name: "timestamp", Type: field.TypeTime},
	}
	// AuditLogsTable holds the schema information for the "audit_logs" table.
	AuditLogsTable = &schema.Table{
		Name:       "audit_logs",
		Columns:    AuditLogsColumns,
		PrimaryKey: []*schema.Column{AuditLogsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "auditlog_entity_uuid",
				Unique:  false,
				Columns: []*schema.Column{AuditLogsColumns[5]},
			},
		},
	}
	// BusinessesColumns holds the columns for the "businesses" table.
	BusinessesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name1", Type: field.TypeString, Size: 2147483647},
		{Name: "name2", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "alias", Type: field.TypeString, Unique: true, Size: 20},
		{Name: "telephone", Type: field.TypeString, Unique: true, Nullable: true, Size: 2147483647},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true, Size: 2147483647},
		{Name: "website", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "active", Type: field.TypeBool, Default: true},
	}
	// BusinessesTable holds the schema information for the "businesses" table.
	BusinessesTable = &schema.Table{
		Name:       "businesses",
		Columns:    BusinessesColumns,
		PrimaryKey: []*schema.Column{BusinessesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "business_name1_name2",
				Unique:  false,
				Columns: []*schema.Column{BusinessesColumns[4], BusinessesColumns[5]},
			},
			{
				Name:    "business_alias",
				Unique:  false,
				Columns: []*schema.Column{BusinessesColumns[6]},
			},
			{
				Name:    "business_email",
				Unique:  false,
				Columns: []*schema.Column{BusinessesColumns[8]},
			},
		},
	}
	// ContentsColumns holds the columns for the "contents" table.
	ContentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "timetable_type", Type: field.TypeEnum, Enums: []string{"DEFAULT", "REGULAR", "CLOSED", "EMERGENCYSERVICE", "HOLIDAY", "SPECIAL"}, Default: "DEFAULT"},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"TEXT", "HTML", "CSS", "OTHER"}, Default: "TEXT"},
		{Name: "locale", Type: field.TypeString, Size: 5, Default: "en_US"},
		{Name: "location", Type: field.TypeEnum, Enums: []string{"TOP", "BOTTOM"}, Default: "TOP"},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"DRAFT", "PUBLISHED"}, Default: "DRAFT"},
		{Name: "published_at", Type: field.TypeTime, Nullable: true},
	}
	// ContentsTable holds the schema information for the "contents" table.
	ContentsTable = &schema.Table{
		Name:       "contents",
		Columns:    ContentsColumns,
		PrimaryKey: []*schema.Column{ContentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "content_status_type_location_locale",
				Unique:  true,
				Columns: []*schema.Column{ContentsColumns[9], ContentsColumns[5], ContentsColumns[7], ContentsColumns[6]},
			},
		},
	}
	// TagsColumns holds the columns for the "tags" table.
	TagsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "name", Type: field.TypeString, Size: 2147483647},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 2147483647},
	}
	// TagsTable holds the schema information for the "tags" table.
	TagsTable = &schema.Table{
		Name:       "tags",
		Columns:    TagsColumns,
		PrimaryKey: []*schema.Column{TagsColumns[0]},
	}
	// TimetablesColumns holds the columns for the "timetables" table.
	TimetablesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "timetable_type", Type: field.TypeEnum, Enums: []string{"DEFAULT", "REGULAR", "CLOSED", "EMERGENCYSERVICE", "HOLIDAY", "SPECIAL"}, Default: "DEFAULT"},
		{Name: "datetime_from", Type: field.TypeTime},
		{Name: "duration", Type: field.TypeInt, Nullable: true, SchemaType: map[string]string{"mysql": "TINYINT", "postgres": "SMALLINT", "sqlite3": "INTEGER"}},
		{Name: "datetime_to", Type: field.TypeTime, Nullable: true},
		{Name: "time_whole_day", Type: field.TypeBool, Default: false},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "availability_by_phone", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "availability_by_email", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "availability_by_sms", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "availability_by_whatsapp", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "address_timetables", Type: field.TypeUUID},
	}
	// TimetablesTable holds the schema information for the "timetables" table.
	TimetablesTable = &schema.Table{
		Name:       "timetables",
		Columns:    TimetablesColumns,
		PrimaryKey: []*schema.Column{TimetablesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timetables_addresses_timetables",
				Columns:    []*schema.Column{TimetablesColumns[14]},
				RefColumns: []*schema.Column{AddressesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "timetable_datetime_from",
				Unique:  false,
				Columns: []*schema.Column{TimetablesColumns[5]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "use_publicapi", Type: field.TypeString, Size: 2147483647, Default: ""},
		{Name: "login", Type: field.TypeString, Size: 2147483647},
		{Name: "surname", Type: field.TypeString, Size: 2147483647},
		{Name: "firstname", Type: field.TypeString, Size: 2147483647},
		{Name: "title", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 2147483647},
		{Name: "passwordhash", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "comment", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "active", Type: field.TypeBool, Default: true},
		{Name: "role", Type: field.TypeString, Default: "8"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_use_publicapi",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[4]},
			},
			{
				Name:    "user_email",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[9]},
			},
		},
	}
	// BusinessTagsColumns holds the columns for the "business_tags" table.
	BusinessTagsColumns = []*schema.Column{
		{Name: "business_id", Type: field.TypeUUID},
		{Name: "tag_id", Type: field.TypeUUID},
	}
	// BusinessTagsTable holds the schema information for the "business_tags" table.
	BusinessTagsTable = &schema.Table{
		Name:       "business_tags",
		Columns:    BusinessTagsColumns,
		PrimaryKey: []*schema.Column{BusinessTagsColumns[0], BusinessTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "business_tags_business_id",
				Columns:    []*schema.Column{BusinessTagsColumns[0]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "business_tags_tag_id",
				Columns:    []*schema.Column{BusinessTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// BusinessUsersColumns holds the columns for the "business_users" table.
	BusinessUsersColumns = []*schema.Column{
		{Name: "business_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// BusinessUsersTable holds the schema information for the "business_users" table.
	BusinessUsersTable = &schema.Table{
		Name:       "business_users",
		Columns:    BusinessUsersColumns,
		PrimaryKey: []*schema.Column{BusinessUsersColumns[0], BusinessUsersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "business_users_business_id",
				Columns:    []*schema.Column{BusinessUsersColumns[0]},
				RefColumns: []*schema.Column{BusinessesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "business_users_user_id",
				Columns:    []*schema.Column{BusinessUsersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TimetableUsersOnDutyColumns holds the columns for the "timetable_users_on_duty" table.
	TimetableUsersOnDutyColumns = []*schema.Column{
		{Name: "timetable_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
	}
	// TimetableUsersOnDutyTable holds the schema information for the "timetable_users_on_duty" table.
	TimetableUsersOnDutyTable = &schema.Table{
		Name:       "timetable_users_on_duty",
		Columns:    TimetableUsersOnDutyColumns,
		PrimaryKey: []*schema.Column{TimetableUsersOnDutyColumns[0], TimetableUsersOnDutyColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "timetable_users_on_duty_timetable_id",
				Columns:    []*schema.Column{TimetableUsersOnDutyColumns[0]},
				RefColumns: []*schema.Column{TimetablesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "timetable_users_on_duty_user_id",
				Columns:    []*schema.Column{TimetableUsersOnDutyColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserTagsColumns holds the columns for the "user_tags" table.
	UserTagsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "tag_id", Type: field.TypeUUID},
	}
	// UserTagsTable holds the schema information for the "user_tags" table.
	UserTagsTable = &schema.Table{
		Name:       "user_tags",
		Columns:    UserTagsColumns,
		PrimaryKey: []*schema.Column{UserTagsColumns[0], UserTagsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_tags_user_id",
				Columns:    []*schema.Column{UserTagsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_tags_tag_id",
				Columns:    []*schema.Column{UserTagsColumns[1]},
				RefColumns: []*schema.Column{TagsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressesTable,
		AuditLogsTable,
		BusinessesTable,
		ContentsTable,
		TagsTable,
		TimetablesTable,
		UsersTable,
		BusinessTagsTable,
		BusinessUsersTable,
		TimetableUsersOnDutyTable,
		UserTagsTable,
	}
)

func init() {
	AddressesTable.ForeignKeys[0].RefTable = BusinessesTable
	TimetablesTable.ForeignKeys[0].RefTable = AddressesTable
	BusinessTagsTable.ForeignKeys[0].RefTable = BusinessesTable
	BusinessTagsTable.ForeignKeys[1].RefTable = TagsTable
	BusinessUsersTable.ForeignKeys[0].RefTable = BusinessesTable
	BusinessUsersTable.ForeignKeys[1].RefTable = UsersTable
	TimetableUsersOnDutyTable.ForeignKeys[0].RefTable = TimetablesTable
	TimetableUsersOnDutyTable.ForeignKeys[1].RefTable = UsersTable
	UserTagsTable.ForeignKeys[0].RefTable = UsersTable
	UserTagsTable.ForeignKeys[1].RefTable = TagsTable
}
