package entities

type Comment struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	ObjectName string `json:"object_name"`
	Text       string `json:"text"`
	Data       string `json:"data"`
}
