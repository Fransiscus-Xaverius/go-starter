package entity

type User struct {
	ID        int32  `json:"id" gorm:"primary_key,column:id"`
	Name      string `json:"name" gorm:"column:name"`
	Email     string `json:"email" gorm:"column:email"`
	Password  string `json:"password" gorm:"column:password"`
	CreatedAt int64  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt int64  `json:"updated_at" gorm:"column:updated_at"`
	CreatedBy int64  `json:"created_by" gorm:"column:created_by"`
	UpdatedBy int64  `json:"updated_by" gorm:"column:updated_by"`
}

func (u User) TableName() string {
	return "user"
}
