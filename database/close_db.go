package database

import (
    "fmt"
    "gorm.io/gorm"
)

func CloseDB(db *gorm.DB) {
    fmt.Println("\n---CloseDB---")
    fmt.Println("Closing DB")
    sqlDB, err := db.DB()
    if err != nil {
        fmt.Printf("Close DB error %v\n", err)
        panic(err)
    }
    sqlDB.Close()
    fmt.Println("Close DB success")
}
