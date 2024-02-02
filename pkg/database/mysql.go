package mysql

import (
	"context"

	"github.com/jmoiron/sqlx"
)

func (mysql *Mysql) CreateByQuery(ctx context.Context, query string, data interface{}) (int64, error) {
    namedQuery, args, err := sqlx.Named(query, data)
    if err != nil {
        mysql.log.Error(ctx, "Error creating named query", err, nil)
        return 0, err
    }

    query = mysql.db.Rebind(namedQuery)
    stmt, err := mysql.db.PreparexContext(ctx, query)
    if err != nil {
        mysql.log.Error(ctx, "Error preparing query", err, nil)
        return 0, err
    }
    defer stmt.Close()

    result, err := stmt.ExecContext(ctx, args...)
    if err != nil {
        mysql.log.Error(ctx, "Error executing query", err, nil)
        return 0, err
    }

    return result.LastInsertId()
}

func (mysql *Mysql) FindByQuery(ctx context.Context, query string, dest interface{}, args ...interface{}) error {
	err := mysql.db.SelectContext(ctx, dest, query, args...)
	if err != nil {
		mysql.log.Error(ctx, "Error retrieving records by query", err, nil)
		return err
	}

	return nil
}
