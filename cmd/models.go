package main

import "gorm.io/gorm"

type Movie struct{
	gorm.Model
	Title string `json:"title" gorm:"text;not null;default:null"`
	Genre string `json:"genre" gorm:"text;not null;default:null"`
}