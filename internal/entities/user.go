package entities

var _table_user = "users"

type UserRegister struct {
	ID         string `json:"id" gorm:"primaryKey;type:varchar(36)" validate:"required,uuid"`
	Username   string `json:"username" gorm:"type:varchar(30);unique" validate:"required,min=3,max=30"`
	Email      string `json:"email" gorm:"type:varchar(100);unique" validate:"required,email"`
	Password   string `json:"password" gorm:"type:varchar(100)" validate:"required,min=6,max=100"`
	LastActive int64  `json:"lastActive"`
	IsActive   bool   `json:"isActive" gorm:"default:true"`
	IsDeleted  bool   `json:"isDeleted" gorm:"default:false"`
	CreatedAt  int64  `json:"createdAt"`
	UpdatedAt  int64  `json:"updatedAt"`
}

func (u *UserRegister) TableName() string {
	return _table_user
}
