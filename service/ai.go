package service

import (
	"new-ec-dashboard/dao/mysql"
	"new-ec-dashboard/models"
)

func CreateAI(ctx []byte, version, classnames string) error {
	return mysql.CreateAIModel(ctx, version, classnames)
}

func GetListAI() ([]models.AI, error) {
	return mysql.GetListAI()
}