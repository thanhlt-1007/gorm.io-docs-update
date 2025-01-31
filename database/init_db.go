package database

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-update/models"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    fmt.Println("\n---InitDB---")
    fmt.Println("Init DB opening gorm.sqlite")
    db, err := gorm.Open(
        sqlite.Open("gorm.sqlite"),
        &gorm.Config{},
    )
    if err != nil {
        fmt.Printf("Init DB open gorm.sqlite error %v\n", err)
        panic(err)
    }
    fmt.Println("Init DB open gorm.sqlite success")

    fmt.Println("Init DB running migrate")
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        fmt.Printf("Init DB run migrate error %v\n", err)
        panic(err)
    }
    fmt.Println("Init DB run migrate success")

    return db
}
