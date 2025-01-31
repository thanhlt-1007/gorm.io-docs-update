package crud

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-update/models"
    "gorm.io/gorm"
)

func UpdateNameActiveUser(db *gorm.DB, user models.User) models.User {
    fmt.Println("\n---UpdateNameActiveUser---")
    fmt.Println("Updating name active user")

    result := db.Model(&user).Where("active = ?", true).Update("name", "Active User")
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
