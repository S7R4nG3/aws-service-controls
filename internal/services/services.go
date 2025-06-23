package services

import (
	"github.com/S7R4nG3/aws-service-controls/internal/llm"
)

type Service struct {
	Short         string
	Long          string
	ControlPrompt llm.LlmPrompt
	ReviewPrompt  llm.LlmPrompt
}

var ModuleServices = []Service{
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
