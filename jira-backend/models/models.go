package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type Users struct {
	gorm.Model
	Name   string
	field2 uint
	field3 uint
	field4 uint
	field5 uint
}

func Print() {
	fmt.Println("Hello")
}
