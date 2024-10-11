package formatters

import (
	"encoding/xml"
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results"
	"github.com/KoNekoD/go-deptrac/pkg/domain/dtos/results/violations_rules"
	"github.com/KoNekoD/go-deptrac/pkg/domain/utils"
	"github.com/KoNekoD/go-deptrac/pkg/infrastructure/services"
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

func (f *JUnitOutputFormatter) Finish(result results.OutputResult, output services.OutputInterface, input OutputFormatterInput) error {
	xmlData, err := f.createXML(result)
	if err != nil {
		return err
	}

	dumpXmlPath := input.OutputPath
	if dumpXmlPath == nil || *dumpXmlPath == "" {
		dumpXmlPath = utils.AsPtr(DefaultJUnitPath)
	}

	if err := os.WriteFile(*dumpXmlPath, []byte(xmlData), 0644); err != nil {
		return err
	}

	output.WriteLineFormatted(services.StringOrArrayOfStrings{String: fmt.Sprintf("<info>JUnit Report dumped to %s</>", filepath.Clean(*dumpXmlPath))})
	return nil
}

func (f *JUnitOutputFormatter) createXML(outputResult results.OutputResult) (string, error) {
	testSuites := TestSuites{}
	if outputResult.HasErrors() {
		errorSuite := TestSuite{
			ID:       "0",
			Name:     "Unmatched skipped violations",
			Hostname: "localhost",
			Errors:   len(outputResult.Errors),
		}
		for _, message := range outputResult.Errors {
			errorSuite.ErrorsList = append(errorSuite.ErrorsList, struct {
				Message string `xml:"message,attr"`
				Type    string `xml:"type,attr"`
			}{
				Message: message.ToString(),
				Type:    "WARNING",
			})
		}
		testSuites.TestSuites = append(testSuites.TestSuites, errorSuite)
	}

	layers := f.groupRulesByLayer(outputResult)
	layerIndex := 0
	for layer, rulesList := range layers {
		layerIndex++
		testSuite := TestSuite{
			ID:       fmt.Sprintf("%d", layerIndex),
			Name:     layer,
			Hostname: "localhost",
		}
		testSuite.Tests = len(rulesList)
		for _, rule := range rulesList {
			testCase := TestCase{
				Name:      fmt.Sprintf("%s - %s", layer, rule.GetDependency().GetDepender().ToString()),
				Classname: rule.GetDependency().GetDepender().ToString(),
			}
			switch r := rule.(type) {
			case *violations_rules.Violation:
				testSuite.Failures++
				testCase.Failures = append(testCase.Failures, struct {
					Message string `xml:"message,attr"`
					Type    string `xml:"type,attr"`
				}{
					Message: fmt.Sprintf("%s:%d must not depend on %s (%s on %s)",
						r.GetDependency().GetDepender().ToString(),
						r.GetDependency().GetContext().FileOccurrence.Line,
						r.GetDependency().GetDependent().ToString(),
						r.GetDependerLayer(),
						r.GetDependentLayer(),
					),
					Type: "WARNING",
				})
			case *violations_rules.SkippedViolation:
				testSuite.Skipped++
				testCase.Skipped = append(testCase.Skipped, struct {
					XMLName xml.Name `xml:"skipped"`
				}{})
			case *violations_rules.Uncovered:
				testCase.Warnings = append(testCase.Warnings, struct {
					Message string `xml:"message,attr"`
					Type    string `xml:"type,attr"`
				}{
					Message: fmt.Sprintf("%s:%d has uncovered dependency_contract on %s (%s)",
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

func (f *JUnitOutputFormatter) groupRulesByLayer(outputResult results.OutputResult) map[string][]violations_rules.RuleInterface {
	layers := make(map[string][]violations_rules.RuleInterface)
	for _, rule := range outputResult.AllRules() {
		switch r := rule.(type) {
		case violations_rules.CoveredRuleInterface:
			layers[r.GetDependerLayer()] = append(layers[r.GetDependerLayer()], rule)
		case *violations_rules.Uncovered:
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
	ErrorsList []struct {
		Message string `xml:"message,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"error"`
}

type TestCase struct {
	Name      string `xml:"name,attr"`
	Classname string `xml:"classname,attr"`
	Failures  []struct {
		Message string `xml:"message,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"failure"`
	Skipped []struct {
		XMLName xml.Name `xml:"skipped"`
	} `xml:"skipped"`
	Warnings []struct {
		Message string `xml:"message,attr"`
		Type    string `xml:"type,attr"`
	} `xml:"warning"`
}
