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
            date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );`

func Connect() error {
    err := godotenv.Load()
    if err != nil {
        log.Fatal(err)
        return err
    }
    dbHost := os.Getenv("DB_HOST")
    fmt.Println(dbHost)
    dbUser := os.Getenv("DB_USER")
    dbPasswd := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")

    // connStr := fmt.Sprintf("host=%s port=5432 user=%s password=%s name=%s sslmode=disable", dbHost, dbPort, dbUser, dbPasswd, dbName)
    connStr := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPasswd + " dbname=" + dbName + " sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
        return err
    }
    defer db.Close()


    err = db.Ping()
    if err != nil {
        log.Fatal(err)
        return err
    }
    _, err = db.Exec(schema)
    if err != nil {
        log.Fatal(err)
        return err
    }
    fmt.Println("Connected, schema generated")
    return nil 
}
