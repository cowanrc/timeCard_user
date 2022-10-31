package employees_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_employees_username = "mysql_employees_username"
	mysql_employees_password = "mysql_employees_password"
	mysql_employees_host     = "mysql_employees_host"
	mysql_employees_schema   = "mysql_employees_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_employees_username)
	password = os.Getenv(mysql_employees_password)
	host     = os.Getenv(mysql_employees_host)
	schema   = os.Getenv(mysql_employees_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, password, host, schema,
	)

	log.Println(username, password, host, schema)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database succussfully configured")

}
