package service

import (
	"new-ec-dashboard/dao/mysql"
	"new-ec-dashboard/models"
)

func InsertIeData(ieParams *models.IeParams) error {
	return mysql.InsertIeData(ieParams)
}
func GetIeDataList(pagenum,pagesize int64) (page *mysql.MyPage, err error) {
	ieDataList, err := mysql.QueryIeDataList(pagenum,pagesize)
	if err != nil {
		return
	}
	total, err := mysql.GetIeDataTotal()
	if err != nil {
		return
	}
	page=&mysql.MyPage{Data: ieDataList,Total:total}
	return
}