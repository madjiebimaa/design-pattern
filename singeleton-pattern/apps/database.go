package apps

import "database/sql"

type DatabaseHelper struct {
	DB *sql.DB
}

func NewDatabaseHelper() *DatabaseHelper {
	return &DatabaseHelper{}
}

func (d *DatabaseHelper) GetConnection() *sql.DB {
	if d.DB == nil {
		dsn := "username:password@/db_name"
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			panic(err)
		}

		d.DB = db
		return db
	}

	return d.DB
}
