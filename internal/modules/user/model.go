package user

type User struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
}

func (User) TableName() string {
	return "users"
}
