package main

import (
    "database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Dependencies struct {
    Database *bun.DB
    TestString *string

}

func initializeAll() *Dependencies {
    inject := Dependencies{
        Database: provideDatabaseConnection(),
        TestString: provideTestString(),
    }

    return &inject
}

func provideDatabaseConnection() *bun.DB {
    dsn := "postgres://" + 
        POSTGRES_USER + ":" + 
        POSTGRES_PASSWORD + "@" + 
        POSTGRES_HOST + ":" + 
        POSTGRES_PORT + "/" + 
        POSTGRES_DATABASE_NAME + "?sslmode=disable"

    sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

    db := bun.NewDB(sqldb, pgdialect.New())

    return db
}

func provideTestString() *string {
    testString := "this should be injected"

    return &testString
}
