package main

import (
	"ToDo/getDB"
	"ToDo/structs"
	"html/template"
	"log"
	"net/http"
)

/*Для ответа о задачах пользователя*/
type taskResponse struct {
	Ready    []structs.MyToDo //Выполненые задачи
	NotReady []structs.MyToDo //Не выполненые задачи
	Tags     []string         //Теги
}

/* / */
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		log.Printf("Don't have page: %v; %s", r.Method, r.URL)
		w.WriteHeader(http.StatusNotFound)
		http.ServeFile(w, r, "templates/404.html")
		return
	}
	http.ServeFile(w, r, "templates/index.html")
}

/* /content */
func getContent(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/content.html")
}

/* GET /user/{login} Query: ?password=123456 */
func getUserLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	have := haveUser(r.PathValue("login"), r.FormValue("password")) //Провнряем есть ли такой пользователь
	if !have {
		log.Printf("Don't have user %v; %v", r.PathValue("login"), r.FormValue("password"))
		w.WriteHeader(404)
		w.Write([]byte("Неправильный логин или пароль"))
		return
	}

	db, err := getDB.GetDB() //Подключаемся к БД
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}
	defer db.Close()

	rows, err := db.Query(myTasks, r.PathValue("login")) //Берём все задачи пользователя
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}

	/*Структура для ответа*/
	data := taskResponse{}

	for rows.Next() { //Цикл проходит по задачам
		row := structs.MyToDo{}
		err = rows.Scan(&row.Text, &row.Status, &row.Tag, &row.Id)
		if err != nil {
			log.Println(err)
			continue
		}
		if row.Status { //Если задача выполнена, то её в data.Ready, иначе в data.NotReady
			data.Ready = append(data.Ready, row)
		} else {
			data.NotReady = append(data.NotReady, row)
		}
	}

	rows, err = db.Query(SQLallTags) //Из БД все теги
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			log.Println(err)
			continue
		}
		data.Tags = append(data.Tags, tag)
	}

	var token string
	tokenRow := db.QueryRow(loginToken, r.PathValue("login"))
	if err = tokenRow.Scan(&token); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}

	tmp, err := template.ParseFiles("templates/user.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		w.Write([]byte("Ошибка на сервере, извините"))
		return
	}

	w.Header().Add("token", token)
	tmp.Execute(w, data)
}
