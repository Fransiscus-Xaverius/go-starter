package entity

type User struct {
	ID       int32  `json:"id" gorm:"primary_key,column:id"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (u User) TableName() string {
	return "user1"
}
