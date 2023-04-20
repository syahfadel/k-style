package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ReviewProduct struct {
	ID          uint   `gorm:"primaryKey;column:ID_REVIEW"`
	ProductID   uint   `gorm:"column:ID_PRODUCT"`
	MemberID    uint   `gorm:"column:ID_MEMBER"`
	DescReview  string `gorm:"not null;column:DESC_REVIEW" valid:"required"`
	LikeReviews []LikeReview
}

func (ReviewProduct) TableName() string {
	return "review_product"
}

func (rp *ReviewProduct) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(rp)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return

}
func (rp *ReviewProduct) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(rp)
	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
