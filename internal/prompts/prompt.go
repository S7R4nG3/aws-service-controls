package prompts

import (
	"bytes"
	"embed"
	"io"
	"os"
	"text/template"

	"github.com/S7R4nG3/aws-service-controls/internal/utils"
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

//go:embed templates
var templates embed.FS

func GenerateControlText(serviceName string, frameworks string) string {
	t := GenerateTemplate{
		ServiceName:  serviceName,
		Frameworks:   frameworks,
		TemplateFile: "templates/generate.tmpl",
	}
	buf := new(bytes.Buffer)
	tmpl, err := template.ParseFS(templates, t.TemplateFile)
	utils.Check(err, "Error parsing control text templates...")
	err = tmpl.Execute(buf, t)
	utils.Check(err, "Error executing control text template...")
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
	tmpl, err := template.ParseFS(templates, r.TemplateFile)
	utils.Check(err, "Error parsing review text templates...")
	err = tmpl.Execute(buf, r)
	utils.Check(err, "Error executing review text template...")
	return buf.String()
}
