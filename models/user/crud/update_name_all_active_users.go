package crud

import (
    "fmt"
    "github.com/thanhlt-1007/gorm.io-docs-update/models"
    "gorm.io/gorm"
    "gorm.io/gorm/clause"
)

func UpdateNameAllActiveUsers(db *gorm.DB) {
    fmt.Println("\n---UpdateNameAllActiveUsers---")
    fmt.Println("Updating name all active users")

    var users []models.User
    result := db.
        Model(&users).
        Clauses(clause.Returning{}).
        Where("active = ?", true).
        Update("name", "Active User")
    err := result.Error

    fmt.Println("\n---Result---")
    fmt.Printf("RowsAffected: %d\n", result.RowsAffected)
    fmt.Printf("Error: %v\n", err)

    if err != nil {
        fmt.Printf("Update users error %v\n", err)
        panic(err)
    }

    fmt.Println("Update users success")
    for _, user := range users {
        user.Println()
    }
}
