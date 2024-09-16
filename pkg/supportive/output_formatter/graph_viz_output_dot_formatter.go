package output_formatter

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/output_formatter/configuration"
	"os"
	"path/filepath"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type GraphVizOutputDotFormatter struct {
	GraphVizOutputFormatter
}

func NewGraphVizOutputDotFormatter(config configuration.FormatterConfiguration) *GraphVizOutputDotFormatter {
	return &GraphVizOutputDotFormatter{
		GraphVizOutputFormatter: *NewGraphVizOutputFormatter(config),
	}
}

func (f *GraphVizOutputDotFormatter) GetName() string {
	return "graphviz-dot"
}

func (f *GraphVizOutputDotFormatter) output(g *graphviz.Graphviz, graph *cgraph.Graph, output output_formatter.OutputInterface, input output_formatter.OutputFormatterInput) error {
	dumpDotPath := input.OutputPath
	if dumpDotPath == nil || *dumpDotPath == "" {
		return fmt.Errorf("no '--output' defined for GraphViz formatter")
	}

	// Render the graph to DOT format
	wr, _ := os.Create(filepath.Clean(*dumpDotPath))
	err := g.Render(graph, graphviz.XDOT, wr)
	if err != nil {
		return fmt.Errorf("unable to render graph to DOT format: %v", err)
	}

	// Write the DOT data to the specified file
	//if err := os.WriteFile(dumpDotPath, dotData, 0644); err != nil {
	//	return fmt.Errorf("unable to write DOT data to file: %v", err)
	//}

	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<info>Script dumped to %s</>", filepath.Clean(*dumpDotPath))})
	return nil
}
