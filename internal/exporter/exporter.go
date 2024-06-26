// Package exporter exports from datasource to destination CSV
package exporter

import (
	csvexport "csvdownloader/internal/csv"
	"csvdownloader/internal/datasource"
	"fmt"
	"strings"
)

// New vreate a database to csv exporter
func New(db datasource.DBer, csv csvexport.CsvExporter) Exporter {
	return &exp{
		db:  db,
		csv: csv,
	}
}

// Exporter is the interface retuned by new exporter
type Exporter interface {
	Export() error
}

type exp struct {
	db       datasource.DBer
	csv      csvexport.CsvExporter
	rowNr    int
	progress int
	rowCount int
}

func (t *exp) Export() error {
	t.rowCount = t.db.RowCount()

	err := t.db.Prepare()
	if err != nil {
		return err
	}

	fieldNames := t.db.GetFieldNames()
	fmt.Printf("Fields\n - ")
	fmt.Println(strings.Join(fieldNames, ", "))

	if err := t.csv.Write(fieldNames); err != nil {
		return err
	}

	for t.db.Next() {
		t.displayProgress()

		row := t.db.Row()
		if err := t.csv.Write(row); err != nil {
			return err
		}
	}

	if t.db.GetLastError() != nil {
		return t.db.GetLastError()
	}

	return nil
}

func (t *exp) displayProgress() {
	t.rowNr++
	newProgress := int(float64(t.rowNr) / float64(t.rowCount) * 100)
	if newProgress != t.progress {
		t.progress = newProgress
		fmt.Printf("Progress %d%%\r", t.progress)
	}
}
