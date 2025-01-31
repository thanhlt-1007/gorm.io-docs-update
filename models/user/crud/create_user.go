package crud

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-update/models"
    "gorm.io/gorm"
)

func CreateUser(db *gorm.DB) models.User {
    fmt.Println("\n---CreateUser---")
    fmt.Println("Creating user")

    user := models.User {
        Name: "User",
        Age: 18,
        Active: true,
    }
    result := db.Create(&user)
    err := result.Error

    fmt.Println("\n---Result---")
    fmt.Printf("RowsAffected: %d\n", result.RowsAffected)
    fmt.Printf("Error: %v\n", err)

    if err != nil {
        fmt.Printf("Create user error %v\n", err)
        panic(err)
    }

    fmt.Println("Create user success")
    user.Println()

    return user
}
