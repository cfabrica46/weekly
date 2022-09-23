--- tasks
DROP TABLE IF EXISTS tasks;

CREATE TABLE IF NOT EXISTS tasks(
    id INT GENERATED ALWAYS AS IDENTITY,
	-- user_id INT REFERENCES users,
    title VARCHAR(64) NOT NULL,
    description TEXT,
	monday BOOLEAN,
	tuesday BOOLEAN,
	wednesday BOOLEAN,
	thursday BOOLEAN,
	friday BOOLEAN,
	saturday BOOLEAN,
	sunday BOOLEAN,
	PRIMARY KEY(id)
);

INSERT INTO tasks(user_id, title, description, monday, tuesday, wednesday, thursday, friday, saturday, sunday)
    VALUES
        (1, 'Title', 'Description', true, true, true, true, true, true, true),
        (2, 'Title Second', 'Description 2', false, true, true, true, true, true, true);

--- users
DROP TABLE IF EXISTS users;

CREATE TABLE IF NOT EXISTS users(
    id INT GENERATED ALWAYS AS IDENTITY,
    username VARCHAR(64) NOT NULL UNIQUE,
    password  VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL UNIQUE,
	PRIMARY KEY(id)
);

INSERT INTO users(username, password,email)
    VALUES
        ('cfabrica46', '01234', 'cfabrica46@gmail.com');


