package model

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsEnable bool   `json:"is_enable"`
}

type Questions struct {
	Id       int64  `json:"id,string" gorm:"primary_key" validate:"required"`
	Question string `json:"question" validate:"required"`
	Answer   string `json:"answer" validate:"required"`
	Tag      string `json:"tag" validate:"required"`
}

type Question struct {
	Id       int64  `json:"id,string" gorm:"primary_key"`
	Question string `json:"question"`
}

func (Question) TableName() string {
	return "question"
}
