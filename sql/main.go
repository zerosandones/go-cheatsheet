package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

type User struct {
	ID           int
	UserName     string
	EmailAddress string
	FirstName    string
	Surname      string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Capture connection properties.
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "host.containers.internal:3306"
	cfg.DBName = "go_tutorial"

	// Get a database handle.
	var dbErr error
	db, dbErr = sql.Open("mysql", cfg.FormatDSN())
	if dbErr != nil {
		log.Fatal(dbErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	users, err := getAllUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Users found: %v\n", users)

	user, err := findUserByUserName("dave")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user found: %v\n", user)

	/*user, err = findUserByUserName("joe")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("user found: %v\n", user)
	*/

	userID, err := addUser(User{
		UserName:     "tester",
		FirstName:    "Ted",
		Surname:      "Tester",
		EmailAddress: "ted@tester.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added user: %v\n", userID)
}

func getAllUsers() ([]User, error) {
	// A user slice to hold data from returned rows.
	var users []User

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("getAllUsers: %v", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var usr User
		if err := rows.Scan(&usr.ID, &usr.UserName, &usr.FirstName, &usr.Surname, &usr.EmailAddress); err != nil {
			return nil, fmt.Errorf("getAllUsers: %v", err)
		}
		users = append(users, usr)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getAllUsers: %v", err)
	}
	return users, nil
}

func findUserByUserName(name string) (User, error) {
	var usr User

	row := db.QueryRow("SELECT * FROM users WHERE UserName = ?", name)

	if err := row.Scan(&usr.ID, &usr.UserName, &usr.FirstName, &usr.Surname, &usr.EmailAddress); err != nil {
		if err == sql.ErrNoRows {
			return usr, fmt.Errorf("getUserByUsrName %d: no such user", name)
		}
		return usr, fmt.Errorf("findUserByUserName %d: %v", name, err)
	}
	return usr, nil
}

func addUser(usr User) (int64, error) {
	result, err := db.Exec("INSERT INTO users (UserName, FirstName, Surname, EmailAddress) VALUES (?, ?, ?, ?)", usr.UserName, usr.FirstName, usr.Surname, usr.EmailAddress)
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addUser: %v", err)
	}
	return id, nil
}
