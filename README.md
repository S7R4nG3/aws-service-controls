# aws-service-controls

Generates Security Control documentation for AWS services using Claude LLM to generate the content, then immediately review the content for accuracy and suggest changes.

Security Controls are created as structured JSON files that are then templated into standard markdown for human readability.

## Usage

Clone the repo down, supply a list of Short and Long names for services as a JSON file (structured in the repo), ensure your Anthropic API key is set as an environment variable and run the program.

JSON content will be aggregated under `results/controls` and documentation markdown tables under `results/docs`.

Generate and review content immediately with `go run . -generate -review`.

### Argument Reference
| Name | Description | Default |
| - | - | - |
| `-generate` | Toggle to activate total regeneration of security controls from the provided frameworks. | `false` |
| `-review` | Toggle to activate a review of the controls from the provided frameworks. | `false` |
| `-docs` | Toggle to enable markdown documentation regeneration for the generated controls. | `true` |