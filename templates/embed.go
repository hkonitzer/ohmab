package templates

import (
	"embed"
	"github.com/hkonitzer/ohmab/internal/pkg/utils"
	"html/template"
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
