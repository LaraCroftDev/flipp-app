package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	host     = "localhost"
	port     = 5432
	dbUser   = "postgres"
	password = "irfan"
	dbName   = "flipp"
)

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
	return db
}

func dbInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, dbUser, password, dbName)
}

// user table
func CreateNewUser(email, password, location string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	if _, err := db.Exec(fmt.Sprintf("INSERT INTO users(email, password, location) VALUES('%s', '%s', '%s');", email, password, location)); err != nil {
		return err
	}
	return nil
}

func UpdateUserPassword(id int64, newPassword string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("UPDATE users SET password = '%s' WHERE id = %d", newPassword, id))
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect != 1 {
		return errors.New("Error: recent update affected more than one column.")
	}
	return nil
}

func UpdateUserEmail(id int64, newEmail string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("UPDATE users SET email = '%s' WHERE id = %d", newEmail, id))
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect != 1 {
		return errors.New("Error: recent update affected more than one column.")
	}
	return nil
}

func UpdateUserLocation(id int64, newLocation string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("UPDATE users SET location = '%s' WHERE id = %d", newLocation, id))
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect != 1 {
		return errors.New("Error: recent update affected more than one column.")
	}
	return nil
}

// shopping list table
func CreateNewShoppingList(userId int64) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	if _, err := db.Exec(fmt.Sprintf("INSERT INTO shopping_lists(user_id) VALUES(%d);", userId)); err != nil {
		return err
	}
	return nil
}

func DeleteShoppingList(shoppingListId int64) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("DELETE FROM shopping_lists WHERE id = %d;", shoppingListId))
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect != 1 {
		return errors.New("Error: recent delete affected more than one column.")
	}
	return nil
}

// products table
func CreateNewProduct(productName, productDescription string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	if _, err := db.Exec(fmt.Sprintf("INSERT INTO products(product_name, product_description) VALUES('%s', '%s');", productName, productDescription)); err != nil {
		return err
	}
	return nil
}

func AddManyProducts(names, descriptions []string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()

	query := "INSERT INTO products(product_name, product_description) VALUES"
	for i := 0; i < len(names)-1; i++ {
		query += fmt.Sprintf("('%s', '%s'),", names[i], descriptions[i])
	}
	query += fmt.Sprintf("('%s', '%s')", names[len(names)-1], descriptions[len(descriptions)-1]) + ";"

	result, err := db.Exec(query)
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect != int64(len(names)) {
		return errors.New("Error: recent bulk insert was not fully successful.")

	}
	return nil
}

func DeleteProduct(productId int64) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec(fmt.Sprintf("DELETE FROM products WHERE id = %d", productId))
	if err != nil {
		return err
	}
	effect, _ := result.RowsAffected()
	if effect != 1 {
		return errors.New("Error: recent delete affected more than one column.")
	}
	return nil
}
