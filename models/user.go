package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name string
    Age uint
}

func (user User) Println() {
    fmt.Println("\n---User---")
    fmt.Printf("ID: %d\n", user.ID)
    fmt.Printf("Name: %s\n", user.Name)
    fmt.Printf("Age: %d\n", user.Age)
}
