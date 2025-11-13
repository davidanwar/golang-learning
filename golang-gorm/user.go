package golang_gorm

import "time"

type User struct {
	ID           string    `gorm:"primary_key;column:id" json:"id"`
	Name         Name      `gorm:"embedded"`
	Password     string    `gorm:"column:password" json:"password"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" json:"updated_at"`
	Wallet       Wallet    `gorm:"foreignKey:user_id;references:id"`
	Addresses    []Address `gorm:"foreignKey:user_id;references:id"`
	LikeProducts []Product `gorm:"many2many:user_like_product;foreignKey:id;joinForeignKey:user_id;references:id;joinReferences:product_id"`
}

func (user *User) TableName() string {
	return "users"
}

type Name struct {
	FirstName  string `gorm:"column:first_name" json:"first_name"`
	MiddleName string `gorm:"column:middle_name" json:"middle_name"`
	LastName   string `gorm:"column:last_name" json:"last_name"`
}

type UserLog struct {
	ID        int    `gorm:"primary_key;column:id;autoIncrement"`
	UserId    string `gorm:"column:user_id"`
	Action    string `gorm:"column:action"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt int64  `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func (l *UserLog) TableName() string {
	return "user_logs"
}
