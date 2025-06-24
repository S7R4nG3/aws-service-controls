# aws-service-controls

Generates Security Control documentation for AWS services using Claude LLM to generate the content, then immediately review the content for accuracy and suggest changes.

## Usage

Clone the repo down, supply a list of Short and Long names for services as a JSON file (structured in the repo), ensure your Anthropic API key is set as an environment variable and run the program.

JSON content will be aggregated under `results/controls` and reviews will be created afterwards under `results/reviews`.