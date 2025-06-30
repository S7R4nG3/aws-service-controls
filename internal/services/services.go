package services

import (
	"encoding/json"
	"errors"
	"io"
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

var ModuleServices = LoadServices()

func LoadServices() []Service {
	wd, _ := os.Getwd()
	filepath := filepath.Join(wd, "services.json")
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return []Service{
			{
				Short: "lambda",
				Long:  "Lambda",
			},
			{
				Short: "ecs",
				Long:  "Elastic Container Service",
			},
			{
				Short: "cloudfront",
				Long:  "CloudFront",
			},
		}
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
