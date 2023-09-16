package mysql

type MyPage struct {
	Data  interface{} `json:"data"`
	Total int         `json:"total"`
}
