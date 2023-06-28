package templates

import (
	"embed"
	"html/template"
	"strings"
)

var (
	//go:embed *
	resources embed.FS

	funcMap = template.FuncMap{
		"convertTel": func(tel string) string {
			if strings.HasPrefix(tel, "+") {
				tel = tel[3:]
			}
			tel = "0" + tel
			return strings.ReplaceAll(tel, " ", "")
		},
	}
	Tmpl = template.Must(template.New("timetables.tmpl").Funcs(funcMap).ParseFS(resources, "*.tmpl"))
)
