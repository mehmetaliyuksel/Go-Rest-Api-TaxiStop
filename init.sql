CREATE TABLE IF NOT EXISTS Users (
                          user_id serial PRIMARY KEY,
                          username VARCHAR ( 50 ) UNIQUE NOT NULL,
                          password VARCHAR ( 50 ) NOT NULL,
                          email VARCHAR ( 255 ) UNIQUE NOT NULL,
                          created_on TIMESTAMP NOT NULL DEFAULT NOW(),
                          last_login TIMESTAMP
);

INSERT INTO Users(username, password, email) VALUES ('test', 'test', 'test');

