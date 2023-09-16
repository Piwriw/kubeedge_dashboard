package service

import (
	"new-ec-dashboard/dao/mysql"
	"new-ec-dashboard/models"
)

//func InsertSvcData2(c *gin.Context) sql.Result {
//	var svcModel model.SvcDataModel
//	all, _ := ioutil.ReadAll(c.Request.Body)
//	json.Unmarshal([]byte(string(all)), &svcModel)
//	db := db.GetDBConn()
//	lenimg := len(svcModel.ImageBase64) - 1
//	svc := svcModel.ImageBase64[0:lenimg]
//	insetSQL := "insert into svc_data(image_base64,class_name,upload_time) values(?,?,?)"
//	data, err := db.Exec(insetSQL, svc, &svcModel.ClassName, time.Now())
//	if err != nil {
//		fmt.Println(err)
//		logger.ZapLogger.Error(err.Error(), zap.String("msg", "InsertSvcData Error!"))
//		return nil
//	}
//	return data
//}

func InsertSvcData(svcDataParams *models.SvcDataParams) error {
	return mysql.InsertSvcData(svcDataParams)
}
func GetSvcDataList() ([]models.SvcModel, error) {
	return mysql.QuerySvcDataList()
}
