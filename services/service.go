package services

import (
	"PackageApi/models"
	"fmt"
	"strings"
)

func Servicefunc1(user *models.User) error {
	user.Name = strings.ToUpper(user.Name)
	if err := models.Db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func Servicefunc2(user *models.User, id string) error {
	if err := models.Db.First(&user, id).Error; err != nil {
		return err
	}
	return nil
}

func Service3func(user *models.User, id string) error {
	if err := models.Db.First(&user, id).Error; err != nil {
		return err
	}
	return nil
}
func Service4func() ([]*models.User, error) {
	if models.Db != nil {
		fmt.Println("DB is nil")
	}
	var users []*models.User
	if err := models.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	for i, user := range users {
		users[i].Name = strings.ToLower(user.Name)
	}
	return users, nil
}
