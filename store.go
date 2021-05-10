package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Extension struct {
	Id   int `gorm:"primary_key"`
	Name string
	Desc string
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Extension{})
}

func main() {
	extension1 := Extension{Name: "a", Desc: "a-a"}
	Db.Create(&extension1)
	extension2 := Extension{Name: "b", Desc: "b-b"}
	Db.Create(&extension2)
	extension3 := Extension{Name: "c", Desc: "c-c"}
	Db.Create(&extension3)
}
