package models

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	UserName string `json:"userName" binding:"required"`
	Email    string `json:"email" binding:"required"`
}