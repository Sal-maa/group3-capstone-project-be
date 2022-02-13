package util

import (
	_config "capstone/be/config"
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func GetDBInstance(config *_config.AppConfig) (*sql.DB, error) {
	if db == nil {
		driverName := config.Database.Driver

		dataSourceName := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			config.Database.Username,
			config.Database.Password,
			config.Database.Host,
			config.Database.Port,
			config.Database.Name,
		)

		dbNewInstance, err := sql.Open(driverName, dataSourceName)

		if err != nil {
			return nil, err
		}

		db = dbNewInstance
	}

	return db, nil
}
