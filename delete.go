package main

import (
	"ToDo/getDB"
	"fmt"
	"log"
	"net/http"
)

/*Удаление профиля DELETE /user/{login}/del*/
func deleteUser(w http.ResponseWriter, r *http.Request) {
	login := r.PathValue("login")

	db, err := getDB.GetDB()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Ошибка на сервере, извините")
		return
	}
	defer db.Close()

	rows, err := db.Exec(deleteUserSQL, login)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Ошибка на сервере, извините")
		return
	}

	rowsAff, err := rows.RowsAffected()
	if rowsAff == 0 || err != nil {
		log.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Ошибка на сервере, извините")
		return
	}

	fmt.Fprint(w, "Профиль успешно удалённ!")
}

/*DELETE /task/{id}/del Удаление задачи*/
func deleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	db, err := getDB.GetDB()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Ошибка на сервере, извините")
		return
	}
	defer db.Close()

	rows, err := db.Exec(deleteTaskSQL, id)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Ошибка на сервере, извините")
		return
	}

	rowsAff, err := rows.RowsAffected()
	if rowsAff == 0 || err != nil {
		log.Println(err)
		w.WriteHeader(500)
		fmt.Fprint(w, "Ошибка на сервере, извините")
		return
	}

	fmt.Fprint(w, "Задача успешно удалена!")
}
