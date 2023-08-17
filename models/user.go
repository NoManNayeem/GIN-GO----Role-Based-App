package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique" json:"username"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

// Set up table name for the User model
func (User) TableName() string {
	return "users"
}
