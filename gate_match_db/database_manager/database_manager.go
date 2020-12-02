package database_manager

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

var Conn *gorm.DB

type DatabaseManager struct {
}

func (*DatabaseManager) ConnectDB(host string, port int, dbName, name, password string) (db *gorm.DB, err error) {
	databaseString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8,utf8mb4&parseTime=true", name, password, host, port, dbName)
	connection, err := gorm.Open("mysql", databaseString)
	if err != nil {
		return
	}
	connection.DB().SetMaxIdleConns(10)
	connection.DB().SetMaxOpenConns(100)

	Conn = connection
	db = Conn
	return
}

func (*DatabaseManager) CloseDB() error {
	return Conn.Close()
}

func (*DatabaseManager) Migrate(migrations []*gormigrate.Migration) error {
	Conn.LogMode(true)
	m := gormigrate.New(Conn, gormigrate.DefaultOptions, migrations)

	err := m.Migrate()
	if err != nil {
		if err := m.RollbackLast(); err != nil {
			return err
		}
		return err
	}
	return err
}
