package main

//go run main.go get.go post.go delete.go sql.go other.go
//go build main.go get.go post.go delete.go sql.go other.go

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func middleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		log.Printf("Request URL: %s; Token: %v", r.URL, token)
		if token == "" || !strings.Contains(token, ":") {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Отказано в доступе. 403"))
			return
		}
		if userHaveToken(token) {
			h.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Отказано в доступе. 403"))
			return
		}
	})
}

func main() {
	fileLog, err := os.OpenFile("app.log", os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer fileLog.Close()

	log.SetOutput(fileLog)

	mainMux := http.NewServeMux()
	APImux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./templates/static/"))
	mainMux.Handle("/JS_CSS/", http.StripPrefix("/JS_CSS", fileServer))

	/*Регистрация маршрутов для с токенами*/
	APImux.HandleFunc("POST /user/{login}/new/task", postNewTask)         //Новая задача
	APImux.HandleFunc("POST /user/{login}/{id}/status", postToggleStatus) //Обновление статуса задачи

	APImux.HandleFunc("DELETE /task/{id}/del", deleteTask)    //Удаление задачи
	APImux.HandleFunc("DELETE /user/{login}/del", deleteUser) //Удаление пользователя

	/*Регистрация APImux по всем маршрутам /api/XXXX*/
	mainMux.Handle("/api/", http.StripPrefix("/api", middleWare(APImux)))

	/*регистрация маршрутов, где нет токена*/
	/*В заголовке token в ответе идёт токен*/
	mainMux.HandleFunc("POST /new/user", postNewUser) //Новый пользователь Query: ?password=123456&login=qwerty
	mainMux.HandleFunc("GET /content", getContent)
	mainMux.HandleFunc("GET /user/{login}", getUserLogin) //Задачи пользователя Query: ?password=123456

	mainMux.HandleFunc("/", index)

	fmt.Println("Listening 5000 . . .")
	http.ListenAndServe("localhost:5000", mainMux)
}
