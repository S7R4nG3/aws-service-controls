package services

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/S7R4nG3/aws-service-controls/internal/llm"
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
			// {
			// 	Short: "ecs",
			// 	Long:  "Elastic Container Service",
			// },
			// {
			// 	Short: "cloudfront",
			// 	Long:  "CloudFront",
			// },
		}
	} else {
		var serviceFile ServicesFile
		file, err := os.Open("services.json")
		if err != nil {
			panic(err)
		}
		contents, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(contents, &serviceFile)
		if err != nil {
			panic(err)
		}
		return serviceFile.Services
	}
}
