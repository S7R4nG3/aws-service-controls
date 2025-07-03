package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/S7R4nG3/aws-service-controls/internal/llm"
	"github.com/S7R4nG3/aws-service-controls/internal/utils"
)

type ServicesFile struct {
	Services []Service `json:"services"`
}

type Service struct {
	Short         string `json:"short"`
	Long          string `json:"long"`
	ControlPrompt llm.LlmPrompt
	ReviewPrompt  llm.LlmPrompt
}

func LoadServices(filename string) []Service {
	wd, _ := os.Getwd()
	filepath := filepath.Join(wd, filename)
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		log.Fatal(err)
		return []Service{}
	} else {
		var serviceFile ServicesFile
		file, err := os.Open("services.json")
		utils.Check(err, "Error opening services.json file...")
		contents, err := io.ReadAll(file)
		utils.Check(err, "Error reading services.json contents...")
		err = json.Unmarshal(contents, &serviceFile)
		utils.Check(err, "Error unmarshalling services.json contents...")
		return serviceFile.Services
	}
}
