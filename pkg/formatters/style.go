package formatters

import (
	"fmt"
	"github.com/KoNekoD/go-deptrac/pkg/results"
	"github.com/gookit/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/schollz/progressbar/v3"
	"strings"
)

type Style struct {
	progressbar *progressbar.ProgressBar
	isVerbose   bool
	isDebug     bool
}

func NewStyle(isVerbose bool, isDebug bool) *Style {
	return &Style{
		isVerbose: isVerbose,
		isDebug:   isDebug,
	}
}

func (s *Style) Title(message string) {
	color.Printf("<comment>%s</>", message)
	color.Printf("<comment>%s</>\n", strings.Repeat("=", len(message)))
}

func (s *Style) Section(message string) {
	color.Printf("<comment>%s</>", message)
	color.Printf("<comment>%s</>\n", strings.Repeat("-", len(message)))
}

func (s *Style) Success(message results.StringOrArrayOfStrings) {
	color.Printf("<success>%s</>", message.ToString())
}

func (s *Style) Error(message results.StringOrArrayOfStrings) {
	color.Printf("<error>%s</>", message.ToString())
}

func (s *Style) Warning(message results.StringOrArrayOfStrings) {
	color.Printf("<warning>%s</>", message.ToString())
}

func (s *Style) Note(message results.StringOrArrayOfStrings) {
	color.Printf("<note>%s</>", message.ToString())
}

func (s *Style) Caution(message results.StringOrArrayOfStrings) {
	color.Printf("<danger>%s</>", message.ToString())
}

func (s *Style) DefinitionList(list []results.StringOrArrayOfStringsOrTableSeparator) {
	headers := make([]string, 0)
	row := make([]string, 0)
	for _, value := range list {
		if value.TableSeparator {
			headers = append(headers, "")
			row = append(row, "<TS>")
			continue
		}

		if value.String != "" {
			headers = append(headers, value.String)
			row = append(row, "")
			continue
		}

		if value.StringsMap == nil {
			panic("Value should be an array, string, or an instance of TableSeparator.")
		}

		for stringsMapKey, stringsMapValue := range value.StringsMap {
			headers = append(headers, stringsMapKey)
			row = append(row, color.Sprintf(stringsMapValue))
		}

	}
	s.Table(headers, [][]string{row})
}

func (s *Style) Table(headers []string, rows [][]string) {
	tw := table.NewWriter()

	for _, rowsRow := range rows {
		for j, row := range rowsRow {
			if row == "<TS>" {
				tw.AppendSeparator()
				continue
			}

			tw.AppendRow(table.Row{headers[j], row})
		}
	}

	style := table.Style{
		Name: "deptrac-original-style",
		Box: table.BoxStyle{
			BottomLeft:       "-",
			BottomRight:      "-",
			BottomSeparator:  " ",
			EmptySeparator:   "",
			Left:             " ",
			LeftSeparator:    "-",
			MiddleHorizontal: "-",
			MiddleSeparator:  " ",
			MiddleVertical:   " ",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			PageSeparator:    "22",
			Right:            " ",
			RightSeparator:   "-",
			TopLeft:          "-",
			TopRight:         "-",
			TopSeparator:     " ",
			UnfinishedRow:    "",
		},
		Color: table.ColorOptions{
			Border: text.Colors{
				text.Reset,
			},
			Footer:      nil,
			Header:      nil,
			IndexColumn: nil,
			Row: text.Colors{
				text.FgGreen,
			},
			RowAlternate: nil,
			Separator: text.Colors{
				text.Reset,
			},
		},
		Format: table.FormatOptions{},
		HTML:   table.HTMLOptions{},
		Options: table.Options{
			DrawBorder:      true,
			SeparateColumns: true,
			SeparateHeader:  true,
		},
		Title: table.TitleOptions{},
	}
	tw.SetStyle(style)
	fmt.Println(tw.Render())
}

func (s *Style) NewLine(count int) {
	for i := 0; i < count; i++ {
		fmt.Println()
	}
}

func (s *Style) ProgressStart(max int) {
	s.progressbar = progressbar.Default(int64(max))
}

func (s *Style) ProgressAdvance(step int) error {
	err := s.progressbar.Add(step)
	if err != nil {
		return err
	}

	return nil
}

func (s *Style) ProgressFinish() error {
	err := s.progressbar.Finish()
	if err != nil {
		return nil // TODO
	}

	return nil
}

func (s *Style) IsVerbose() bool {
	return s.isVerbose
}

func (s *Style) IsDebug() bool {
	return s.isDebug
}
