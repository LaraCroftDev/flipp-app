package main

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"

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
	if _, err := db.Exec("INSERT INTO users(email, password, location) VALUES($1,$2,$3);", email, password, location); err != nil {
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
	result, err := db.Exec("UPDATE users SET password=$1 WHERE id=$2", newPassword, id)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("%v rows were updated", effect)
	return nil
}

func UpdateUserEmail(id int64, newEmail string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec("UPDATE users SET email=$1 WHERE id=$2", newEmail, id)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("%v rows were updated", effect)
	return nil
}

func UpdateUserLocation(id int64, newLocation string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec("UPDATE users SET location=$1 WHERE id=$2", newLocation, id)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("%v rows were updated", effect)
	return nil
}

// shopping list table
func CreateNewShoppingList(userId int64) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	if _, err := db.Exec("INSERT INTO shopping_lists(user_id) VALUES($1);", userId); err != nil {
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
	result, err := db.Exec("DELETE FROM shopping_lists WHERE id=$1;", shoppingListId)
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
	if _, err := db.Exec("INSERT INTO products(product_name, product_description) VALUES($1,$2);", productName, productDescription); err != nil {
		return err
	}
	return nil
}

func DeleteProduct(productId int64) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()
	result, err := db.Exec("DELETE FROM products WHERE id=$1", productId)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("%v rows were updated", effect)
	return nil
}

func AddManyProducts(data []map[string]string) error {
	db := dbConnect()
	if db == nil {
		return errors.New("Error: db connection error.")
	}
	defer db.Close()

	query := "INSERT INTO products(product_name, product_description) VALUES"
	values := []interface{}{}
	n := 0
	for _, row := range data {
		query += "($" + strconv.Itoa(n+1) + ",$" + strconv.Itoa(n+2) + "),"
		values = append(values, row["name"], row["description"])
		n += 2
	}
	query = query[0:len(query)-1] + ";"
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return err
	}
	effect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Printf("%v rows were updated", effect)
	return nil
}
