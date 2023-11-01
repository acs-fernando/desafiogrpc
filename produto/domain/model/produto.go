package model

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Produto struct {
	Base        `valid:"required"`
	Name        string  `json:"name" gorm:"type:varchar(255)" valid:"notnull"`
	Description string  `json:"description" gorm:"type:varchar(20)" valid:"notnull"`
	Price       float32 `json:"price" gorm:"type:float" valid:"notnull"`
}
