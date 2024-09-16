package output_formatter

import (
	"encoding/xml"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/contract/output_formatter"
	result2 "github.com/KoNekoD/go-deptrac/pkg/contract/result"
	"github.com/KoNekoD/go-deptrac/pkg/contract/result/output_result"
	"github.com/KoNekoD/go-deptrac/pkg/util"
	"os"
	"path/filepath"
)

type JUnitOutputFormatter struct{}

const DefaultJUnitPath = "./junit-report.xml"

func NewJUnitOutputFormatter() *JUnitOutputFormatter {
	return &JUnitOutputFormatter{}
}

func (f *JUnitOutputFormatter) GetName() string {
	return "junit"
}

func (f *JUnitOutputFormatter) Finish(result output_result.OutputResult, output output_formatter.OutputInterface, input output_formatter.OutputFormatterInput) error {
	xmlData, err := f.createXML(result)
	if err != nil {
		return err
	}

	dumpXmlPath := input.OutputPath
	if dumpXmlPath == nil || *dumpXmlPath == "" {
		dumpXmlPath = util.AsPtr(DefaultJUnitPath)
	}

	if err := os.WriteFile(*dumpXmlPath, []byte(xmlData), 0644); err != nil {
		return err
	}

	output.WriteLineFormatted(output_formatter.StringOrArrayOfStrings{String: fmt.Sprintf("<info>JUnit Report dumped to %s</>", filepath.Clean(*dumpXmlPath))})
	return nil
}

func (f *JUnitOutputFormatter) createXML(outputResult output_result.OutputResult) (string, error) {
	testSuites := TestSuites{}
	if outputResult.HasErrors() {
		errorSuite := TestSuite{
			ID:       "0",
			Name:     "Unmatched skipped violations",
			Hostname: "localhost",
			Errors:   len(outputResult.Errors),
		}
		for _, message := range outputResult.Errors {
			errorSuite.ErrorsList = append(errorSuite.ErrorsList, Error{
				Message: message.ToString(),
				Type:    "WARNING",
			})
		}
		testSuites.TestSuites = append(testSuites.TestSuites, errorSuite)
	}

	layers := f.groupRulesByLayer(outputResult)
	layerIndex := 0
	for layer, rules := range layers {
		layerIndex++
		testSuite := TestSuite{
			ID:       fmt.Sprintf("%d", layerIndex),
			Name:     layer,
			Hostname: "localhost",
		}
		testSuite.Tests = len(rules)
		for _, rule := range rules {
			testCase := TestCase{
				Name:      fmt.Sprintf("%s - %s", layer, rule.GetDependency().GetDepender().ToString()),
				Classname: rule.GetDependency().GetDepender().ToString(),
			}
			switch r := rule.(type) {
			case *result2.Violation:
				testSuite.Failures++
				testCase.Failures = append(testCase.Failures, Failure{
					Message: fmt.Sprintf("%s:%d must not depend on %s (%s on %s)",
						r.GetDependency().GetDepender().ToString(),
						r.GetDependency().GetContext().FileOccurrence.Line,
						r.GetDependency().GetDependent().ToString(),
						r.GetDependerLayer(),
						r.GetDependentLayer(),
					),
					Type: "WARNING",
				})
			case *result2.SkippedViolation:
				testSuite.Skipped++
				testCase.Skipped = append(testCase.Skipped, Skipped{})
			case *result2.Uncovered:
				testCase.Warnings = append(testCase.Warnings, Warning{
					Message: fmt.Sprintf("%s:%d has uncovered dependency on %s (%s)",
						r.GetDependency().GetDepender().ToString(),
						r.GetDependency().GetContext().FileOccurrence.Line,
						r.GetDependency().GetDependent().ToString(),
						r.Layer,
					),
					Type: "WARNING",
				})
			}
			testSuite.TestCases = append(testSuite.TestCases, testCase)
		}
		testSuites.TestSuites = append(testSuites.TestSuites, testSuite)
	}

	xmlData, err := xml.MarshalIndent(testSuites, "", "  ")
	if err != nil {
		return "", fmt.Errorf("unable to create XML: %v", err)
	}

	return xml.Header + string(xmlData), nil
}

func (f *JUnitOutputFormatter) groupRulesByLayer(outputResult output_result.OutputResult) map[string][]result2.RuleInterface {
	layers := make(map[string][]result2.RuleInterface)
	for _, rule := range outputResult.AllRules() {
		switch r := rule.(type) {
		case result2.CoveredRuleInterface:
			layers[r.GetDependerLayer()] = append(layers[r.GetDependerLayer()], rule)
		case *result2.Uncovered:
			layers[r.Layer] = append(layers[r.Layer], rule)
		}
	}
	return layers
}

type TestSuites struct {
	XMLName    xml.Name    `xml:"testsuites"`
	TestSuites []TestSuite `xml:"testsuite"`
}

type TestSuite struct {
	ID         string     `xml:"id,attr"`
	Name       string     `xml:"name,attr"`
	Hostname   string     `xml:"hostname,attr"`
	Tests      int        `xml:"tests,attr"`
	Failures   int        `xml:"failures,attr"`
	Skipped    int        `xml:"skipped,attr"`
	Errors     int        `xml:"errors,attr"`
	TestCases  []TestCase `xml:"testcase"`
	ErrorsList []Error    `xml:"error"`
}

type TestCase struct {
	Name      string    `xml:"name,attr"`
	Classname string    `xml:"classname,attr"`
	Failures  []Failure `xml:"failure"`
	Skipped   []Skipped `xml:"skipped"`
	Warnings  []Warning `xml:"warning"`
}

type Failure struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
}

type Skipped struct {
	XMLName xml.Name `xml:"skipped"`
}

type Warning struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
}

type Error struct {
	Message string `xml:"message,attr"`
	Type    string `xml:"type,attr"`
}
