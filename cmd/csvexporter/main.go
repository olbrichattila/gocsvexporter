// Package main is the main entry point
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/olbrichattila/gocsvexporter/internal/cargs"
	csvmanager "github.com/olbrichattila/gocsvexporter/internal/csv"
	"github.com/olbrichattila/gocsvexporter/internal/datasource"
	"github.com/olbrichattila/gocsvexporter/internal/environment"
	"github.com/olbrichattila/gocsvexporter/internal/exporter"
)

func main() {
	help := flag.Bool("help", false, "Display help")
	flag.Parse()

	if *help {
		displayHelp()
		os.Exit(0)
	}

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

	connectionString, driverName, quoteSign, err := datasource.GetDbConnectionParams(env)
	if err != nil {
		fmt.Println(err)
		return
	}

	ds := datasource.New(tableName, connectionString, driverName, quoteSign)
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

	fmt.Printf("\nDone\n")
}
