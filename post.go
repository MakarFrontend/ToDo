package main

import (
	"ToDo/getDB"
	"ToDo/structs"
	"encoding/json"
	"io"
	"net/http"
)

/*POST /new/user ?password=123456&login=qwerty Создание пользователя*/
func postNewUser(w http.ResponseWriter, r *http.Request) {
	log := r.URL.Query().Get("login")
	pass := r.URL.Query().Get("password")

	have := haveNick(log)
	if have {
		w.WriteHeader(403)
		w.Write([]byte("Такой никнейм уже занят"))
		return
	} else {
		db, err := getDB.GetDB()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка на сервере, извините"))
			return
		}
		defer db.Close()

		res, err := db.Exec(appendUser, log, pass)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка на сервере, извините"))
			return
		}

		rowsAff, err := res.RowsAffected() //Если количество добавленных строк равно нулю, значит произошла ошибка
		if err != nil || rowsAff == 0 {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка на сервере, извините"))
			return
		}
	}

	w.Write([]byte("Вы успешно зарегистрированы!"))
}

/*Отправка новой задачи, данные приходят в формате JSON в виде структуры structs.MyToDo*/
func postNewTask(w http.ResponseWriter, r *http.Request) {
	rByteData, err := io.ReadAll(r.Body)
	if err != nil { //Не смогли прочитать тело запроса
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}

	/*Структура для приёма ответа*/
	var newToDo structs.MyToDo
	err = json.Unmarshal(rByteData, &newToDo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}

	//Провека на наличие тега и текста задачи
	if !newToDo.Status && newToDo.Tag != "" && newToDo.Text != "" {
		err := newToDo.Insert(r.PathValue("login"))
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка на сервере, извините"))
			return
		}

		w.Write([]byte("Успех"))
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Ошибка на сервере, извините"))
}

/*POST /user/{login}/{id}/status*/
func postToggleStatus(w http.ResponseWriter, r *http.Request) {
	db, err := getDB.GetDB()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}
	defer db.Close()

	login := r.PathValue("login")
	todoId := r.PathValue("id")

	res, err := db.Exec(toggleTaskStatus, login, todoId)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}

	rowsAff, err := res.RowsAffected()
	if err != nil || rowsAff == 0 {
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}

	w.Write([]byte("Успех"))
}
