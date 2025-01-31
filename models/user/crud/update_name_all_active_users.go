package crud

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-update/models"
    "gorm.io/gorm"
)

func UpdateNameAllActiveUsers(db *gorm.DB) {
    fmt.Println("\n---UpdateNameAllActiveUsers---")
    fmt.Println("Updating name all active users")

    result := db.Model(&models.User{}).Where("active = ?", true).Update("name", "Active User")
    err := result.Error

    fmt.Println("\n---Result---")
    fmt.Printf("RowsAffected: %d\n", result.RowsAffected)
    fmt.Printf("Error: %v\n", err)

    if err != nil {
        fmt.Printf("Update users error %v\n", err)
        panic(err)
    }

    fmt.Println("Update users success")
}
