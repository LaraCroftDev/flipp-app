package main

import (
	"database/sql"
	"errors"
	"fmt"
)

const (
	dbDriver = "postgres"
	host     = "localhost"
	port     = 5432
	dbUser   = "postgres"
	password = "irfan"
	dbName   = "" // update later
)

func CreateNewUser(user, password, location string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: new user creation failed.")
	}
	defer db.Close()
	if _, err := db.Exec(""); err != nil {
		return err
	}
	return nil
}

func dbConnect() *sql.DB {
	db, err := sql.Open(dbDriver, dbInfo())
	if err != nil {
		fmt.Println("Error db connection: ", err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("Error db connection interruption: ", err)
		return nil
	}
	fmt.Println("Successful db connection.")
	return db
}

func dbInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, dbUser, password, dbName)
}
