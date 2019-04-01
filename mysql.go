package config

import (
	"context"
	"database/sql"

	"github.com/angadn/tabular"
)

// FromMySQL gives us a Source for MySQL.
func FromMySQL(
	ctx context.Context, db *sql.DB,
) (sourceImpl SourceImpl, err error) {
	if err = db.PingContext(ctx); err != nil {
		return
	}

	sourceImpl = MySQLSource{db}
	return
}

// MySQLSource is an implementation of SourceImpl using MySQL.
type MySQLSource struct {
	db *sql.DB
}

// Get a key from the database.
func (src MySQLSource) Get(ctx context.Context, key string) (value string, err error) {
	err = src.db.QueryRowContext(ctx, table.Selection(
		"SELECT %s FROM `configurations` WHERE `name` = ?",
	), key).Scan(&tabular.Scapegoat{}, &value)

	return
}

// Set a key in the database.
func (src MySQLSource) Set(ctx context.Context, key string, value string) (err error) {
	_, err = src.db.ExecContext(
		ctx,
		table.Insertion("%s ON DUPLICATE KEY UPDATE `value` = VALUES(`value`)"),
		key,
		value,
	)

	return
}

// table is a tabular representation of Configurations, and helps us persist them in an
// SQL database.
var table = tabular.New(
	"configurations",

	"name",
	"value",
)
