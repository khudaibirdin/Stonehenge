package main

import (
	"fmt"

	db "github.com/khudaibirdin/GoLangModules/database_actions"
)

func VerifyUser(login string, password string) (bool, User) {
	user := User{}
	db := db.Db{Path: "./database.db"}
	condition := fmt.Sprintf(`login = "%s"`, login)
	db.GetRowByCondition("user", "*", condition, &user)
	return user.Password == password, user
}
