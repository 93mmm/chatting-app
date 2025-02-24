package helpers

import (
	"fmt"

	"github.com/93mmm/chatting-app/app/pkg/env"
)

func GetDistinct(arr []int) []int {
    set := map[int]bool{}
    for _, val := range arr {
        set[val] = true
    }
    arr = make([]int, 0, len(set))
    for k := range set {
        arr = append(arr, k)
    }
    return arr
}

func GetDatabaseUrl() string {
    return fmt.Sprintf("postgres://%v:%v@%v:5432/%v",
                       env.PostgresDatabase.User,
                       env.PostgresDatabase.Password,
                       env.PostgresDatabase.ContName,
                       env.PostgresDatabase.DbName,
                       )
}

