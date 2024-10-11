package formatters

import (
	"fmt"
	services2 "github.com/KoNekoD/go-deptrac/pkg/application/services"
	"os"
	"path/filepath"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type GraphVizOutputImageFormatter struct {
	GraphVizOutputFormatter
}

func NewGraphVizOutputImageFormatter(config FormatterConfiguration) *GraphVizOutputImageFormatter {
	return &GraphVizOutputImageFormatter{
		GraphVizOutputFormatter: *NewGraphVizOutputFormatter(config),
	}
}

func (f *GraphVizOutputImageFormatter) GetName() string {
	return "graphviz-image"
}

func (f *GraphVizOutputImageFormatter) output(g *graphviz.Graphviz, graph *cgraph.Graph, output services2.OutputInterface, input OutputFormatterInput) error {
	dumpImagePath := input.OutputPath
	if dumpImagePath == nil || *dumpImagePath == "" {
		return fmt.Errorf("no '--output' defined for GraphViz formatter")
	}

	imagePathInfo := filepath.Dir(*dumpImagePath)
	if _, err := os.Stat(imagePathInfo); os.IsNotExist(err) {
		return fmt.Errorf("unable to dump image: Path \"%s\" does not exist or is not writable", imagePathInfo)
	}

	if err := g.RenderFilename(graph, f.getImageFormat(*dumpImagePath), *dumpImagePath); err != nil {
		return fmt.Errorf("unable to display output: %v", err)
	}

	output.WriteLineFormatted(services2.StringOrArrayOfStrings{String: fmt.Sprintf("<info>Image dumped to %s</>", *dumpImagePath)})
	return nil
}

func (f *GraphVizOutputImageFormatter) getImageFormat(filePath string) graphviz.Format {
	extension := filepath.Ext(filePath)
	if extension == "" {
		return graphviz.PNG
	}
	// Removing the dot from extension (e.g., ".png" -> "png")
	extension = extension[1:]

	switch extension {
	case "png":
		return graphviz.PNG
	case "jpg", "jpeg":
		return graphviz.JPG
	case "svg":
		return graphviz.SVG
	default:
		return graphviz.PNG
	}
}
