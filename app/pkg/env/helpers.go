package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func getInt(env string) (int, error) {
    return strconv.Atoi(os.Getenv(env))
}

func getString(env string) (string, error) {
    value := os.Getenv(env)
    if value == "" {
        return value, errors.New(fmt.Sprintf("Env variable %v doesn't exists", env))
    }
    return value, nil
}

func getCWD() error {
    cwd, err := getString("CHATTING_APP_WORKING_DIRECTORY")
    if err != nil {
        return err
    }
    cwd += "/"
    CWD = cwd
    return nil
}

func initGoMainServer() (error) {
    port, err := getInt("SERVER_PORT")
    if err != nil {
        return err
    }

    name, err := getString("SERVER_CONTAINER_NAME")
    if err != nil {
        return err
    }

    GoMainServer = &server{
        Port: port,
        ContName: name,
    }
    return nil
}

func initDbServer() error {
    port, err := getInt("DB_SERVER_PORT")
    if err != nil {
        return err
    }

    name, err := getString("DB_SERVER_CONTAINER_NAME")
    if err != nil {
        return err
    }

    GoDatabaseServer = &server{
        Port: port,
        ContName: name,
    }
    return nil
}

func initPostgresDb() error {
    user, err := getString("POSTGRES_USER")
    if err != nil {
        return err
    }

    dbName, err := getString("POSTGRES_DB")
    if err != nil {
        return err
    }

    password, err := getString("POSTGRES_PASSWORD")
    if err != nil {
        return err
    }

    contName, err := getString("POSTGRES_CONTAINER_NAME")
    if err != nil {
        return err
    }

    PostgresDatabase = &database{
        User: user,
        DbName: dbName,
        Password: password,
        ContName: contName,
    }
    return nil
}
