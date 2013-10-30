package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getDB(username, userpwd, dbname string) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s?charset=utf8", username, userpwd, dbname)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Println(err.Error()) //仅仅是显示异常
		return nil, err
	}
	return db, nil
}
