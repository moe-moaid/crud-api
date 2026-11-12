package services

import "example/crud-api/internal/models"

type CustomUsers struct {
	users []models.User 
}

func NewUserHandler() *CustomUsers{
	return &CustomUsers{}
}

func List(s *CustomUsers) []models.User {
	return s.users
}