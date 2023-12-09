-- db/migrations/20231208120000_create_urls_table.up.sql

CREATE TABLE urls (
                      id SERIAL PRIMARY KEY,
                      longurl VARCHAR(255) UNIQUE NOT NULL,
                      shorturl VARCHAR(10) UNIQUE NOT NULL
);