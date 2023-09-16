package models

type AI struct {
	Id        int    `json:"id" db:"id"`
	File      []byte `json:"file" db:"file"`
	Version   string `json:"version" db:"version"`
	ClassName string `json:"class_names" db:"class_names"`
	FileMd5   string `json:"file_md5" db:"file_md5"`
}
