package docgen

import (
	"bufio"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/template"

	"github.com/S7R4nG3/aws-service-controls/internal/services"
	"github.com/S7R4nG3/aws-service-controls/internal/utils"
)

type Controls struct {
	Service  string `json:"service"`
	Security []struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		Controls []struct {
			Identifier     string `json:"identifier"`
			Title          string `json:"title"`
			Severity       string `json:"severity"`
			Description    string `json:"description"`
			Implementation struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Code        string `json:"code"`
			} `json:"implementation"`
		} `json:"controls"`
	} `json:"security"`
	Operational []struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		Controls []struct {
			Identifier     string `json:"identifier"`
			Title          string `json:"title"`
			Severity       string `json:"severity"`
			Description    string `json:"description"`
			Implementation struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Code        string `json:"code"`
			} `json:"implementation"`
		} `json:"controls"`
	} `json:"operational"`
	Cost []struct {
		Name     string `json:"name"`
		Version  string `json:"version"`
		Controls []struct {
			Identifier     string `json:"identifier"`
			Title          string `json:"title"`
			Severity       string `json:"severity"`
			Description    string `json:"description"`
			Implementation struct {
				Type        string `json:"type"`
				Description string `json:"description"`
				Code        string `json:"code"`
			} `json:"implementation"`
		} `json:"controls"`
	} `json:"cost"`
}

//go:embed templates
var templates embed.FS

func Generate(service services.Service) {
	fmt.Printf("Generating docs for %s -> ", service.Long)
	var tmplFileName = "templates/table.tmpl"
	var control = "results/controls/" + service.Short + ".json"
	var mkdnFileName = "results/docs/" + service.Short + ".md"
	file, err := os.Create(mkdnFileName)
	utils.Check(err, "Error creating markdown file...")
	defer file.Close()
	writer := bufio.NewWriter(file)

	controls := Controls{}
	ctrlFile, err := os.Open(control)
	utils.Check(err, "Error opening control file...")
	defer ctrlFile.Close()
	contents, err := io.ReadAll(ctrlFile)
	utils.Check(err, "Error reading control file contents...")
	json.Unmarshal(contents, &controls)

	tmpl, err := template.ParseFS(templates, tmplFileName)
	utils.Check(err, "Error templating control contents...")
	err = tmpl.Execute(writer, controls)
	utils.Check(err, "Error executing control contents...")
	writer.Flush()
	fmt.Print("DONE..")
}
