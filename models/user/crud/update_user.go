package crud

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-update/models"
    "gorm.io/gorm"
)

func UpdateUser(db *gorm.DB, user models.User) models.User {
    fmt.Println("\n---UpdateUser---")
    fmt.Println("Updating user")

    user.Name = "Update User"
    user.Age = 30
    result := db.Save(&user)
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
