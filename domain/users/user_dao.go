package users

import (
	"fmt"
	"github.com/AhmedSafwat1/micro_service_go-api/datasources/mysql/users_db"
	"github.com/AhmedSafwat1/micro_service_go-api/utils/date_utils"
	"github.com/AhmedSafwat1/micro_service_go-api/utils/errors"
	"strings"
)

const (
	err_not_found   = "no rows in result set"
	querGetUser		= "select * from users where id = ?"
	queryInsertUser = "INSERT INTO `users`( `first_name`, `last_name`, `email`, `date_created`) VALUES (?, ?, ?, ?);"
)

var  (
	usersDB = make(map[int64]*User)
)

func (user *User) Get()  *errors.RestError  {
	x    := users_db.Connect()
	fmt.Println(users_db.Client, x)
	if err := x.Ping(); err != nil {
		panic(err)
	}

	result := x.QueryRow(querGetUser,user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		if strings.Contains(err.Error(), err_not_found) {
			return errors.NotFoundError("not dound user")
		}
		return errors.NewInteenalServerError("error in get")

	}


	return nil
}

func (user *User) Save()  *errors.RestError  {
	user.DateCreated = date_utils.GetNowString()
	x := users_db.Connect()
	if err := x.Ping(); err != nil {
		fmt.Println(err,"ind ")
		panic("bigggg")
	}
	stm , err := x.Prepare(queryInsertUser)
	if err != nil {
		errors.NewInteenalServerError("Error in perpare")
	  }
	fmt.Println("hi", user)
	defer  stm.Close()
	inserResult , err := stm.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		errors.NewInteenalServerError("Error in insert")
	}
	userId , err :=  inserResult.LastInsertId()
	if err != nil {
		errors.NewInteenalServerError("Error in get lat id")
	}
	user.Id = userId
	fmt.Println(userId)
	//if usersDB[user.Id] != nil {
	//	return errors.NewBadRequest(fmt.Sprintf("user %d alerad exis", user.Id))
	//}
	//
	////usersDB[user.Id] = user
	return  nil;
}
