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
		"convertWebsite": func(website string) string {
			if strings.HasPrefix(website, "http://") {
				return website[7:]
			}
			if strings.HasPrefix(website, "https://") {
				return website[8:]
			}
			return website
		},
	}

	FullTimeTableHTMLTemplate = template.Must(
		template.New("timetables.tmpl").Funcs(funcMap).ParseFS(resources, "*.tmpl"))

	TableTimeTableHTMLTemplate = template.Must(
		template.New("table").Funcs(funcMap).ParseFS(resources, "timetablestable.tmpl"))
)
