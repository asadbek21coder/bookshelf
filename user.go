package bookshelf

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name" binding:"required"`
	Key    string `json:"key" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}
