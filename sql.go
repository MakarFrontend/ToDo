package main

/*----------------------------------Запросы на выборку---------------------------------------------*/

/*Все задачи пользователя*/
const myTasks string = `SELECT txt, status, tag, todo_id
FROM
	ToDo
	JOIN Tags
	ON ToDo.tag_id = Tags.tag_id
	JOIN Users
	ON Users.user_id = ToDo.user_id
WHERE login = $1
ORDER BY todo_id;`

/*Взятие токена*/
const loginToken string = `SELECT CONCAT(login, ':', token)
FROM Users
WHERE login = $1`

/*Все теги*/
const SQLallTags string = `SELECT tag
FROM Tags;`

/*Есть ли такой пользователь*/
const haveSQL string = `SELECT user_id
FROM Users
WHERE password = $1 AND login = $2`

/*Есть ли такой никнейм*/
const haveNickname string = `SELECT user_id
FROM Users
WHERE login = $1`

/*Запросы на обновление данных*/

/*Изменение статуса задачи*/
const toggleTaskStatus string = `UPDATE ToDo
SET status = NOT status
WHERE user_id IN (
	SELECT user_id
	FROM Users
	WHERE login = $1
) AND todo_id = $2;`

/*----------------------Запросы на добавление данных------------------------------*/

/*Добавление пользователя*/
const appendUser string = `INSERT INTO Users(login, password, token) VALUES($1, $2, $3)`

/*-----------------------------Запросы на удаление--------------------------------*/

/*Удаление пользователя*/
const deleteUserSQL string = `DELETE FROM Users
WHERE login = $1`

/*Удаление задачи*/
const deleteTaskSQL string = `DELETE FROM ToDo
WHERE todo_id = $1;`
