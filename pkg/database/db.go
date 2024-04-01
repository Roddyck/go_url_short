package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitDB(name string) error {
    connInfo := "user=postgres password=postgres host=localhost sslmode=disable"
    db, err := sql.Open("postgres", connInfo)

    if err != nil {
        return err
    }

    _, err = db.Exec("CREATE TABLE IF NOT EXISTS urls (id SERIAL, short_url varchar(20), origin_url varchar(400), PRIMARY KEY (id))")
    if err != nil {
        return err
    }

    Db = db

    return nil
}

func AddUrl(short, origin string) error {
    stmt, err := Db.Prepare("INSERT INTO urls (short_url, origin_url) VALUES ($1, $2)")

    if err != nil {
        return err
    }

    if _, err := stmt.Exec(short, origin); err != nil {
        return err
    }

    return nil
}

func GetUrl(short string) (string, error) {
    stmt, err := Db.Prepare("SELECT origin_url FROM urls WHERE short_url=$1")
    if err != nil {
        return "", err
    }

    var origin string
    if err := stmt.QueryRow(short).Scan(&origin); err != nil {
        return "", err
    }

    return origin, nil
}
