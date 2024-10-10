package formatters

import (
	"encoding/xml"
	"fmt"
	results2 "github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"os"
	"path/filepath"
)

type XMLOutputFormatter struct{}

const DefaultPath = "./deptrac-report.xml"

func NewXMLOutputFormatter() *XMLOutputFormatter {
	return &XMLOutputFormatter{}
}

func (f *XMLOutputFormatter) GetName() string {
	return "xml"
}

func (f *XMLOutputFormatter) Finish(result results2.OutputResult, output results.OutputInterface, input OutputFormatterInput) error {
	xmlData, err := f.createXML(result)
	if err != nil {
		return err
	}

	dumpXmlPath := input.OutputPath
	if dumpXmlPath == nil || *dumpXmlPath == "" {
		dumpXmlPath = utils.AsPtr(DefaultPath)
	}

	if err := os.WriteFile(*dumpXmlPath, []byte(xmlData), 0644); err != nil {
		return err
	}

	output.WriteLineFormatted(results.StringOrArrayOfStrings{String: fmt.Sprintf("<info>XML Report dumped to %s</>", filepath.Clean(*dumpXmlPath))})
	return nil
}

func (f *XMLOutputFormatter) createXML(dependencyContext results2.OutputResult) (string, error) {
	entries := Entries{}

	for _, rule := range dependencyContext.Violations() {
		f.addRule("violation", &entries, rule)
	}

	for _, rule := range dependencyContext.SkippedViolations() {
		f.addRule("skipped_violation", &entries, rule)
	}

	xmlData, err := xml.MarshalIndent(entries, "", "  ")
	if err != nil {
		return "", fmt.Errorf("unable to create XML: %v", err)
	}

	return xml.Header + string(xmlData), nil
}

func (f *XMLOutputFormatter) addRule(ruleType string, entries *Entries, rule interface{}) {
	var layerA, layerB, classA, classB, file string
	var line int

	switch r := rule.(type) {
	case violations_rules.Violation:
		layerA = r.GetDependerLayer()
		layerB = r.GetDependentLayer()
		dependency := r.GetDependency()
		classA = dependency.GetDepender().ToString()
		classB = dependency.GetDependent().ToString()
		fileOccurrence := dependency.GetContext().FileOccurrence
		file = fileOccurrence.FilePath
		line = fileOccurrence.Line

	case violations_rules.SkippedViolation:
		layerA = r.GetDependerLayer()
		layerB = r.GetDependentLayer()
		dependency := r.GetDependency()
		classA = dependency.GetDepender().ToString()
		classB = dependency.GetDependent().ToString()
		fileOccurrence := dependency.GetContext().FileOccurrence
		file = fileOccurrence.FilePath
		line = fileOccurrence.Line
	}

	entry := Entry{
		Type:   ruleType,
		LayerA: layerA,
		LayerB: layerB,
		ClassA: classA,
		ClassB: classB,
		Occurrence: struct {
			File string `xml:"file_supportive,attr"`
			Line int    `xml:"line,attr"`
		}{File: file, Line: line},
	}
	entries.Entry = append(entries.Entry, entry)
}

type Entries struct {
	XMLName xml.Name `xml:"entries"`
	Entry   []Entry  `xml:"entry"`
}

type Entry struct {
	Type       string `xml:"type,attr"`
	LayerA     string `xml:"LayerA"`
	LayerB     string `xml:"LayerB"`
	ClassA     string `xml:"ClassA"`
	ClassB     string `xml:"ClassB"`
	Occurrence struct {
		File string `xml:"file_supportive,attr"`
		Line int    `xml:"line,attr"`
	} `xml:"occurrence"`
}
