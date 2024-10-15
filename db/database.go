package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const schema = `
    CREATE TABLE IF NOT EXISTS job_listings (
            id SERIAL PRIMARY KEY,
            url TEXT NOT NULL,
            tech_stack TEXT[],
            misc_info TEXT[],
            salary TEXT,
            date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`

func InitDB() (*sql.DB, error) {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    dbHost := os.Getenv("DB_HOST")
    fmt.Println(dbHost)
    dbUser := os.Getenv("DB_USER")
    dbPasswd := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPasswd + " dbname=" + dbName + " sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    defer db.Close()


    err = db.Ping()
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    _, err = db.Exec(schema)
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    fmt.Println("Connected, schema generated")
    return db, nil
}


