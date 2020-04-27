CREATE TABLE db.users (
    id INT PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL
);

INSERT INTO db.users (id, email, name) VALUES
(1, 'email1@mysql.com', 'user1'),
(2, 'email2@mysql.com', 'user2'),
(3, 'email3@mysql.com', 'user3')
