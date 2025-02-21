package env

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type server struct {
    Port        int
    ContName    string
}

type database struct {
    User        string
    DbName      string
    Password    string
    ContName    string
}

var CWD                 string
var JsonTestFiles       string
var GoMainServer        *server
var GoDatabaseServer    *server
var PostgresDatabase    *database

func Init() error {
    if len(CWD) != 0 && GoMainServer != nil && GoDatabaseServer != nil && PostgresDatabase != nil {
        return nil
    }
    err := getCWD()
    if err != nil {
        return err
    }

    if _, err = os.Stat(CWD + ".env"); errors.Is(err, os.ErrNotExist) {
        return err;
    }

    err = godotenv.Load(CWD + ".env")
    if err != nil {
        return err
    }

    err = getJsonPath()
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
