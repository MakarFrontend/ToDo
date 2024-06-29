package structs

import (
	"ToDo/getDB"
	"log"
)

/*Задача пользователя*/
type MyToDo struct {
	Text   string `json:"txt"`    //Текст задачи
	Status bool   `json:"status"` //Статус
	Tag    string `json:"tag"`    //Тег
	Id     int    //ID задачи
}

/*Добавляет задачу в базу данных, на вход принимает логин пользователя*/
func (m MyToDo) Insert(login string) error {
	db, err := getDB.GetDB()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	/*Для добавления задачи*/
	const req string = `INSERT INTO ToDo (tag_id, status, txt, user_id) 
VALUES((SELECT tag_id FROM Tags WHERE tag = $1),
false,
$2,
(SELECT user_id FROM Users WHERE login = $3));`

	res, err := db.Exec(req, m.Tag, m.Text, login)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	rowsAff, err := res.RowsAffected()
	if err != nil && rowsAff == 0 {
		log.Println(err.Error())
		return err
	}

	return nil
}
