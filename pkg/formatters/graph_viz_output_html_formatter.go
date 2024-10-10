package formatters

import (
	"encoding/base64"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"os"
	"path/filepath"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type GraphVizOutputHtmlFormatter struct {
	GraphVizOutputFormatter
}

func NewGraphVizOutputHtmlFormatter(config FormatterConfiguration) *GraphVizOutputHtmlFormatter {
	return &GraphVizOutputHtmlFormatter{
		GraphVizOutputFormatter: *NewGraphVizOutputFormatter(config),
	}
}

func (f *GraphVizOutputHtmlFormatter) GetName() string {
	return "graphviz-html"
}

func (f *GraphVizOutputHtmlFormatter) output(g *graphviz.Graphviz, graph *cgraph.Graph, output results.OutputInterface, input OutputFormatterInput) error {
	dumpHtmlPath := input.OutputPath
	if dumpHtmlPath == nil || *dumpHtmlPath == "" {
		return fmt.Errorf("no '--output' defined for GraphViz formatter")
	}

	// Generate a temporary image file_supportive
	filename, err := f.getTempImage(g, graph)
	if err != nil {
		return fmt.Errorf("unable to create temp file_supportive for output: %v", err)
	}

	// Ensure the temporary file_supportive is removed after processing
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			fmt.Printf("unable to remove temp file_supportive: %v", err)
		}
	}(filename)

	// Read the image data
	imageData, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("unable to read temp image file_supportive: %v", err)
	}

	// Encode the image data to base64
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	// Create the HTML content with the embedded base64 image
	htmlContent := fmt.Sprintf(`<img src="data:image/png;base64,%s" />`, base64Image)

	// Write the HTML content to the specified file_supportive
	if err := os.WriteFile(*dumpHtmlPath, []byte(htmlContent), 0644); err != nil {
		return fmt.Errorf("unable to write HTML file_supportive: %v", err)
	}

	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<info>HTML dumped to %s</>", filepath.Clean(*dumpHtmlPath))})
	return nil
}
