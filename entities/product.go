package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID          uint   `gorm:"primaryKey;column:ID_PRODUCT"`
	NameProduct string `gorm:"not null;column:NAME_PRODUCT" valid:"required"`
	Price       uint   `gorm:"not null;column:PRICE" valid:"required"`
}

func (Product) TableName() string {
	return "product"
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
