package formatters

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/formatters_configs"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"os"
	"strings"
)

type MermaidJSOutputFormatter struct {
	config *formatters_configs.MermaidJsConfig
}

const (
	GraphTypeFormat      = "flowchart %s;\n"
	GraphEnd             = "  end;\n"
	SubgraphFormat       = "  subgraph %sGroup;\n"
	LayerFormat          = "    %s;\n"
	GraphNodeFormat      = "    %s -->|%d| %s;\n"
	ViolationStyleFormat = "    linkStyle %d stroke:red,stroke-width:4px;\n"
	DefaultOutputPath    = "./mermaid-graph.md"
)

func NewMermaidJSOutputFormatter(config FormatterConfiguration) *MermaidJSOutputFormatter {
	extractedConfig := config.GetConfigFor("mermaidjs").(interface{}).(*formatters_configs.MermaidJsConfig)
	return &MermaidJSOutputFormatter{config: extractedConfig}
}

func (f *MermaidJSOutputFormatter) GetName() string {
	return "mermaidjs"
}

func (f *MermaidJSOutputFormatter) Finish(result results.OutputResult, output results.OutputInterface, input OutputFormatterInput) error {
	graph := f.parseResults(result)
	violations := result.Violations()
	var buffer strings.Builder

	buffer.WriteString(fmt.Sprintf(GraphTypeFormat, f.config.Direction))

	for subGraphName, layers := range f.config.Groups {
		buffer.WriteString(fmt.Sprintf(SubgraphFormat, subGraphName))
		for _, layer := range layers {
			buffer.WriteString(fmt.Sprintf(LayerFormat, layer.Name))
		}
		buffer.WriteString(GraphEnd)
	}

	linkCount := 0
	violationsLinks := make(map[string]map[string]int)
	violationGraphLinks := make([]int, 0)

	for _, violation := range violations {
		dependerLayer := violation.GetDependerLayer()
		dependentLayer := violation.GetDependentLayer()

		if violationsLinks[dependerLayer] == nil {
			violationsLinks[dependerLayer] = make(map[string]int)
		}

		violationsLinks[dependerLayer][dependentLayer]++
	}

	for dependerLayer, layers := range violationsLinks {
		for dependentLayer, count := range layers {
			buffer.WriteString(fmt.Sprintf(GraphNodeFormat, dependerLayer, count, dependentLayer))
			violationGraphLinks = append(violationGraphLinks, linkCount)
			linkCount++
		}
	}

	for dependerLayer, layers := range graph {
		for dependentLayer, count := range layers {
			if _, exists := violationsLinks[dependerLayer][dependentLayer]; !exists {
				buffer.WriteString(fmt.Sprintf(GraphNodeFormat, dependerLayer, count, dependentLayer))
			}
		}
	}

	for _, linkNumber := range violationGraphLinks {
		buffer.WriteString(fmt.Sprintf(ViolationStyleFormat, linkNumber))
	}

	if input.OutputPath != nil && *input.OutputPath != "" {
		if err := os.WriteFile(*input.OutputPath, []byte(buffer.String()), 0644); err != nil {
			return err
		}
	} else {
		output.WriteRaw(buffer.String())
	}

	return nil
}

func (f *MermaidJSOutputFormatter) parseResults(result results.OutputResult) map[string]map[string]int {
	graph := make(map[string]map[string]int)
	for _, rule := range result.Allowed() {
		dependerLayer := rule.GetDependerLayer()
		dependentLayer := rule.GetDependentLayer()

		if graph[dependerLayer] == nil {
			graph[dependerLayer] = make(map[string]int)
		}

		graph[dependerLayer][dependentLayer]++
	}
	return graph
}
