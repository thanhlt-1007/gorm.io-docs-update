package database

import (
    "fmt"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    fmt.Println("---InitDB---")
    fmt.Println("Init DB opening gorm.sqlite")
    db, err := gorm.Open(
        sqlite.Open("gorm.sqlite"),
        &gorm.Config{},
    )
    if err != nil {
        fmt.Printf("Init DB open gorm.sqlite error %v\n", err)
        panic(err)
    }
    fmt.Println("InitDB success")

    return db
}
