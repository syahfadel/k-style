package services

import (
	"errors"
	"fmt"
	"unnispick/entities"

	"gorm.io/gorm"
)

type UnnispickService struct {
	DB *gorm.DB
}

func (us *UnnispickService) CreateMember(member entities.Member) (entities.Member, error) {
	if err := us.DB.Debug().Create(&member).Error; err != nil {
		return entities.Member{}, err
	}
	return member, nil
}

func (us *UnnispickService) UpdateMember(member entities.Member) (entities.Member, error) {
	res := us.DB.Debug().Model(&member).Where("ID_MEMBER = ?", member.ID).Updates(&member)
	if res.Error != nil {
		return entities.Member{}, res.Error
	}

	if res.RowsAffected == 0 {
		return entities.Member{}, errors.New("no data updated")
	}
	return member, nil
}

func (us *UnnispickService) DeleteMember(MemberId uint) error {
	res := us.DB.Debug().Where("ID_MEMBER = ?", MemberId).Delete(&entities.Member{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("member with ID_MEMBER %v not available", MemberId))
	}
	return nil
}

func (us *UnnispickService) GetAllMember() ([]entities.Member, error) {
	var members []entities.Member
	if err := us.DB.Debug().Find(&members).Error; err != nil {
		return []entities.Member{}, err
	}
	return members, nil
}

func (us *UnnispickService) GetProductById(productId uint) ([]map[string]interface{}, error) {
	var (
		// member        entities.Member
		// product       entities.Product
		// reviewProduct entities.ReviewProduct
		// likeReview    entities.LikeReview
		results []map[string]interface{}
	)

	err := us.DB.Debug().Table("product").
		Select("member.USERNAME as username, member.GENDER as gender, member.SKINTYPE as skin_type, member.SKINCOLOR as skin_color, review_product.DESC_REVIEW as desc_review ,COUNT(like_review.ID_MEMBER) as 'like'").
		Joins("JOIN review_product ON product.ID_PRODUCT = review_product.ID_PRODUCT").
		Joins("LEFT JOIN like_review ON review_product.ID_REVIEW = like_review.ID_REVIEW").
		Joins("JOIN member ON review_product.ID_MEMBER = member.ID_MEMBER").
		Where("product.ID_PRODUCT = ?", productId).
		Group("review_product.ID_REVIEW").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (us *UnnispickService) GetLikeReview(like entities.LikeReview) bool {
	var count int64
	us.DB.Debug().Model(&entities.LikeReview{}).Where("ID_MEMBER = ? AND ID_REVIEW = ?", like.MemberID, like.ReviewProductID).Count(&count)
	return count != 0
}

func (us *UnnispickService) DeleteLike(like entities.LikeReview) error {
	res := us.DB.Debug().Where("ID_MEMBER = ? AND ID_REVIEW = ?", like.MemberID, like.ReviewProductID).Delete(&entities.LikeReview{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New(fmt.Sprintf("member with ID_MEMBER %v and ID_REVIEW %v not available", like.MemberID, like.ReviewProductID))
	}
	return nil
}

func (us *UnnispickService) CreateLike(like entities.LikeReview) error {
	if err := us.DB.Debug().Create(&like).Error; err != nil {
		return err
	}
	return nil
}
