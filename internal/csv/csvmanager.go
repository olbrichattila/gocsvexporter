// Package csvmanager writes the csv file
package csvmanager

import (
	"encoding/csv"
	"os"
)

// New creates a new csv file writer
func New(fileName string, separator rune) (CsvExporter, error) {
	exporter := &csvexport{
		fileName:  fileName,
		separator: separator,
	}
	err := exporter.Open()
	if err != nil {
		return nil, err
	}

	return exporter, nil
}

// CsvExporter returns with this interface
type CsvExporter interface {
	Open() error
	Close() error
	Write([]string) error
}

type csvexport struct {
	fileName  string
	separator rune
	writer    *csv.Writer
	file      *os.File
}

func (t *csvexport) Open() error {
	file, err := os.Create(t.fileName)
	if err != nil {
		return err
	}

	t.writer = csv.NewWriter(file)
	t.writer.Comma = t.separator

	return nil
}

func (t *csvexport) Close() error {
	if t.writer != nil {
		t.writer.Flush()
	}

	if t.file != nil {
		return t.file.Close()
	}

	return nil
}

func (t *csvexport) Write(line []string) error {
	return t.writer.Write(line)
}
