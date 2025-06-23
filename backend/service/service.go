package service

import (
	"github.com/google/uuid"
	"github.com/sharukh010/url-shortner/repository"
)

func GenerateShortCode() string {
	shortCode := uuid.New().String()[0:8] 
	for repository.IsLinkExists(shortCode) {
		shortCode = uuid.New().String()[0:8] 
	}
	return shortCode
}