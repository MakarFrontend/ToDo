package main

import (
	"ToDo/getDB"
	"log"
	"strings"
)

/*Есть ли такой пользователь*/
func haveUser(login, password string) bool {
	db, err := getDB.GetDB()
	if err != nil {
		return false
	}
	defer db.Close()

	row := db.QueryRow(haveSQL, password, login)
	var id int
	err = row.Scan((&id))
	return err == nil
}

/*Есть ли такой Ник*/
func haveNick(login string) bool {
	db, err := getDB.GetDB()
	if err != nil {
		return false
	}
	defer db.Close()

	row := db.QueryRow(haveNickname, login)
	var id int
	err = row.Scan((&id))
	return err == nil
}

/*Если токен соответствует пользователю*/
func userHaveToken(token string) bool {
	have := false

	tokenSlice := strings.Split(token, ":")
	if len(tokenSlice) != 2 {
		return have
	}

	db, err := getDB.GetDB()
	if err != nil {
		log.Println(err)
		return have
	}
	defer db.Close()

	row := db.QueryRow("SELECT token FROM Users WHERE login = $1", tokenSlice[0])

	var tokenDB string
	err = row.Scan(&tokenDB)
	if err != nil || tokenDB != tokenSlice[1] {
		log.Println(err)
		return have
	}

	have = true
	return have
}
