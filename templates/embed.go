package templates

import (
	"embed"
	"html/template"
	"hynie.de/ohmab/internal/pkg/utils"
)

var (
	//go:embed *
	resources embed.FS

	funcMap = template.FuncMap{
		"convertTel":     utils.ConvertTel,
		"convertWebsite": utils.ConvertWebsite,
	}

	FullTimeTableHTMLTemplate = template.Must(
		template.New("timetables.tmpl").Funcs(funcMap).ParseFS(resources, "*.tmpl"))

	TableTimeTableHTMLTemplate = template.Must(
		template.New("table").Funcs(funcMap).ParseFS(resources, "timetablestable.tmpl"))
)
