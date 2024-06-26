// Package datasource is responsible to get the data from database and serve it
package datasource

import (
	"database/sql"
	"fmt"
)

// DBer is the database manager
type DBer interface {
	Open() error
	Close() error
	RowCount() int
	Prepare() error
	Next() bool
	Row() []string
	GetFieldNames() []string
	GetLastError() error
}

type datasoure struct {
	tableName       string
	connectinString string
	driverName      string
	quoteSign       string
	rowCount        int
	db              *sql.DB
	stmt            *sql.Stmt
	rows            *sql.Rows
	fieldCount      int
	fieldNames      []string
	fields          []interface{}
	row             []string
	lastError       error
}

// New creates a new database manager
func New(
	tableName string,
	connectinString string,
	driverName string,
	quoteSign string,
) DBer {
	return &datasoure{
		tableName:       tableName,
		connectinString: connectinString,
		driverName:      driverName,
		quoteSign:       quoteSign,
	}
}

func (t *datasoure) Open() error {
	var err error
	// TODO this dependency should be inversed
	// connectinString, driverName, quoteSign, err := getDbConnectionParams(environment.New())
	// if err != nil {
	// 	return err
	// }
	/// t.quoteSign = quoteSign

	t.db, err = sql.Open(t.driverName, t.connectinString)
	if err != nil {
		return err
	}

	t.rowCount, err = t.count()
	return err
}

func (t *datasoure) Prepare() error {
	var err error
	t.lastError = nil
	sql := fmt.Sprintf("SELECT * FROM %s%s%s", t.quoteSign, t.tableName, t.quoteSign)
	t.stmt, err = t.db.Prepare(sql)
	if err != nil {
		return err
	}

	t.rows, err = t.stmt.Query()
	if err != nil {
		return err
	}

	return t.initHeader()
}

func (t *datasoure) Close() error {
	stmtError := t.stmt.Close()
	closeError := t.db.Close()

	if stmtError != nil && closeError != nil {
		return fmt.Errorf("%s, %s", stmtError.Error(), closeError.Error())
	}

	if stmtError != nil {
		return stmtError
	}

	if closeError != nil {
		return closeError
	}

	return nil
}

func (t *datasoure) Next() bool {
	hasNextRow := t.rows.Next()
	if !hasNextRow {
		return false
	}

	valuePtrs := make([]interface{}, t.fieldCount)
	for i := range valuePtrs {
		valuePtrs[i] = &t.fields[i]
	}

	err := t.rows.Scan(valuePtrs...)
	if err != nil {
		t.lastError = err
		return false
	}

	return true
}

func (t *datasoure) Row() []string {

	for i, field := range t.fields {
		if field == nil {
			t.row[i] = ""
			continue
		}

		t.row[i] = fmt.Sprintf("%v", field)
	}

	return t.row
}

func (t *datasoure) GetFieldNames() []string {
	return t.fieldNames
}

func (t *datasoure) GetLastError() error {
	return t.lastError
}

func (t *datasoure) initHeader() error {
	var err error
	t.fieldNames, err = t.rows.Columns()
	if err != nil {
		return err
	}

	t.fieldCount = len(t.fieldNames)
	t.row = make([]string, t.fieldCount)
	t.fields = make([]interface{}, t.fieldCount)

	return nil
}

func (t *datasoure) count() (int, error) {
	sql := fmt.Sprintf("SELECT count(1) as res FROM %s%s%s", t.quoteSign, t.tableName, t.quoteSign)
	stmt, err := t.db.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var cnt int

	row := stmt.QueryRow()
	err = row.Scan(&cnt)
	if err != nil {
		return 0, err
	}

	return cnt, nil
}

func (t *datasoure) RowCount() int {
	return t.rowCount
}
