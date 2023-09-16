package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"new-ec-dashboard/models"
	"new-ec-dashboard/pkg/md5"
)

func QueryAIModel(version string) bool {
	sqlStr := `select * from ai_model where version=?`
	aiModel := new(models.AI)
	err := db.Get(&aiModel, sqlStr, version)
	if err != nil {
		return false
	}
	return true
}

func CreateAIModel(ctx []byte, version, classNames string) (err error) {
	isExist := QueryAIModel(version)
	if !isExist {
		return ErrorAiHasExist
	}
	sum, err := md5.MD5sum(ctx)
	if err != nil {
		return
	}
	aiDB := new(models.AI)
	aiDB.Version = version
	aiDB.ClassName = classNames
	aiDB.File = ctx
	aiDB.FileMd5 = sum

	return InsertAIModel(aiDB)
}

func InsertAIModel(aiModel *models.AI) (err error) {
	insertSql := "insert into ai_model(file,file_md5,class_names,version) values(?,?,?,?)"
	_, err = db.Exec(insertSql, aiModel.File, aiModel.FileMd5, aiModel.ClassName, aiModel.Version)
	if err != nil {
		return err
	}
	return
}

func GetListAI() (aiModels []models.AI, err error) {
	sqlStr := `select * from ai_model`
	if err:= db.Select(&aiModels, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Error("there is no aiModel in db", zap.Error(err))
			err=nil
		}
	}
	return
}
