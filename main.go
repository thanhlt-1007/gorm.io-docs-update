package main

import (
    "github.com/thanhlt-1007/gorm.io-docs-update/database"
)

func main() {
    db := database.InitDB()
    database.CloseDB(db)
}
