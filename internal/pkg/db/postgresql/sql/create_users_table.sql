CREATE TABLE IF NOT EXISTS Users(
    ID INT NOT NULL UNIQUE SERIAL PRIMARY KEY,
    Username VARCHAR (127) NOT NULL UNIQUE,
    Password VARCHAR (127) NOT NULL,
)