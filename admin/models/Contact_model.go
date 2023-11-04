package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	NameSurname, Email, Message string
}

func (contact Contact) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.AutoMigrate(&contact)
}
func (contact Contact) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Create(&contact)
}
func (contact Contact) GetAll(where ...interface{}) []Contact {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	var contacts []Contact
	db.Find(&contacts, where...)
	return contacts
}
func (contact Contact) Get(where ...interface{}) Contact {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return contact
	}
	db.First(&contact, where...)
	return contact
}
func (contact Contact) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Delete(&contact, contact.ID)
}
