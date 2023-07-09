package utils

// source: https://siongui.github.io/+
import (
	"strconv"
	"strings"
)

type LangQ struct {
	Lang   string
	Q      float64
	Locale string
}

func ParseAcceptLanguage(acptLang string) []LangQ {
	var lqs []LangQ

	langQStrs := strings.Split(acptLang, ",")
	for _, langQStr := range langQStrs {
		trimedLangQStr := strings.Trim(langQStr, " ")

		langQ := strings.Split(trimedLangQStr, ";")
		if len(langQ) == 1 {
			lq := createLangQ(langQ[0], 1)
			lqs = append(lqs, lq)
		} else {
			qp := strings.Split(langQ[1], "=")
			q, err := strconv.ParseFloat(qp[1], 64)
			if err != nil {
				panic(err)
			}
			lq := createLangQ(langQ[0], q)
			lqs = append(lqs, lq)
		}
	}
	return lqs
}

func createLangQ(lang string, q float64) LangQ {
	locale := lang
	if strings.Contains(lang, "-") {
		locale = strings.Replace(lang, "-", "_", 1)
	}
	l := strings.Split(locale, "_")
	if len(l) == 1 {
		locale = l[0] + "_" + strings.ToUpper(l[0])
	}
	return LangQ{lang, q, locale}
}
