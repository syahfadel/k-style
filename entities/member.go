package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Member struct {
	ID             uint   `gorm:"primaryKey;column:ID_MEMBER"`
	Username       string `gorm:"not null;uniqueIndex;column:USERNAME" valid:"required"`
	Gender         string `gorm:"not null;column:GENDER" valid:"required"`
	Skintype       string `gorm:"column:SKINTYPE"`
	Skincolor      string `gorm:"column:SKINCOLOR"`
	ReviewProducts []ReviewProduct
	LikeReviews    []LikeReview
}

func (Member) TableName() string {
	return "member"
}

func (m *Member) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
func (m *Member) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
