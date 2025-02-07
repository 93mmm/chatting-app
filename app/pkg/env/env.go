package env

import (
    "github.com/joho/godotenv"
)

type server struct {
    Port        int
    ContName    string
}

type database struct {
    Port        int
    Password    string
    ContName    string
}

var CWD                 string
var GoMainServer        *server
var GoDatabaseServer    *server
var PostgresDatabase    *database

func Init() error {
    // TODO: add logic: if unable to open .env file -> file missing
    //                  if one of variables is NULL -> variable missing

    err := getCWD()
    if err != nil {
        return err
    }

    err = godotenv.Load(CWD + ".env")
    if err != nil {
        return err
    }

    err = initGoMainServer()
    if err != nil {
        return err
    }

    err = initDbServer()
    if err != nil {
        return err
    }

    err = initPostgresDb()
    if err != nil {
        return err
    }

    return nil
}
