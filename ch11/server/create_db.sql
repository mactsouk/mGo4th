DROP TABLE IF EXISTS users;

CREATE TABLE users (
    UserID INTEGER PRIMARY KEY,
	username TEXT NOT NULL,
	password TEXT NOT NULL,
	lastlogin INTEGER,
	admin INTEGER,
	active INTEGER
);


INSERT INTO users (username, password, lastlogin, admin, active) VALUES ('admin', 'admin', 1620922454, 1, 1);
