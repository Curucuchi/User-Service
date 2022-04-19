package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Person struct {
	PersonID  int    `db:"personID"`
	FirstName string `db:"FirstName"`
	LastName  string `db:"LastName"`
	UserName  string `db:"userName"`
	Email     string `db:"Email"`
	Password  string `db:"password"`
}

const DSN = "root:easy2remember@tcp(localhost:3000)/users"

func Connect() {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal("There was an issue opening the DB: ", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to ping DB: ", err)
	}
}

func CreateUser(firstName, lastName, email, userName, password string) {
	ctx := context.Background()

	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal("There was an issue opening the DB: ", err)
	}

	defer db.Close()

	uniqueU, uniqueE := UniqueUserEM(strings.ToLower(userName), strings.ToLower(email))

	if uniqueU && uniqueE {
		res, err := db.ExecContext(ctx, `INSERT INTO Person VALUES (0,?,?,?,?,PASSWORD(?))`, firstName, lastName, strings.ToLower(email), strings.ToLower(userName), password)
		if err != nil {
			log.Fatal("There was an issue creating user: "+userName, err)
		}
		ID, err := res.LastInsertId()
		if err != nil {
			log.Fatal("There was an issue getting LastInsertID(): ", err)
		}

		fmt.Println("UserID:User", ID, userName+" Was created successfully")
	} else if !uniqueU {
		fmt.Println("Username already in use.")
	} else if !uniqueE {
		fmt.Println("Email already in use.")
	}

}

func DeleteUser(userName, password string) {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal("There was an issue opening the DB: ", err)
	}

	defer db.Close()

	verified := VerifyCreds(userName, password)
	if verified {
		query := "DELETE FROM Person WHERE userName = " + `"` + strings.ToLower(userName) + `"`

		res, err := db.Exec(query)
		if err != nil {
			log.Fatal("There was an issue deleting user: "+userName, err)
		}
		check, err := res.RowsAffected()
		if err != nil {
			log.Fatal("Error checking rowsAffected()", err)
		}

		if check == 0 {
			log.Fatal("No Rows were affected")
		} else {
			fmt.Println("User: " + userName + " was successfully deleted!")
		}
	} else {
		fmt.Println("Incorrect username or password!")
	}

}

func SignIn(userName, password string) {

}

// true == unique false == !unique
//this function makes sure person singing up is using a unique username and email
func UniqueUserEM(userName, email string) (U, E bool) {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal("There was an issue opening the DB: ", err)
	}

	defer db.Close()

	rows, err := db.Query(`SELECT userName,Email FROM Person`)
	if err != nil {
		log.Fatal("There was an issue querying userName: ", err)
	}

	for rows.Next() {
		var p Person
		err := rows.Scan(&p.UserName, &p.Email)
		if err != nil {
			log.Fatal("There was an issue scanning rows: ", err)
		}

		if userName == p.UserName && email != p.Email {
			return false, true
		} else if userName != p.UserName && email == p.Email {
			return true, false
		}
	}

	return true, true
}

func VerifyCreds(userName, password string) bool {
	db, err := sql.Open("mysql", DSN)
	if err != nil {
		log.Fatal("There was an issue opening the DB: ", err)
	}

	defer db.Close()

	rows, err := db.Query(`SELECT userName,password FROM Person`)
	if err != nil {
		log.Fatal("There was an issue querying userName: ", err)
	}

	for rows.Next() {
		var p Person
		err := rows.Scan(&p.UserName, &p.Password)
		if err != nil {
			log.Fatal("There was an issue scanning rows: ", err)
		}

		if userName == p.UserName && password == p.Password {
			fmt.Println(p.UserName, p.Password)
			return true
		}
	}
	return false
}
