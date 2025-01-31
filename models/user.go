package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string
    Age uint
    Active bool
}

func (user User) Println() {
    fmt.Println("\n---User---")
    fmt.Printf("ID: %d\n", user.ID)
    fmt.Printf("Name: %s\n", user.Name)
    fmt.Printf("Age: %d\n", user.Age)
    fmt.Printf("Active: %t\n", user.Active)
}
