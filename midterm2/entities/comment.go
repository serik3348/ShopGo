package entities

type Comment struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	ObjectName string `json:"objectname"`
	Text       string `json:"text"`
	Date       string `json:"date"`
}
