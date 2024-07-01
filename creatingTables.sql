/*В файле хранятся SQL запросы, которыми создавались таблицы*/

-- Таблица с юзерами
CREATE TABLE Users (
	user_id SERIAL PRIMARY KEY,
	password TEXT,
	login TEXT
);

-- Таблица с тегами
CREATE TABLE Tags (
	tag_id SERIAL PRIMARY KEY,
	tag TEXT
);

-- Таблица с задачами
CREATE TABLE ToDo (
	todo_id SERIAL PRIMARY KEY,
	user_id INT,
	tag_id INT,
	status BOOL,
	txt TEXT,
	FOREIGN KEY (user_id) REFERENCES Users (user_id) ON DELETE CASCADE,
	FOREIGN KEY (tag_id) REFERENCES Tags (tag_id) ON DELETE SET NULL
);

/*В таблицу с пользователями похже был добавлен столбец token*/
ALTER TABLE Users
ADD token text NOT NULL DEFAULT 'login:333';
