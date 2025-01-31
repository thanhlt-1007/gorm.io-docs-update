package main

import (
    "github.com/thanhlt-1007/gorm.io-docs-update/database"
    UserCrud "github.com/thanhlt-1007/gorm.io-docs-update/models/user/crud"
)

func main() {
    db := database.InitDB()
    defer database.CloseDB(db)

    UserCrud.CreateUser(db)
}
