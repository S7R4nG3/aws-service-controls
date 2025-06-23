package main

import (
	"bufio"
	"encoding/json"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
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

func main() {
	var tmplFileName = "table.tmpl"
	var ctrlDir = "security/"
	var mkdnFileName = "controls.md"

	file, err := os.Create(mkdnFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	filepath.Walk(ctrlDir, func(path string, info fs.FileInfo, err error) error {
		if strings.Contains(info.Name(), ".json") {
			controls := Controls{}
			ctrlFile, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer ctrlFile.Close()
			contents, err := io.ReadAll(ctrlFile)
			json.Unmarshal(contents, &controls)

			tmpl, err := template.New(tmplFileName).ParseFiles(tmplFileName)
			if err != nil {
				panic(err)
			}
			err = tmpl.Execute(writer, controls)
			if err != nil {
				panic(err)
			}
		}
		return nil
	})
	writer.Flush()
}
