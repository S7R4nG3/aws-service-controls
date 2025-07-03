package main

import (
	"flag"
	"fmt"
	"sync"

	docgen "github.com/S7R4nG3/aws-service-controls/internal/doc-gen"
	"github.com/S7R4nG3/aws-service-controls/internal/llm"
	"github.com/S7R4nG3/aws-service-controls/internal/prompts"
	"github.com/S7R4nG3/aws-service-controls/internal/services"
)

const frameworks string = "CSA v5, NIST 800-53, AWS Foundational Security Best Practices, and AWS Security Hub"

func main() {
	servicesFile := flag.String("services-json", "services.json", "A JSON file of services to generate and review controls.")
	generation := flag.Bool("generate", false, "Toggle to enable generation of service controls.")
	reviewEnabled := flag.Bool("review", false, "Toggle to enable review of generated service controls.")
	docgenEnabled := flag.Bool("docs", true, "Toggle to enable Markdown doc generation of controls.")
	flag.Parse()

	moduleServices := services.LoadServices(*servicesFile)
	initServices := make(chan services.Service, len(moduleServices))
	controls := make(chan services.Service, len(moduleServices))
	reviews := make(chan string, len(moduleServices))
	// workerCount := len(moduleServices)
	controlWorkers := 2
	reviewWorkers := 2
	var wg sync.WaitGroup

	for i := 0; i < controlWorkers; i++ {
		go generateControls(initServices, controls, &wg, *generation, *docgenEnabled)
	}

	for t := 0; t < reviewWorkers; t++ {
		go reviewControls(controls, reviews, *reviewEnabled, *docgenEnabled)
	}

	wg.Add(len(moduleServices))
	for _, s := range moduleServices {
		fmt.Printf("Loading %s\n", s.Short)
		s.ControlPrompt.Outfile = "results/controls/" + s.Short + ".json"
		s.ReviewPrompt.Outfile = "results/reviews/" + s.Short + ".json"
		initServices <- s
	}
	close(initServices)
	wg.Wait()

	for r := 0; r < len(moduleServices); r++ {
		fmt.Println(<-reviews)
	}
}

func generateControls(services chan services.Service, controls chan services.Service, wg *sync.WaitGroup, enabled bool, docsEnabled bool) {
	for s := range services {
		if enabled {
			p := llm.NewLlmPrompt(
				llm.WithText(
					prompts.GenerateControlText(s.Long, frameworks),
				),
				llm.WithOutfile(s.ControlPrompt.Outfile),
				llm.WithPrefill("{"),
			)
			fmt.Printf("Generating controls for %s -> ", s.Short)
			p.Run()
		}
		if docsEnabled {
			docgen.Generate(s)
		}
		controls <- s
		wg.Done()
	}
}

func reviewControls(controls chan services.Service, reviews chan string, enabled bool, docsEnabled bool) {
	for c := range controls {
		if enabled {
			p := llm.NewLlmPrompt(
				llm.WithText(
					prompts.GenerateReviewText(c.ControlPrompt.Outfile),
				),
				llm.WithOutfile(c.ControlPrompt.Outfile),
				llm.WithPrefill("{"),
			)
			fmt.Printf("Reviewing controls for %s", c.Short)
			p.Run()
			reviews <- c.ReviewPrompt.Response
		} else {
			reviews <- "DONE"
		}
		if docsEnabled {
			docgen.Generate(c)
		}
	}
}
