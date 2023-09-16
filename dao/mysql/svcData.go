package mysql

import (
	"database/sql"
	"new-ec-dashboard/models"
	"time"
)


func InsertSvcData(svcDataParams *models.SvcDataParams) (err error) {
	insetSQL := `insert into svc_data(image_base64,class_name,upload_time) values(?,?,?)`
	if err != nil {
		return
	}
	_, err = db.Exec(insetSQL, []byte(svcDataParams.ImageBase64[:len(svcDataParams.ImageBase64)-1]), svcDataParams.ClassName, time.Now())
	if err != nil {
		return
	}
	return
}

func QuerySvcDataList() ([]models.SvcModel, error) {
	querySQL := `select * from svc_data`
	var svc []models.SvcModel
	err := db.Select(&svc, querySQL)
	//if err == sql.ErrNoRows {
	//	return
	//}
	return svc, err
}

func InsertIeData(ie *models.IeParams) (err error) {
	insetSQL := `insert into ie_data(edge_device_id,voltage,electric_current,power,power_factor,frequency,
                label,upload_time,electricity_consumption) values(?,?,?,?,?,?,?,?,?)`
	_, err = db.Exec(insetSQL, ie.EdgeDeviceId, ie.Voltage, ie.ElectricCurrent, ie.Power, ie.
		PowerFactor, ie.Frequency, ie.Label, time.Now(), ie.ElectricityConsumption)
	if err != nil {
		return
	}
	return
}
func QueryIeDataList(pagenum, pagesize int64) (ieList []models.IeModel, err error) {
	querySQL := `select * from ie_data  order by upload_time desc limit ?,?`
	ieList = make([]models.IeModel, 0)
	err = db.Select(&ieList, querySQL, (pagenum-1)*pagesize, pagesize)
	if err == sql.ErrNoRows {
		return
	}
	return
}
func GetIeDataTotal() (int, error) {
	sqlStr := `select count(id) from ie_data`
	var count int
	err := db.Get(&count, sqlStr)
	if err != nil {
		return 0, err
	}
	return count, err
}
