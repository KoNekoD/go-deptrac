package output_formatter

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result/output_result"
	"github.com/KoNekoD/go-deptrac/pkg/supportive/output_formatter/configuration"
	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
	"os"
	"path/filepath"
	"slices"
)

type GraphVizOutputFormatter struct {
	config configuration.ConfigurationGraphViz
}

func NewGraphVizOutputFormatter(config configuration.FormatterConfiguration) *GraphVizOutputFormatter {
	extractedConfig := config.GetConfigFor("graphviz").(interface{}).(configuration.ConfigurationGraphViz)
	return &GraphVizOutputFormatter{config: extractedConfig}
}

func (f *GraphVizOutputFormatter) Finish(result output_result.OutputResult, output output_formatter.OutputInterface, input output_formatter.OutputFormatterInput) error {
	layerViolations := f.calculateViolations(result.Violations())
	layersDependOnLayers := f.calculateLayerDependencies(result.AllRules())

	outputConfig := f.config
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		return err
	}
	defer func(graph *cgraph.Graph) {
		err := graph.Close()
		if err != nil {
			fmt.Println("failed to close graph", err)
		}
	}(graph)

	nodes := f.createNodes(outputConfig, layersDependOnLayers, graph)
	f.addNodesToGraph(graph, nodes, outputConfig)
	f.connectEdges(graph, nodes, outputConfig, layersDependOnLayers, layerViolations)

	if err := f.output(g, graph, output, input); err != nil {
		return err
	}
	return nil
}

func (f *GraphVizOutputFormatter) calculateViolations(violations []*result.Violation) map[string]map[string]int {
	layerViolations := make(map[string]map[string]int)
	for _, violation := range violations {
		dependerLayer := violation.GetDependerLayer()
		dependentLayer := violation.GetDependentLayer()

		if layerViolations[dependerLayer] == nil {
			layerViolations[dependerLayer] = make(map[string]int)
		}

		layerViolations[dependerLayer][dependentLayer]++
	}
	return layerViolations
}

func (f *GraphVizOutputFormatter) calculateLayerDependencies(rules []result.RuleInterface) map[string]map[string]int {
	layersDependOnLayers := make(map[string]map[string]int)
	for _, rule := range rules {
		switch r := rule.(type) {
		case result.CoveredRuleInterface:
			layerA := r.GetDependerLayer()
			layerB := r.GetDependentLayer()

			if layersDependOnLayers[layerA] == nil {
				layersDependOnLayers[layerA] = make(map[string]int)
			}

			layersDependOnLayers[layerA][layerB]++
		case *result.Uncovered:
			if layersDependOnLayers[r.Layer] == nil {
				layersDependOnLayers[r.Layer] = make(map[string]int)
			}
		}
	}
	return layersDependOnLayers
}

func (f *GraphVizOutputFormatter) createNodes(outputConfig configuration.ConfigurationGraphViz, layersDependOnLayers map[string]map[string]int, graph *cgraph.Graph) map[string]*cgraph.Node {
	nodes := make(map[string]*cgraph.Node)
	for layer, layersDependOn := range layersDependOnLayers {
		if slices.Contains(outputConfig.HiddenLayers, layer) {
			continue
		}
		if nodes[layer] == nil {
			nodes[layer], _ = graph.CreateNode(layer)
		}
		for layerDependOn := range layersDependOn {
			if slices.Contains(outputConfig.HiddenLayers, layerDependOn) {
				continue
			}
			if nodes[layerDependOn] == nil {
				nodes[layerDependOn], _ = graph.CreateNode(layerDependOn)
			}
		}
	}
	return nodes
}

func (f *GraphVizOutputFormatter) connectEdges(graph *cgraph.Graph, nodes map[string]*cgraph.Node, outputConfig configuration.ConfigurationGraphViz, layersDependOnLayers, layerViolations map[string]map[string]int) {
	for layer, layersDependOn := range layersDependOnLayers {
		if slices.Contains(outputConfig.HiddenLayers, layer) {
			continue
		}
		for layerDependOn, layerDependOnCount := range layersDependOn {
			if slices.Contains(outputConfig.HiddenLayers, layerDependOn) {
				continue
			}
			edge, _ := graph.CreateEdge(fmt.Sprintf("%s->%s", layer, layerDependOn), nodes[layer], nodes[layerDependOn])
			if outputConfig.PointToGroups && graph.SubGraph(f.getSubgraphName(layerDependOn), 0) != nil {
				edge.Set("lhead", f.getSubgraphName(layerDependOn))
			}
			if count, ok := layerViolations[layer][layerDependOn]; ok {
				edge.SetLabel(fmt.Sprintf("%d", count))
				edge.Set("color", "red")
			} else {
				edge.SetLabel(fmt.Sprintf("%d", layerDependOnCount))
			}
		}
	}
}

func (f *GraphVizOutputFormatter) addNodesToGraph(graph *cgraph.Graph, nodes map[string]*cgraph.Node, outputConfig configuration.ConfigurationGraphViz) {
	for groupName, groupLayerNames := range outputConfig.GroupsLayerMap {
		subgraph := graph.SubGraph(f.getSubgraphName(groupName), 1)
		subgraph.SetLabel(groupName)
		for _, groupLayerName := range groupLayerNames {
			if node, exists := nodes[groupLayerName]; exists {
				subgraph.NextNode(node)
				node.Set("group", groupName)
				delete(nodes, groupLayerName)
			}
		}
	}

	for _, node := range nodes {
		graph.NextNode(node)
	}
}

func (f *GraphVizOutputFormatter) output(g *graphviz.Graphviz, graph *cgraph.Graph, output output_formatter.OutputInterface, input output_formatter.OutputFormatterInput) error {
	filename, err := f.getTempImage(g, graph)
	if err != nil {
		return fmt.Errorf("unable to create temp file for output: %v", err)
	}

	if input.OutputPath != nil && *input.OutputPath != "" {
		if err := os.Rename(filename, *input.OutputPath); err != nil {
			return fmt.Errorf("unable to move temp file to output path: %v", err)
		}
		output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<info>GraphViz Report saved to %s</>", filepath.Clean(*input.OutputPath))})
		return nil
	}

	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<info>GraphViz temp image created at %s</>", filename)})
	return nil
}

func (f *GraphVizOutputFormatter) getTempImage(g *graphviz.Graphviz, graph *cgraph.Graph) (string, error) {
	tempFile, err := os.CreateTemp("", "deptrac-*.png")
	if err != nil {
		return "", fmt.Errorf("unable to create temp file: %v", err)
	}
	defer func(tempFile *os.File) {
		err := tempFile.Close()
		if err != nil {
			fmt.Printf("unable to close temp file: %v", err)
		}
	}(tempFile)

	filename := tempFile.Name()
	if err := g.RenderFilename(graph, graphviz.PNG, filename); err != nil {
		return "", fmt.Errorf("unable to export graph to image: %v", err)
	}

	return filename, nil
}

func (f *GraphVizOutputFormatter) getSubgraphName(groupName string) string {
	return "cluster_" + groupName
}
