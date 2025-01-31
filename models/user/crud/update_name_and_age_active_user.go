package crud

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-update/models"
    "gorm.io/gorm"
)

func UpdateNameAndAgeActiveUser(db *gorm.DB, user models.User) models.User {
    fmt.Println("\n---UpdateNameAndAgeActiveUser---")
    fmt.Println("Updating name and age active user")

    result := db.Model(&user).
        Where("active = ?", true).
        Updates(models.User{Name: "Active User", Age: 30})
    err := result.Error

    fmt.Println("\n---Result---")
    fmt.Printf("RowsAffected: %d\n", result.RowsAffected)
    fmt.Printf("Error: %v\n", err)

    if err != nil {
        fmt.Printf("Update user error %v\n", err)
        panic(err)
    }

    fmt.Println("Update user success")
    user.Println()

    return user
}
