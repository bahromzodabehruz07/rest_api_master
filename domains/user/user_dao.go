package user

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rest_api_master/domains/error"
	"net/http"
	"strconv"
)

func CreateConnection() *sql.DB {

	connStr := "user=postgres password=2404 dbname=test_db sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	// Open the connection

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// return the connection
	return db
}

func (user *User) SaveUser() *error.Error {
	insertStatement := "Insert into users(firstname,lastname,email) values($1,$2,$3) RETURNING id"
	var db = CreateConnection()
	// the inserted id will store in this id
	var id int64
	err := db.QueryRow(insertStatement, user.FirstName, user.LastName, user.Email).Scan(&id)
	if err != nil {
		return &error.Error{
			Status:  http.StatusMethodNotAllowed,
			Message: "Error while inserting new user",
			Error:   err.Error(),
		}
	}
	user.Id = id
	return nil
}

func (user *User) GetUserByID() (*User, *error.Error) {
	var db = CreateConnection()
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, user.Id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.UserToken)

	if row != nil {
		return nil, &error.Error{
			Message: "Error while query is run",
			Status:  http.StatusNotFound,
			Error:   "bad date",
		}
	}
	return user, nil

}

func (user *User) GetUser() *error.Error {
	var db = CreateConnection()
	fmt.Println(user.Id)
	sqlStatement := `SELECT * FROM users WHERE id=$1`

	// execute the sql statement
	row := db.QueryRow(sqlStatement, user.Id)

	// unmarshal the row object to user
	err := row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.UserToken)

	if err != nil {
		return &error.Error{
			Message: "User with this id" + " " + strconv.Itoa(int(user.Id)),
			Status:  http.StatusNotFound,
			Error:   "not found",
		}
	}
	return nil
}

func (user *User) Delete() *error.Error {
	var db = CreateConnection()
	sqlStatement := `DELETE FROM users WHERE id=$1`
	// execute the sql statement
	row, err := db.Exec(sqlStatement, user.Id)

	if err != nil {
		panic(err)
		return &error.Error{
			Message: "User with this id" + " " + strconv.Itoa(int(user.Id)),
			Status:  http.StatusNotFound,
			Error:   "not found",
		}
	}
	fmt.Println(row.RowsAffected())
	return nil
}

func (user *User) UpdateUser() *error.Error {
	var db = CreateConnection()
	fmt.Println(user)
	sqlStatement := `UPDATE users SET firstname=$1, lastname=$2, email=$3,userToken=$4 WHERE id=$5`

	_, err := db.Exec(sqlStatement, &user.FirstName, user.LastName, user.Email, user.UserToken, user.Id)
	if err != nil {
		panic(err)
	}
	return nil
}
