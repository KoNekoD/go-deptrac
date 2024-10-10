package formatters

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg"
	"os/exec"
	"runtime"
	"time"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

type GraphVizOutputDisplayFormatter struct {
	GraphVizOutputFormatter
}

const DelayOpen = 2 * time.Second

func NewGraphVizOutputDisplayFormatter(config FormatterConfiguration) *GraphVizOutputDisplayFormatter {
	return &GraphVizOutputDisplayFormatter{
		GraphVizOutputFormatter: *NewGraphVizOutputFormatter(config),
	}
}

func (f *GraphVizOutputDisplayFormatter) GetName() string {
	return "graphviz-display"
}

func (f *GraphVizOutputDisplayFormatter) output(g *graphviz.Graphviz, graph *cgraph.Graph, output pkg.OutputInterface, input OutputFormatterInput) error {
	filename, err := f.getTempImage(g, graph)
	if err != nil {
		return fmt.Errorf("unable to create temp file_supportive for output: %v", err)
	}

	staticNext := time.Now()
	if time.Now().Before(staticNext) {
		time.Sleep(DelayOpen)
	}

	var openCmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		openCmd = exec.Command("cmd", "/C", "start", "", filename)
	case "darwin":
		openCmd = exec.Command("open", filename)
	default: // Assuming Linux or other Unix-like OS
		openCmd = exec.Command("xdg-open", filename)
	}

	if err := openCmd.Start(); err != nil {
		return fmt.Errorf("unable to display output: %v", err)
	}

	staticNext = time.Now().Add(DelayOpen)
	return nil
}
