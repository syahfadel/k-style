package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type LikeReview struct {
	ReviewProductID uint `gorm:"column:ID_REVIEW" valid:"required"`
	MemberID        uint `gorm:"column:ID_MEMBER" valid:"required"`
}

func (LikeReview) TableName() string {
	return "like_review"
}

func (lr *LikeReview) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(lr)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
func (lr *LikeReview) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(lr)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
