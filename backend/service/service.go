package service

import "github.com/google/uuid"

func GenerateShortCode() string {
	return uuid.NewString()[1:6]
}