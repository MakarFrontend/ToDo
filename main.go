package main

//go run main.go get.go post.go delete.go sql.go other.go
//go build main.go get.go post.go delete.go sql.go other.go

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fileLog, err := os.OpenFile("app.log", os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fileLog.Close()

	log.SetOutput(fileLog)

	http.HandleFunc("GET /user/{login}", getUserLogin) //Задачи пользователя Query: ?password=123456
	http.HandleFunc("GET /content", getContent)

	http.HandleFunc("POST /new/user", postNewUser)                      //Новый пользователь Query: ?password=123456&login=qwerty
	http.HandleFunc("POST /user/{login}/new/task", postNewTask)         //Новая задача
	http.HandleFunc("POST /user/{login}/{id}/status", postToggleStatus) //Обновление статуса задачи

	http.HandleFunc("DELETE /task/{id}/del", deleteTask)    //Удаление задачи
	http.HandleFunc("DELETE /user/{login}/del", deleteUser) //Удаление пользователя

	fileServer := http.FileServer(http.Dir("./templates/static/"))
	http.Handle("/JS_CSS/", http.StripPrefix("/JS_CSS", fileServer))

	http.HandleFunc("/", index)

	fmt.Println("Listening 5000 . . .")
	http.ListenAndServe("localhost:5000", nil)
}
