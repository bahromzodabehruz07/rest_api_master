package service

import (
	"github.com/rest_api_master/domains/error"
	"github.com/rest_api_master/domains/user"
	"net/http"
)

func GetAllUsers() ([]user.User, *error.Error) {
	db := user.CreateConnection()

	// close the db connection
	defer db.Close()

	var users []user.User

	// create the select sql query
	sqlStatement := `SELECT * FROM users`

	row, err := db.Query(sqlStatement)
	if err != nil {
		return nil, &error.Error{
			Status:  http.StatusNotFound,
			Message: err.Error(),
			Error:   "Error while query is run",
		}

	}
	for row.Next() {
		var user user.User
		err = row.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.UserToken)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return users, nil

}

func GetUserById(user user.User) (*user.User, *error.Error) {
	result, err := user.GetUserByID()
	if err != nil {
		panic(err)
	}
	return result, nil
}
