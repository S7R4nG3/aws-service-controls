package prompts

import (
	"bytes"
	"io"
	"os"
	"text/template"
)

type GenerateTemplate struct {
	ServiceName  string
	Frameworks   string
	TemplateFile string
}

type ReviewTemplate struct {
	FileContent  string
	TemplateFile string
}

func GenerateControlText(serviceName string, frameworks string) string {
	t := GenerateTemplate{
		ServiceName:  serviceName,
		Frameworks:   frameworks,
		TemplateFile: "./templates/generate.tmpl",
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.ParseFiles(t.TemplateFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func GenerateReviewText(filename string) string {
	file, _ := os.Open(filename)
	c, _ := io.ReadAll(file)

	buf := new(bytes.Buffer)
	r := ReviewTemplate{
		FileContent:  string(c),
		TemplateFile: "templates/review.tmpl",
	}
	tmpl, err := template.ParseFiles(r.TemplateFile)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(buf, r)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
