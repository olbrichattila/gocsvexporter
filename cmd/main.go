// Package main is the main entry point
package main

import (
	"csvdownloader/internal/cargs"
	csvmanager "csvdownloader/internal/csv"
	"csvdownloader/internal/datasource"
	"csvdownloader/internal/environment"
	"csvdownloader/internal/exporter"
	"fmt"
)

func main() {
	cargs := cargs.New()
	tableName, outuptFileName, separator, err := cargs.Args()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Processing...")
	csv, err := csvmanager.New(outuptFileName, separator)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csv.Close()
	env, err := environment.New()
	if err != nil {
		fmt.Println(err)
		return
	}

	connectinString, driverName, quoteSign, err := datasource.GetDbConnectionParams(env)
	if err != nil {
		fmt.Println(err)
		return
	}

	ds := datasource.New(tableName, connectinString, driverName, quoteSign)
	err = ds.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ds.Close()

	fmt.Printf("Found %d rows\n", ds.RowCount())

	exporter := exporter.New(ds, csv)
	err = exporter.Export()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Done")
}
