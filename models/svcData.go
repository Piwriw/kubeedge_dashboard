package models

import "time"

type SvcModel struct {
	Id          int       `json:"id" db:"id"`
	ImageBase64 string    `json:"image_base64" db:"image_base64"`
	ClassName   string    `json:"class_name" db:"class_name"`
	UpLoadTime  time.Time `json:"upload_time" db:"upload_time"`
}
type IeModel struct {
	Id                     int       `json:"id" db:"id"`
	EdgeDeviceId           string    `json:"edge_device_id" db:"edge_device_id"`
	Voltage                float32   `json:"voltage" db:"voltage"`
	ElectricCurrent        float32   `json:"electric_current" db:"electric_current"`
	Power                  float32   `json:"power" db:"power"`
	PowerFactor            float32   `json:"power_factor" db:"power_factor"`
	Frequency              float32   `json:"frequency" db:"frequency"`
	ElectricityConsumption float32   `json:"electricity_consumption" db:"electricity_consumption"`
	UploadTime             time.Time `json:"upload_time" db:"upload_time"`
	Label                  string    `json:"label" db:"label"`
}
