CREATE DATABASE jobs;

CREATE TABLE IF NOT EXISTS joblist (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    date_scraped DATE,
    tech_stack TEXT[],
    misc_info TEXT[]
);

