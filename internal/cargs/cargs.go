// Package cargs parses command line arguments
package cargs

import (
	"fmt"
	"os"
)

// New creates a new argument parser
func New() Arger {
	return &arg{}
}

// Arger is the interface
type Arger interface {
	Args() (string, string, rune, error)
}

type arg struct {
}

// Args returns table name, target file name, separator or error
func (*arg) Args() (string, string, rune, error) {
	separator := ','
	if len(os.Args) < 3 {
		return "", "", ',', fmt.Errorf("need least 2 command line arguments, <tableName> <outputFileName> <separator> where separator is optional and default to comma")
	}

	if len(os.Args) > 3 {
		separator = []rune(os.Args[3])[0]
	}

	return os.Args[1], os.Args[2], separator, nil
}
