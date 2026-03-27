package db

import (

)

type User struct{
	login string
	passwd string
	uName string
}

var dbUsers = []User{}
/*
func InitDb() []User {
	return users
}
*/
func AddUser(l string, pwd string, un string) string {
	if _, b := checkLogin(l); !b {
		dbUsers = append(dbUsers, User{l,pwd, un})
		return "registry"
	}else{
		return "login is exist"
	}
}

func CheckUser(l string, pwd string) bool {
	if i, b := checkLogin(l); b{
		if dbUser[i].passwd == pwd {
			return true
		}
		return false
	}
	return false
}

func checkLogin(l string) int, bool{
	for i, user := range dbUsers {
		if user.login == l {
			return i, true
		}
	}
	return -1, false
}
