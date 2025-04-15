package utils

import "github.com/jmoiron/sqlx"

// Translate a named query with bind vars as :namedVar to $1, $2... syntax
func TranslateNamedQuery(namedSql string, parameters any) (query string, args []any, err error) {
	sql, args, err := sqlx.Named(namedSql, parameters)
	if err != nil {
		return "", nil, err
	}

	sql, args, err = sqlx.In(sql, args...)
	if err != nil {
		return "", nil, err
	}

	sql = sqlx.Rebind(sqlx.DOLLAR, sql)
	return sql, args, nil
}
