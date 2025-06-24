package main

import (
	"fmt"
	"sync"

	"github.com/S7R4nG3/aws-service-controls/internal/llm"
	"github.com/S7R4nG3/aws-service-controls/internal/prompts"
	"github.com/S7R4nG3/aws-service-controls/internal/services"
)

const frameworks string = "CSA v5, NIST 800-53, AWS Foundational Security Best Practices, and AWS Security Hub"

func main() {
	initServices := make(chan services.Service, len(services.ModuleServices))
	controls := make(chan services.Service, len(services.ModuleServices))
	reviews := make(chan string, len(services.ModuleServices))
	workerCount := len(services.ModuleServices)
	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		go generateControls(initServices, controls, &wg)
	}

	for t := 0; t < workerCount; t++ {
		go reviewControls(controls, reviews)
	}

	wg.Add(len(services.ModuleServices))
	for _, s := range services.ModuleServices {
		fmt.Printf("Loading %s\n", s.Short)
		s.ControlPrompt.Outfile = "controls/" + s.Short + ".json"
		s.ReviewPrompt.Outfile = "reviews/" + s.Short + ".md"
		initServices <- s
	}
	close(initServices)
	wg.Wait()

	for r := 0; r < len(services.ModuleServices); r++ {
		fmt.Println(<-reviews)
	}
}

func generateControls(services chan services.Service, controls chan services.Service, wg *sync.WaitGroup) {
	for s := range services {
		p := llm.NewLlmPrompt(
			llm.WithText(
				prompts.GenerateControlText(s.Long, frameworks),
			),
			llm.WithOutfile(s.ControlPrompt.Outfile),
			llm.WithPrefill("{"),
		)
		fmt.Printf("Generating controls for %s -> ", s.Short)
		p.Run()
		controls <- s
		wg.Done()
	}
}

func reviewControls(controls chan services.Service, reviews chan string) {
	for c := range controls {
		p := llm.NewLlmPrompt(
			llm.WithText(
				prompts.GenerateReviewText(c.ControlPrompt.Outfile),
			),
			llm.WithOutfile(c.ReviewPrompt.Outfile),
		)
		fmt.Printf("Reviewing controls for %s", c.Short)
		p.Run()
		reviews <- c.ReviewPrompt.Response
	}
}
