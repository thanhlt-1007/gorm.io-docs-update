package main

import (
    "github.com/thanhlt-1007/gorm.io-docs-update/database"
    UserCrud "github.com/thanhlt-1007/gorm.io-docs-update/models/user/crud"
)

func main() {
    db := database.InitDB()
    defer database.CloseDB(db)

    user := UserCrud.CreateUser(db)
    UserCrud.UpdateUser(db, user)
    UserCrud.UpdateNameAllActiveUsers(db)
}
