package tablerow

type SQLPrimaryKeys map[string]interface{}

type SQLTableRow interface {
	PrimaryKeys() SQLPrimaryKeys
	Sql() string
	CreateTableSql() (string, string)
}

type Consumer interface {
	Consume(string) error
}

type SQLTableCreator interface {
	CreateTableSql() string
}
