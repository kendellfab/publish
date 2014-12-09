package usecases

import (
	"github.com/russross/blackfriday"
	"html/template"
	"time"
)

// Mon Jan 2 15:04:05 MST 2006

func FormatDate(input time.Time, fmts ...string) string {
	format := "Mon Jan 2, 2006"
	if len(fmts) > 0 {
		format = fmts[0]
	}
	return input.Format(format)
}

func FormatBool(input bool) string {
	if input {
		return "Yes"
	}
	return "No"
}

func RenderMarkdown(input string) template.HTML {
	return template.HTML(blackfriday.MarkdownCommon([]byte(input)))
}
