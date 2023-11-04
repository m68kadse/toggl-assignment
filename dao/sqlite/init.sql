-- this can be run on start as it will not change the db schema if it already exists


CREATE TABLE IF NOT EXISTS question(
    id PRIMARY KEY,
    body TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS "option"(
    id PRIMARY KEY,
    fk_question INTEGER NOT NULL,
    body TEXT NOT NULL,
    correct INTEGER NOT NULL,
    FOREIGN KEY(fk_question) REFERENCES question(id) ON DELETE CASCADE
);