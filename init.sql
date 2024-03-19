CREATE USER docker;
CREATE DATABASE docker;
GRANT ALL PRIVILEGES ON DATABASE docker TO docker;
Create Table IF NOT EXISTS users(
    id serial PRIMARY KEY,
    name VARCHAR(255),
    balance int
);
Create Table IF NOT EXISTS quest(
    id serial PRIMARY KEY,
    name VARCHAR(255),
    cost int,
    stages int
);
Create Table IF NOT EXISTS completed_quests(
    id serial PRIMARY KEY,
    user_id int,
    quest_id int,
    stages int,
    completion_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
