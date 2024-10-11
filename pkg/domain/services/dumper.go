package services

import (
	"errors"
	"os"
)

type Dumper struct {
	templateFile string
}

func NewDumper(templateFile string) *Dumper {
	return &Dumper{
		templateFile: templateFile,
	}
}

func (d *Dumper) Dump(file string) error {
	if _, err := os.Stat(file); !errors.Is(err, os.ErrNotExist) {
		return errors.New("file_supportive already exists")
	}

	templateFileContent, err := os.ReadFile(d.templateFile)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, templateFileContent, 0666)
	if err != nil {
		return err
	}

	return nil
}
