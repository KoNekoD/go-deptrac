package output_formatter

import (
	"encoding/base64"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/output_formatter/configuration"
	"os"
	"path/filepath"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type GraphVizOutputHtmlFormatter struct {
	GraphVizOutputFormatter
}

func NewGraphVizOutputHtmlFormatter(config configuration.FormatterConfiguration) *GraphVizOutputHtmlFormatter {
	return &GraphVizOutputHtmlFormatter{
		GraphVizOutputFormatter: *NewGraphVizOutputFormatter(config),
	}
}

func (f *GraphVizOutputHtmlFormatter) GetName() string {
	return "graphviz-html"
}

func (f *GraphVizOutputHtmlFormatter) output(g *graphviz.Graphviz, graph *cgraph.Graph, output output_formatter.OutputInterface, input output_formatter.OutputFormatterInput) error {
	dumpHtmlPath := input.OutputPath
	if dumpHtmlPath == nil || *dumpHtmlPath == "" {
		return fmt.Errorf("no '--output' defined for GraphViz formatter")
	}

	// Generate a temporary image file
	filename, err := f.getTempImage(g, graph)
	if err != nil {
		return fmt.Errorf("unable to create temp file for output: %v", err)
	}

	// Ensure the temporary file is removed after processing
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Printf("unable to remove temp file: %v", err)
		}
	}(filename)

	// Read the image data
	imageData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read temp image file: %v", err)
	}

	// Encode the image data to base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	// Create the HTML content with the embedded base64 image
	htmlContent := fmt.Sprintf(`<img src="data:image/png;base64,%s" />`, base64Image)

	// Write the HTML content to the specified file
	if err := os.WriteFile(*dumpHtmlPath, []byte(htmlContent), 0644); err != nil {
		return fmt.Errorf("unable to write HTML file: %v", err)
	}

	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<info>HTML dumped to %s</>", filepath.Clean(*dumpHtmlPath))})
	return nil
}
