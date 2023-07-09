package utils

import "strings"

func ConvertTel(tel string) string {
	if strings.HasPrefix(tel, "+") {
		tel = tel[3:]
	}
	tel = "0" + tel
	return strings.ReplaceAll(tel, " ", "")
}

func ConvertWebsite(website string) string {
	if strings.HasPrefix(website, "http://") {
		return website[7:]
	}
	if strings.HasPrefix(website, "https://") {
		return website[8:]
	}
	return website
}
