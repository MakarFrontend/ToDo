package main

import "ToDo/getDB"

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
